package kserver

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/ishumei/krpc/logging"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	"github.com/samber/do"
)

func init() {
	do.Provide(Injector, func(i *do.Injector) (*registry_zookeeper.ZookeeperResolver, error) {
		c := do.MustInvoke[*FrameConfig](Injector).Registry
		logger := do.MustInvoke[*logging.Logger](logging.Injector)
		return registry_zookeeper.NewZookeeperResolverWithConf(
			c,
			registry_zookeeper.WithLogger(logger),
		)
	})
	do.Provide(Injector, func(i *do.Injector) (discovery.Resolver, error) {
		return do.Invoke[*registry_zookeeper.ZookeeperResolver](Injector)
	})
}
