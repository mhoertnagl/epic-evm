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
	R10 = 0xA
	R11 = 0xB
	R12 = 0xC
	R13 = 0xD
	FP  = 0xD
	R14 = 0xE
	SP  = 0xE
	R15 = 0xF
	IP  = 0xF
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
		return "fp "
	case R14:
		return "sp "
	case R15:
		return "ip "
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
	AluOpSra = 0xA
	AluOpRor = 0xB // Better Lift LFT
)
