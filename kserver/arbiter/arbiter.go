package arbiter

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"github.com/samber/do"
)

type ArbiterService struct {
	*kserver.Kservice
}

func NewArbiterService(i *do.Injector) (*ArbiterService, error) {
	opts := do.MustInvoke[*kserver.ServerOptions](kserver.Injector)

	p := do.MustInvoke[service.Predictor](kserver.Injector)

	return &ArbiterService{
		Kservice: kserver.MustNewKservice(i, predictor.NewServer(p, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(kserver.Injector, NewArbiterService)
}
