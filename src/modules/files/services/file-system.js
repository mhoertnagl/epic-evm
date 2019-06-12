const path = require('path')
const fs = require('fs')

/**
 * Generator function that lists all files and sub-directories in a folder 
 * recursively in a synchronous fashion.
 *
 * @param {String} folder - Folder to start with.
 * @param {Number} recurseLevel - Number of times to recurse folders.
 * @returns {IterableIterator<String>}
 */
export function * walkFileSystem (folder, recurseLevel = 0) {
  try {
    const files = fs.readdirSync(folder)

    for (const file of files) {
      try {
        const pathToFile = path.join(folder, file)
        const stat = fs.statSync(pathToFile)
        const isDirectory = stat.isDirectory()
        if (isDirectory && recurseLevel > 0) {
          yield * walkFolders(pathToFile, recurseLevel - 1)
        }
        else {
          yield {
            rootDir: folder,
            fileName: file,
            isDir: isDirectory,
            stat: stat
          }
        }
      }
      catch (err) {
        yield {
          rootDir: folder,
          fileName: file,
          error: err
        }
      }
    }
  }
  catch (err) {
    yield {
      rootDir: folder,
      error: err
    }
  }
}

/**
 * Generator function that lists all sub-directories in a folder recursively in 
 * a synchronous fashion.
 *
 * @param {String} folder - Folder to start with.
 * @param {Number} recurseLevel - Number of times to recurse folders.
 * @returns {IterableIterator<String>}
 */
export function * walkFolders (folder, recurseLevel = 0) {
  for (const fileInfo of walkFileSystem(folder, recurseLevel)) {
    if (fileInfo.isDir) {
      yield fileInfo
    }
  }
}
