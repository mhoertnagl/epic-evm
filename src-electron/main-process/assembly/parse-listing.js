export default function parseListingLine(line) {
  let res = {
      isIns: false
      line: line,
      addr: 0
      ins: 0,
  }
  let insRegex = /^(0x[0-9a-fA-F]{8})\s+(0x[0-9a-fA-F]{8}).*$/
  let match = insRegex.match(line)
  console.log(match)
  if (match !== null) {
    res.isIns = true
    // res.addr = parseInt()
    // res.ins = parseInt()
  }
  return res
}
