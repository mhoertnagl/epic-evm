package vm

func Alu(op uint32, va uint32, vb uint32) uint32 {
	switch op {
	case OpADD:
		return uint32(int32(va) + int32(vb))
	case OpSUB:
		return uint32(int32(va) - int32(vb))
	case OpMUL:
		return uint32(int32(va) * int32(vb))
	case OpDIV:
		return uint32(int32(va) / int32(vb))
	case OpAND:
		return va & vb
	case OpOOR:
		return va | vb
	case OpXOR:
		return va ^ vb
	case OpNOR:
		return ^(va | vb)
	case OpADU:
		return va + vb
	case OpSBU:
		return va - vb
	case OpMLU:
		return va * vb
	case OpDVU:
		return va / vb
	case OpCMP:
	case OpCPU:
	case OpTST:
	case OpMOV:
		return vb
	}
	return 0
}
