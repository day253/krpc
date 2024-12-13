package models

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/boilerplate/pkg/conf"
	"github.com/day253/krpc/objects"
)

type FeaturesHook interface {
	Run(context.Context, map[string]interface{}) map[string]interface{}
}

// features代表上游传过来的
type FeaturesHooksChain []FeaturesHook

func (r FeaturesHooksChain) Run(ctx context.Context, features map[string]interface{}) map[string]interface{} {
	for _, hook := range r {
		hook := hook
		if conf.Debug() {
			before := objects.Bytes(features)
			features = hook.Run(ctx, features)
			klog.CtxInfof(ctx, diff(before, objects.Bytes(features)))
		} else {
			features = hook.Run(ctx, features)
		}
	}
	return features
}
