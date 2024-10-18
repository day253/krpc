package main

import (
	"context"
	"os"

	"github.com/ishumei/krpc/kclient"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

func main() {
	os.Setenv("ENV_ROLE", "test")
	kclient.MustNewMultiClientConf("./", "clients", "yaml")
	kclient.ArbiterIns("arbiter").Predict(context.Background(), &service.PredictRequest{})
}
