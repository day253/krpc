package prometheus

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Labels
const (
	labelKeyCaller = "caller"
	labelKeyCallee = "callee"
	labelKeyDetail = "detail"
	labelKeyMethod = "method"
	labelKeyStatus = "status"
	labelKeyRetry  = "retry"

	// status
	statusSucceed = "succeed"
	statusError   = "error"

	unknownLabelValue = "unknown"
)

var (
	globalClientHandledCounter = prom.NewCounterVec(
		prom.CounterOpts{
			Name: "arch_client_throughput",
			Help: "Total number of RPCs completed by the client, regardless of success or failure.",
		},
		[]string{labelKeyCaller, labelKeyCallee, labelKeyDetail, labelKeyMethod, labelKeyStatus, labelKeyRetry},
	)
	globalClientHandledHistogram = prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "arch_client_latency_us",
			Help:    "Latency (microseconds) of the RPC until it is finished.",
			Buckets: []float64{5000, 10000, 25000, 50000, 100000, 250000, 500000, 1000000},
		},
		[]string{labelKeyCaller, labelKeyCallee, labelKeyDetail, labelKeyMethod, labelKeyStatus, labelKeyRetry},
	)
	globalServerHandledCounter = prom.NewCounterVec(
		prom.CounterOpts{
			Name: "arch_server_throughput",
			Help: "Total number of RPCs completed by the server, regardless of success or failure.",
		},
		[]string{labelKeyCaller, labelKeyCallee, labelKeyDetail, labelKeyMethod, labelKeyStatus, labelKeyRetry},
	)
	globalServerHandledHistogram = prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "arch_server_latency_us",
			Help:    "Latency (microseconds) of RPC that had been application-level handled by the server.",
			Buckets: []float64{5000, 10000, 25000, 50000, 100000, 250000, 500000, 1000000},
		},
		[]string{labelKeyCaller, labelKeyCallee, labelKeyDetail, labelKeyMethod, labelKeyStatus, labelKeyRetry},
	)
)

type errDetail struct {
	rpcinfo.RPCInfo
}

func (e errDetail) error() error {
	err := e.Stats().Error()
	if err == nil {
		return nil
	}
	if v, ok := err.(*kerrors.DetailedError); ok {
		return v.ErrorType()
	}
	// maybe baseError
	if kerrors.IsKitexError(err) {
		return err
	}
	// unknown error
	return kerrors.ErrInternalException
}

func (e errDetail) Detail() string {
	err := e.error()
	if err == nil {
		return statusSucceed
	}
	return err.Error()
}

// genLabels make labels values.
func genLabels(ri rpcinfo.RPCInfo) prom.Labels {
	var (
		labels = make(prom.Labels)

		caller = ri.From()
		callee = ri.To()
	)
	labels[labelKeyCaller] = defaultValIfEmpty(caller.ServiceName(), unknownLabelValue)
	labels[labelKeyCallee] = defaultValIfEmpty(callee.ServiceName(), unknownLabelValue)
	labels[labelKeyMethod] = defaultValIfEmpty(callee.Method(), unknownLabelValue)

	labels[labelKeyStatus] = statusSucceed
	if ri.Stats().Error() != nil {
		labels[labelKeyStatus] = statusError
	}
	labels[labelKeyDetail] = errDetail{ri}.Detail()

	labels[labelKeyRetry] = "0"
	if retriedCnt, ok := callee.Tag(rpcinfo.RetryTag); ok {
		labels[labelKeyRetry] = retriedCnt
	}

	return labels
}

type clientTracer struct {
	clientHandledCounter   *prom.CounterVec
	clientHandledHistogram *prom.HistogramVec
}

// Start record the beginning of an RPC invocation.
func (c *clientTracer) Start(ctx context.Context) context.Context {
	return ctx
}

