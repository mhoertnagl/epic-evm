package evm

import (
  "fmt"
  "evm/isa"
)

var (
  ip   uint32
  regs [16]uint32
)

var (
  code  = [...]uint32{0x10000010}
  heap  []uint32
  stack []uint32
)

func main() {
  for {
    decode(code[ip])
    ip++
  }
}

var opTable = [...](func(uint32)){execDP, undefined}

func decode(instr uint32) {
  opTable[Op(instr)](instr)
}

func execDP(instr uint32) {
  if IsImm(instr) {

  }
}

func compute(aluOp uint32, left uint32, right uint32) uint32 {
  switch aluOp & 0xF {
  case AluOpAdd:
    return left + right
  }
  return 0
}

func undefined(instr uint32) {
  fmt.Printf("Undefined op code instruction '%v'\n", instr)
}
