package vm

import (
	"encoding/binary"
	"math/bits"
)

// TODO: Do I need a mvs - move signed
// TODO: Cannot move 32 bit value with two moves!!!

type VM struct {
	cir  uint32
	csr  uint32
	regs [32]uint32
	mem  []byte

	rd    uint32
	ra    uint32
	ma    bool
	ld    bool
	lnk   bool
	aop   uint32
	sc    bool
	sn    bool
	wb    bool
	br    bool
	inc   bool
	va    uint32
	vb    uint32
	vd    uint32
	sop   uint32
	shamt uint32
}

func NewVM(mem []byte) *VM {
	return &VM{mem: mem}
}

func (m *VM) Run() {
	len := uint32(len(m.mem) >> 2)
	for m.regs[IP] < len {
		ins := m.ins(m.mem)
		m.decode(ins)
		m.execute()
	}
}

func (m *VM) decode(ins uint32) {
	op := op(ins)
	cp := m.condPassed(ins)

	if op == OpBRA {
		m.rd = IP
		m.ra = IP
	} else {
		m.rd = rd(ins)
		m.ra = ra(ins)
	}

	aop := aluop(ins)
	switch op {
	case OpMEM:
		m.aop = OpADD
	case OpBRA:
		m.aop = OpADD
	default:
		if aop == OpCPS || aop == OpCPU {
			m.aop = OpSUB
		} else {
			m.aop = aop
		}
	}

	if op == OpMEM || op == OpBRA {
		m.sn = false
	} else {
		m.sn = aop == OpCPS
	}

	switch op {
	case OpMEM:
		m.sc = false
		m.wb = cp && isLoad(ins)
	case OpBRA:
		m.sc = false
		m.wb = false
	default:
		switch aop {
		case OpCPS:
			m.sc = cp
			m.wb = false
		case OpCPU:
			m.sc = cp
			m.wb = false
		default:
			m.sc = false
			m.wb = cp
		}
	}

	m.br = cp && (op == OpBRA || m.wb && m.rd == IP)

	if op == OpDPR || op == OpMEM {
		m.sop = sop(ins)
		m.shamt = shamt(ins)
	} else if op == OpI16 {
		if isHigh(ins) {
			m.shamt = 16
		} else {
			m.shamt = 0
		}
	} else {
		m.sop = OpSLL
		m.shamt = 0
	}

	m.va = m.regs[m.ra]

	if op == OpDPR || op == OpMEM {
		m.vb = m.regs[rb(ins)]
	} else if op == OpD12 || op == OpM12 {
		if m.sn {
			m.vb = Sext(imm12(ins), 12)
		} else {
			m.vb = imm12(ins)
		}
	} else if op == OpI16 {
		if m.sn {
			m.vb = Sext(imm16(ins), 16)
		} else {
			m.vb = imm16(ins)
		}
	} else if op == OpBRA {
		m.vb = Sext(imm25(ins), 25)
	} else {
		m.vb = 0
	}

	m.vd = m.regs[m.rd]

	m.lnk = op == OpBRA && isLink(ins)
	m.ma = op == OpMEM
	m.ld = op == OpMEM && isLoad(ins)
}

func (m *VM) execute() {

	vb := shift(m.vb, m.sop, m.shamt)
	rs, c := alu(m.aop, m.va, vb)

	if m.ma {
		if m.ld {
			rs = read32(m.mem, rs)
		} else {
			write32(m.mem, rs, m.vd)
		}
	}

	if m.sc {
		m.setEqualFlag(rs == 0)
		if m.sn {
			m.setLessFlag(Bit(rs, 31) == 1)
		} else {
			m.setLessFlag(c == 0)
		}
	}

	if m.lnk {
		m.regs[RP] = m.regs[IP] + 1
	}

	if m.wb {
		m.regs[m.rd] = rs
	}

	if m.br == false {
		m.regs[IP]++
	}
}

func alu(op uint32, va uint32, vb uint32) (uint32, uint32) {
	switch op {
	case OpADD:
		return bits.Add32(va, vb, 0)
	case OpSUB:
		return bits.Sub32(va, vb, 0)
	case OpMUL:
		return bits.Mul32(va, vb)
	case OpDIV:
		return bits.Div32(0, va, vb)
	case OpAND:
		return va & vb, 0
	case OpOOR:
		return va | vb, 0
	case OpXOR:
		return va ^ vb, 0
	case OpNOR:
		return ^(va | vb), 0
	case OpMOV:
		return vb, 0
	}
	panic("unsupported alu operation")
}

func shift(vb uint32, op SOp, shamt uint32) uint32 {
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
	panic("unsupported shift operation")
}

func (m *VM) condPassed(ins uint32) bool {
	csrE := Bit(m.csr, 26)
	csrL := Bit(m.csr, 27)
	csrG := ^csrL & ^csrE

	insE := Bit(ins, 26)
	insL := Bit(ins, 27)
	insG := Bit(ins, 28)

	return (csrG&insG | csrL&insL | csrE&insE | insG&insL&insE) == 1
}

func (m *VM) setEqualFlag(eq bool) {
	m.csr = SetBool(m.csr, 26, eq)
}

func (m *VM) setLessFlag(lt bool) {
	m.csr = SetBool(m.csr, 27, lt)
}

func (m *VM) ins(code []byte) uint32 {
	// Instructions access is word aligned only.
	return read32(code, m.regs[IP])
}

func read32(data []byte, a uint32) uint32 {
	a = a << 2
	return binary.BigEndian.Uint32(data[a : a+4])
}

func write32(data []byte, a uint32, v uint32) {
	a = a << 2
	binary.BigEndian.PutUint32(data[a:a+4], v)
}
