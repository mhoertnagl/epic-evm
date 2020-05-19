package vm

func Sext(word uint32, len uint32) uint32 {
	d := 32 - len
	return uint32((int32(word) << d) >> d)
}

func Sext64(word uint32) uint64 {
	return uint64((int64(word) << 32) >> 32)
}

func Bit(word uint64, pos uint32) uint32 {
	return uint32((word << (31 - pos)) >> 31)
}

func SetBool(word uint32, pos uint32, bit bool) uint32 {
	if bit {
		return Set(word, pos, 1)
	}
	return Set(word, pos, 0)
}

func Set(word uint32, pos uint32, bit uint32) uint32 {
	return (word & ^(1 << pos)) | (bit & 1 << pos)
}
