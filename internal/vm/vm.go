package vm

import (
	"encoding/binary"
)

type VM struct {
	cir  uint32
	csr  uint32
	regs [16]uint32
	mem  []byte
}

func NewVM(mem []byte) *VM {
	return &VM{mem: mem}
}

func (m *VM) Run() {
	for {
		ins := m.ins(m.mem)
		if m.condPassed(ins) {
			switch op(ins) {
			case OpDPR:
				m.execDPR(ins)
				if doesNotWriteIP(ins) {
					m.regs[IP] += 4
				}
			case OpIMM:
				m.execIMM(ins)
				if doesNotWriteIP(ins) {
					m.regs[IP] += 4
				}
			case OpMEM:
				m.regs[IP] += 4
			case OpNA3:
				panic("unsopported NA3 operation")
			case OpNA4:
				panic("unsopported NA4 operation")
			case OpNA5:
				panic("unsopported NA5 operation")
			case OpNA6:
				panic("unsopported NA6 operation")
			case OpBRA:
			}
		} else {
			m.regs[IP] += 4
		}
	}
}

func (m *VM) execDPR(ins uint32) {
	ra := ra(ins)
	va := m.regs[ra]
	var vb uint32
	if isImm12(ins) {
		vb = imm12(ins)
		if ShoudSignExtend(ins) {
			vb = Sext(vb, 12)
		}
	} else {
		vb = m.regs[rb(ins)]
		vb = Shift(vb, sop(ins), shamt(ins))
	}
	m.alu(ins, va, vb)
}

func (m *VM) execIMM(ins uint32) {
	ra := ra(ins)
	va := m.regs[ra]
	vb := imm16(ins)
	if ShoudSignExtend(ins) {
		vb = Sext(vb, 16)
	}
	if isSll16(ins) {
		vb = vb << 16
	}
	m.alu(ins, va, vb)
}

func (m *VM) execMEM(ins uint32) {
	rd := rd(ins)
	ra := ra(ins)
	va := m.regs[ra]
	var vb uint32
	if isImm12(ins) {
		vb = imm12(ins)
		vb = Sext(vb, 12)
	} else {
		vb = m.regs[rb(ins)]
		vb = Shift(vb, sop(ins), shamt(ins))
	}
	// a := uint32(int64(uint64(va)) + int64(Sext64(vb)))
	a := uint32(int64(va) + int64(Sext64(vb)))
	if isLoad(ins) {
		m.regs[rd] = read32(m.mem, a)
	} else {
		write32(m.mem, a, m.regs[rd])
	}

	// Sign extend the value in register RD and set condition flags accordingly.
	if ShouldSetCond(ins) {
		m.setCond(Sext64(m.regs[rd]))
	}
}

func (m *VM) alu(ins uint32, va uint32, vb uint32) {
	rd := rd(ins)
	rs := Alu(aluop(ins), va, vb)

	if ShouldWriteBack(ins) {
		m.regs[rd] = uint32(rs)
	}

	if ShouldSetCond(ins) {
		m.setCond(rs)
	}
}

func (m *VM) condPassed(ins uint32) bool {
	gt := Bit(m.csr, 28) & Bit(ins, 28)
	lt := Bit(m.csr, 27) & Bit(ins, 27)
	eq := Bit(m.csr, 26) & Bit(ins, 26)
	al := Bit(ins, 28) & Bit(ins, 27) & Bit(ins, 26)
	return (gt | lt | eq | al) == 1
}

func (m *VM) setCond(rs uint64) {
	eq := rs == 0
	lt := Bit64(rs, 32) == 1
	gt := !eq && !lt

	m.csr = SetBool(m.csr, 26, eq)
	m.csr = SetBool(m.csr, 27, lt)
	m.csr = SetBool(m.csr, 28, gt)
}

func (m *VM) ins(code []byte) uint32 {
	return read32(code, m.regs[IP])
}

func doesNotWriteIP(ins uint32) bool {
	return (ShouldWriteBack(ins) && rd(ins) == IP) == false
}

func read32(data []byte, a uint32) uint32 {
	return binary.BigEndian.Uint32(data[a : a+4])
}

func write32(data []byte, a uint32, v uint32) {
	binary.BigEndian.PutUint32(data[a:a+4], v)
}
