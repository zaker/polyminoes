package polymino

import (
	"github.com/zaker/polyminoes/binarymatrix"
)

type Polymino struct {
	Size    uint8
	array2d *binarymatrix.BinaryMatrix[uint8]
}

func (p Polymino) String() string {
	return p.array2d.String()
}
func (p Polymino) Code() string {
	return p.array2d.Code()
}
func (p Polymino) Hash() string {
	return p.array2d.Hash()
}

func (p Polymino) Value() uint8 {
	return p.array2d.Value()
}

func (p Polymino) SetBit(x, y uint8, b bool) {
	p.array2d.SetBit(x, y, b)
}
func (p Polymino) Dup() Polymino {
	np := New(p.Size)
	np.array2d = p.array2d.Dup()
	return np
}

func Load(n uint8, s string) Polymino {
	p := New(n)

	p.array2d = binarymatrix.Load[uint8](n, s)
	return p
}

func (p *Polymino) SetToMinRotation() {
	minRot := p.array2d

	rot1 := minRot.Rot90().Crop(p.Size)
	rot2 := rot1.Rot90().Crop(p.Size)
	rot3 := rot2.Rot90().Crop(p.Size)
	if rot1.Less(minRot) {
		minRot = rot1
	}

	if rot2.Less(minRot) {
		minRot = rot2
	}
	if rot3.Less(minRot) {
		minRot = rot3
	}

	p.array2d = minRot
}

func New(n uint8) Polymino {
	return Polymino{n, binarymatrix.New[uint8](n, n)}
}

func (p *Polymino) blit(op Polymino, xOffsett, yOffset uint8) {
	sz := int(op.Size * op.Size)
	c := 0
	for i := 0; i < sz; i++ {
		if c >= int(op.Size) {
			return
		}
		x := uint8(i % int(op.Size))
		y := uint8(i / int(op.Size))
		if op.array2d.GetBit(x, y) {
			p.SetBit(x+xOffsett, y+yOffset, true)
			c++
		}

	}
}

func (p Polymino) Expand(expPoly map[string]string) {

	expansions := []Polymino{}
	sz := int(p.Size * p.Size)
	for i := 0; i < sz; i++ {

		x := uint8(i % int(p.Size))
		y := uint8(i / int(p.Size))
		last := p.Size - 1
		curBit := p.array2d.GetBit(x, y)
		// U
		if y == 0 && curBit {
			np := New(p.Size + 1)
			np.SetBit(x, y, true)
			np.blit(p, 0, 1)
			expansions = append(expansions, np)
		} else if y > 0 && curBit && !p.array2d.GetBit(x, y-1) {
			np := New(p.Size + 1)
			np.SetBit(x, y-1, true)
			np.blit(p, 0, 0)
			expansions = append(expansions, np)
		}

		// L
		if x == 0 && curBit {
			np := New(p.Size + 1)
			np.SetBit(x, y, true)
			np.blit(p, 1, 0)
			expansions = append(expansions, np)
		} else if x > 0 && curBit && !p.array2d.GetBit(x-1, y) {
			np := New(p.Size + 1)
			np.SetBit(x-1, y, true)
			np.blit(p, 0, 0)
			expansions = append(expansions, np)
		}

		// D
		if y == last && curBit {
			np := New(p.Size + 1)
			np.SetBit(x, y+1, true)
			np.blit(p, 0, 0)
			expansions = append(expansions, np)
		} else if y < last && curBit && !p.array2d.GetBit(x, y+1) {
			np := New(p.Size + 1)
			np.SetBit(x, y+1, true)
			np.blit(p, 0, 0)
			expansions = append(expansions, np)
		}

		// R
		if x == last && curBit {
			np := New(p.Size + 1)
			np.SetBit(x+1, y, true)
			np.blit(p, 0, 0)
			expansions = append(expansions, np)
		} else if x < last && curBit && !p.array2d.GetBit(x+1, y) {
			np := New(p.Size + 1)
			np.SetBit(x+1, y, true)
			np.blit(p, 0, 0)
			expansions = append(expansions, np)
		}

	}
	for _, v := range expansions {
		v.SetToMinRotation()
		expPoly[v.Hash()] = v.Code()
	}

}
