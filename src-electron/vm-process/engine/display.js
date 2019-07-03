import Device from './device'

export default class Display extends Device {

  constructor () {
    super()
    this.init()
  }

  get cols () {
    return 80
  }

  get rows () {
    return 30
  }

  init () {
    this.mem = []
    for (let r = 0; r < this.rows; r++) {
      let row = []
      for (let c = 0; c < this.cols; c++) {
        row.push(0)
      }
      this.mem.push(row)
    }
  }

  get memory () {
    return this.mem
  }

  read (addr) {
    let r = Math.floor(addr / this.cols)
    let c = addr % this.cols
    return this.mem[r][c]
  }

  write (addr, val) {
    let r = Math.floor(addr / this.cols)
    let c = addr % this.cols
    this.mem[r][c] = val
  }
}
