import { dialog } from 'electron'

const options = {
  properties: ['openFile'],
  filters: [{ name: 'Epic Assembly', extensions: ['lst'] }]
}

export default function openFile(callback) {
  dialog.showOpenDialog(options, fileNames => {
    if (fileNames) {
      callback(fileNames[0])
    }
  })
}
