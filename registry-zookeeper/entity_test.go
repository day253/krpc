package registry_zookeeper

import (
	"reflect"
	"testing"
)

func TestNodeEntity_Content(t *testing.T) {
	type fields struct {
		Host   string
		Port   int
		Weight int
		Tags   map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				Host:   "127.0.0.1",
				Port:   80,
				Weight: 1,
				Tags: map[string]string{
					"k1": "v1",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeInfo{
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
				Weight: tt.fields.Weight,
				Tags:   tt.fields.Tags,
			}
			if got := string(n.Data()); !reflect.DeepEqual(got, tt.want) {
				t.Logf("NodeEntity.Content() = %v, want %v", got, tt.want)
			}
		})
	}
}
