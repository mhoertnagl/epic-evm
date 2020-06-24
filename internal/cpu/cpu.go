package cpu

import (
	"encoding/binary"
	"math/bits"
)

// stw
// ldw
// bra
// brl

// console UI
// a bus
// video ram
// interrupts
// kbd input

type Cpu struct {
	cir  uint32
	csr  uint32
	regs [32]uint32
	mem  []byte
	len  uint32
}

func NewCpu(mem []byte) *Cpu {
	return &Cpu{
		mem: mem,
		len: uint32(len(mem) >> 2),
	}
}

func (m *Cpu) Running() bool {
	return m.regs[IP] < m.len
}

func (m *Cpu) Run() {
	for m.Running() {
		m.Step()
	}
}

func (m *Cpu) Step() {
	ins := m.ins(m.mem)
	if m.condPassed(ins) {
		m.run(ins)
	} else {
		m.regs[IP]++
	}
}

func (m *Cpu) run(ins uint32) {
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

func (m *Cpu) runDPR(ins uint32) {
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

func (m *Cpu) runDI8(ins uint32) {
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

func (m *Cpu) runMEM(ins uint32) {
	rb := rb(ins)
	sop := sop(ins)
	shamt := shamt(ins)

	vb := m.regs[rb]
	vb = shift(vb, sop, shamt)

	m.accessMem(ins, vb)
}

func (m *Cpu) runMI8(ins uint32) {
	shamt := shamt(ins)
	vb := Sext(imm8(ins), 8)
	vb = shift(vb, OpROL, shamt&0x1E)
	m.accessMem(ins, vb)
}

func (m *Cpu) runD16(ins uint32) {
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

func (m *Cpu) runBRA(ins uint32) {
	if isLink(ins) {
		m.regs[RP] = m.regs[IP] + 1
	}

	of := Sext(imm25(ins), 25)
	ip, _ := bits.Add32(m.regs[IP], of, 0)
	m.regs[IP] = ip
}

func (m *Cpu) writeRegs(rd uint32, va uint32, vb uint32, aop uint32) {
	// vr, c := alu(aop, va, vb)
	vr, _ := alu(aop, va, vb)
	switch aop {
	case OpCPS:
		m.setEqualFlag(vr == 0)
		m.setLessFlag(int32(va) < int32(vb))
		// m.setLessFlag(c == 1)
		// m.setVCFlags(c == 1, Bit(vr, 31) == 1)
		m.regs[IP]++
	case OpCPU:
		m.setEqualFlag(vr == 0)
		m.setLessFlag(va < vb)
		// m.setLessFlag(c != Bit(vr, 31))
		// m.setVCFlags(c == 1, Bit(vr, 31) == 1)
		m.regs[IP]++
	default:
		m.regs[rd] = vr
		if rd != IP {
			m.regs[IP]++
		}
	}
}

func (m *Cpu) accessMem(ins uint32, vb uint32) {
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
		return uint32(int32(va) * int32(vb)), 0
	case OpDIV:
		return uint32(int32(va) / int32(vb)), 0
	case OpAND:
		return va & vb, 0
	case OpOOR:
		return va | vb, 0
	case OpXOR:
		return va ^ vb, 0
	case OpNOR:
		return ^(va | vb), 0
	case OpCPS:
		// v, _ := bits.Sub64(uint64(int32(va)), uint64(int32(vb)), 0)
		// return uint32(v), uint32(v & 0x100000000 >> 32)
		return bits.Sub32(va, vb, 0)
	case OpCPU:
		// v, _ := bits.Sub64(uint64(va), uint64(vb), 0)
		// return uint32(v), uint32(v & 0x100000000 >> 32)
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

func (m *Cpu) condPassed(ins uint32) bool {
	csrE := Bit(m.csr, 26)
	csrL := Bit(m.csr, 27)
	csrG := ^csrL & ^csrE

	insE := Bit(ins, 26)
	insL := Bit(ins, 27)
	insG := Bit(ins, 28)

	return (csrG&insG | csrL&insL | csrE&insE | insG&insL&insE) == 1
}

func isSignedAluOp(aop uint32) bool {
	return aop == OpMUL || aop == OpDIV || aop == OpCPS
}

func (m *Cpu) setEqualFlag(eq bool) {
	m.csr = SetBool(m.csr, 26, eq)
}

func (m *Cpu) setLessFlag(lt bool) {
	m.csr = SetBool(m.csr, 27, lt)
}

// func (m *Cpu) setVCFlags(cr bool, ov bool) {
// 	m.csr = SetBool(m.csr, 28, cr)
// 	m.csr = SetBool(m.csr, 29, ov)
// }

func (m *Cpu) ins(code []byte) uint32 {
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
