package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/ishumei/krpc/autolimit"
	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/kconf"
	"github.com/ishumei/krpc/kframe/sconfig"
	"github.com/ishumei/krpc/objects"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	"go.uber.org/ratelimit"
)

var (
	hostports   = flag.String("hostports", "", "hostports")
	metabase    = flag.String("metabase", "127.0.0.1:2181", "metabase")
	qps         = flag.Int("qps", 1, "qps")
	requestFile = flag.String("request_file", "request.json", "request file")
	serviceName = flag.String("service_name", "/test/models/boilerplate", "service name")
	timeout     = flag.Int("timeout", 1000, "timeout")
)

func main() {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		klog.Info(f.Name, ": ", f.Value)
	})
	client := kclient.MustNewArbiterClient(func() *kclient.SingleClientConf {
		c := &sconfig.FrameConfig{}
		err := kconf.LoadDefaultConf(c, "frame", "overwrite.yaml")
		klog.Info(objects.String(c))
		return &kclient.SingleClientConf{
			ResolverConf: kclient.ResolverConf{
				Hostports: func() []string {
					if *hostports != "" {
						return strings.Split(*hostports, ",")
					}
					return []string{}
				}(),
				Resolver: func() registry_zookeeper.Conf {
					if err == nil {
						return c.Registry
					}
					return registry_zookeeper.Conf{
						Metabase:  *metabase,
						TimeoutMs: 10000,
					}
				}(),
			},
			ClientConf: kclient.ClientConf{
				ServiceName: func() string {
					if err == nil {
						return c.ServiceName
					}
					return *serviceName
				}(),
				Retries:   0,
				TimeoutMs: *timeout,
			},
		}
	}())
	client.Start()
	defer func() {
		_ = client.Shutdown()
	}()
	r := &service.PredictRequest{}
	bs, err := os.ReadFile(*requestFile)
	if err != nil {
		klog.Error("Read request file error:", err)
	}
	if err := sonic.Unmarshal(bs, r); err != nil {
		klog.Error("Unmarshal request file error:", err)
	}
	rl := ratelimit.New(*qps)
	for {
		rl.Take()
		gopool.Go(func() {
			_, err := client.Predict(context.Background(), r)
			if err != nil {
				klog.Error(err)
			}
		})
	}
}
