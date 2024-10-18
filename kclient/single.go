package kclient

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/conf"
	"github.com/ishumei/krpc/objects"
	"github.com/samber/do"
)

type SingleClientConf struct {
	ClientName   string
	ResolverConf `mapstructure:",squash"`
	ClientConf   `mapstructure:",squash"`
}

func MustNewSingleClientConf(path, file, suffix string) {
	c := &SingleClientConf{}
	conf.MustLoadConf(c, path, file, suffix)
	klog.Info("load: ", objects.String(c))
	InjectClientFromSingleClientConf(c)
}

func InjectClientFromSingleClientConf(c *SingleClientConf) {
	switch c.Type {
	case ClientTypeAudio:
		do.OverrideNamedValue(Injector, ClientTypeAudio, MustNewAudioClient(c))
	case ClientTypeEvent:
		do.OverrideNamedValue(Injector, ClientTypeEvent, MustNewEventClient(c))
	case ClientTypeImage:
		do.OverrideNamedValue(Injector, ClientTypeImage, MustNewImageClient(c))
	case ClientTypeText:
		do.OverrideNamedValue(Injector, ClientTypeText, MustNewTextClient(c))
	case ClientTypeArbiter:
		fallthrough
	default:
		do.OverrideNamedValue(Injector, ClientTypeArbiter, MustNewArbiterClient(c))
	}
}
