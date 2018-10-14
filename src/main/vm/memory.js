import Device from './device'
/**
 * https://nodejs.org/api/buffer.html#buffer_buf_readuint32be_offset
 * https://nodejs.org/api/fs.html#fs_fs_readfilesync_path_options
 * https://nodejs.org/api/path.html
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Typed_arrays
 */
export default class Memory extends Device {
  /**
   * Creates an new block of memory of length {size} words (4 bytes). Each
   * memory cell will be initialized as 0.
   *
   * @param {Number} size - The length of the memory in words (4 bytes).
   */
  constructor (size) {
    this.size = size
    this.mem = Buffer.alloc(size << 2, 0)
  }

  /**
   * Reads the value at address {addr} on the device.
   *
   * @param {Number} addr - The address.
   * @return {Number} The value at address {addr}.
   */
  read (addr) {
    this.validateAddress(addr)
    return this.mem.readUInt32BE(addr)
  }

  /**
   * Writes the value {val} to address {addr} on the device.
   *
   * @param {Number} addr - The address.
   * @param {Number} val  - The value to write.
   */
  write (addr, val) {
    this.validateAddress(addr)
    this.mem.writeInt32BE(val, addr)
  }

  /**
   * Validates the given address {addr} if it is within the allowed memory
   * range. Throws an error if not.
   *
   * @param {Number} addr - The address.
   */
  validateAddress (addr) {
    if (addr < 0) {
      throw new Error(`Cannot address memory below 0.`)
    }
    if (addr > this.size - 1) {
      throw new Error(`Cannot address memory above [${this.size}].`)
    }
  }
}