import Device from './device'

export default class Display extends Device {
  rows = 30
  cols = 80
  mem = []

  constructor () {
    super()
    this.initDisplay()
  }

  initDisplay () {
    this.mem = []
    for (let r = 0; r < this.rows; r++) {
      let row = []
      for (let c = 0; c < this.cols; c++) {
        row.push(String.fromCharCode(66))
      }
      this.mem.push(row)
    }
  }

  get memory () {
    return this.mem
  }

  write (addr, val) {
    throw new Error('Not implemented!')
  }
}
