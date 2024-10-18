package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re"
)

func event() {
	klog.Info(kclient.MustNewEventClient(&kclient.SingleClientConf{
		ResolverConf: kclient.ResolverConf{
			Hostports: []string{
				fmt.Sprintf("%v:%v", *host, *port),
			},
		},
		ClientConf: kclient.ClientConf{
			ServiceName: kclient.ClientTypeEvent,
		},
	}).Predict(context.Background(), &re.EventPredictRequest{
		RequestId:    requestId,
		Organization: organization,
	}))
}
