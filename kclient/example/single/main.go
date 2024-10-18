package main

import (
	"context"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

func main() {
	os.Setenv("ENV_ROLE", "test")
	kclient.MustNewSingleClientConf("./", "client", "yaml")
	klog.Info(kclient.ArbiterClientIns().Predict(context.Background(), &service.PredictRequest{}))
}
