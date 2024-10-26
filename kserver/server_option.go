package kserver

import (
	"net"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/system"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/kitex"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	_ "github.com/ishumei/krpc/autolimit"
	"github.com/ishumei/krpc/monitor"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/samber/do"
	"github.com/samber/lo"
)

type ServerOptions struct {
	sOpts []server.Option
}

func (s *ServerOptions) Options() []server.Option {
	return s.sOpts
}

func NewServerOptions(i *do.Injector) (*ServerOptions, error) {
	c := do.MustInvoke[*FrameConfig](Injector)

	addr, _ := net.ResolveTCPAddr("tcp", c.Address())

	options := []server.Option{
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.ServiceName}),
		server.WithTracer(monitor.NewServerTracerWithoutExport()),
		server.WithExitSignal(DefaultUserExitSignal),
	}

	if c.Registry.Enabled {
		iRegistry := do.MustInvoke[registry.Registry](Injector)
		options = append(
			options,
			server.WithRegistry(iRegistry),
		)
	}

	if c.OpenTelemetry.Enabled {
		options = append(
			options,
			server.WithSuite(tracing.NewServerSuite()),
		)
	}

	if c.Sentinel.Enabled {
		options = append(
			options,
			server.WithMiddleware(sentinel.SentinelServerMiddleware()),
		)
		lo.Must0(api.InitDefault())
		_, err := system.LoadRules(c.Sentinel.ToSystemRules())
		lo.Must0(err)
	}

	return &ServerOptions{options}, nil
}

func init() {
	do.Provide(Injector, NewServerOptions)
}
