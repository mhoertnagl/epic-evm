package main

import "fmt"

// Set of valid registers.
const (
	R00 = 0x0
	R01 = 0x1
	R02 = 0x2
	R03 = 0x3
	R04 = 0x4
	R05 = 0x5
	R06 = 0x6
	R07 = 0x7
	R08 = 0x8
	R09 = 0x9
	R10 = 0xa
	R11 = 0xb
	R12 = 0xc
	R13 = 0xd
	R14 = 0xe
	R15 = 0xf
)

func regName(reg uint32) string {
	switch reg {
	case R00:
		return "r0 "
	case R01:
		return "r1 "
	case R02:
		return "r2 "
	case R03:
		return "r3 "
	case R04:
		return "r4 "
	case R05:
		return "r5 "
	case R06:
		return "r6 "
	case R07:
		return "r7 "
	case R08:
		return "r8 "
	case R09:
		return "r9 "
	case R10:
		return "r10"
	case R11:
		return "r11"
	case R12:
		return "r12"
	case R13:
		return "r13"
	case R14:
		return "r14"
	case R15:
		return "r15"
	}
	return fmt.Sprintf("?%v", reg)
}

func printRegs(regs [16]uint32) {
	for i := 0; i < 16; i++ {
		fmt.Printf("%v = %v\n", regName(uint32(i)), regs[i])
	}
}

const (
	OpDP uint32 = 0x0
	OpU1 uint32 = 0x1
	OpME uint32 = 0x2
	OpU3 uint32 = 0x3
	OpU4 uint32 = 0x4
	OpBR uint32 = 0x5
	OpU6 uint32 = 0x6
	OpU7 uint32 = 0x7
)

// Set of valid alu operations.
const (
	AluOpAdd = 0x0
	AluOpSub = 0x1
	AluOpMul = 0x2
	AluOpDiv = 0x3

	AluOpAnd = 0x4
	AluOpOor = 0x5
	AluOpNor = 0x6 // Or Nand?
	AluOpXor = 0x7

	AluOpSll = 0x8
	AluOpSrl = 0x9
	AluOpSra = 0xa
	AluOpLft = 0xb // Lift shifts left ry by 16 bit

	AluOpTst = 0xc // Test
	AluOpCmp = 0xd // Compare
	AluOpMov = 0xe // Move
	AluOpMvn = 0xf //
)

func aluOpName(aluOp uint32) string {
	switch aluOp {
	case AluOpAdd:
		return "+"
	case AluOpSub:
		return "-"
	case AluOpMul:
		return "*"
	case AluOpDiv:
		return "/"
	case AluOpAnd:
		return "&"
	case AluOpOor:
		return "|"
	case AluOpNor:
		return "~|"
	case AluOpXor:
		return "^"
	case AluOpSll:
		return "<<"
	case AluOpSrl:
		return ">>"
	case AluOpSra:
		return ">>>"
	case AluOpRor:
		return "<>>"
	}
	return fmt.Sprintf("?%v", aluOp)
}
