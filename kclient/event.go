package kclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re/eventpredictor"
	"github.com/samber/do"
)

type EventClient struct {
	*Kclient
	eventpredictor.Client
}

func MustNewEventClient(c *SingleClientConf) *EventClient {
	k := MustNewKclient(c)
	return &EventClient{
		Kclient: k,
		Client: eventpredictor.MustNewClient(
			c.ServiceName,
			client.WithSuite(k),
		),
	}
}

func EventClientIns() *EventClient {
	return do.MustInvokeNamed[*EventClient](Injector, ClientTypeEvent)
}

func MustNewEventClientWithInjector(m ClientConf) eventpredictor.Client {
	k := MustNewClientOptionsWithoutResolver(
		&SingleClientConf{
			ClientConf: m,
		},
		client.WithResolver(do.MustInvoke[discovery.Resolver](Injector)),
	)
	return eventpredictor.MustNewClient(
		m.ServiceName,
		client.WithSuite(k),
	)
}

func EventNgIns(name string) eventpredictor.Client {
	return do.MustInvokeNamed[eventpredictor.Client](Injector, name)
}

func EventNgDefaultIns() eventpredictor.Client {
	return do.MustInvokeNamed[eventpredictor.Client](Injector, ClientTypeEvent)
}
