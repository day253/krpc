package audio

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/kserver/kservice"
	"github.com/ishumei/krpc/kserver/sconfig"
	"github.com/ishumei/krpc/kserver/ssuite"
	"github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re/audiopredictor"
	"github.com/samber/do"
)

type AudioService struct {
	*kservice.Kservice
}

func NewAudioService(i *do.Injector) (*AudioService, error) {
	opts := do.MustInvoke[*ssuite.ServerOptions](sconfig.Injector)

	predictor := do.MustInvoke[re.AudioPredictor](sconfig.Injector)

	return &AudioService{
		Kservice: kservice.MustNewKservice(i, audiopredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(sconfig.Injector, NewAudioService)
}
