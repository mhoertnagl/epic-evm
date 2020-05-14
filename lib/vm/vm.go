package vm

import (
	"encoding/binary"
	"math/bits"
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
	if isImm12(ins) {
		vb = imm12(ins)
	} else {
		vb = m.regs[rb(ins)]
	}
	vb = m.shift(vb, sop(ins), shamt(ins))
	m.alu(ins, va, vb)
}

func (m *VM) imm(ins uint32) {
	ra := ra(ins)
	va := m.regs[ra]
	vb := imm16(ins)
	if isSll16(ins) {
		vb = vb << 16
	}
	m.alu(ins, va, vb)
}

func (m *VM) alu(ins uint32, va uint32, vb uint32) {
	rd := rd(ins)
	switch aluop(ins) {
	case OpADD:
		m.regs[rd] = uint32(int32(va) + int32(vb))
	case OpSUB:
		m.regs[rd] = uint32(int32(va) - int32(vb))
	case OpMUL:
		m.regs[rd] = uint32(int32(va) * int32(vb))
	case OpDIV:
		m.regs[rd] = uint32(int32(va) / int32(vb))
	case OpAND:
		m.regs[rd] = va & vb
	case OpOOR:
		m.regs[rd] = va | vb
	case OpXOR:
		m.regs[rd] = va ^ vb
	case OpNOR:
		m.regs[rd] = ^(va | vb)
	case OpADU:
		m.regs[rd] = va + vb
	case OpSBU:
		m.regs[rd] = va - vb
	case OpMLU:
		m.regs[rd] = va * vb
	case OpDVU:
		m.regs[rd] = va / vb
	case OpCMP:
	case OpCPU:
	case OpTST:
	case OpMOV:
		m.regs[rd] = vb
	}
}

func (m *VM) shift(vb uint32, op SOp, shamt uint32) uint32 {
	switch op {
	case OpSLL:
		return vb << shamt
	case OpROL:
		return bits.RotateLeft32(vb, int(shamt))
	case OpSRL:
		return vb >> shamt
	case OpSRA:
		return uint32(int32(vb) >> shamt)
	}
	return 0
}

func (m *VM) ins(code []byte) uint32 {
	ip := m.regs[IP]
	return binary.BigEndian.Uint32(code[ip : ip+4])
}

// func (m *VM) ip() uint32 {
// 	return m.regs[IP]
// }
