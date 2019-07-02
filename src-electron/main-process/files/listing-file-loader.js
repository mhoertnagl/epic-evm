
import FileLoader from './file-loader'

export default class ListingFileLoader extends FileLoader {
  
  constructor () {
    super()
  }
    
  processLine (file, line) {
    let res = this.parseListingLine(line)
    file.lines.push(res.line);
    if (res.isIns) {
      file.code.push(res.ins)
    }    
  }
  
  parseListingLine (line) {
    let res = {
        isIns: false,
        line: line,
        addr: 0,
        ins: 0,
    }
    let insRegex = /^(0x[0-9a-fA-F]{8})\s+(0x[0-9a-fA-F]{8})\s+(.*)$/
    let match = line.match(insRegex)

    if (match !== null) {
      res.isIns = true
      // res.addr = parseInt()
      // res.ins = parseInt()
    }
    return res
  }
}
