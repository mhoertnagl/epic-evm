package vm

import (
	"encoding/binary"
	"math/bits"
)

type VM struct {
	cir  uint32
	csr  uint32
	regs [32]uint32
	mem  []byte
}

func NewVM(mem []byte) *VM {
	return &VM{mem: mem}
}

func (m *VM) Run() {
	len := uint32(len(m.mem) >> 2)
	for m.regs[IP] < len {
		ins := m.ins(m.mem)
		if m.condPassed(ins) {
			m.run(ins)
		} else {
			m.regs[IP]++
		}
	}
}

func (m *VM) run(ins uint32) {
	switch op(ins) {
	case OpDPR:
		m.runDPR(ins)
	case OpDI8:
		m.runDI8(ins)
	case OpMEM:
		m.runMEM(ins)
	case OpMI8:
		m.runMI8(ins)
	case OpI16:
		m.runD16(ins)
	case OpBRA:
		m.runBRA(ins)
	}
}

func (m *VM) runDPR(ins uint32) {
	rd := rd(ins)
	ra := ra(ins)
	rb := rb(ins)
	sop := sop(ins)
	shamt := shamt(ins)
	aop := aluop(ins)

	va := m.regs[ra]
	vb := m.regs[rb]
	vb = shift(vb, sop, shamt)

	m.writeRegs(rd, va, vb, aop)
}

func (m *VM) runDI8(ins uint32) {
	rd := rd(ins)
	ra := ra(ins)
	aop := aluop(ins)
	shamt := shamt(ins)

	va := m.regs[ra]
	vb := imm8(ins)

	if isSignedAluOp(aop) {
		vb = Sext(imm8(ins), 8)
	}

	vb = shift(vb, OpROL, shamt&0x1E)

	m.writeRegs(rd, va, vb, aop)
}

func (m *VM) runMEM(ins uint32) {
	rb := rb(ins)
	sop := sop(ins)
	shamt := shamt(ins)

	vb := m.regs[rb]
	vb = shift(vb, sop, shamt)

	m.accessMem(ins, vb)
}

func (m *VM) runMI8(ins uint32) {
	shamt := shamt(ins)
	vb := Sext(imm8(ins), 8)
	vb = shift(vb, OpROL, shamt&0x1E)
	m.accessMem(ins, vb)
}

func (m *VM) runD16(ins uint32) {
	rd := rd(ins)
	aop := aluop(ins)

	va := m.regs[rd]
	vb := imm16(ins)

	if isSignedAluOp(aop) {
		vb = Sext(vb, 16)
	}

	if isHigh(ins) {
		vb = shift(vb, OpSLL, 16)
	}

	m.writeRegs(rd, va, vb, aop)
}

func (m *VM) runBRA(ins uint32) {
	if isLink(ins) {
		m.regs[RP] = m.regs[IP] + 1
	}

	of := Sext(imm25(ins), 25)
	ip, _ := bits.Add32(m.regs[IP], of, 0)
	m.regs[IP] = ip
}

func (m *VM) writeRegs(rd uint32, va uint32, vb uint32, aop uint32) {
	vr, c := alu(aop, va, vb)
	ov := Bit(vr, 31) == 1
	cr := c == 1

	switch aop {
	case OpCPS:
		m.setEqualFlag(vr == 0)
		m.setLessFlag(ov)
		m.setVCFlags(cr, ov)
		m.regs[IP]++
	case OpCPU:
		m.setEqualFlag(vr == 0)
		m.setLessFlag(cr == false)
		m.setVCFlags(cr, ov)
		m.regs[IP]++
	default:
		m.regs[rd] = vr
		if rd != IP {
			m.regs[IP]++
		}
	}
}

func (m *VM) accessMem(ins uint32, vb uint32) {
	rd := rd(ins)
	ra := ra(ins)

	vd := m.regs[rd]
	va := m.regs[ra]

	ad, _ := alu(OpADD, va, vb)

	if isLoad(ins) {
		m.regs[rd] = read32(m.mem, ad)
	} else {
		write32(m.mem, ad, vd)
	}

	if rd != IP {
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
		hi, lo := bits.Mul32(va, vb)
		return lo, hi & 1
	case OpDIV:
		quo, _ := bits.Div32(0, va, vb)
		return quo, 0
	case OpAND:
		return va & vb, 0
	case OpOOR:
		return va | vb, 0
	case OpXOR:
		return va ^ vb, 0
	case OpNOR:
		return ^(va | vb), 0
	case OpCPS:
		return bits.Sub32(va, vb, 0)
	case OpCPU:
		return bits.Sub32(va, vb, 0)
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

func isSignedAluOp(aop uint32) bool {
	return aop == OpMUL || aop == OpCPS
}

func (m *VM) setEqualFlag(eq bool) {
	m.csr = SetBool(m.csr, 26, eq)
}

func (m *VM) setLessFlag(lt bool) {
	m.csr = SetBool(m.csr, 27, lt)
}

func (m *VM) setVCFlags(cr bool, ov bool) {
	m.csr = SetBool(m.csr, 28, cr)
	m.csr = SetBool(m.csr, 29, ov)
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
