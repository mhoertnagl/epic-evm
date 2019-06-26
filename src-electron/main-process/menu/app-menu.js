import { app, BrowserWindow, Menu } from 'electron'
import loadFile from '../files/load-file'

export default function appMenu(app, view) {
  return Menu.buildFromTemplate([
    {
      label: 'File',
      submenu: [
        { label:'Load binary ...', click() { loadFile(app, view) } },
        { type:'separator' },
        { label:'Exit', click() { app.quit() } }
      ]
    }
  ])
}
