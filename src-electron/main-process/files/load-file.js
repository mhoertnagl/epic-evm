import openFile from './file-dialog'
import processFile from './process-file'

export default function loadFile(mn, ui) {
  openFile((filePath) => {
    processFile(filePath, (file) => {
      // vm.send('mn->vm[load]', vm, file.code)
      ui.send('mn->ui[load]', file)
    })
  })
}
