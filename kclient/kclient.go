package kclient

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/logging"
	"github.com/day253/krpc/zookeeper"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/samber/do"
)

type Kclient struct {
	provider provider.OtelProvider
	client.Suite
}

func (c *Kclient) Start() {
	defer klog.Info("starting kclient")
	klog.SetLevel(klog.LevelDebug)
}

func (c *Kclient) Shutdown() error {
	defer klog.Info("shutdown kclient")
	var err error
	if c.provider != nil {
		err = c.provider.Shutdown(context.Background())
		if err != nil {
			return err
		}
	}
	return err
}

func MustNewKclient(c *SingleClientConf) *Kclient {
	var suits client.Suite
	if len(c.Hostports) > 0 {
		suits = MustNewClientOptionsWithoutResolver(c, client.WithHostPorts(c.Hostports...))
	} else {
		do.Override(Injector, func(i *do.Injector) (discovery.Resolver, error) {
			logger, err := do.Invoke[*logging.Logger](logging.Injector)
			if err == nil {
				return zookeeper.NewZookeeperResolverWithConf(
					c.ResolverConf.Resolver,
					zookeeper.WithLogger(logger),
				)
			} else {
				return zookeeper.NewZookeeperResolverWithConf(
					c.ResolverConf.Resolver,
				)
			}
		})
		suits = MustNewClientOptionsWithoutResolver(c, client.WithResolver(do.MustInvoke[discovery.Resolver](Injector)))
	}
	var otelprovider provider.OtelProvider
	if c.OpenTelemetry.Enabled {
		otelprovider = provider.NewOpenTelemetryProvider(
			provider.WithServiceName(c.ServiceName),
			provider.WithExportEndpoint(c.OpenTelemetry.Address),
			provider.WithInsecure(),
		)
	}
	return &Kclient{
		provider: otelprovider,
		Suite:    suits,
	}
}
