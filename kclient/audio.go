package kclient

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re/audiopredictor"
	"github.com/samber/do"
)

type AudioClient struct {
	*Kclient
	audiopredictor.Client
}

func MustNewAudioClient(c *SingleClientConf) *AudioClient {
	k := MustNewKclient(c)
	return &AudioClient{
		Kclient: k,
		Client: audiopredictor.MustNewClient(
			c.ServiceName,
			client.WithSuite(k),
		),
	}
}

func AudioClientIns() *AudioClient {
	return do.MustInvokeNamed[*AudioClient](Injector, ClientTypeAudio)
}

func MustNewAudioClientWithInjector(m ClientConf) audiopredictor.Client {
	k := MustNewClientOptionsWithoutResolver(
		&SingleClientConf{
			ClientConf: m,
		},
		client.WithResolver(do.MustInvoke[discovery.Resolver](Injector)),
	)
	return audiopredictor.MustNewClient(
		m.ServiceName,
		client.WithSuite(k),
	)
}

func AudioNgIns(name string) audiopredictor.Client {
	return do.MustInvokeNamed[audiopredictor.Client](Injector, name)
}

func AudioNgDefaultIns() audiopredictor.Client {
	return do.MustInvokeNamed[audiopredictor.Client](Injector, ClientTypeAudio)
}
