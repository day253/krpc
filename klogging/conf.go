package klogging

import "github.com/samber/do"

var Injector = do.New()

type LogConfig struct {
	Dir         string `default:"log"`
	FunctionKey string
	Loggers     []LoggerConfig
}

type LoggerConfig struct {
	Name       string `default:"default"`
	Level      string `default:"info"`
	Buffered   bool
	MaxSize    int `default:"2048"`
	MaxAge     int `default:"7"`
	MaxBackups int `default:"20"`
	Compress   bool
	Outputs    []string
}
