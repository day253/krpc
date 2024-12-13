package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/kclient"
	"github.com/day253/krpc/protocols/text/kitex_gen/shumei/strategy/re"
)

func text() {
	klog.Info(kclient.MustNewTextClient(&kclient.SingleClientConf{
		ResolverConf: kclient.ResolverConf{
			Hostports: []string{
				fmt.Sprintf("%v:%v", *host, *port),
			},
		},
		ClientConf: kclient.ClientConf{
			ServiceName: kclient.ClientTypeText,
		},
	}).Predict(context.Background(), &re.TextPredictRequest{
		RequestId:    requestId,
		Organization: organization,
	}))
}
