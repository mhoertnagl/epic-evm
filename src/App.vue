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
    ipc.on('init-memory', (event, file) => {
      self.$store.commit('setFile', file)
    })
  }
}
</script>

<style>
</style>
