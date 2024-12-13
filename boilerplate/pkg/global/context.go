package global

import (
	"context"

	"github.com/day253/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

type strategyKeyType struct{}

var strategyKey strategyKeyType

type StrategyContext interface {
	GetServicePredictRequest(map[string]interface{}) *service.PredictRequest
}

func GetStrategyContext(ctx context.Context) StrategyContext {
	if ctx != nil {
		if val, ok := ctx.Value(strategyKey).(StrategyContext); ok {
			return val
		}
	}
	return nil
}

func WithStrategyContext(ctx context.Context, n StrategyContext) context.Context {
	if n == nil {
		return ctx
	}
	return context.WithValue(ctx, strategyKey, n)
}

func HasStrategyContext(ctx context.Context) bool {
	return GetStrategyContext(ctx) != nil
}
