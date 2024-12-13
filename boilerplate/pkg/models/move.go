package models

import (
	"context"

	"github.com/day253/krpc/boilerplate/pkg/conf"
	"github.com/day253/krpc/objects"
)

type Move struct {
	MoveFeatures     bool
	MoveFeaturesKeys []conf.FeaturesKeys // 拷贝逻辑，从一个地方复制到另外一个地方
}

func (s *Move) Run(ctx context.Context, details map[string]interface{}) map[string]interface{} {
	if details == nil {
		details = make(map[string]interface{})
	}
	if !s.MoveFeatures {
		return details
	}
	m := objects.Map(details)
	for _, featuresKeys := range s.MoveFeaturesKeys {
		fromVal := m.GetBySegs(featuresKeys.From)
		if fromVal == nil || fromVal == "" {
			continue
		}
		if len(featuresKeys.To) == 0 {
			continue
		}
		m.SetBySegs(featuresKeys.To, fromVal)
	}
	return m.MSI()
}
