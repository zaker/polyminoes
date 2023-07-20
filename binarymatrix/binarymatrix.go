package binarymatrix

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type BinaryMatrix[T Unsigned] struct {
	Height, Width uint8
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
		Height: Height,
		Width:  Width,
		data:   make([]T, ((Height*Width)/(sizofBT))+1),
	}
}

func (bm *BinaryMatrix[T]) SetBytes(bs []byte) error {

	return nil
}

func (bm *BinaryMatrix[T]) SetBit(x, y uint8, b bool) error {
	if x > bm.Width {
		return fmt.Errorf("setting bit where x at %d is out of range [%d]", x, bm.Width)
	}
	if y > bm.Height {
		return fmt.Errorf("setting bit where y at %d is out of range [%d]", x, bm.Height)
	}
	sizeofBT := sizeOf(bm.data[0])
	o := (y*bm.Width + x)
	i := (int)(o / sizeofBT)
	m := (T)(1 << (o % sizeofBT))

	if b {
		bm.data[i] |= m
	} else {
		bm.data[i] &= ^m
	}

	return nil
}

func (bm *BinaryMatrix[T]) getBit(x, y uint8) bool {
	sizeofBT := sizeOf(bm.data[0])
	o := (y*bm.Width + x)
	i := (int)(o / sizeofBT)
	m := (T)(1 << (o % sizeofBT))
	c := bm.data[i]&(m) > 0
	return c
}

func (bm *BinaryMatrix[T]) Less(cm *BinaryMatrix[T]) bool {
	if len(bm.data) != len(cm.data) {
		return false
	}
	for i := len(bm.data); i >= 0; i-- {
		if bm.data[i] < cm.data[i] {
			return true
		}
	}
	return false
}

func castBTToBytes[T Unsigned](vals []T) []byte {
	// length := len(ints) * int(sizofBT)
	// hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	// return *(*[]byte)(unsafe.Pointer(&hdr))
	buf := new(bytes.Buffer)
	for _, v := range vals {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			panic("binary.Write failed:" + err.Error())
		}
	}

	bs := buf.Bytes()
	return bytes.TrimRight(bs, "\x00")
}

func (bm *BinaryMatrix[T]) Code() string {
	d := []byte{bm.Width, bm.Height}
	d = append(d, castBTToBytes(bm.data)...)

	return base64.RawStdEncoding.EncodeToString(d)
}
