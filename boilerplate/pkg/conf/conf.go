package conf

import (
	"github.com/creasty/defaults"
	"github.com/ishumei/krpc/conf"
	"github.com/ishumei/krpc/zookeeper"
	"github.com/samber/do"
)

var Injector = do.New()

type FeaturesKeys struct {
	To   []string `json:"to,omitempty"`
	From []string `json:"from,omitempty"`
}

type Model struct {
	BasePath string
}

type MoveConf struct {
	MoveFeatures     bool
	MoveFeaturesKeys []FeaturesKeys
}

type Strategy struct {
	BasePath string
}

type Config struct {
	Debug    bool
	Registry zookeeper.Conf
	Model    Model
	Strategy Strategy
}

func defaultConfig() (*Config, error) {
	c := &Config{}
	if err := defaults.Set(c); err != nil {
		return nil, err
	}
	return c, nil
}

func mustGetConfig() *Config {
	if c, err := do.Invoke[*Config](Injector); err == nil {
		return c
	}
	if c, err := defaultConfig(); err == nil {
		return c
	}
	return &Config{}
}

func Debug() bool {
	return mustGetConfig().Debug
}

func NewConfig(i *do.Injector) (*Config, error) {
	reConfig, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	return reConfig, conf.LoadDefaultConf(reConfig, "frame", "overwrite")
}

func init() {
	do.Provide(Injector, NewConfig)
}
