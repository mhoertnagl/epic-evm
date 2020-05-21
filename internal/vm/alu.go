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
	case OpTST:
		return true
	default:
		return isSetCond(ins)
	}
}

func ShouldWriteBack(ins uint32) bool {
	switch aluop(ins) {
	case OpCMP:
		return false
	case OpCPU:
		return false
	case OpTST:
		return false
	default:
		return true
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

// TODO: Use bits.Add
// TODO: Unsigned operations obsolete (except for cpu)?
func Alu(op uint32, va uint32, vb uint32) uint64 {
	switch op {
	case OpADD:
		return uint64(int64(Sext64(va)) + int64(Sext64(vb)))
	case OpSUB:
		return uint64(int64(Sext64(va)) - int64(Sext64(vb)))
	case OpMUL:
		return uint64(int64(Sext64(va)) * int64(Sext64(vb)))
	case OpDIV:
		return uint64(int64(Sext64(va)) / int64(Sext64(vb)))
	case OpAND:
		return uint64(va & vb)
	case OpOOR:
		return uint64(va | vb)
	case OpXOR:
		return uint64(va ^ vb)
	case OpNOR:
		return uint64(^(va | vb))
	case OpADU:
		return uint64(va) + uint64(vb)
	case OpSBU:
		return uint64(va) - uint64(vb)
	case OpMLU:
		return uint64(va) * uint64(vb)
	case OpDVU:
		return uint64(va) / uint64(vb)
	case OpCMP:
		return uint64(int64(Sext64(va)) - int64(Sext64(vb)))
	case OpCPU:
		return uint64(va) - uint64(vb)
	case OpTST:
		return uint64(va & vb)
	case OpMOV:
		return uint64(vb)
	}
	panic("unsupported alu operation")
}
