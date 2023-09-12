package polymino

import (
	"fmt"
	"testing"
)

func TestPolymino_Expand(t *testing.T) {
	p3 := Load(3, "AwML")
	map4 := map[string]string{}
	p3.Expand(map4)

	for k, ps4 := range map4 {
		p4 := Load(4, ps4)
		fmt.Println("-4-", k, ps4)
		fmt.Println(p4)
		fmt.Println("-4-")

	}

	p2 := Load(2, "AgID")
	if !p2.array2d.GetBit(0, 0) && !p2.array2d.GetBit(1, 0) {
		t.Errorf("Not a valid p2")
	}
}

func TestPolymino_SetToMinRotation(t *testing.T) {

	tests := []struct {
		name string
		size uint8
		want string
	}{
		{"BAQRAw", 4, "BAQX"},
		{"BAQX", 4, "BAQX"},
		{"BQXj", 5, "BQXH"},
		{"BQXH", 5, "BQXH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Load(tt.size, tt.name)
			p.SetToMinRotation()
			if p.Code() != tt.want {
				t.Errorf("Polymino.StrideCount() = %v, want %v", p.Code(), tt.want)
			}
		})
	}
}
