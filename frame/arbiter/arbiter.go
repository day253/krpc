package arbiter

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/frame/kservice"
	"github.com/ishumei/krpc/frame/sconfig"
	"github.com/ishumei/krpc/frame/ssuite"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	arbiterpredictor "github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"github.com/samber/do"
)

type ArbiterService struct {
	*kservice.Kservice
}

func NewArbiterService(i *do.Injector) (*ArbiterService, error) {
	opts := do.MustInvoke[*ssuite.ServerOptions](sconfig.Injector)

	predictor := do.MustInvoke[service.Predictor](sconfig.Injector)

	return &ArbiterService{
		Kservice: kservice.MustNewKservice(i, arbiterpredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(sconfig.Injector, NewArbiterService)
}
