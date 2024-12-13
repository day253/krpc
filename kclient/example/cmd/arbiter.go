package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/kclient"
	"github.com/day253/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

func arbiter() {
	klog.Info(kclient.MustNewArbiterClient(&kclient.SingleClientConf{
		ResolverConf: kclient.ResolverConf{
			Hostports: []string{
				fmt.Sprintf("%v:%v", *host, *port),
			},
		},
		ClientConf: kclient.ClientConf{
			ServiceName: kclient.ClientTypeArbiter,
		},
	}).Predict(context.Background(), &service.PredictRequest{
		RequestId:    requestId,
		Organization: organization,
		Data:         data,
	}))
}
