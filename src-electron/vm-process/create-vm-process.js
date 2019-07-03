const { fork } = require('child_process');

export default function createVmProcess() {
  const process = fork('./src-electron/vm-process//vm-process-esm.js');
}
