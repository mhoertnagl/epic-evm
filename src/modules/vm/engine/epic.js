import Bus from './bus'
import { RangeBusLocation, VoidLocation } from './location'
import Display from './display'
import Memory from './memory'
import Cpu from './cpu'

export default class Epic {
//module.exports = class Epic {
  
  constructor () {
    this.bus = new Bus()
    this.display = new Display()
    // 16 kB of memory.
    this.mem = new Memory(0x00010000)
    this.cpu = new Cpu()
    
    const memoryLoc = new RangeBusLocation(0x00000000, 0x0000FFFF)
    const displayLoc = new RangeBusLocation(0xFFF00000, 0xFFF0FFFF)
    const cpuLoc = new VoidLocation()
    
    this.bus.connect(memoryLoc, this.mem)
    this.bus.connect(displayLoc, this.display)
    this.bus.connect(cpuLoc, this.cpu)
    
    console.log('VM instance initialized.')
  }

  get displayMemory () {
    return this.display.memory
  }
  
  reset (data) {
    this.mem.init(data)
    this.cpu.reset()
  }
  
  step () {
    this.cpu.step()
  }
}
