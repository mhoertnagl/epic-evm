/**
 * A BusLocation defines input and output ports of a device.
 */
export class BusLocation {

  /**
   * Returns true iff the address {addr} matches any of the location's
   * accepted addresses.
   *
   * @param {Number} addr - The address.
   * @return {Boolean} True iff the address matches.
   */
  matches (addr) {
    throw new Error("Not implemented!")
  }

  /**
   * Maps the bus address {addr} to the correct local device address.
   *
   * @param {Number} addr - The address.
   */
  map (addr) {
    throw new Error("Not implemented!")
  }
}

/**
 * A RangeBusLocation assignes a contiguous range of addresses to a device.
 *
 * @see BusLocation
 */
export class RangeBusLocation extends BusLocation {
  
  /**
   * Constructs a new RangeBusLocation that assignes a contiguous range of
   * addresses to a device. The valid addresses start from {min} until {max}
   * inclusive.
   * The mapped device address will start at address 0.
   *
   * @param {Number} min - The minimum address.
   * @param {Number} max - The maximum address inclusive.
   */
  constructor (min, max) {
    super()
    this.min = min
    this.max = max
  }

  matches (addr) {
    return this.min <= addr && addr <= this.max
  }

  map (addr) {
    return addr - this.min
  }
}

/**
 * A VoidLocation never matches.
 *
 * @see BusLocation
 */
export class VoidLocation extends BusLocation {
  
  matches (addr) {
    return false
  }

  map (addr) {
    return addr
  }  
}
