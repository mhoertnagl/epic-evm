import Display from './display'

class Epic {
  constructor () {
    this.display = new Display()
  }

  get displayMemory () {
    return this.display.memory
  }
}

const VM = new Epic()

export default VM
