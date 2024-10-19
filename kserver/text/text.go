package text

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re/textpredictor"
	"github.com/samber/do"
)

type TextService struct {
	*kserver.Kservice
}

func NewTextService(i *do.Injector) (*TextService, error) {
	opts := do.MustInvoke[*kserver.ServerOptions](kserver.Injector)

	predictor := do.MustInvoke[re.TextPredictor](kserver.Injector)

	return &TextService{
		Kservice: kserver.MustNewKservice(i, textpredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(kserver.Injector, NewTextService)
}
