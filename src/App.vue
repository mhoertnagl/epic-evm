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
    var self = this
    ipc.on('read-file', function(event, file) {
      self.$store.commit('setFile', file)
    })
  }
}
</script>

<style>
</style>
