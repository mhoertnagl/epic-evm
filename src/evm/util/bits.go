package bits

func bits(instr uint32, offset uint8, length uint8) uint32 {
	return (instr >> offset) & ((1 << length) - 1)
}

func set(instr uint32, offset uint8) bool {
	return 1 == (instr & (1 << offset))
}
