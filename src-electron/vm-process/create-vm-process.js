const { fork } = require('child_process');

export default function createVmProcess() {
  const prog = './src-electron/vm-process//vm-process-esm.js'
  const params = []
  const opts = { stdio: [ 'inherit', 'inherit', 'inherit', 'ipc' ] }
  // const opts = { stdio: [ 'pipe', 'pipe', 'pipe', 'ipc' ] }
  const vm = fork(prog, params, opts)
  
  vm.on('message', message => {
    console.log('message from child:', message);
    vm.send('Hi');
  })
  
  return vm
}
