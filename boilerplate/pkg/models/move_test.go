package models

import (
	"context"
	"reflect"
	"testing"

	"github.com/day253/krpc/boilerplate/pkg/conf"
)

func TestMove_Run(t *testing.T) {
	type fields struct {
		MoveFeatures     bool
		MoveFeaturesKeys []conf.FeaturesKeys
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
			name: "move features",
			fields: fields{
				MoveFeatures: true,
				MoveFeaturesKeys: []conf.FeaturesKeys{
					{
						From: []string{
							"key1",
						},
						To: []string{
							"key3",
							"key31",
							"key311",
						},
					},
				},
			},
			args: args{
				features: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": map[string]interface{}{
					"key31": map[string]interface{}{
						"key311": "value1",
					},
				},
			},
		},
		{
			name: "move features 2",
			fields: fields{
				MoveFeatures: true,
				MoveFeaturesKeys: []conf.FeaturesKeys{
					{
						From: []string{
							"key1",
						},
						To: []string{
							"key3",
							"key31",
							"key311",
						},
					},
					{
						From: []string{
							"key2",
						},
						To: []string{
							"key3",
						},
					},
				},
			},
			args: args{
				features: map[string]interface{}{
					"key1": "",
					"key2": "value2",
				},
			},
			want: map[string]interface{}{
				"key1": "",
				"key2": "value2",
				"key3": "value2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Move{
				MoveFeatures:     tt.fields.MoveFeatures,
				MoveFeaturesKeys: tt.fields.MoveFeaturesKeys,
			}
			if got := s.Run(context.Background(), tt.args.features); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Move.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
