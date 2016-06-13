package main

import "fmt"

var (
	regs [16]uint32
)

var (
	code  = [...]uint32{0x10000010}
	heap  []uint32
	stack []uint32
)

func main() {
	for regs[IP] < uint32(len(code)) {
		decode(code[regs[IP]])
		regs[IP]++
		printRegs(regs)
	}
	fmt.Printf("Execution completed.\n")
}

func decode(instr uint32) {
	fmt.Printf("OP = %v\n", Op(instr))
	switch Op(instr) {
	case OpDP:
		execDP(instr)
	case OpU1:
		undefined(instr)
	case OpME:
	case OpU3:
		undefined(instr)
	case OpU4:
		undefined(instr)
	case OpBR:
	case OpU6:
		undefined(instr)
	case OpU7:
		undefined(instr)
	}
}

func execDP(instr uint32) {
	fmt.Printf("Data processing.\n")
	if IsImm(instr) {
		regs[Rz(instr)] = compute(AluOp(instr), regs[Rx(instr)], Imm16(instr))
		fmt.Printf("%v = %v ? %v\n", regName(Rz(instr)), regName(Rx(instr)), Imm16(instr))
	} else {
		regs[Rz(instr)] = compute(AluOp(instr), regs[Rx(instr)], regs[Ry(instr)])
		fmt.Printf("%v = %v ? %v\n", regName(Rz(instr)), regName(Rx(instr)), regName(Ry(instr)))
	}
}

func compute(aluOp uint32, left uint32, right uint32) uint32 {
	switch aluOp {
	case AluOpAdd:
		return left + right
	case AluOpSub:
		return left - right
	case AluOpMul:
		return left * right
	case AluOpDiv:
		return left / right
	case AluOpAnd:
		return left & right
	case AluOpOor:
		return left | right
	case AluOpNor:
		return ^(left | right)
	case AluOpXor:
		return left ^ right
	case AluOpSll:
		return left << right
	case AluOpSrl:
		return left >> right
	case AluOpSra:
		return uint32(int(left) >> right)
	case AluOpRor:
		return (left << (32 - right)) | (left >> right)
	}
	return 0
}

func undefined(instr uint32) {
	fmt.Printf("Undefined op code instruction '%v'\n", instr)
}
