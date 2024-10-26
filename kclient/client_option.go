package kclient

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/ishumei/krpc/monitor"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

const (
	DefaultClientName = "unknown"
)

type ClientOptions struct {
	copts []client.Option
}

func (c *ClientOptions) Options() []client.Option {
	return c.copts
}

func MustNewClientOptionsWithoutResolver(c *SingleClientConf, opts ...client.Option) *ClientOptions {
	options := []client.Option{
		client.WithTransportProtocol(transport.Framed),
		//lint:ignore SA1019 The server is only supported Framed ignore the deprecation warnings
		//nolint:staticcheck // SA1019 ignore the deprecation warnings
		client.WithSuite(tracing.NewFramedClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: func() string {
			if c.ClientName != "" {
				return c.ClientName
			}
			return DefaultClientName
		}()}),
		client.WithTracer(monitor.NewClientTracerWithoutExport()),
	}

	options = append(options, opts...)

	if c.TimeoutMs > 0 {
		options = append(
			options,
			client.WithRPCTimeout(time.Duration(c.TimeoutMs)*time.Millisecond),
		)
	}

	if c.ConnectTimeoutMs > 0 {
		options = append(
			options,
			client.WithConnectTimeout(time.Duration(c.ConnectTimeoutMs)*time.Millisecond),
		)
	}

	if c.Retries >= 0 {
		failurePolicy := retry.NewFailurePolicy()
		failurePolicy.WithMaxRetryTimes(c.Retries)
		retryBreakerErrorRate := 0.2
		if c.ErrorRate > 0 && c.ErrorRate <= 0.3 {
			retryBreakerErrorRate = c.ErrorRate
		}
		failurePolicy.WithRetryBreaker(retryBreakerErrorRate)
		options = append(
			options,
			client.WithFailureRetry(failurePolicy),
			client.WithSpecifiedResultRetry(
				retry.AllErrorRetry(),
			),
		)
	}

	if c.LongConnection.Enabled {
		options = append(
			options,
			client.WithLongConnection(c.LongConnection.IdleConfig),
		)
	} else {
		options = append(
			options,
			client.WithShortConnection(),
		)
	}

	if c.OpenTelemetry.Enabled {
		options = append(
			options,
			client.WithSuite(tracing.NewClientSuite()),
		)
	}

	return &ClientOptions{options}
}
