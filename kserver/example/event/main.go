package main

import (
	"context"
	"os"

	"github.com/day253/krpc/kserver"
	"github.com/day253/krpc/kserver/event"
	"github.com/day253/krpc/protocols/event/kitex_gen/shumei/strategy/re"
	"github.com/samber/do"
)

type predictorImpl struct{}

func (s *predictorImpl) Predict(ctx context.Context, request *re.EventPredictRequest) (resp *re.EventPredictResult_, err error) {
	return nil, nil
}

func (s *predictorImpl) Health(ctx context.Context) (resp bool, err error) {
	return true, nil
}

func main() {
	os.Setenv("ENV_ROLE", "test")
	injector := kserver.Injector
	eventService := do.MustInvoke[*event.EventService](injector)
	defer func() { _ = eventService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	eventService.Start()
}

func init() {
	do.Provide(kserver.Injector, func(i *do.Injector) (re.EventPredictor, error) {
		return new(predictorImpl), nil
	})
}
