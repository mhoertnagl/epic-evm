import fs from 'fs'
import readline from 'readline'
//import parseListingLine from '../assembly/parse-listing'

// TODO: Ordentliche Kapselung. FileLoader interface, ListingFileLoader
// TODO: Ordentliches VM setup.
// TODO: Engine in den main process.
export default function readAssembly(filePath, callback) {

  var file = {
    code: [],
    lines: [],
  }

  var rd = readline.createInterface({
    input: fs.createReadStream(filePath),
    //output: process.stdout,
    console: false
  })

  rd.on('line', function(line) {
    let res = parseListingLine(line)
    file.lines.push(res.line);
    if (res.isIns) {
      file.code.push(res.ins)
    }
  })

  rd.on('close', function() {
    callback(file);
  })
}
