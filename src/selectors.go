package main

type Instr uint32

func Op(instr Instr) uint32 {
	return Bits(instr, 29, 3)
}

func IsImm(instr Instr) bool {
	return IsSet(instr, 28)
}

func IsLinkBranch(instr Instr) bool {
	return IsSet(instr, 28)
}

func Rz(instr Instr) uint32 {
	return Bits(instr, 24, 4)
}

func Rx(instr Instr) uint32 {
	return Bits(instr, 20, 4)
}

func Ry(instr Instr) uint32 {
	return Bits(instr, 16, 4)
}

func Imm16(instr Instr) uint32 {
	return Bits(instr, 4, 16)
}

func AluOp(instr Instr) uint32 {
	return Bits(instr, 0, 4)
}

func Offset(instr Instr) int32 {
	return int32(Bits(instr, 0, 25))
}
