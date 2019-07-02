import { app, BrowserWindow, Menu } from 'electron'
import loadFile from '../files/load-file'

export default function appMenu(app, ui) {
  return Menu.buildFromTemplate([
    {
      label: 'File',
      submenu: [
        { label:'Load binary ...', click() { loadFile(app, ui) } },
        { type:'separator' },
        { label:'Exit', click() { app.quit() } }
      ]
    }
  ])
}
