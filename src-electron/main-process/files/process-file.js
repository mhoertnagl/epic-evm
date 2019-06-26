import ListingFileLoader from './listing-file-loader'

export default function processFile(path, callback) {
  switch (fileExt(path)) {
    case 'lst':
      let loader = new ListingFileLoader()
      loader.load(path, callback)
      break
    default:
      console.error(`Unsupported file extension: ${path}`)
      break
  }
}

function fileExt(path) {
  return path.substring(path.lastIndexOf('.') + 1, path.length)
}
