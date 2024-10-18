package sconfig

import (
	"fmt"
	"runtime"

	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/creasty/defaults"
	"github.com/ishumei/krpc/conf"
	"github.com/ishumei/krpc/logging"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	"github.com/samber/do"
)

var Injector = do.New()

type HttpConfig struct {
	Enabled bool
	Port    int
}

type OpenTelemetry struct {
	Enabled bool
	Address string `default:":4317"`
}

type SentinelSystemRule struct {
	system.Rule        `mapstructure:",squash"`
	TriggerCountFactor float64
}

func (s *SentinelSystemRule) ToSystemRule() *system.Rule {
	return &system.Rule{
		ID:         s.ID,
		MetricType: s.MetricType,
		TriggerCount: func() float64 {
			if s.MetricType == 0 && s.TriggerCount == 0 && s.TriggerCountFactor > 0 {
				return s.TriggerCountFactor * float64(runtime.NumCPU())
			}
			return s.TriggerCount
		}(),
		Strategy: s.Strategy,
	}
}

type Sentinel struct {
	Enabled     bool
	SystemRules []*SentinelSystemRule
}

func (s Sentinel) ToSystemRules() []*system.Rule {
	r := make([]*system.Rule, len(s.SystemRules))
	for _, rule := range s.SystemRules {
		r = append(r, rule.ToSystemRule())
	}
	return r
}

type Overload struct {
	Enabled    bool
	CpuPercent int
	MemPercent int
}

type FrameConfig struct {
	Addr          string
	Port          int    `default:"8888"`
	ServiceName   string `default:"server"`
	Registry      registry_zookeeper.Conf
	Http          HttpConfig
	OpenTelemetry OpenTelemetry
	Sentinel      Sentinel
	Overload      Overload
	Log           logging.LogConfig
}

func (c FrameConfig) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c FrameConfig) HttpAddress() string {
	httpPort := c.Http.Port
	if httpPort == 0 {
		httpPort = c.Port + 1
	}
	return fmt.Sprintf(":%d", httpPort)
}

func NewFrameConfig(i *do.Injector) (*FrameConfig, error) {
	frameConfig := &FrameConfig{}
	err := defaults.Set(frameConfig)
	if err != nil {
		return nil, err
	}
	err = conf.LoadDefaultConf(frameConfig, "frame", "overwrite")
	if err != nil {
		return nil, err
	}
	return frameConfig, nil
}

func NewFrameConfigWithLogger(i *do.Injector) (*FrameConfig, error) {
	frameConfig, err := NewFrameConfig(i)
	if err != nil {
		return nil, err
	}
	do.ProvideValue(logging.Injector, frameConfig.Log)
	logger, err := do.Invoke[*logging.Logger](logging.Injector)
	if err != nil {
		return nil, err
	}
	klog.SetLogger(logger)
	return frameConfig, nil
}

func init() {
	do.Provide(Injector, NewFrameConfigWithLogger)
}
