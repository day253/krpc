package handler

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/ishumei/krpc/frame/governance"
	"github.com/ishumei/krpc/frame/sconfig"
	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/objects"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"github.com/ishumei/krpc/registry-zookeeper/registry"
	"github.com/jncornett/doublebuf"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"github.com/wI2L/jsondiff"
	"go.uber.org/ratelimit"
)

var (
	clientsCache *doublebuf.DoubleBuffer[*mirrorClients]
	requests     chan *service.PredictRequest
	qps          = flag.Int("qps", 0, "qps")
)

type mirrorClients []predictor.Client

func (s *mirrorClients) mirror(ctx context.Context, request *service.PredictRequest) []*service.PredictResult_ {
	p := pool.NewWithResults[*service.PredictResult_]()
	for _, client := range *s {
		client := client
		p.Go(func() *service.PredictResult_ {
			res, err := client.Predict(ctx, request)
			if err != nil {
				klog.Error(err)
				return &service.PredictResult_{}
			}
			return res
		})
	}
	return p.Wait()
}

func unmarshalMap(source []byte) map[string]interface{} {
	res := make(map[string]interface{})
	_ = sonic.Unmarshal(source, &res)
	return res
}

func (s *mirrorClients) diff(results []*service.PredictResult_) [][]jsondiff.Operation {
	resultsString := lo.Map(results, func(result *service.PredictResult_, _ int) []byte {
		resultMap := unmarshalMap(objects.Bytes(result))
		resultMap["detail"] = unmarshalMap([]byte(result.GetDetail()))
		return objects.Bytes(resultMap)
	})
	if len(resultsString) < 2 {
		return [][]jsondiff.Operation{}
	}
	res := make([][]jsondiff.Operation, 0)
	for i := 1; i < len(resultsString); i++ {
		res = append(
			res,
			func() []jsondiff.Operation {
				patches, _ := jsondiff.CompareJSON(resultsString[0], resultsString[i])
				return patches
			}(),
		)
	}
	return res
}

func (s *mirrorClients) count(diffs [][]jsondiff.Operation) {
	for _, diff := range diffs {
		if len(diff) > 0 {
			klog.Info(objects.String(diff))
		}
	}
}

func (s *mirrorClients) Predict(ctx context.Context, request *service.PredictRequest) (*service.PredictResult_, error) {
	results := s.mirror(ctx, request)
	count := len(results)
	switch {
	case count >= 2:
		s.count(s.diff(results))
		fallthrough
	case count >= 1:
		return results[0], nil
	default:
		return &service.PredictResult_{}, nil
	}
}

func MustNew() *mirrorClients {
	zkConn := do.MustInvoke[*registry.ZookeeperRegistry](sconfig.Injector)
	sConf := do.MustInvoke[*sconfig.FrameConfig](sconfig.Injector)
	childNodes, _, err := zkConn.Children(sConf.ServiceName)
	lo.Must0(err)
	localIp, err := registry.GetLocalIp("")
	lo.Must0(err)
	localIpPort := fmt.Sprintf("%s:%d", localIp, sConf.Port)
	klog.Info("selfIpPort: ", localIpPort)
	var clients mirrorClients = make([]predictor.Client, 0)
	for _, childNode := range childNodes {
		if childNode == localIpPort {
			continue
		}
		klog.Info(sConf.ServiceName, ": ", childNode)
		clients = append(
			clients,
			kclient.MustNewArbiterClient(
				&kclient.SingleClientConf{
					ResolverConf: kclient.ResolverConf{
						Hostports: []string{childNode},
					},
					ClientConf: kclient.ClientConf{
						ServiceName:      sConf.ServiceName,
						Retries:          0,
						ConnectTimeoutMs: 1000,
						TimeoutMs:        3000,
					},
				},
			),
		)
	}
	return &(clients)
}

type ArbiterPredictorImpl struct{}

func (s *ArbiterPredictorImpl) Predict(ctx context.Context, request *service.PredictRequest) (resp *service.PredictResult_, err error) {
	gopool.Go(func() {
		for {
			select {
			case requests <- request:
				continue
			case <-time.After(time.Second * 1):
				return
			}
		}
	})
	return clientsCache.Front().Predict(ctx, request)
}

func (s *ArbiterPredictorImpl) Health(ctx context.Context) (resp bool, err error) {
	return true, nil
}

func BackgroundTask() {
	klog.Info("QPS: ", *qps)
	if *qps <= 0 {
		return
	}
	maxProcs := runtime.GOMAXPROCS(0)
	requests = make(chan *service.PredictRequest, (*qps)*maxProcs)
	for i := 0; i < maxProcs; i++ {
		go func() {
			rl := ratelimit.New(*qps / maxProcs)
			for request := range requests {
				rl.Take()
				c := clientsCache.Front()
				c.count(c.diff(c.mirror(context.Background(), request)))
			}
		}()
	}
}

func init() {
	do.Provide(sconfig.Injector, func(i *do.Injector) (service.Predictor, error) {
		return new(ArbiterPredictorImpl), nil
	})
	clientsCache = doublebuf.New(MustNew(), MustNew())
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			back, err := clientsCache.Back(context.Background())
			if err != nil {
				klog.Error(err)
				continue
			}
			*back = MustNew()
			clientsCache.Ready()
			clientsCache.Next()
		}
	}()
}
