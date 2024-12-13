package models

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/transport"
	"github.com/day253/krpc/objects"
	"github.com/samber/do"
	"github.com/samber/lo"
)

const (
	DefaultRegistrySeparater = ","
	DefaultTransportProtocol = transport.Framed
	DefaultConnectTimeout    = 500 * time.Millisecond
)

// Model实例容器
var Injector = do.New()

type ModelInput interface {
	Input() map[string]interface{}
	Json() string
}

type ModelResult interface {
	Result() map[string]interface{}
	Json() string
}

type MapModelResult map[string]interface{}

func (m *MapModelResult) Result() map[string]interface{} {
	return *m
}

func (m *MapModelResult) Json() string {
	return objects.StringIndent(m.Result())
}

func NewMapModelResult(data map[string]interface{}) *MapModelResult {
	return lo.ToPtr(MapModelResult(data))
}

type Model interface {
	Name() string
	Clone() (Model, error) // 浅拷贝就好
	Run(context.Context, ModelInput) (ModelResult, error)
}
