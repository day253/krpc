package main

import (
	"context"
	"os"

	"github.com/day253/krpc/kclient"
	"github.com/day253/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

func main() {
	os.Setenv("ENV_ROLE", "test")
	kclient.MustNewMultiClientConf("./", "clients", "yaml")
	kclient.ArbiterIns("arbiter").Predict(context.Background(), &service.PredictRequest{})
}
