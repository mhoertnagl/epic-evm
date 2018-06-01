class Cpu extends Device {

    constructor() {
        this.regs = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }

    step() {
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
                break
            case 6:
                throw new Error(`Reserved OP code [${op}].`)
                break
            case 7:
                // BRA
                break
            default:
                throw new Error(`Unsupported OP code [${op}].`)
                break
        }
    }

    read(addr) {
        throw new Error("Not implemented!")
    }

    write(addr, val) {
        throw new Error("Not implemented!")
    }

    extract(ins, start, len) {
        return (ins >> start) & ((1 << len) - 1)
    }
}