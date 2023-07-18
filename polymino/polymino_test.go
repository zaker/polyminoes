package polymino

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPolymino_StrideCount(t *testing.T) {
	tests := []struct {
		name string
		p    Polymino
		want *big.Int
	}{
		{"1B", func() Polymino {
			p := New(4)
			p[0][0] = true
			return p
		}(), func() *big.Int {
			z := &big.Int{}
			z.SetString("1B", 62)
			return z
		}()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.StrideCount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polymino.StrideCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
