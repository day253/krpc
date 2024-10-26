package main

import (
	"context"
	"os"

	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/kserver/audio"
	"github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re"
	"github.com/samber/do"
)

type predictorImpl struct{}

func (s *predictorImpl) Predict(ctx context.Context, request *re.AudioPredictRequest) (resp *re.AudioPredictResult_, err error) {
	return nil, nil
}

func (s *predictorImpl) Health(ctx context.Context) (resp bool, err error) {
	return true, nil
}

func main() {
	os.Setenv("ENV_ROLE", "test")
	injector := kserver.Injector
	audioService := do.MustInvoke[*audio.AudioService](injector)
	defer func() { _ = audioService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	audioService.Start()
}

func init() {
	do.Provide(kserver.Injector, func(i *do.Injector) (re.AudioPredictor, error) {
		return new(predictorImpl), nil
	})
}
