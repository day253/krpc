package governance

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/ishumei/krpc/frame/sconfig"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestResolver(t *testing.T) {
	assert.NotNil(t, do.MustInvoke[discovery.Resolver](sconfig.Injector))
}
