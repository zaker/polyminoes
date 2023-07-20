package polymino

import (
	"testing"
)

func TestPolymino_StrideCount(t *testing.T) {
	tests := []struct {
		name string
		size uint
	}{
		{"7", 3},
		{"1B", 4},
		{"S", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := LoadPolymino(tt.size, tt.name)

			if p.StrideCount().Cmp(p.StrideCount2()) != 0 {
				t.Errorf("Polymino.StrideCount() = %v, want %v", p.StrideCount().Text(62), p.StrideCount2().Text(62))
			}
		})
	}
}
