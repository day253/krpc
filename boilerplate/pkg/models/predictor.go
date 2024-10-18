package models

import (
	"context"
	"errors"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/boilerplate/pkg/conf"
	"github.com/ishumei/krpc/boilerplate/pkg/global"
	"github.com/ishumei/krpc/deepcopy"
	"github.com/ishumei/krpc/objects"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"github.com/samber/do"
)

type Option interface {
	apply(*Predictor)
}

type PredictorBuilder func(name string, opts ...Option) Model

type Predictor struct {
	name   string
	client predictor.Client
	send   FeaturesHooksChain
	move   DetailsHooksChain
}

func (p *Predictor) Name() string {
	return p.name
}

func (p *Predictor) predict(ctx context.Context, request *service.PredictRequest) (*service.PredictResult_, error) {
	return p.client.Predict(ctx, request)
}

func (p *Predictor) Clone() (Model, error) {
	return &Predictor{
		name:   p.name,
		client: p.client,
		send:   p.send,
		move:   p.move,
	}, nil
}

func (p *Predictor) Run(ctx context.Context, input ModelInput) (ModelResult, error) {
	if !global.HasStrategyContext(ctx) {
		klog.CtxErrorf(ctx, "%s no strategy context", p.name)
		return nil, errors.New("no strategy context")
	}
	features := deepcopy.Clone(input.Input()) // 可能产生并发问题因此拷贝
	req := global.GetStrategyContext(ctx).GetServicePredictRequest(
		p.send.Run(ctx, features),
	)
	if conf.Debug() {
		klog.CtxInfof(ctx, "%s request: %s", p.name, objects.String(req))
		klog.CtxInfof(ctx, "%s requestData: %s", p.name, req.GetData())
	}
	resp, err := p.predict(ctx, req)
	if err != nil || resp == nil {
		klog.CtxErrorf(ctx, "%s: %v", p.name, err)
		return nil, err
	}
	detail := make(map[string]interface{})
	err = sonic.Unmarshal([]byte(resp.GetDetail()), &detail)
	if err != nil {
		klog.CtxErrorf(ctx, "%s: %v", p.name, err)
		return nil, err
	}
	if detail == nil {
		detail = make(map[string]interface{})
	}
	if conf.Debug() {
		klog.CtxInfof(ctx, "%s responseDetail: %s", p.name, objects.StringIndent(detail))
	} else {
		klog.CtxInfof(ctx, "%s responseDetail: %s", p.name, objects.String(detail))
	}
	return NewMapModelResult(
		p.move.Run(ctx, detail),
	), err
}

func MustNewPredictor(name string, opts ...Option) Model {
	predictor := &Predictor{
		name: name,
		send: FeaturesHooksChain{},
		move: DetailsHooksChain{},
	}
	for _, o := range opts {
		o.apply(predictor)
	}
	return predictor
}

func NewPredictor(name string, opts ...Option) (Model, error) {
	return MustNewPredictor(name, opts...), nil
}

func GetPredictor(name string) (Model, error) {
	model, err := do.InvokeNamed[Model](Injector, name)
	if err != nil {
		return nil, err
	}
	return model.Clone()
}
