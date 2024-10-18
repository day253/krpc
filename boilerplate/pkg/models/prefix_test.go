package models

import (
	"context"
	"reflect"
	"testing"
)

func TestObjectPrefix_Run(t *testing.T) {
	type fields struct {
		VariablePrefix string
	}
	type args struct {
		details map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		{
			name: "case1",
			fields: fields{
				VariablePrefix: "prefix",
			},
			args: args{
				details: map[string]interface{}{
					"key1": "value1",
				},
			},
			want: map[string]interface{}{
				"prefix": map[string]interface{}{
					"key1": "value1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ObjectPrefix{
				VariablePrefix: tt.fields.VariablePrefix,
			}
			if got := s.Run(context.Background(), tt.args.details); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjectPrefix.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
