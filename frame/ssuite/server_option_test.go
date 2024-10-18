package ssuite

import (
	"testing"

	"github.com/ishumei/krpc/frame/sconfig"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestServerOptions(t *testing.T) {
	assert.NotNil(t, do.MustInvoke[*ServerOptions](sconfig.Injector))
}
