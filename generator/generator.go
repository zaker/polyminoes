package generator

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/zaker/polyminoes/polymino"
)

type Generator struct {
	cache map[int]map[string]struct{}
}

func New() *Generator {
	return &Generator{cache: map[int]map[string]struct{}{}}
}

func LoadPolyminoes(n uint8, m map[string]struct{}) map[string]string {
	pm := map[string]string{}
	for k := range m {
		pm[k] = polymino.Load(n, k).Code()
	}
	return pm
}
func (g *Generator) LoadCache(n int) (map[string]string, bool) {
	if n <= len(g.cache) {
		return LoadPolyminoes((uint8)(n), g.cache[n-1]), true
	}
	fileName := fmt.Sprintf("poly_%d.txt", n)
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return map[string]string{}, false
	}
	f, err := os.Open(fileName)
	if err != nil {
		log.Panicln("open file:", fileName, err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)
	nCache := map[string]struct{}{}
	for fileScanner.Scan() {
		nCache[fileScanner.Text()] = struct{}{}
	}
	g.cache[n] = nCache
	return LoadPolyminoes((uint8)(n), g.cache[n]), true

}

func (g *Generator) StoreCache(n int, polyminoes map[string]polymino.Polymino) {
	m := map[string]struct{}{}
	f, err := os.Create(fmt.Sprintf("poly_%d.txt", n))
	if err != nil {
		log.Fatal(err)
	}
	for k := range polyminoes {
		f.WriteString((k + "\n"))
		m[k] = struct{}{}
	}
	g.cache[n] = m
}

func (g *Generator) GeneratePolyminoes(n uint8) map[string]string {

	polyminoes := map[string]string{}

	if n < 1 {
		return polyminoes
	}
	if n == 1 {

		polymino := polymino.New(n)
		polymino.SetBit(0, 0, true)

		polyminoes[polymino.Hash()] = polymino.Code()

		return polyminoes
	}
	if n == 2 {

		polymino := polymino.New(n)
		polymino.SetBit(0, 0, true)
		polymino.SetBit(1, 0, true)
		polyminoes[polymino.Hash()] = polymino.Code()
		return polyminoes
	}

	if polyminoes, ok := g.LoadCache((int)(n)); ok {
		return polyminoes
	}

	for _, ps := range g.GeneratePolyminoes(n - 1) {
		p := polymino.Load(n-1, ps)
		p.Expand(polyminoes)
	}

	return polyminoes

}
