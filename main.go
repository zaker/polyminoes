package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/zaker/polyminoes/generator"
)

func main() {

	g := generator.New()

	f, _ := os.Create("polyminoes.pprof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	n := 11
	startTime := time.Now()
	polyminoes := g.GeneratePolyminoes((uint8)(n))
	// for k, p := range polyminoes {
	// 	fmt.Print(k, "\n", polymino.Load((uint8)(n), p))
	// 	fmt.Println("--------------")
	// }
	fmt.Println(len(polyminoes), time.Since(startTime))

	// for i := uint8(1); i < 15; i++ {
	// 	fmt.Println("Polyminos:", i)
	// 	fmt.Println("--------------")
	// 	startTime := time.Now()
	// 	polyminoes := g.GeneratePolyminoes(i)
	// 	// for k, p := range polyminoes {
	// 	// 	fmt.Print(k, "\n", p)
	// 	// 	fmt.Println("--------------")
	// 	// }
	// 	fmt.Println(len(polyminoes), time.Since(startTime))
	// 	// g.StoreCache(i, polyminoes)
	// }

}
