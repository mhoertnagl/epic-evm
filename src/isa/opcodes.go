package isa

// Set of valid registers.
const (
  R00 = 0x0
  R01 = 0x1
  R02 = 0x2
  R03 = 0x3
  R04 = 0x4
  R05 = 0x5
  R06 = 0x6
  R07 = 0x7
  R08 = 0x8
  R09 = 0x9
  R10 = 0xA
  R11 = 0xB
  R12 = 0xC
  R13 = 0xD
  R14 = 0xE
  R15 = 0xF
)

// Set of valid alu operations.
const (
  AluOpAdd = 0x0
  AluOpSub = 0x1
  AluOpMul = 0x2
  AluOpDiv = 0x3

  AluOpAnd = 0x4
  AluOpOor = 0x5
  AluOpNor = 0x6
  AluOpXor = 0x7

  AluOpSll = 0x8
  AluOpSrl = 0x9
  AluOpSra = 0xA
  AluOpRor = 0xB
)
