package klogging

import "context"

const (
	organizationKey = "organization"
)

type organizationKeyType struct{}

var organizationKeyVal organizationKeyType

func GetOrganization(ctx context.Context) string {
	if ctx != nil {
		if val, ok := ctx.Value(organizationKeyVal).(string); ok {
			return val
		}
	}
	return ""
}

func WithOrganization(ctx context.Context, organization string) context.Context {
	if organization == "" {
		return ctx
	}
	return context.WithValue(ctx, organizationKeyVal, organization)
}

func HasOrganization(ctx context.Context) bool {
	return GetOrganization(ctx) != ""
}
