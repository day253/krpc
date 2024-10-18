// Code generated by Kitex v0.5.2. DO NOT EDIT.

package textpredictor

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	re "github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Predict(ctx context.Context, request *re.TextPredictRequest, callOptions ...callopt.Option) (r *re.TextPredictResult_, err error)
	Health(ctx context.Context, callOptions ...callopt.Option) (r bool, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTextPredictorClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTextPredictorClient struct {
	*kClient
}

func (p *kTextPredictorClient) Predict(ctx context.Context, request *re.TextPredictRequest, callOptions ...callopt.Option) (r *re.TextPredictResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Predict(ctx, request)
}

func (p *kTextPredictorClient) Health(ctx context.Context, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Health(ctx)
}
