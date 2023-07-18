package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/zaker/polyminos/polymino"
)

func (g Generator) GeneratePolyminoes(n uint) map[string]polymino.Polymino {

	polyminoes := map[string]polymino.Polymino{}

	if n < 1 {
		return polyminoes
	}
	if n == 1 {

		polymino := polymino.New(n)
		polymino[0][0] = true
		polyminoes[polymino.StrideCount().Text(62)] = polymino

		return polyminoes
	}
	if n == 2 {

		polymino := polymino.New(n)
		polymino[0] = []bool{true, true}
		polyminoes[polymino.StrideCount().Text(62)] = polymino
		return polyminoes
	}

	if polyminoes, ok := g.LoadCache((int)(n)); ok {
		return polyminoes
	}

	for _, p := range g.GeneratePolyminoes(n - 1) {
		p.ExpandPolymino(polyminoes)
	}

	return polyminoes

}

type Generator struct {
	cache map[int]map[string]struct{}
}

func (g *Generator) LoadCache(n int) (map[string]polymino.Polymino, bool) {
	if n <= len(g.cache) {
		return LoadPolyminoes((uint)(n), g.cache[n-1]), true
	}
	fileName := fmt.Sprintf("poly_%d.txt", n)
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return map[string]polymino.Polymino{}, false
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
	return LoadPolyminoes((uint)(n), g.cache[n]), true

}

func LoadPolyminoes(n uint, m map[string]struct{}) map[string]polymino.Polymino {
	pm := map[string]polymino.Polymino{}
	for k := range m {
		pm[k] = polymino.LoadPolymino(n, k)
	}
	return pm
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

func main() {

	g := Generator{cache: map[int]map[string]struct{}{}}

	f, _ := os.Create("polyminoes.pprof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	startTime := time.Now()
	polyminoes := g.GeneratePolyminoes((uint)(15))
	// for k, p := range polyminoes {
	// 	fmt.Print(k, "\n", p)
	// 	fmt.Println("--------------")
	// }
	fmt.Println(len(polyminoes), time.Since(startTime))

	// for i := 1; i < 14; i++ {
	// 	fmt.Println("Polyminos:", i)
	// 	fmt.Println("--------------")
	// 	startTime := time.Now()
	// 	polyminoes := g.GeneratePolyminoes((uint)(i))
	// 	// for k, p := range polyminoes {
	// 	// 	fmt.Print(k, "\n", p)
	// 	// 	fmt.Println("--------------")
	// 	// }
	// 	fmt.Println(len(polyminoes), time.Since(startTime))
	// 	g.StoreCache(i, polyminoes)
	// }

}
