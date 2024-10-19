package event

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re/eventpredictor"
	"github.com/samber/do"
)

type EventService struct {
	*kserver.Kservice
}

func NewEventService(i *do.Injector) (*EventService, error) {
	opts := do.MustInvoke[*kserver.ServerOptions](kserver.Injector)

	predictor := do.MustInvoke[re.EventPredictor](kserver.Injector)

	return &EventService{
		Kservice: kserver.MustNewKservice(i, eventpredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(kserver.Injector, NewEventService)
}
