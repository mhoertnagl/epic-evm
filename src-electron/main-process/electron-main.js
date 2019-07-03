import { app, BrowserWindow, Menu } from 'electron'
import appMenu from './menu/app-menu'
import createUiProcess from './create-ui-process'
import createVmProcess from '../vm-process/create-vm-process'

/**
 * Set `__statics` path to static files in production;
 * The reason we are setting it here is that the path needs to be evaluated at 
 * runtime.
 */
if (process.env.PROD) {
  global.__statics = require('path')
    .join(__dirname, 'statics')
    .replace(/\\/g, '\\\\')
}

let ui
let vm

function createWindow() {
  ui = createUiProcess(app, process.env.APP_URL)
  vm = createVmProcess()
  Menu.setApplicationMenu(appMenu(app, ui)); 
}

app.on('ready', createWindow)

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', () => {
  if (ui === null) {
    createWindow()
  }
})
