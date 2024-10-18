package governance

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/ishumei/krpc/kserver/sconfig"
	"github.com/ishumei/krpc/logging"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	zkregistry "github.com/ishumei/krpc/registry-zookeeper/registry"
	"github.com/samber/do"
)

func init() {
	do.Provide(sconfig.Injector, func(i *do.Injector) (*zkregistry.ZookeeperRegistry, error) {
		f := do.MustInvoke[*sconfig.FrameConfig](sconfig.Injector)
		c := f.Registry
		logger := do.MustInvoke[*logging.Logger](logging.Injector)
		return zkregistry.NewZookeeperRegistryWithConf(
			c,
			f.Addr,
			registry_zookeeper.WithLogger(logger),
		)
	})
	do.Provide(sconfig.Injector, func(i *do.Injector) (registry.Registry, error) {
		return do.Invoke[*zkregistry.ZookeeperRegistry](sconfig.Injector)
	})
}
