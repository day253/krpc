package governance

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/ishumei/krpc/kserver/sconfig"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	assert.NotNil(t, do.MustInvoke[registry.Registry](sconfig.Injector))
}
