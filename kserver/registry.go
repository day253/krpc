package kserver

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/ishumei/krpc/logging"
	"github.com/ishumei/krpc/zookeeper"
	"github.com/samber/do"
)

func init() {
	do.Provide(Injector, func(i *do.Injector) (*zookeeper.ZookeeperRegistry, error) {
		f := do.MustInvoke[*FrameConfig](Injector)
		c := f.Registry
		logger := do.MustInvoke[*logging.Logger](logging.Injector)
		return zookeeper.NewZookeeperRegistryWithConf(
			c,
			f.Addr,
			zookeeper.WithLogger(logger),
		)
	})
	do.Provide(Injector, func(i *do.Injector) (registry.Registry, error) {
		return do.Invoke[*zookeeper.ZookeeperRegistry](Injector)
	})
}
