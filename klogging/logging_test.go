package klogging

import (
	"testing"

	"github.com/creasty/defaults"
	"github.com/samber/do"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	do.OverrideValue(Injector, func() LogConfig {
		v := viper.New()
		v.AddConfigPath("./")
		v.SetConfigName("log")
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		logConfig := LogConfig{}
		if err := defaults.Set(&logConfig); err != nil {
			panic(err)
		}
		err := v.Unmarshal(&logConfig)
		if err != nil {
			panic(err)
		}
		return logConfig
	}())
	assert.NotNil(t, do.MustInvoke[*Logger](Injector))
}
