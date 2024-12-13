package kclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/day253/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"github.com/samber/do"
)

type ArbiterClient struct {
	*Kclient
	predictor.Client
}

func MustNewArbiterClient(c *SingleClientConf) *ArbiterClient {
	k := MustNewKclient(c)
	return &ArbiterClient{
		Kclient: k,
		Client: predictor.MustNewClient(
			c.ServiceName,
			client.WithSuite(k),
		),
	}
}

func ArbiterClientIns() *ArbiterClient {
	return do.MustInvokeNamed[*ArbiterClient](Injector, ClientTypeArbiter)
}

func MustNewArbiterClientWithInjector(m ClientConf) predictor.Client {
	k := MustNewClientOptionsWithoutResolver(
		&SingleClientConf{
			ClientConf: m,
		},
		client.WithResolver(do.MustInvoke[discovery.Resolver](Injector)),
	)
	return predictor.MustNewClient(
		m.ServiceName,
		client.WithSuite(k),
	)
}

func ArbiterIns(name string) predictor.Client {
	return do.MustInvokeNamed[predictor.Client](Injector, name)
}

func ArbiterDefaultIns() predictor.Client {
	return do.MustInvokeNamed[predictor.Client](Injector, ClientTypeArbiter)
}
