package kclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/day253/krpc/protocols/text/kitex_gen/shumei/strategy/re/textpredictor"
	"github.com/samber/do"
)

type TextClient struct {
	*Kclient
	textpredictor.Client
}

func MustNewTextClient(c *SingleClientConf) *TextClient {
	k := MustNewKclient(c)
	return &TextClient{
		Kclient: k,
		Client: textpredictor.MustNewClient(
			c.ServiceName,
			client.WithSuite(k),
		),
	}
}

func TextClientIns() *TextClient {
	return do.MustInvokeNamed[*TextClient](Injector, ClientTypeText)
}

func MustNewTextClientWithInjector(m ClientConf) textpredictor.Client {
	k := MustNewClientOptionsWithoutResolver(
		&SingleClientConf{
			ClientConf: m,
		},
		client.WithResolver(do.MustInvoke[discovery.Resolver](Injector)),
	)
	return textpredictor.MustNewClient(
		m.ServiceName,
		client.WithSuite(k),
	)
}

func TextNgIns(name string) textpredictor.Client {
	return do.MustInvokeNamed[textpredictor.Client](Injector, name)
}

func TextNgDefaultIns() textpredictor.Client {
	return do.MustInvokeNamed[textpredictor.Client](Injector, ClientTypeText)
}
