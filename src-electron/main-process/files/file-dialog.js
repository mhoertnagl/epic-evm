const { dialog } = require('electron')

const options = {
  properties: ['openFile'],
  filters: [{ name: 'Epic Assembly', extensions: ['bin', 'lst'] }]
}

export default function openAssemblyFile(callback) {
  dialog.showOpenDialog(options, fileNames => {
    if (fileNames) {
      callback(fileNames[0])
    }
  })
}
