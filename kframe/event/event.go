package event

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kframe/kservice"
	"github.com/ishumei/krpc/kframe/sconfig"
	"github.com/ishumei/krpc/kframe/ssuite"
	"github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re/eventpredictor"
	"github.com/samber/do"
)

type EventService struct {
	*kservice.Kservice
}

func NewEventService(i *do.Injector) (*EventService, error) {
	opts := do.MustInvoke[*ssuite.ServerOptions](sconfig.Injector)

	predictor := do.MustInvoke[re.EventPredictor](sconfig.Injector)

	return &EventService{
		Kservice: kservice.MustNewKservice(i, eventpredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(sconfig.Injector, NewEventService)
}
