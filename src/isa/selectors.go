package isa

import "evm/util"

func Op(instr uint32) uint32 {
  return Bits(instr, 29, 3)
}

func IsImm(instr uint32) bool {
  return Set(instr, 28)
}

func Rz(instr uint32) uint32 {
  return Bits(instr, 24, 4)
}

func Rx(instr uint32) uint32 {
  return Bits(instr, 20, 4)
}

func Ry(instr uint32) uint32 {
  return Bits(instr, 16, 4)
}

func Imm16(instr uint32) uint32 {
  return Bits(instr, 4, 16)
}

func AluOp(instr uint32) uint32 {
  return Bits(instr, 0, 4)
}
