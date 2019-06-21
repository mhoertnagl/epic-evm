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
    ipc.on('read-line', function(event, line) {
      console.log(line);
    })
  }
}
</script>

<style>
</style>
