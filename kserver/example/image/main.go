package main

import (
	"context"
	"os"

	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/kserver/image"
	re "github.com/ishumei/krpc/protocols/image/kitex_gen/shumei/strategy/re"
	"github.com/samber/do"
)

type predictorImpl struct{}

func (s *predictorImpl) Predict(ctx context.Context, request *re.ImagePredictRequest) (resp *re.ImagePredictResult_, err error) {
	return nil, nil
}

func (s *predictorImpl) Health(ctx context.Context) (resp bool, err error) {
	return true, nil
}

func main() {
	os.Setenv("ENV_ROLE", "test")
	injector := kserver.Injector
	imageService := do.MustInvoke[*image.ImageService](injector)
	defer func() { _ = imageService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	imageService.Start()
}

func init() {
	do.Provide(kserver.Injector, func(i *do.Injector) (re.ImagePredictor, error) {
		return new(predictorImpl), nil
	})
}
