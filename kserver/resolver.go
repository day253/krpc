package kserver

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/ishumei/krpc/logging"
	"github.com/ishumei/krpc/zookeeper"
	"github.com/samber/do"
)

func init() {
	do.Provide(Injector, func(i *do.Injector) (*zookeeper.ZookeeperResolver, error) {
		c := do.MustInvoke[*FrameConfig](Injector).Registry
		logger := do.MustInvoke[*logging.Logger](logging.Injector)
		return zookeeper.NewZookeeperResolverWithConf(
			c,
			zookeeper.WithLogger(logger),
		)
	})
	do.Provide(Injector, func(i *do.Injector) (discovery.Resolver, error) {
		return do.Invoke[*zookeeper.ZookeeperResolver](Injector)
	})
}
