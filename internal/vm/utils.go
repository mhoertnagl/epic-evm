package vm

func Sext(word uint32, len uint32) uint32 {
	d := 32 - len
	return uint32(int32(word<<d) >> d)
}
