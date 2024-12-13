package logging

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/day253/krpc/asyncwriter"
	"github.com/day253/lumberjack"
	"github.com/samber/do"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLogExt = ".log"
)

const (
	OutputFile    string = "File"
	OutputConsole string = "Console"
)

func newWriteSyncer(dir string, output string, cfg LoggerConfig) zapcore.WriteSyncer {
	switch output {
	case OutputConsole:
		return newConsoleWriteSyncer()
	case OutputFile:
		return newFileWriteSyncer(dir, cfg)
	default:
		return nil
	}
}

func newConsoleWriteSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func newFileWriteSyncer(dir string, cfg LoggerConfig) zapcore.WriteSyncer {
	if dir == "" || cfg.Name == "" {
		return nil
	}
	fileName := filepath.Join(dir, cfg.Name)
	if !strings.HasPrefix(fileName, defaultLogExt) {
		fileName += defaultLogExt
	}
	return zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			LocalTime:  true,
			Compress:   cfg.Compress,
		},
	)
}

func newEncoder(functionKey string) zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.FunctionKey = functionKey
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(config)
}

func newAtomicLevel(text string) zap.AtomicLevel {
	lvl, _ := zap.ParseAtomicLevel(text)
	return lvl
}

func newZapLogger(logConfig *LogConfig) *Logger {
	var cores []zapcore.Core
	var stopers []Stopable
	for _, cfg := range logConfig.Loggers {
		for _, output := range cfg.Outputs {
			writeSyncer := newWriteSyncer(logConfig.Dir, output, cfg)
			if writeSyncer == nil {
				continue
			}
			if output != OutputConsole && cfg.Buffered {
				bufferedWriteSyncer := asyncwriter.NewDefaultAsyncBufferWriteSyncer(writeSyncer)
				writeSyncer = bufferedWriteSyncer
				stopers = append(stopers, bufferedWriteSyncer)
			}
			cores = append(cores, zapcore.NewCore(newEncoder(logConfig.FunctionKey), writeSyncer, newAtomicLevel(cfg.Level)))
		}
	}
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(3))
	config := defaultConfig()
	return &Logger{
		l:       logger.Sugar(),
		config:  config,
		stopers: stopers,
	}
}

func init() {
	do.Provide(Injector, func(i *do.Injector) (*Logger, error) {
		c := do.MustInvoke[LogConfig](Injector)
		return newZapLogger(&c), nil
	})
}
