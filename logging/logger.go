package logging

import (
	"context"
	"errors"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ klog.FullLogger = (*Logger)(nil)

// Ref to https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/README.md#json-formats
const (
	traceIDKey    = "trace_id"
	spanIDKey     = "span_id"
	traceFlagsKey = "trace_flags"
	logEventKey   = "log"
)

var (
	logSeverityTextKey = attribute.Key("otel.log.severity.text")
	logMessageKey      = attribute.Key("otel.log.message")
)

type Stopable interface {
	Stop() error
}

type Logger struct {
	l       *zap.SugaredLogger
	config  *config
	stopers []Stopable
}

func NewLogger(opts ...Option) *Logger {
	config := defaultConfig()

	// apply options
	for _, opt := range opts {
		opt.apply(config)
	}

	logger := zap.New(
		zapcore.NewCore(config.coreConfig.enc, config.coreConfig.ws, config.coreConfig.lvl),
		config.zapOpts...)

	return &Logger{
		l:      logger.Sugar(),
		config: config,
	}
}

func (l *Logger) Log(level klog.Level, kvs ...interface{}) {
	switch level {
	case klog.LevelTrace, klog.LevelDebug:
		l.l.Debug(kvs...)
	case klog.LevelInfo:
		l.l.Info(kvs...)
	case klog.LevelNotice, klog.LevelWarn:
		l.l.Warn(kvs...)
	case klog.LevelError:
		l.l.Error(kvs...)
	case klog.LevelFatal:
		l.l.Fatal(kvs...)
	default:
		l.l.Warn(kvs...)
	}
}

func (l *Logger) Logf(level klog.Level, format string, kvs ...interface{}) {
	logger := l.l.With()
	switch level {
	case klog.LevelTrace, klog.LevelDebug:
		logger.Debugf(format, kvs...)
	case klog.LevelInfo:
		logger.Infof(format, kvs...)
	case klog.LevelNotice, klog.LevelWarn:
		logger.Warnf(format, kvs...)
	case klog.LevelError:
		logger.Errorf(format, kvs...)
	case klog.LevelFatal:
		logger.Fatalf(format, kvs...)
	default:
		logger.Warnf(format, kvs...)
	}
}

func (l *Logger) CtxLogf(level klog.Level, ctx context.Context, format string, kvs ...interface{}) {
	var zlevel zapcore.Level
	span := trace.SpanFromContext(ctx)

	var args []interface{}
	args = append(args, requestIdKey, GetRequestId(ctx), organizationKey, GetOrganization(ctx))
	if span.IsRecording() {
		args = append(args, traceIDKey, span.SpanContext().TraceID(), spanIDKey, span.SpanContext().SpanID(), traceFlagsKey, span.SpanContext().TraceFlags())
	}
	sl := l.l.With(args...)
	switch level {
	case klog.LevelDebug, klog.LevelTrace:
		zlevel = zap.DebugLevel
		sl.Debugf(format, kvs...)
	case klog.LevelInfo:
		zlevel = zap.InfoLevel
		sl.Infof(format, kvs...)
	case klog.LevelNotice, klog.LevelWarn:
		zlevel = zap.WarnLevel
		sl.Warnf(format, kvs...)
	case klog.LevelError:
		zlevel = zap.ErrorLevel
		sl.Errorf(format, kvs...)
	case klog.LevelFatal:
		zlevel = zap.FatalLevel
		sl.Fatalf(format, kvs...)
	default:
		zlevel = zap.WarnLevel
		sl.Warnf(format, kvs...)
	}

	if !span.IsRecording() {
		// l.Logf(level, format, kvs...)
		return
	}

	msg := getMessage(format, kvs)
	attrs := []attribute.KeyValue{
		logMessageKey.String(msg),
		logSeverityTextKey.String(OtelSeverityText(zlevel)),
	}
	span.AddEvent(logEventKey, trace.WithAttributes(attrs...))

	// set span status
	if zlevel <= l.config.traceConfig.errorSpanLevel {
		span.SetStatus(codes.Error, msg)
		span.RecordError(errors.New(msg), trace.WithStackTrace(l.config.traceConfig.recordStackTraceInSpan))
	}
}

func (l *Logger) Trace(v ...interface{}) {
	l.Log(klog.LevelTrace, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Log(klog.LevelDebug, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Log(klog.LevelInfo, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.Log(klog.LevelNotice, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Log(klog.LevelWarn, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Log(klog.LevelError, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Log(klog.LevelFatal, v...)
}

// Printf implements zk.Logger.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Logf(klog.LevelInfo, format, v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Logf(klog.LevelTrace, format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(klog.LevelDebug, format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(klog.LevelInfo, format, v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.Logf(klog.LevelInfo, format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logf(klog.LevelWarn, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(klog.LevelError, format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(klog.LevelFatal, format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelInfo, ctx, format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelWarn, ctx, format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelWarn, ctx, format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelFatal, ctx, format, v...)
}

func (l *Logger) SetLevel(level klog.Level) {
	var lvl zapcore.Level
	switch level {
	case klog.LevelTrace, klog.LevelDebug:
		lvl = zap.DebugLevel
	case klog.LevelInfo:
		lvl = zap.InfoLevel
	case klog.LevelWarn, klog.LevelNotice:
		lvl = zap.WarnLevel
	case klog.LevelError:
		lvl = zap.ErrorLevel
	case klog.LevelFatal:
		lvl = zap.FatalLevel
	default:
		lvl = zap.WarnLevel
	}
	l.config.coreConfig.lvl.SetLevel(lvl)
}

func (l *Logger) SetOutput(writer io.Writer) {
	ws := zapcore.AddSync(writer)
	log := zap.New(
		zapcore.NewCore(l.config.coreConfig.enc, ws, l.config.coreConfig.lvl),
		l.config.zapOpts...,
	).Sugar()
	l.config.coreConfig.ws = ws
	l.l = log
}

func (l *Logger) Sync() {
	_ = l.l.Sync()
}

func (l *Logger) Stop() {
	for _, stoper := range l.stopers {
		_ = stoper.Stop()
	}
}
