import { BrowserWindow } from 'electron'

export default function createUiProcess(app, appUrl) {
  const ui = new BrowserWindow({
    width: 1200,
    height: 800,
    useContentSize: true,
    //frame: false
    webPreferences: {
      nodeIntegrationInWorker: true
    }
  })

  ui.loadURL(appUrl)

  ui.on('closed', () => {
    app.quit()
  })
  
  return ui
}
