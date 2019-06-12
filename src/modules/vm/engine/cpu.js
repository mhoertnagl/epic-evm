import Device from './device'

export default class Cpu extends Device {
  /**
   * Initializes the CPU.
   */
  constructor () {
    super()

    // Initialize the registers r0 - r14. On reset the contents of these
    // registers is undefined and can take on any values.
    // The register r15 is the IP and will start execution at address 0.

    /** 32bit general purpose registers and the IP. */
    this.regs = new Uint32Array(16)

    /** CPU state register. */
    this.csr = 0x1D000001
  }

  get gp_regs () {
    return this.regs
  }

  get cs_reg () {
    return this.csr
  }

  ip () {
    return this.regs[15]
  }

  step () {
    let ins = this.bus.read(this.ip())
    //let op = ins >> 28
    let op = this.extract(ins, 3, 28)
    switch (op) {
      case 0:
        // DAT-REG
        break
      case 1:
        // DAT-I16
        break
      case 2:
        // MEM
        break
      case 3:
        // COP
        break
      case 4:
        // SYS
        break
      case 5:
        throw new Error(`Reserved OP code [${op}].`)
      case 6:
        throw new Error(`Reserved OP code [${op}].`)
      case 7:
        // BRA
        break
      default:
        throw new Error(`Unsupported OP code [${op}].`)
    }
  }

  dataInstr (ins) {
    let cnd = this.extract(ins, 3, 26)
    let set = this.extract(ins, 1, 25)
    let wb = true
    let imm = this.extract(ins, 1, 24)
    let rd = this.extract(ins, 4, 20)
    let ra = this.extract(ins, 4, 16)
    let rb = this.extract(ins, 4, 12)
    let sop = this.extract(ins, 2, 9)
    let smt = this.extract(ins, 5, 4)
    let aop = this.extract(ins, 4, 0)

    let vd = 0
    let va = this.regs[ra]
    // TODO: Hier muss immediate wert gelesen werden.
    let ub = imm ? this.extract(ins, 12, 4) : this.regs[rb]
    let vb = this.shift(ub, sop, smt)

    // Unsigned / Signed.

    // TODO: Test conditions.
    // if () {

    // }

    switch (aop) {
      case 0x0:
        vd = va + vb
        break
      case 0x1:
        vd = va - vb
        break
      case 0x2:
        vd = (va * vb) & 0xFFFFFFFF
        break
      case 0x3:
        vd = va / vb
        break
      case 0x4:
        vd = va & vb
        break
      case 0x5:
        vd = va | vb
        break
      case 0x6:
        vd = va ^ vb
        break
      case 0x7:
        vd = ~(va | vb)
        break
      // case 0x8:
      //   break
      // case 0x9:
      //   break
      // case 0xA:
      //   break
      // case 0xB:
      //   break
      case 0xC:
        // TODO: Signed.
        vd = va - vb
        set = true
        wb = false
        break
      case 0xD:
        // TODO: Unsigned.
        vd = va - vb
        set = true
        wb = false
        break
      case 0xE:
        vd = va & vb
        set = true
        wb = false
        break
      case 0xF:
        vd = vb
        break
      default:
        throw new Error(`Unsupported data instruction code [${aop}].`)
    }

    if (set) {
      this.setCond(vd)
    }

    if (wb) {
      this.regs[rd] = vd
    }
  }

  shift (val, sop, smt) {
    switch (sop) {
      case 0:
        return val << smt
      case 1:
        return val >> smt
      case 2:
        return val >>> smt
      case 3:
        return (val << smt) | (val >> (32 - smt)) // rol
      default:
        break
    }
  }

  setCond (vd) {
    let eq = vd === 0
    let lt = vd < 0
    let gt = vd > 0
    // TODO: Set condition flags.
  }

  extract (ins, len, start) {
    return (ins >> start) & ((1 << len) - 1)
  }

  // TODO: Short instruction to set a bit.
  set (val, pos, bit) {
    if (bit) {
      return val ^ (1 << pos)
    }
    return val & ~(1 << pos)
  }
}
