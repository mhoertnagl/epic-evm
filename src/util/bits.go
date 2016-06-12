package util

// Bits extracts a consecutive portion of length and offest from the LSB of an
// unsigned integer.
func Bits(instr uint32, offset uint8, length uint8) uint32 {
  return (instr >> offset) & ((1 << length) - 1)
}

// Set returns true if the bit at the specified offset is set.
func Set(instr uint32, offset uint8) bool {
  return 1 == (instr & (1 << offset))
}
