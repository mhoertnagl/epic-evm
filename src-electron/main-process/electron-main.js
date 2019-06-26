import { app, BrowserWindow, Menu } from 'electron'
import appMenu from './menu/app-menu'
import Epic from './engine/epic'

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

let mainWindow
let epic

function createWindow() {
  // Initial window options.
  mainWindow = new BrowserWindow({
    width: 1200,
    height: 800,
    useContentSize: true,
    //frame: false
  })

  mainWindow.loadURL(process.env.APP_URL)

  mainWindow.on('closed', () => {
    mainWindow = null
  })
  
  // Setup main menu.
  Menu.setApplicationMenu(appMenu(app, mainWindow)); 
  
  // Bootstrap virtual machine.
  epic = new Epic()
}

app.on('ready', createWindow)

app.on('reset', epic.reset)
//app.on('run', )
app.on('step', epic.step)
//app.on('pause', )
//app.on('stop', )

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', () => {
  if (mainWindow === null) {
    createWindow()
  }
})
