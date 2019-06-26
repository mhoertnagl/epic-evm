/**
 * A Device is a logical component connected to other devices via a bus.
 *
 * @see Bus
 */
export default class Device {
  
  constructor () {
    this.bus = null
  }

  connect (b) {
    this.bus = b
  }

  /**
   * Reads the value at address {addr} on the device.
   *
   * @param {Number} addr - The address.
   * @return {Number} The value at address {addr}.
   */
  read (addr) {
    throw new Error('Not implemented!')
  }

  /**
   * Writes the value {val} to address {addr} on the device.
   *
   * @param {*} addr - The address.
   * @param {*} val  - The value to write.
   */
  write (addr, val) {
    throw new Error('Not implemented!')
  }
}
