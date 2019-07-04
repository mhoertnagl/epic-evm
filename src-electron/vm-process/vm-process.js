//const Epic = require('./engine/epic')
import Epic from './engine/epic'

console.log('VM process initialized.');

const epic = new Epic()

process.send('Hallo')

process.on('message', message => {
  console.log('message from parent:', message);
})
