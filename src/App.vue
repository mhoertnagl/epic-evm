<template>
  <div id="q-app">
    <router-view />
  </div>
</template>

<script>
// Eigenen window process für die vm.
// https://dzone.com/articles/learning-electron-4-things-i-wish-i-knew-sooner
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
