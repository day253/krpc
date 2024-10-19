package main

import (
	"context"
	"os"

	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/kserver/arbiter"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	"github.com/samber/do"
)

type predictorImpl struct{}

func (s *predictorImpl) Predict(ctx context.Context, request *service.PredictRequest) (resp *service.PredictResult_, err error) {
	return nil, nil
}

func (s *predictorImpl) Health(ctx context.Context) (resp bool, err error) {
	return true, nil
}

func main() {
	os.Setenv("ENV_ROLE", "test")
	injector := kserver.Injector
	arbiterService := do.MustInvoke[*arbiter.ArbiterService](injector)
	defer func() { _ = arbiterService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	arbiterService.Start()
}

func init() {
	do.Provide(kserver.Injector, func(i *do.Injector) (service.Predictor, error) {
		return new(predictorImpl), nil
	})
}
