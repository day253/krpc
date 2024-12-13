package kclient

import (
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/day253/krpc/zookeeper"
	"github.com/samber/do"
)

var Injector = do.New()

type OpenTelemetry struct {
	Enabled bool
	Address string
}

type LongConnection struct {
	Enabled             bool
	connpool.IdleConfig `mapstructure:",squash"`
}

type ClientConf struct {
	Type             string
	ServiceName      string
	ConnectTimeoutMs int
	Retries          int
	TimeoutMs        int
	ErrorRate        float64
	LongConnection   LongConnection
	OpenTelemetry    OpenTelemetry
}

type ResolverConf struct {
	Hostports []string
	Resolver  zookeeper.Conf
}
