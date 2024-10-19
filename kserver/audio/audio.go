package audio

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re/audiopredictor"
	"github.com/samber/do"
)

type AudioService struct {
	*kserver.Kservice
}

func NewAudioService(i *do.Injector) (*AudioService, error) {
	opts := do.MustInvoke[*kserver.ServerOptions](kserver.Injector)

	predictor := do.MustInvoke[re.AudioPredictor](kserver.Injector)

	return &AudioService{
		Kservice: kserver.MustNewKservice(i, audiopredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(kserver.Injector, NewAudioService)
}
