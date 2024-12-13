package strategies

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/boilerplate/pkg/models"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestNewStrategyDAG(t *testing.T) {
	stages := &Stages{
		Stages: [][]string{
			{"text-preprocess"},
			{"text-models", "text-list"},
			{"evaluation"},
		},
	}
	for _, layers := range stages.Stages {
		for _, name := range layers {
			placeholder, err := models.GetPlaceholder(name)
			assert.NoError(t, err)
			do.OverrideNamedValue(models.Injector, name, placeholder)
		}
	}
	d, err := NewStrategyDAG(stages)
	assert.NoError(t, err)
	assert.NotNil(t, d)
	klog.Infof("dag: %s", d.Dag.String())
}
