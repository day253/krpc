package models

import (
	"time"

	json "github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/ishumei/krpc/boilerplate/pkg/conf"
	"github.com/ishumei/krpc/kserver"
	prometheus "github.com/ishumei/krpc/monitor-prometheus"
	"github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/samber/do"
)

type PredictorOption struct {
	NodePath         string              `json:"-"`
	KeepLive         bool                `json:"keep_alive,omitempty"`
	TimeoutMs        int                 `json:"timeout_ms,omitempty"`
	MaxTimeoutMs     int                 `json:"max_timeout_ms,omitempty"`
	MinTimeoutMs     int                 `json:"min_timeout_ms,omitempty"`
	Retries          int                 `json:"retries,omitempty"`
	VariablePrefix   string              `json:"key,omitempty"`
	Derived          bool                `json:"derived,omitempty"`
	DerivedFrom      string              `json:"derived_from,omitempty"`
	SendData         bool                `json:"send_data,omitempty"`
	SendFeatures     bool                `json:"send_features,omitempty"`
	SendAllFeatures  bool                `json:"send_all_features,omitempty"`
	SendFeaturesKeys []conf.FeaturesKeys `json:"send_features_keys,omitempty"`
	MoveFeaturesKeys []conf.FeaturesKeys `json:"move_features_keys,omitempty"`
}

func (m *PredictorOption) apply(p *Predictor) {
	p.client = func() predictor.Client {
		failurePolicy := retry.NewFailurePolicy()
		failurePolicy.WithMaxRetryTimes(m.Retries)
		failurePolicy.WithRetryBreaker(0.1)
		c, _ := predictor.NewClient(
			m.NodePath,
			client.WithShortConnection(),
			client.WithResolver(do.MustInvoke[discovery.Resolver](kserver.Injector)),
			client.WithTransportProtocol(DefaultTransportProtocol),
			client.WithConnectTimeout(DefaultConnectTimeout),
			client.WithRPCTimeout(time.Duration(m.TimeoutMs)*time.Millisecond),
			client.WithFailureRetry(failurePolicy),
			client.WithSpecifiedResultRetry(retry.AllErrorRetry()),
			//lint:ignore SA1019 The server is only supported Framed ignore the deprecation warnings
			//nolint:staticcheck // SA1019 ignore the deprecation warnings
			client.WithSuite(tracing.NewFramedClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: m.NodePath}),
			client.WithTracer(prometheus.NewClientTracerWithoutExport()),
		)
		return c
	}()
	p.send = FeaturesHooksChain{
		&Send{
			SendData:         m.SendData,
			SendAllFeatures:  m.SendAllFeatures,
			SendFeatures:     m.SendFeatures,
			SendFeaturesKeys: m.SendFeaturesKeys,
		},
	}
	p.move = DetailsHooksChain{
		&ObjectPrefix{
			VariablePrefix: func() string {
				if m.VariablePrefix == "" {
					return p.name
				}
				return m.VariablePrefix
			}(),
		},
		&Move{
			MoveFeatures:     len(m.MoveFeaturesKeys) != 0,
			MoveFeaturesKeys: m.MoveFeaturesKeys,
		},
	}
}

func NewPredictorOption(path string, content []byte) (*PredictorOption, error) {
	opt := &PredictorOption{}
	if err := json.Unmarshal(content, opt); err != nil {
		return nil, err
	}
	opt.NodePath = path
	return opt, nil
}
