package main

import "fmt"

/**
 *
 */
type Cpu struct {
	// Instruction Pointer (IP) indicates the position of the current Instruction
	// word, which is 4 bytes long. Incrementing the IP by one is equal to
	// advancing for a single instruction. Epic supports a 34 bit address space.
	// If we copy IP to another register it will NOT be shifted to the left by 2
	// bits.
	Regs [16]int32
}

func (cpu Cpu) Execute(code []Instr) {
	codeLength := len(code)
	for uint32(cpu.Regs[IP]) < uint32(codeLength) {
		cpu.step(code[uint32(cpu.Regs[IP])])
	}
	fmt.Printf("Execution completed.\n")
}

func (cpu Cpu) step(instr Instr) {
	fmt.Printf("OP = %v\n", Op(instr))
	switch Op(instr) {
	case OpDP:
		cpu.executeDP(instr)
	case OpU1:
		cpu.undefined(instr)
	case OpME:
		cpu.executeME(instr)
	case OpU3:
		cpu.undefined(instr)
	case OpU4:
		cpu.undefined(instr)
	case OpBR:
		cpu.executeBR(instr)
	case OpU6:
		cpu.undefined(instr)
	case OpU7:
		cpu.undefined(instr)
	}
	// Step forward to the next instruction word. Branches will decrement the IP
	// to compensate for the extra increment here.
	cpu.Regs[IP]++
}

func (cpu Cpu) executeDP(instr Instr) {
	fmt.Printf("Data processing.\n")
	rz := Rz(instr)
	rx := Rx(instr)
	ry := Ry(instr)
	op := AluOp(instr)
	vx := cpu.Regs[rx]
	var vy int32
	if IsImm(instr) {
		vy = Imm16(instr)
	} else {
		vy = cpu.Regs[ry]
	}
	cpu.Regs[rz] = compute(op, vx, vy)
	if IsImm(instr) {
		fmt.Printf("%v = %v %v %v\n", regName(rz), regName(rx), aluOpName(op), vy)
	} else {
		fmt.Printf("%v = %v %v %v\n", regName(rz), regName(rx), aluOpName(op), regName(ry))
	}
}

func (cpu Cpu) executeME(instr Instr) {
	fmt.Printf("Memory operation.\n")
	rz := Rz(instr)
	rx := Rx(instr)
	ry := Ry(instr)
	vx := cpu.Regs[rx]
	var vy int32
	if IsImm(instr) {
		vy = Imm16(instr) << 2
	} else {
		vy = cpu.Regs[ry]
	}
	if IsLoad(instr) {
		cpu.Regs[rz] = cpu.Mem[vx+vy] // LDW rz rx[ry] oder LDW rz rx[8] <-- Kontext klärt ob immediate oder register.
	} else {
		cpu.Mem[vx+vy] = cpu.Regs[rz] // STW rz rx[ry] oder STW rz rx[8]
	}
	// Print instruction.
}

func (cpu Cpu) executeBR(instr Instr) {
	fmt.Printf("Branch instruction.\n")
	// Save the adress of the next instruction to the return pointer (RP) if it
	// is a linking branch.
	if IsLinkBranch(instr) {
		cpu.Regs[RP] = cpu.Regs[IP] + 1
	}
	// Add the offset to the instruction pointer (IP) and decrement by one as the
	// cpu will advance it in every cycle.
	cpu.Regs[IP] += Offset(instr) - 1
}

func (cpu Cpu) undefined(instr Instr) {
	fmt.Printf("%v: Undefined op code instruction '%v'\n", cpu.Regs[IP], instr)
}
