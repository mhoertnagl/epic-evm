import Device from './device'

/**
 * The Bus connects the various parts of the virtual machine.
 */
export default class Bus {

  constructor () {
    this.connections = [];
  }

  /**
   * Connects a {device} to the bus at location {loc}.
   *
   * @param {BusLocation} loc
   * @param {Device} dev
   */
  connect (loc, dev) {
    dev.connect(this)
    this.connections.push(new Connection(loc, dev))
  }

  /**
   * Reads the data stored at address {addr}.
   *
   * @param {Number} addr - The address.
   * @return {Number} The value at address {addr}.
   */
  read (addr) {
    for (var con of this.connections) {
      const loc = con.loc
      const dev = con.dev
      if (loc.matches(addr)) {
        return dev.read(loc.map(addr))
      }
    }
    throw new Error(`Cannot read at address [${addr}].`)
  }

  /**
   * Writes the value {value} to address {addr}.
   *
   * @param {Number} addr - The address.
   * @param {Number} val  - The value.
   */
  write (addr, val) {
    for (let con of this.connections) {
      const loc = con.loc
      const dev = con.dev
      if (loc.matches(addr)) {
        return dev.write(loc.map(addr), val)
      }
    }
    throw new Error(`Cannot write value [${val}] to address [${addr}].`)
  }
}

/**
 * A Connection links a device with a address range.
 */
class Connection {

  constructor (loc, dev) {
    this.loc = loc
    this.dev = dev
  }

  get loc () {
    return this.loc
  }

  get dev () {
    return this.dev
  }
}

/**
 * A BusLocation defines input and output ports of a device.
 */
class BusLocation {

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
class RangeBusLocation extends BusLocation {
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