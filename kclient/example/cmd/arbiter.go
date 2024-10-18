package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
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
