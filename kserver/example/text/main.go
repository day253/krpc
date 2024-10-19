package main

import (
	"context"
	"os"

	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/kserver/text"
	re "github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re"
	"github.com/samber/do"
)

type predictorImpl struct{}

func (s *predictorImpl) Predict(ctx context.Context, request *re.TextPredictRequest) (resp *re.TextPredictResult_, err error) {
	return nil, nil
}

func (s *predictorImpl) Health(ctx context.Context) (resp bool, err error) {
	return true, nil
}

func main() {
	os.Setenv("ENV_ROLE", "test")
	injector := kserver.Injector
	textService := do.MustInvoke[*text.TextService](injector)
	defer func() { _ = textService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	textService.Start()
}

func init() {
	do.Provide(kserver.Injector, func(i *do.Injector) (re.TextPredictor, error) {
		return new(predictorImpl), nil
	})
}
