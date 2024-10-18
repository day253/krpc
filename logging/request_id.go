package logging

import "context"

const (
	requestIdKey = "request_id"
)

type requestIdKeyType struct{}

var requestIdKeyVal requestIdKeyType

func GetRequestId(ctx context.Context) string {
	if ctx != nil {
		if val, ok := ctx.Value(requestIdKeyVal).(string); ok {
			return val
		}
	}
	return ""
}

func WithRequestId(ctx context.Context, requestId string) context.Context {
	if requestId == "" {
		return ctx
	}
	return context.WithValue(ctx, requestIdKeyVal, requestId)
}

func HasRequestId(ctx context.Context) bool {
	return GetRequestId(ctx) != ""
}
