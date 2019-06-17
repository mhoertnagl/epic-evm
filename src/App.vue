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
    ipc.on('load-binary', function(event, filePath) {
      console.log(filePath);
    })
  }
}
</script>

<style>
</style>
