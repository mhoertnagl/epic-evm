import { app, BrowserWindow, Menu } from 'electron'
import openAssemblyFile from '../files/file-dialog'
import readAssembly from '../files/file-read'

export default function appMenu(mainWindow) {
  return Menu.buildFromTemplate([
    {
      label: 'File',
      submenu: [
        { label:'Load binary ...', click() { loadBinary(mainWindow) } },
        { type:'separator' },
        { label:'Exit', click() { app.quit() } }
      ]
    }
  ])
}

function loadBinary(mainWindow) {
  openAssemblyFile((filePath) => {
    readAssembly(filePath, (line) => {
      mainWindow.send('read-line', line);
    })
  })
}
