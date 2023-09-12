package generator

import (
	"reflect"
	"testing"
)

func TestGenerator_GeneratePolyminoes(t *testing.T) {
	type fields struct {
		cache map[int]map[string]struct{}
	}
	type args struct {
		n uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{"", fields{cache: map[int]map[string]struct{}{}}, args{n: 3}, map[string]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				cache: tt.fields.cache,
			}
			if got := g.GeneratePolyminoes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.GeneratePolyminoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
