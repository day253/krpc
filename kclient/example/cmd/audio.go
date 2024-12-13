package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/kclient"
	"github.com/day253/krpc/protocols/audio/kitex_gen/shumei/strategy/re"
)

func audio() {
	klog.Info(kclient.MustNewAudioClient(&kclient.SingleClientConf{
		ResolverConf: kclient.ResolverConf{
			Hostports: []string{
				fmt.Sprintf("%v:%v", *host, *port),
			},
		},
		ClientConf: kclient.ClientConf{
			ServiceName: kclient.ClientTypeAudio,
		},
	}).Predict(context.Background(), &re.AudioPredictRequest{
		RequestId:    requestId,
		Organization: organization,
	}))
}
