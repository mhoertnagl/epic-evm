package vm

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
	case "r16":
		return m.regs[16]
	case "r17":
		return m.regs[17]
	case "r18":
		return m.regs[18]
	case "r19":
		return m.regs[19]
	case "r20":
		return m.regs[20]
	case "r21":
		return m.regs[21]
	case "r22":
		return m.regs[22]
	case "r23":
		return m.regs[23]
	case "r24":
		return m.regs[24]
	case "r25":
		return m.regs[25]
	case "r26":
		return m.regs[26]
	case "r27":
		return m.regs[27]
	case "r28":
		return m.regs[28]
	case "r29":
		return m.regs[29]
	case "r30":
		return m.regs[30]
	case "r31":
		return m.regs[31]
	case "sp":
		return m.regs[28]
	case "fp":
		return m.regs[29]
	case "rp":
		return m.regs[30]
	case "ip":
		return m.regs[31]
	case "csr":
		return m.csr
	}
	panic("undefined register id")
}