// Finish record after receiving the response of server.
func (c *clientTracer) Finish(ctx context.Context) {
	ri := rpcinfo.GetRPCInfo(ctx)
	if ri.Stats().Level() == stats.LevelDisabled {
		return
	}
	rpcStart := ri.Stats().GetEvent(stats.RPCStart)
	rpcFinish := ri.Stats().GetEvent(stats.RPCFinish)
	cost := rpcFinish.Time().Sub(rpcStart.Time())

	extraLabels := make(prom.Labels)
	extraLabels[labelKeyStatus] = statusSucceed
	if ri.Stats().Error() != nil {
		extraLabels[labelKeyStatus] = statusError
	}

	err := counterAdd(c.clientHandledCounter, 1, genLabels(ri))
	if err != nil {
		klog.CtxDebugf(ctx, "%s", err.Error())
	}
	err = histogramObserve(c.clientHandledHistogram, cost, genLabels(ri))
	if err != nil {
		klog.CtxDebugf(ctx, "%s", err.Error())
	}
}

// NewClientTracer provide tracer for client call, addr and path is the scrape_configs for prometheus server.
func NewClientTracer(addr, path string) stats.Tracer {
	http.Handle(path, promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal("Unable to start a promhttp server, err: " + err.Error())
		}
	}()
	return NewClientTracerWithoutExport()
}

// NewClientTracerWithoutExport provide tracer for client call.
func NewClientTracerWithoutExport() stats.Tracer {
	return &clientTracer{
		clientHandledCounter:   globalClientHandledCounter,
		clientHandledHistogram: globalClientHandledHistogram,
	}
}

type serverTracer struct {
	serverHandledCounter   *prom.CounterVec
	serverHandledHistogram *prom.HistogramVec
}

// Start record the beginning of server handling request from client.
func (c *serverTracer) Start(ctx context.Context) context.Context {
	return ctx
}

// Finish record the ending of server handling request from client.
func (c *serverTracer) Finish(ctx context.Context) {
	ri := rpcinfo.GetRPCInfo(ctx)
	if ri.Stats().Level() == stats.LevelDisabled {
		return
	}

	rpcStart := ri.Stats().GetEvent(stats.RPCStart)
	rpcFinish := ri.Stats().GetEvent(stats.RPCFinish)
	cost := rpcFinish.Time().Sub(rpcStart.Time())

	extraLabels := make(prom.Labels)
	extraLabels[labelKeyStatus] = statusSucceed
	if ri.Stats().Error() != nil {
		extraLabels[labelKeyStatus] = statusError
	}

	err := counterAdd(c.serverHandledCounter, 1, genLabels(ri))
	if err != nil {
		klog.CtxDebugf(ctx, "%s", err.Error())
	}
	err = histogramObserve(c.serverHandledHistogram, cost, genLabels(ri))
	if err != nil {
		klog.CtxDebugf(ctx, "%s", err.Error())
	}
}

// NewServerTracer provides tracer for server access, addr and path is the scrape_configs for prometheus server.
func NewServerTracer(addr, path string) stats.Tracer {
	http.Handle(path, promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal("Unable to start a promhttp server, err: " + err.Error())
		}
	}()
	return NewServerTracerWithoutExport()
}

// NewServerTracer provides tracer for server access, addr and path is the scrape_configs for prometheus server.
func NewServerTracerWithoutExport() stats.Tracer {
	return &serverTracer{
		serverHandledCounter:   globalServerHandledCounter,
		serverHandledHistogram: globalServerHandledHistogram,
	}
}

func defaultValIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func removeDynamicDetail(val, def string) string {
	patterns := []string{
		`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`,
	}
	for _, pattern := range patterns {
		match, _ := regexp.MatchString(pattern, val)
		if match {
			return pattern
		}
	}
	return defaultValIfEmpty(val, def)
}

func init() {
	prom.MustRegister(globalClientHandledCounter)
	prom.MustRegister(globalClientHandledHistogram)
	prom.MustRegister(globalServerHandledCounter)
	prom.MustRegister(globalServerHandledHistogram)
}
