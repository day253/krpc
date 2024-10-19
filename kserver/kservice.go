package kserver

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/logging"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/samber/do"
)

type Kservice struct {
	provider   provider.OtelProvider
	server     server.Server
	logger     *logging.Logger
	httpServer *HttpServer
	exitSignal func()
}

func (c *Kservice) Start() {
	klog.Info("starting kservice")
	if c.httpServer != nil {
		go func() {
			c.httpServer.Start()
		}()
	}
	c.exitSignal()
	if err := c.server.Run(); err != nil {
		panic(err)
	}
}

func (c *Kservice) Shutdown() error {
	klog.Info("shutdown kservice")
	var err error
	if c.provider != nil {
		err = c.provider.Shutdown(context.Background())
		if err != nil {
			return err
		}
	}
	return err
}

func MustNewKservice(i *do.Injector, server server.Server) *Kservice {
	var otelprovider provider.OtelProvider

	c := do.MustInvoke[*FrameConfig](Injector)

	if c.OpenTelemetry.Enabled {
		otelprovider = provider.NewOpenTelemetryProvider(
			provider.WithServiceName(c.ServiceName),
			provider.WithExportEndpoint(c.OpenTelemetry.Address),
			provider.WithInsecure(),
		)
	}

	logger := do.MustInvoke[*logging.Logger](logging.Injector)

	var httpServer *HttpServer
	if c.Http.Enabled {
		httpServer = NewHttpServer(c.HttpAddress())
	}

	return &Kservice{
		provider:   otelprovider,
		server:     server,
		logger:     logger,
		httpServer: httpServer,
		exitSignal: DefaultDeregisterSignal,
	}
}
