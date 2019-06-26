import openFile from './file-dialog'
import processFile from './process-file'

function loadFile(app, view) {
  openFile((filePath) => {
    processFile(filePath, (file) => {
      app.send('reset', file.code)
      view.send('init-memory', file)
    })
  })
}
