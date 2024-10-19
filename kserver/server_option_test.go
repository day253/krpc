package kserver

import (
	"testing"

	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestServerOptions(t *testing.T) {
	assert.NotNil(t, do.MustInvoke[*ServerOptions](Injector))
}
