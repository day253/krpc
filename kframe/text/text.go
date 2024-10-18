package text

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kframe/kservice"
	"github.com/ishumei/krpc/kframe/sconfig"
	"github.com/ishumei/krpc/kframe/ssuite"
	"github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re/textpredictor"
	"github.com/samber/do"
)

type TextService struct {
	*kservice.Kservice
}

func NewTextService(i *do.Injector) (*TextService, error) {
	opts := do.MustInvoke[*ssuite.ServerOptions](sconfig.Injector)

	predictor := do.MustInvoke[re.TextPredictor](sconfig.Injector)

	return &TextService{
		Kservice: kservice.MustNewKservice(i, textpredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(sconfig.Injector, NewTextService)
}
