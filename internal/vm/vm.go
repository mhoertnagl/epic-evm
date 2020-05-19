package vm

import (
	"encoding/binary"
)

// TODO: Sign extension is a problem. Define some unit tests.
// TODO: signed() - Turns a bit number into an integer
// TODO: cond flags
// TODO: set cond

type VM struct {
	cir  uint32
	csr  uint32
	regs [16]uint32
}

func NewVM() *VM {
	return &VM{}
}

func (m *VM) Run(code []byte) {
	for {
		ins := m.ins(code)
		switch op(ins) {
		case OpDPR:
			m.dpr(ins)
			if rd(ins) != IP {
				m.regs[IP] += 4
			}
		case OpIMM:
			m.imm(ins)
			if rd(ins) != IP {
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
	}
}

func (m *VM) dpr(ins uint32) {
	ra := ra(ins)
	va := m.regs[ra]
	var vb uint32
	// TODO: Je nach Operation sign oder zero extend.
	if isImm12(ins) {
		vb = imm12(ins)
	} else {
		vb = m.regs[rb(ins)]
	}
	vb = Shift(vb, sop(ins), shamt(ins))
	m.alu(ins, va, vb)
}

func (m *VM) imm(ins uint32) {
	ra := ra(ins)
	va := m.regs[ra]
	// TODO: Je nach Operation sign oder zero extend.
	vb := imm16(ins)
	if isSll16(ins) {
		vb = vb << 16
	}
	m.alu(ins, va, vb)
}

func (m *VM) alu(ins uint32, va uint32, vb uint32) {
	rd := rd(ins)
	rs := Alu(aluop(ins), va, vb)
	// TODO: Return value if cond are met
	m.regs[rd] = uint32(rs)
	// TODO: Some operations return a value and some always change cond
	if isSetCond(ins) {
		m.setCond(rs)
	}
}

func (m *VM) setCond(rs uint64) {
	eq := rs == 0
	lt := Bit(rs, 32) == 1
	gt := !eq && !lt

	m.csr = SetBool(m.csr, 26, eq)
	m.csr = SetBool(m.csr, 27, lt)
	m.csr = SetBool(m.csr, 28, gt)
}

func (m *VM) ins(code []byte) uint32 {
	ip := m.regs[IP]
	in := code[ip : ip+4]
	return binary.BigEndian.Uint32(in)
}

// func (m *VM) ip() uint32 {
// 	return m.regs[IP]
// }
