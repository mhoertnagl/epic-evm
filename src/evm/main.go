package evm

import "fmt"

var regs [16]uint32

func op(instr uint32) uint32 {
	return bits(instr, 29, 3)
}

func main() {
	for i, v := range regs {
		fmt.Printf("r%v = %v\n", i, v)
	}
}
