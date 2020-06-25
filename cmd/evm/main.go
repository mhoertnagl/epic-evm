package main

import "github.com/mhoertnagl/epic-evm/internal/vm"

// import (
// 	"flag"
// 	"os"
// )

func main() {
	// flag.Parse()
	//
	// for _, inFileName := range flag.Args() {
	// 	inFile, err := os.Open(inFileName)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	vm := vm.NewVM()

	// for vm.Continue() {
	// 	for x := 0; x < 10; x++ {
	// 		for y := 0; y < 10; y++ {
	// 			vm.SetContent(x, y, 'R')
	// 		}
	// 	}
	// }

	vm.Finalize()
}
