package polymino

import (
	"math/big"
	"strings"
	"unicode/utf8"
)

type Polymino [][]bool

func (p Polymino) String() string {
	s := ""
	for _, r := range p {
		l := ""
		for _, b := range r {
			if b {
				l += "#"
			} else {
				l += " "
			}
		}
		s += l + "\n"
	}
	return s
}
func (p Polymino) Size() uint {
	return (uint)(len(p))
}
func (p Polymino) Dup() Polymino {
	tmp := New(p.Size())
	for y, l := range p {

		copy(tmp[y], l)

	}
	return tmp
}
func (p Polymino) Crop(n uint) Polymino {
	tmp := p.Dup()
	moreWork := true
	for moreWork {
		moreWork = false
		spaceUp := false
		spaceLeft := false
		spaceRight := false
		spaceDown := false
		for i := 0; i < len(tmp); i++ {
			lr := len(tmp[0]) - 1
			ld := len(tmp) - 1
			spaceLeft = spaceLeft || tmp[i][0]
			spaceUp = spaceUp || tmp[0][i]
			spaceRight = spaceRight || tmp[lr][0]
			spaceDown = spaceDown || tmp[0][ld]
		}
		if !spaceUp && len(tmp) > (int)(n) {
			tmp = tmp[1:]
			moreWork = true
		}
		if !spaceLeft && len(tmp[0]) > (int)(n) {
			for i := range tmp {
				tmp[i] = tmp[i][1:]
			}

			moreWork = true
		}

		if !spaceDown && len(tmp) > (int)(n) {
			tmp = tmp[:len(tmp)-1]
			moreWork = true
		}

		if !spaceRight && len(tmp[0]) > (int)(n) {
			for i := range tmp {
				tmp[i] = tmp[i][:len(tmp[i])-1]
			}

			moreWork = true
		}

	}
	return tmp.Dup()
}

func (p Polymino) StrideCount() *big.Int {
	b := &strings.Builder{}
	b.Grow(len(p) ^ 2)
	for _, l := range p {
		for _, c := range l {
			if c {
				b.WriteString("1")
			} else {
				b.WriteString("0")
			}
		}
	}
	z := &big.Int{}
	c := ReverseString(b.String())
	z, _ = z.SetString(c, 2)

	return z

}

func LoadPolymino(n uint, s string) Polymino {
	p := New(n)
	z := big.NewInt(0)
	z, _ = z.SetString(s, 62)
	st := ReverseString(z.Text(2))

	for y, l := range p {
		ystride := y * (int)(n)
		if ystride > len(st) {
			break
		}
		for x := range l {
			xstride := ystride + x
			if xstride > len(st)-1 {
				break
			}
			p[y][x] = st[xstride] == (byte)('1')
		}
	}
	return p
}

func (p Polymino) Rot90() Polymino {
	np := New(p.Size())
	lenp := len(p)
	for y := 0; y < lenp; y++ {
		for x := 0; x < lenp; x++ {
			np[x][y] = p[y][x]
		}
	}
	for y, l := range np {
		np[y] = ReverseBytes(l)
	}

	moreWork := true
	for moreWork {
		moreWork = false
		spaceLeft := false
		for i := 0; i < lenp; i++ {
			spaceLeft = spaceLeft || np[i][0]
		}
		if !spaceLeft {
			for i := range np {
				np[i] = append(np[i][1:], false)
			}

			moreWork = true
		}
	}
	return np
}

func (p Polymino) BaseRotation(sc *big.Int, cache map[string]Polymino) (Polymino, bool) {
	baseStrideCount := sc
	baseRot := p
	for i := 0; i < 3; i++ {
		np := p.Rot90()
		sc := np.StrideCount()
		if _, ok := cache[sc.Text(62)]; ok {
			return baseRot, false
		}
		if sc.Cmp(baseStrideCount) < 0 {
			baseStrideCount = sc
			baseRot = np
		}
		p = np
	}
	return baseRot, true
}

func ReverseString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func ReverseBytes(b []bool) []bool {
	size := len(b) - 1
	buf := make([]bool, size+1)
	for start := 0; start <= size; start++ {
		buf[size-start] = b[start]
	}
	return buf
}

func New(n uint) Polymino {
	p := make([][]bool, n)
	for i := range p {
		p[i] = make([]bool, n)
	}
	return p
}

func (p Polymino) ExpandPolymino(expPoly map[string]Polymino) {

	tmp := New((uint)(len(p)) + 4)
	for y, l := range p {
		for x, b := range l {
			tmp[y+2][x+2] = b
		}
	}

	for y := 1; y < len(tmp)-1; y++ {

		for x := 1; x < len(tmp)-1; x++ {
			if tmp[y][x] {
				continue
			}
			if !(tmp[y-1][x] ||
				tmp[y+1][x] ||
				tmp[y][x-1] ||
				tmp[y][x+1]) {
				continue
			}

			cp := tmp.Dup()
			cp[y][x] = true
			cp = cp.Crop((uint)(len(p)) + 1)
			sc := cp.StrideCount()
			if br, ok := cp.BaseRotation(sc, expPoly); ok {

				expPoly[br.StrideCount().Text(62)] = br
			} else {
				continue
			}
		}
	}

	return

}
