package models

import (
	"context"

	"github.com/ishumei/krpc/boilerplate/pkg/conf"
	"github.com/ishumei/krpc/objects"
	"github.com/jinzhu/copier"
)

type Send struct {
	SendData         bool // 发送所有features.data字段
	SendAllFeatures  bool // 发送所有features字段
	SendFeatures     bool // 只移动SendFeaturesKeys指定的字段
	SendFeaturesKeys []conf.FeaturesKeys
}

func (s *Send) Run(ctx context.Context, features map[string]interface{}) map[string]interface{} {
	if features == nil {
		features = make(map[string]interface{})
	}
	// 所以全都没配置的默认行为是SendData
	if !s.SendFeatures && !s.SendAllFeatures {
		switch v := features["data"].(type) {
		case map[string]interface{}:
			return v
		default:
			return make(map[string]interface{})
		}
	}
	// ae的场景应该是即SendAllFeatures又SendData这里支持不了需要同时设置SendFeatures
	if !s.SendFeatures && s.SendAllFeatures {
		return features
	}
	result := make(map[string]interface{})
	if s.SendAllFeatures {
		result = features
	}
	// SendData和SendFeatures可以共存
	if s.SendData {
		if _, ok := features["data"].(map[string]interface{}); ok {
			_ = copier.Copy(&result, features["data"])
		}
	}
	from := objects.Map(features)
	to := objects.Map(result)
	for _, featuresKeys := range s.SendFeaturesKeys {
		fromVal := from.GetBySegs(featuresKeys.From)
		if fromVal == nil || fromVal == "" {
			continue
		}
		if len(featuresKeys.To) == 0 {
			continue
		}
		to.SetBySegs(featuresKeys.To, fromVal)
	}
	return to.MSI()
}
