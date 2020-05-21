package vm

import "math/bits"

func ShoudSignExtend(ins uint32) bool {
	switch aluop(ins) {
	case OpADD:
		return true
	case OpSUB:
		return true
	case OpMUL:
		return true
	case OpDIV:
		return true
	case OpCMP:
		return true
	default:
		return false
	}
}

// TODO: Not necessary, assembler will set S flag.
func ShouldSetCond(ins uint32) bool {
	switch aluop(ins) {
	case OpCMP:
		return true
	case OpCPU:
		return true
	default:
		return false
	}
}

func Shift(vb uint32, op SOp, shamt uint32) uint32 {
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

func Alu(op uint32, va uint32, vb uint32) (uint32, uint32) {
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
	case OpXX8:
		panic("unsupported alu operation")
	case OpXX9:
		panic("unsupported alu operation")
	case OpXXA:
		panic("unsupported alu operation")
	case OpXXB:
		panic("unsupported alu operation")
	case OpCMP:
		return bits.Sub32(va, vb, 0)
	case OpCPU:
		return bits.Sub32(va, vb, 0)
	case OpXXE:
		panic("unsupported alu operation")
	case OpMOV:
		return vb, 0
	}
	panic("unsupported alu operation")
}
