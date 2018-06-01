/**
 * A Device is a logical component connected to other devices via a bus.
 * @see Bus
 */
class Device {

    bus = null

    constructor() {

    }

    set bus(b) {
        this.bus = b
    }

    read(addr) {
        throw new Error("Not implemented!")
    }

    write(addr, val) {
        throw new Error("Not implemented!")
    }
}