package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()

	for _, inFileName := range flag.Args() {
		inFile, err := os.Open(inFileName)
		if err != nil {
			panic(err)
		}
	}
}
