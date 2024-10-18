package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

const (
	defaultConfigSuffix = "yaml"
)

var (
	configPathList = []string{
		"./",
		"./conf",
		"../",
		"../conf",
	}
)

func LoadDefaultConf[T any](conf *T, confFile, mergeFile string) error {
	v := viper.New()
	for _, path := range configPathList {
		v.AddConfigPath(path)
	}
	v.SetConfigName(confFile)
	v.SetConfigType(defaultConfigSuffix)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.SetConfigName(mergeFile)
	if err := v.MergeInConfig(); err != nil {
		klog.Error(err)
	}
	return v.Unmarshal(conf, viper.DecodeHook(StringRenderTextTemplateHookFunc()))
}

func MustLoadConf[T any](conf *T, path, file, suffix string) {
	viper.SetConfigName(file)
	viper.SetConfigType(suffix)
	viper.AddConfigPath(path)
	lo.Must0(viper.ReadInConfig())
	lo.Must0(viper.Unmarshal(conf, viper.DecodeHook(StringRenderTextTemplateHookFunc())))
}
