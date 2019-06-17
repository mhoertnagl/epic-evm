<template>
  <div id="q-app">
    <router-view />
  </div>
</template>

<script>
var ipc = require('electron').ipcRenderer; 
// import openAssemblyFile from '@/modules/files/services/file-dialog'

export default {
  name: 'App',

  created() {
    ipc.on('load-binary', function(event, filePath) {
      console.log(filePath);
    })
  }
}
</script>

<style>
</style>
