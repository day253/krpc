// Code generated by Kitex v0.11.3. DO NOT EDIT.

package predictor

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	service "github.com/day253/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Predict(ctx context.Context, request *service.PredictRequest, callOptions ...callopt.Option) (r *service.PredictResult_, err error)
	Health(ctx context.Context, callOptions ...callopt.Option) (r bool, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kPredictorClient{
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

type kPredictorClient struct {
	*kClient
}

func (p *kPredictorClient) Predict(ctx context.Context, request *service.PredictRequest, callOptions ...callopt.Option) (r *service.PredictResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Predict(ctx, request)
}

func (p *kPredictorClient) Health(ctx context.Context, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Health(ctx)
}
