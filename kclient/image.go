package kclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/day253/krpc/protocols/image/kitex_gen/shumei/strategy/re/imagepredictor"
	"github.com/samber/do"
)

type ImageClient struct {
	*Kclient
	imagepredictor.Client
}

func MustNewImageClient(c *SingleClientConf) *ImageClient {
	k := MustNewKclient(c)
	return &ImageClient{
		Kclient: k,
		Client: imagepredictor.MustNewClient(
			c.ServiceName,
			client.WithSuite(k),
		),
	}
}

func ImageClientIns() *ImageClient {
	return do.MustInvokeNamed[*ImageClient](Injector, ClientTypeImage)
}

func MustNewImageClientWithInjector(m ClientConf) imagepredictor.Client {
	k := MustNewClientOptionsWithoutResolver(
		&SingleClientConf{
			ClientConf: m,
		},
		client.WithResolver(do.MustInvoke[discovery.Resolver](Injector)),
	)
	return imagepredictor.MustNewClient(
		m.ServiceName,
		client.WithSuite(k),
	)
}

func ImageNgIns(name string) imagepredictor.Client {
	return do.MustInvokeNamed[imagepredictor.Client](Injector, name)
}

func ImageNgDefaultIns() imagepredictor.Client {
	return do.MustInvokeNamed[imagepredictor.Client](Injector, ClientTypeImage)
}
