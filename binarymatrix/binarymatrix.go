package binarymatrix

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
)

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type BinaryMatrix[T Unsigned] struct {
	Height, Width uint8
	binarySize    uint8
	data          []T
}

func sizeOf[T Unsigned](value T) uint8 {
	switch any(value).(type) {
	case uint8:
		return 8
	case uint16:
		return 16
	case uint32:
		return 32
	case uint64:
		return 64
	default:
		return 0
	}

}

func New[T Unsigned](Height, Width uint8) *BinaryMatrix[T] {
	bt := T(0)
	sizofBT := sizeOf(bt)

	return &BinaryMatrix[T]{
		Height:     Height,
		Width:      Width,
		binarySize: sizofBT,
		data:       make([]T, ((Height*Width)/(sizofBT))+1),
	}
}

func (bm *BinaryMatrix[T]) SetBytes(bs []byte) error {

	return nil
}

func (bm *BinaryMatrix[T]) SetBit(x, y uint8, b bool) error {
	if x >= bm.Width {
		return fmt.Errorf("setting bit where x at %d is out of range [%d]", x, bm.Width)
	}
	if y >= bm.Height {
		return fmt.Errorf("setting bit where y at %d is out of range [%d]", x, bm.Height)
	}

	o := (y*bm.Width + x)
	i := (int)(o / bm.binarySize)
	m := (T)(1 << (o % bm.binarySize))

	if b {
		bm.data[i] |= m
	} else {
		bm.data[i] &= ^m
	}

	return nil
}

func (bm *BinaryMatrix[T]) GetBit(x, y uint8) bool {

	bs := bm.binarySize
	o := (y*bm.Width + x)
	i := (int)(o / bs)
	m := (T)(1 << (o % bs))
	c := bm.data[i]&(m) > 0
	return c
}

func (bm *BinaryMatrix[T]) Value() uint8 {

	var c uint8 = 0
	for y := uint8(0); y < bm.Height; y++ {
		for x := uint8(0); x < bm.Width; x++ {
			if bm.GetBit(x, y) {
				c++
			}
		}
	}
	return c
}

func (bm *BinaryMatrix[T]) Less(cm *BinaryMatrix[T]) bool {
	if len(bm.data) != len(cm.data) {
		return false
	}
	for i := len(bm.data) - 1; i >= 0; i-- {
		if bm.data[i] < cm.data[i] {
			return true
		}

		if bm.data[i] > cm.data[i] {
			return false
		}
	}
	return false
}

func castBTToBytes[T Unsigned](vals []T) []byte {
	buf := &bytes.Buffer{}
	for _, v := range vals {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			panic("binary.Write failed:" + err.Error())
		}
	}

	bs := buf.Bytes()
	return bytes.TrimRight(bs, "\x00")
}

func readFromBytes[T Unsigned](bs []byte, vs []T) {
	buf := bytes.NewReader(bs)

	err := binary.Read(buf, binary.LittleEndian, &vs)
	if err != nil {
		panic("binary.Read failed:" + err.Error())
	}

}

func (bm *BinaryMatrix[T]) Code() string {
	d := []byte{bm.Width, bm.Height}
	d = append(d, castBTToBytes(bm.data)...)

	return base64.RawStdEncoding.EncodeToString(d)
}

func (bm *BinaryMatrix[T]) Hash() string {
	d := []byte{bm.Width, bm.Height}
	d = append(d, castBTToBytes(bm.data)...)

	return base64.RawStdEncoding.EncodeToString(d)
}

func (bm *BinaryMatrix[T]) String() string {
	sb := strings.Builder{}
	sb.Grow((int(bm.Height) * int(bm.Width)) + int(bm.Height))
	for y := uint8(0); y < bm.Height; y++ {
		for x := uint8(0); x < bm.Width; x++ {

			if bm.GetBit(x, y) {
				sb.WriteRune('0')
			} else {
				sb.WriteRune('_')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (bm *BinaryMatrix[T]) Dup() *BinaryMatrix[T] {
	d := []T{}
	copy(d, bm.data)
	return &BinaryMatrix[T]{Width: bm.Width, Height: bm.Height, binarySize: bm.binarySize, data: d}
}
func (bm *BinaryMatrix[T]) Rot90() *BinaryMatrix[T] {
	bmr := New[T](bm.Width, bm.Height)
	cnt := 1
	for y := uint8(0); y < bm.Height; y++ {
		for x := uint8(0); x < bm.Width; x++ {
			if bm.GetBit(x, y) {
				bmr.SetBit(bmr.Width-y-1, x, true)
				cnt++
			}
			if cnt > int(bmr.Width) {
				goto end
			}
		}
	}
end:
	return bmr
}

func (bm *BinaryMatrix[T]) Crop(n uint8) *BinaryMatrix[T] {

	work := true
	for work {
		work = false
		xShift := true
		yShift := true
		for x := uint8(0); x < bm.Width; x++ {
			yShift = !bm.GetBit(x, 0) && yShift
		}

		for y := uint8(0); y < bm.Height; y++ {
			xShift = !bm.GetBit(0, y) && xShift
		}

		if yShift {
			yc := uint8(0)
			nbm := New[T](bm.Height, bm.Width)
			for y := uint8(1); y < nbm.Height; y++ {
				if yc >= n {
					break
				}
				for x := uint8(0); x < nbm.Width; x++ {
					if bm.GetBit(x, y) {
						nbm.SetBit(x, y-1, true)
						yc++
					}
				}
			}
			bm = nbm
			work = true
		}

		if xShift {
			xc := uint8(0)
			nbm := New[T](bm.Height, bm.Width)
			for y := uint8(0); y < nbm.Height; y++ {
				if xc >= n {
					break
				}
				for x := uint8(1); x < nbm.Width; x++ {
					if bm.GetBit(x, y) {
						nbm.SetBit(x-1, y, true)
						xc++
					}

				}
			}
			bm = nbm
			work = true
		}

	}

	return bm
}

func Load[T Unsigned](n uint8, s string) *BinaryMatrix[T] {
	bm := New[T](n, n)
	db, _ := base64.RawStdEncoding.DecodeString(s)
	sz := int(db[0] * db[1])
	db = append(db, make([]byte, sz-len(db[2:]))...)
	readFromBytes(db[2:], bm.data)
	return bm
}
