export default class FileLoader {
  
  load (filePath, callback) {

    let file = {
      code: [],
      lines: [],
    }

    let rd = readline.createInterface({
      input: fs.createReadStream(filePath),
      console: false
    })

    rd.on('line', function(line) {
      processLine(file, line)
    })

    rd.on('close', function() {
      callback(file);
    })
  }
  
  processLine (file, line) {
    throw new Error('Not implemented!')
  }
}
