import { walkFolders } from './file-system'

export default function getFolders(absolutePath) {
  let folders = []

  // check incoming arg
  if (!absolutePath || typeof absolutePath !== 'string') {
    return folders
  }

  for (const fileInfo of walkFolders(absolutePath, false)) {
    if ('error' in fileInfo) {
      console.error(`Error: ${fileInfo.rootDir} - ${fileInfo.error}`)
      continue
    }
    console.log(fileInfo)
  }
  return folders
}
