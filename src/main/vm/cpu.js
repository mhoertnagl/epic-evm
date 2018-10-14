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
    this.regs = Array(16).fill(0)
  }

  step () {
    let ins = read(this.regs[15])
    let op = ins >> 28
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

  // read (addr) {
  //     throw new Error("Not implemented!")
  // }

  // write (addr, val) {
  //     throw new Error("Not implemented!")
  // }

  extract (ins, start, len) {
    return (ins >> start) & ((1 << len) - 1)
  }
}