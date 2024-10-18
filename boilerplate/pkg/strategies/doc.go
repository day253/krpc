package strategies

import (
	"path/filepath"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/boilerplate/pkg/conf"
	"github.com/ishumei/krpc/boilerplate/pkg/models"
	"github.com/ishumei/krpc/kserver/sconfig"
	"github.com/ishumei/krpc/objects"
	"github.com/ishumei/krpc/registry-zookeeper/resolver"
	"github.com/samber/do"
)

func Init() error {
	c, err := do.Invoke[*conf.Config](conf.Injector)
	if err != nil {
		return err
	}
	zookeeperConn, err := do.Invoke[*resolver.ZookeeperResolver](sconfig.Injector)
	if err != nil {
		return err
	}
	modelNames, _, err := zookeeperConn.Children(c.Model.BasePath)
	if err != nil {
		return err
	}
	for _, modelName := range modelNames {
		modelName := modelName
		if modelName == models.ApolloName {
			continue
		}
		path := filepath.Join(c.Model.BasePath, modelName)
		content, _, err := zookeeperConn.Get(path)
		if err != nil {
			return err
		}
		opts, err := models.NewPredictorOption(path, content)
		if err != nil {
			klog.Error("load model: ", modelName, " from zk, err: ", err)
			continue
		}
		klog.Info("load model: ", modelName, " from zk, predictorOpt: ", objects.String(opts))
		do.ProvideNamed(models.Injector, modelName, func(i *do.Injector) (models.Model, error) {
			return models.MustNewPredictor(modelName, opts), nil
		})
	}
	strategyNames, _, err := zookeeperConn.Children(c.Strategy.BasePath)
	if err != nil {
		return err
	}
	for _, strategyName := range strategyNames {
		path := filepath.Join(c.Strategy.BasePath, strategyName)
		content, _, err := zookeeperConn.Get(path)
		if err != nil {
			return err
		}
		klog.Info("load strategy: ", strategyName, " from zk, content: ", string(content))
		stages, err := NewStages(content)
		if err != nil {
			klog.Error("load strategy: ", strategyName, " from zk, err: ", err)
			return err
		}
		dag, err := NewStrategyDAG(stages)
		if err != nil {
			klog.Error("load strategy: ", strategyName, " from zk, err: ", err)
			continue
		}
		klog.Info("load strategy: ", strategyName, " from zk, dag: ", dag.String())
		do.ProvideNamed(Injector, strategyName, func(i *do.Injector) (*Stages, error) {
			return stages, nil
		})
	}
	return nil
}
