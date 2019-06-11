import Device from './device'
import { BusLocation } from './location'

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
        dev.write(loc.map(addr), val)
        return
      }
    }
    throw new Error(`Cannot write value [${val}] to address [${addr}].`)
  }
}

/**
 * A Connection links a device with an address range.
 */
class Connection {

  constructor (loc, dev) {
    this.loc = loc
    this.dev = dev
  }
}
