package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/protocols/image/kitex_gen/shumei/strategy/re"
)

func image() {
	klog.Info(kclient.MustNewImageClient(&kclient.SingleClientConf{
		ResolverConf: kclient.ResolverConf{
			Hostports: []string{
				fmt.Sprintf("%v:%v", *host, *port),
			},
		},
		ClientConf: kclient.ClientConf{
			ServiceName: kclient.ClientTypeImage,
		},
	}).Predict(context.Background(), &re.ImagePredictRequest{
		RequestId:    requestId,
		Organization: organization,
	}))
}
