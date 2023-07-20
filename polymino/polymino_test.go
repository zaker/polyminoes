package polymino

import (
	"fmt"
	"testing"
)

// func TestPolymino_StrideCount(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		size uint
// 	}{
// 		{"7", 3},
// 		{"1B", 4},
// 		{"S", 4},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := LoadPolymino(tt.size, tt.name)

// 			if p.StrideCount().Cmp(p.StrideCount2()) != 0 {
// 				t.Errorf("Polymino.StrideCount() = %v, want %v", p.StrideCount().Text(62), p.StrideCount2().Text(62))
// 			}
// 		})
// 	}
// }

func TestPolymino_Expand(t *testing.T) {
	// type args struct {
	// 	expPoly map[string]string
	// }

	// n2 := "AgID"
	// p2 := polymino.Load(2, n2)
	// map3 := map[string]string{}
	// p2.Expand(map3)

	// for k, ps3 := range map3 {
	// 	p3 := polymino.Load(3, ps3)
	// 	fmt.Println("-3-", k, ps3)
	// 	fmt.Println(p3)
	// 	fmt.Println("-3-")

	// }

	p3 := Load(3, "AwML")
	map4 := map[string]string{}
	p3.Expand(map4)

	for k, ps4 := range map4 {
		p4 := Load(4, ps4)
		fmt.Println("-4-", k, ps4)
		fmt.Println(p4)
		fmt.Println("-4-")

	}
	// tests := []struct {
	// 	name string
	// 	p    Polymino
	// 	args args
	// }{
	// 	// TODO: Add test cases.

	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		tt.p.Expand(tt.args.expPoly)
	// 	})
	// }
}

func TestPolymino_SetToMinRotation(t *testing.T) {

	tests := []struct {
		name string
		size uint8
		want string
	}{
		{"BAQRAw", 4, "BAQX"},
		{"BAQX", 4, "BAQX"},
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
