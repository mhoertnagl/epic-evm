package vm

type Op = uint32

const (
	OpDPR Op = 0x0
	OpD12    = 0x1
	OpMEM    = 0x2
	OpM12    = 0x3
	OpI16    = 0x4
	OpNA5    = 0x5
	OpNA6    = 0x6
	OpBRA    = 0x7
)

const (
	VbREG  = 0x0
	VbS12  = 0x1
	VbU12  = 0x2
	VbSL16 = 0x3
	VbUL16 = 0x4
	VbH16  = 0x5
	VbS25  = 0x6
)

type AOp = uint32

const (
	OpADD AOp = 0x0
	OpSUB     = 0x1
	OpMUL     = 0x2
	OpDIV     = 0x3
	OpAND     = 0x4
	OpOOR     = 0x5
	OpXOR     = 0x6
	OpNOR     = 0x7
	OpXX8     = 0x8
	OpXX9     = 0x9
	OpXXA     = 0xA
	OpXXB     = 0xB
	OpCPS     = 0xC
	OpCPU     = 0xD
	OpXXE     = 0xE
	OpMOV     = 0xF
)

type SOp = uint32

const (
	OpSLL SOp = 0x0
	OpROL     = 0x1
	OpSRL     = 0x2
	OpSRA     = 0x3
)

const (
	RP uint32 = 0x1E
	IP uint32 = 0x1F
)

const (
	CndNone     = 0x0
	CndSigned   = 0x1
	CndUnsigned = 0x2
)

const (
	CmbRegReg = 0x0
	CmbRegImm = 0x1
	CmbImmReg = 0x2
)

func op(ins uint32) Op {
	return bitsAt(ins, 3, 29)
}

func aluop(ins uint32) AOp {
	return bitsAt(ins, 4, 0)
}

func rd(ins uint32) uint32 {
	return bitsAt(ins, 4, 20)
}

func ra(ins uint32) uint32 {
	return bitsAt(ins, 4, 16)
}

func rb(ins uint32) uint32 {
	return bitsAt(ins, 4, 12)
}

func sop(ins uint32) SOp {
	return bitsAt(ins, 2, 9)
}

func shamt(ins uint32) SOp {
	return bitsAt(ins, 5, 4)
}

func imm12(ins uint32) uint32 {
	return bitsAt(ins, 12, 4)
}

func imm16(ins uint32) uint32 {
	return bitsAt(ins, 16, 4)
}

func imm25(ins uint32) uint32 {
	return bitsAt(ins, 25, 0)
}

func isLoad(ins uint32) bool {
	return bitSet(ins, 0)
}

func isHigh(ins uint32) bool {
	return bitSet(ins, 20)
}

func isLink(ins uint32) bool {
	return bitSet(ins, 25)
}

func bitsAt(val uint32, p uint8, s uint8) uint32 {
	return (val >> s) & ((1 << p) - 1)
}

func bitSet(val uint32, s uint8) bool {
	return ((val >> s) & 1) == 1
}
