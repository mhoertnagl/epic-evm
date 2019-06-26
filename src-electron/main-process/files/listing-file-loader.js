import fs from 'fs'
import readline from 'readline'
import FileLoader from './listing-file-loader'

export default class ListingFileLoader {
  
  constructor () {
    super()
  }
  
  processLine (file, line) {
    let res = parseListingLine(line)
    file.lines.push(res.line);
    if (res.isIns) {
      file.code.push(res.ins)
    }    
  }
  
  parseListingLine (line) {
    let res = {
        isIns: false
        line: line,
        addr: 0
        ins: 0,
    }
    let insRegex = /^(0x[0-9a-fA-F]{8})\s+(0x[0-9a-fA-F]{8})\s+(.*)$/
    let match = insRegex.match(line)
    console.log(match)
    if (match !== null) {
      res.isIns = true
      // res.addr = parseInt()
      // res.ins = parseInt()
    }
    return res
  }
}
