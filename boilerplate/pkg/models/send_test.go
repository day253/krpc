package models

import (
	"context"
	"reflect"
	"testing"

	"github.com/day253/krpc/boilerplate/pkg/conf"
)

func TestSend_Run(t *testing.T) {
	type fields struct {
		SendData         bool
		SendAllFeatures  bool
		SendFeatures     bool
		SendFeaturesKeys []conf.FeaturesKeys
	}
	type args struct {
		features map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		{
			name: "Test_SendAllFeatures",
			fields: fields{
				SendData:        false,
				SendAllFeatures: true,
				SendFeatures:    false,
				SendFeaturesKeys: []conf.FeaturesKeys{
					{
						From: []string{
							"audio-asr",
							"audioText",
							"asr_text",
						},
						To: []string{
							"data",
							"text",
						},
					},
				},
			},
			args: args{
				features: map[string]interface{}{
					"data": map[string]interface{}{
						"__product__": "value1",
					},
					"audio-asr": map[string]interface{}{
						"audioText": map[string]interface{}{
							"asr_text": "value2",
						},
					},
				},
			},
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"__product__": "value1",
				},
				"audio-asr": map[string]interface{}{
					"audioText": map[string]interface{}{
						"asr_text": "value2",
					},
				},
			},
		},
		{
			name: "Test_SendData",
			fields: fields{
				SendData:        true,
				SendAllFeatures: false,
				SendFeatures:    false,
				SendFeaturesKeys: []conf.FeaturesKeys{
					{
						From: []string{
							"audio-asr",
							"audioText",
							"asr_text",
						},
						To: []string{
							"data",
							"text",
						},
					},
				},
			},
			args: args{
				features: map[string]interface{}{
					"data": map[string]interface{}{
						"__product__": "value1",
					},
					"audio-asr": map[string]interface{}{
						"audioText": map[string]interface{}{
							"asr_text": "value2",
						},
					},
				},
			},
			want: map[string]interface{}{
				"__product__": "value1",
			},
		},
		{
			name: "Test_SendFeatures",
			fields: fields{
				SendData:        false,
				SendAllFeatures: false,
				SendFeatures:    true,
				SendFeaturesKeys: []conf.FeaturesKeys{
					{
						From: []string{
							"audio-asr",
							"audioText",
							"asr_text",
						},
						To: []string{
							"data",
							"text",
						},
					},
				},
			},
			args: args{
				features: map[string]interface{}{
					"data": map[string]interface{}{
						"__product__": "value1",
					},
					"audio-asr": map[string]interface{}{
						"audioText": map[string]interface{}{
							"asr_text": "value2",
						},
					},
				},
			},
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"text": "value2",
				},
			},
		},
		{
			name: "Test_SendFeatures",
			fields: fields{
				SendData:        false,
				SendAllFeatures: false,
				SendFeatures:    true,
				SendFeaturesKeys: []conf.FeaturesKeys{
					{
						From: []string{
							"audio-asr",
							"audioText",
							"asr_text",
						},
						To: []string{
							"data",
							"text",
						},
					},
					{
						From: []string{
							"data",
							"__product__",
						},
						To: []string{
							"__product__",
						},
					},
				},
			},
			args: args{
				features: map[string]interface{}{
					"data": map[string]interface{}{
						"__product__": "value1",
					},
					"audio-asr": map[string]interface{}{
						"audioText": map[string]interface{}{
							"asr_text": "",
						},
					},
				},
			},
			want: map[string]interface{}{
				"__product__": "value1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Send{
				SendData:         tt.fields.SendData,
				SendAllFeatures:  tt.fields.SendAllFeatures,
				SendFeatures:     tt.fields.SendFeatures,
				SendFeaturesKeys: tt.fields.SendFeaturesKeys,
			}
			if got := s.Run(context.Background(), tt.args.features); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
