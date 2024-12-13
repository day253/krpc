package models

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/boilerplate/pkg/conf"
	"github.com/day253/krpc/objects"
)

type DetailsHook interface {
	Run(context.Context, map[string]interface{}) map[string]interface{}
}

// details代表返回结果
type DetailsHooksChain []DetailsHook

func (r DetailsHooksChain) Run(ctx context.Context, details map[string]interface{}) map[string]interface{} {
	for _, hook := range r {
		hook := hook
		if conf.Debug() {
			before := objects.Bytes(details)
			details = hook.Run(ctx, details)
			klog.CtxInfof(ctx, diff(before, objects.Bytes(details)))
		} else {
			details = hook.Run(ctx, details)
		}
	}
	return details
}
