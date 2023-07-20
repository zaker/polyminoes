package binarymatrix

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		Height, Width uint8
	}
	tests := []struct {
		name string
		args args
		want *BinaryMatrix[uint64]
	}{
		{"1x1", args{1, 1}, &BinaryMatrix[uint64]{1, 1, 64, make([]uint64, ((1*1)/64)+1)}},
		{"3x3", args{3, 3}, &BinaryMatrix[uint64]{3, 3, 64, make([]uint64, ((3*3)/64)+1)}},
		{"5x5", args{5, 5}, &BinaryMatrix[uint64]{5, 5, 64, make([]uint64, ((5*5)/64)+1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[uint64](tt.args.Height, tt.args.Width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBinaryMatrix_SetBit(t *testing.T) {
	type args struct {
		x, y uint8
		b    bool
	}
	tests := []struct {
		name    string
		bm      *BinaryMatrix[uint64]
		args    args
		wantErr bool
	}{
		{"1x1 1", New[uint64](1, 1), args{0, 0, true}, false},
		{"2x2 8", New[uint64](2, 2), args{1, 1, true}, false},
		{"3x3 16-00", New[uint64](3, 3), args{1, 1, true}, false},
		{"3x3 00-01", New[uint64](3, 3), args{2, 2, true}, false},
		{"5x5 00-16-00-00", New[uint64](5, 5), args{2, 2, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bm.SetBit(tt.args.x, tt.args.y, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("BinaryMatrix.SetBit() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := tt.bm.GetBit(tt.args.x, tt.args.y)
			if !got {
				t.Errorf("getBit(%d,%d) == false , want true",
					tt.args.x,
					tt.args.y)
			}
		})
	}
}

func TestBinaryMatrix_Code(t *testing.T) {

	bm1 := New[uint64](8, 8)
	bm1.SetBit(4, 4, true)

	t.Run("base 64 8x8 middle", func(t *testing.T) {
		if got := bm1.Code(); got != "CAgAAAAAEA" {
			t.Errorf("BinaryMatrix.Code() = %v, want %v", got, "CAgAAAAAEA")
		}
	})

	bm2 := New[uint32](20, 20)
	bm2.SetBit(4, 4, true)
	bm2.SetBit(4, 5, true)
	bm2.SetBit(2, 5, true)
	t.Run("base 32 20x20 middle", func(t *testing.T) {
		if got := bm2.Code(); got != "FBQAAAAAAAAAAAAAEABAAQ" {
			t.Errorf("BinaryMatrix.Code() = %v, want %v", got, "FBQAAAAAAAAAAAAAEABAAQ")
		}
	})

}

func TestBinaryMatrix_Less(t *testing.T) {

	bm := New[uint8](3, 3)
	bm.SetBit(0, 0, true)

	bm1 := New[uint8](3, 3)
	bm1.SetBit(2, 2, true)
	if !bm.Less(bm1) {
		t.Errorf("BinaryMatrix.Code() = %v, want %v", false, true)
	}
	bm2 := New[uint8](3, 3)
	bm2.SetBit(0, 2, true)
	if !bm.Less(bm2) {
		t.Errorf("BinaryMatrix.Code() = %v, want %v", false, true)
	}

}

func TestLoad(t *testing.T) {

	gen4 := func() *BinaryMatrix[uint8] {
		t := New[uint8](4, 4)
		t.SetBit(0, 0, true)
		t.SetBit(1, 0, true)
		t.SetBit(2, 0, true)
		t.SetBit(1, 1, true)
		return t
	}

	type args struct {
		n uint8
		s string
	}
	tests := []struct {
		name string
		args args
		want *BinaryMatrix[uint8]
	}{
		// TODO: Add test cases.
		{"1x1", args{1, New[uint8](1, 1).Code()}, New[uint8](1, 1)},
		{"2x2", args{2, New[uint8](2, 2).Code()}, New[uint8](2, 2)},
		{"3x3", args{3, New[uint8](3, 3).Code()}, New[uint8](3, 3)},
		{"4x4", args{4, gen4().Code()}, gen4()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Load[uint8](tt.args.n, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
