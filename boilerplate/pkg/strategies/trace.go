package strategies

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	dagTraceTextKey = attribute.Key("vertex.elapsed")
)

type Entry struct {
	Context context.Context
	ID      string
	Elapsed int64
}

func Trace(entry *Entry) error {
	if entry.Context == nil {
		return nil
	}

	span := trace.SpanFromContext(entry.Context)
	if !span.IsRecording() {
		return nil
	}

	attrs := []attribute.KeyValue{
		dagTraceTextKey.Int64(entry.Elapsed),
	}

	span.AddEvent(entry.ID,
		trace.WithTimestamp(time.Now()),
		trace.WithAttributes(attrs...),
	)

	return nil
}
