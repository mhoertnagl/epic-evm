package main

import "fmt"

type Cpu struct {
	Regs [16]uint32
}

func (cpu Cpu) Execute(code []Instr) {
	ip := cpu.Regs[IP]
	for ip < uint32(len(code)) {
		cpu.executeOne(code[ip])
	}
	fmt.Printf("Execution completed.\n")
}

func (cpu Cpu) executeOne(instr Instr) {
	fmt.Printf("OP = %v\n", Op(instr))
	switch Op(instr) {
	case OpDP:
		cpu.executeDP(instr)
	case OpU1:
		cpu.undefined(instr)
	case OpME:
	case OpU3:
		cpu.undefined(instr)
	case OpU4:
		cpu.undefined(instr)
	case OpBR:
	case OpU6:
		cpu.undefined(instr)
	case OpU7:
		cpu.undefined(instr)
	}
	cpu.Regs[IP]++
}

func (cpu Cpu) executeDP(instr Instr) {
	fmt.Printf("Data processing.\n")
	if IsImm(instr) {
		cpu.Regs[Rz(instr)] = compute(AluOp(instr), cpu.Regs[Rx(instr)], Imm16(instr))
		fmt.Printf("%v = %v ? %v\n", regName(Rz(instr)), regName(Rx(instr)), Imm16(instr))
	} else {
		cpu.Regs[Rz(instr)] = compute(AluOp(instr), cpu.Regs[Rx(instr)], cpu.Regs[Ry(instr)])
		fmt.Printf("%v = %v ? %v\n", regName(Rz(instr)), regName(Rx(instr)), regName(Ry(instr)))
	}
}

func (cpu Cpu) executeBR(instr Instr) {
	if IsLinkBranch(instr) {
		cpu.Regs[RT] = cpu.Regs[IP] + 4
	}
	cpu.Regs[IP] = cpu.Regs[IP] + Offset(instr)
}

func (cpu Cpu) undefined(instr Instr) {
	fmt.Printf("%v: Undefined op code instruction '%v'\n", cpu.Regs[IP], instr)
}
