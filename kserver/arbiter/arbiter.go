package arbiter

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kserver/kservice"
	"github.com/ishumei/krpc/kserver/sconfig"
	"github.com/ishumei/krpc/kserver/ssuite"
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
