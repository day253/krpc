package image

import (
	"github.com/cloudwego/kitex/server"
	"github.com/day253/krpc/kserver"
	"github.com/day253/krpc/protocols/image/kitex_gen/shumei/strategy/re"
	"github.com/day253/krpc/protocols/image/kitex_gen/shumei/strategy/re/imagepredictor"
	"github.com/samber/do"
)

type ImageService struct {
	*kserver.Kservice
}

func NewImageService(i *do.Injector) (*ImageService, error) {
	opts := do.MustInvoke[*kserver.ServerOptions](kserver.Injector)

	predictor := do.MustInvoke[re.ImagePredictor](kserver.Injector)

	return &ImageService{
		Kservice: kserver.MustNewKservice(i, imagepredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(kserver.Injector, NewImageService)
}
