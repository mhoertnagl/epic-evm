import fs from 'fs'
import readline from 'readline'

export default function readAssembly(filePath, callback) {

  var file = []

  var rd = readline.createInterface({
    input: fs.createReadStream(filePath),
    //output: process.stdout,
    console: false
  });

  rd.on('line', function(line) {
    file.push(line);
  });

  rd.on('close', function() {
    callback(file);
  });
}