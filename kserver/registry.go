package kserver

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/ishumei/krpc/logging"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	"github.com/samber/do"
)

func init() {
	do.Provide(Injector, func(i *do.Injector) (*registry_zookeeper.ZookeeperRegistry, error) {
		f := do.MustInvoke[*FrameConfig](Injector)
		c := f.Registry
		logger := do.MustInvoke[*logging.Logger](logging.Injector)
		return registry_zookeeper.NewZookeeperRegistryWithConf(
			c,
			f.Addr,
			registry_zookeeper.WithLogger(logger),
		)
	})
	do.Provide(Injector, func(i *do.Injector) (registry.Registry, error) {
		return do.Invoke[*registry_zookeeper.ZookeeperRegistry](Injector)
	})
}
