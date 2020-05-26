package vm

import (
	"encoding/binary"
)

// TODO: Do I need a mvs - move signed
// TODO: Cannot move 32 bit value with two moves!!!

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
	len := uint32(len(m.mem) >> 2)
	// fmt.Printf("Code length: %d\n", len)
	for m.regs[IP] < len {
		ins := m.ins(m.mem)
		if m.condPassed(ins) {
			switch op(ins) {
			case OpDPR:
				m.execDPR(ins)
				if doesNotWriteIP(ins) {
					m.regs[IP]++
				}
			case OpMEM:
				m.execMEM(ins)
				if doesNotWriteIP(ins) {
					m.regs[IP]++
				}
			case OpNA2:
				panic("unsopported NA2 operation")
			case OpNA3:
				panic("unsopported NA3 operation")
			case OpNA4:
				panic("unsopported NA4 operation")
			case OpNA5:
				panic("unsopported NA5 operation")
			case OpNA6:
				panic("unsopported NA6 operation")
			case OpBRA:
				m.execBRA(ins)
			}
		} else {
			m.regs[IP]++
		}
	}
}

func (m *VM) Reg(id string) uint32 {
	switch id {
	case "r0":
		return m.regs[0]
	case "r1":
		return m.regs[1]
	case "r2":
		return m.regs[2]
	case "r3":
		return m.regs[3]
	case "r4":
		return m.regs[4]
	case "r5":
		return m.regs[5]
	case "r6":
		return m.regs[6]
	case "r7":
		return m.regs[7]
	case "r8":
		return m.regs[8]
	case "r9":
		return m.regs[9]
	case "r10":
		return m.regs[10]
	case "r11":
		return m.regs[11]
	case "r12":
		return m.regs[12]
	case "r13":
		return m.regs[13]
	case "r14":
		return m.regs[14]
	case "r15":
		return m.regs[15]
	case "sp":
		return m.regs[13]
	case "rp":
		return m.regs[14]
	case "ip":
		return m.regs[15]
	}
	panic("undefined register id")
}

func (m *VM) execDPR(ins uint32) {
	rd := rd(ins)
	ra := ra(ins)
	aop := aluop(ins)

	va := m.regs[ra]
	vb := m.computeVB(ins)
	rs, c := Alu(aop, va, vb)

	switch aop {
	case OpCMP:
		m.setEqualFlag(rs == 0)
		m.setLessFlag(Bit(rs, 31) == 1)
	case OpCPU:
		m.setEqualFlag(rs == 0)
		m.setLessFlag(c == 0)
	default:
		m.regs[rd] = uint32(rs)
	}
}

func (m *VM) execMEM(ins uint32) {
	rd := rd(ins)
	ra := ra(ins)

	va := m.regs[ra]
	vb := m.computeVB(ins)
	adr, _ := Alu(OpADD, va, vb)

	if isLoad(ins) {
		m.regs[rd] = read32(m.mem, adr)
	} else {
		write32(m.mem, adr, m.regs[rd])
	}
}

func (m *VM) execBRA(ins uint32) {
	if isLink(ins) {
		m.regs[RP] = m.regs[IP] + 1
	}
	m.regs[IP] += Sext(imm25(ins), 25)
}

func (m *VM) computeVB(ins uint32) uint32 {
	switch iop(ins) {
	case OpREG:
		return Shift(m.regs[rb(ins)], sop(ins), shamt(ins))
	case OpI12:
		// TODO: Do not sign extend?
		return Sext(imm12(ins), 12)
	case OpL16:
		// TODO: Do not sign extend?
		return Sext(imm16(ins), 16)
	case OpU16:
		return Shift(imm16(ins), OpSLL, 16)
	}
	panic("unsupported RB operation")
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

func doesNotWriteIP(ins uint32) bool {
	aop := aluop(ins)
	return aop == OpCMP || aop == OpCPU || rd(ins) != IP
}

func read32(data []byte, a uint32) uint32 {
	a = a << 2
	return binary.BigEndian.Uint32(data[a : a+4])
}

func write32(data []byte, a uint32, v uint32) {
	a = a << 2
	binary.BigEndian.PutUint32(data[a:a+4], v)
}
