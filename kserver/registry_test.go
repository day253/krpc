package kserver

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	assert.NotNil(t, do.MustInvoke[registry.Registry](Injector))
}
