package governance

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/ishumei/krpc/kserver/sconfig"
	"github.com/ishumei/krpc/logging"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	"github.com/samber/do"
)

func init() {
	do.Provide(sconfig.Injector, func(i *do.Injector) (*registry_zookeeper.ZookeeperResolver, error) {
		c := do.MustInvoke[*sconfig.FrameConfig](sconfig.Injector).Registry
		logger := do.MustInvoke[*logging.Logger](logging.Injector)
		return registry_zookeeper.NewZookeeperResolverWithConf(
			c,
			registry_zookeeper.WithLogger(logger),
		)
	})
	do.Provide(sconfig.Injector, func(i *do.Injector) (discovery.Resolver, error) {
		return do.Invoke[*registry_zookeeper.ZookeeperResolver](sconfig.Injector)
	})
}
