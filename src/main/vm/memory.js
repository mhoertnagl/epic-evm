/**
 * https://nodejs.org/api/buffer.html#buffer_buf_readuint32be_offset
 * https://nodejs.org/api/fs.html#fs_fs_readfilesync_path_options
 * https://nodejs.org/api/path.html
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Typed_arrays
 */
class Memory extends Device {

    /**
     * Creates an new block of memory of length {len} words (4 bytes). Each
     * memory cell will be initialized as 0.
     * 
     * @param {Number} len - The length of the memory in words (4 bytes).
     */
    constructor(len) {
        this.mem = Array(len).fill(0) // Buffer
    }

    /**
     * Reads the value at address {addr} on the device.
     * 
     * @param {Number} addr - The address.
     * @return {Number} The value at address {addr}.
     */
    read(addr) {
        throw new Error("Not implemented!")
    }

    /**
     * Writes the value {val} to address {addr} on the device.
     * 
     * @param {Number} addr - The address.
     * @param {Number} val  - The value to write.
     */
    write(addr, val) {
        throw new Error("Not implemented!")
    }
}