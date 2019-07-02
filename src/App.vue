<template>
  <div id="q-app">
    <router-view />
  </div>
</template>

<script>
var ipc = require('electron').ipcRenderer

export default {
  name: 'App',

  created() {
    const self = this
    ipc.on('mn->ui[load]', (e, file) => {
      self.$store.commit('setFile', file)
    })
    
    const vm = new Worker('modules/vm/worker/vm-worker.js');
  }
}
</script>

<style>
</style>
