<template>
  <div id="q-app">
    <router-view />
  </div>
</template>

<script>
// https://medium.com/quasar-framework/building-an-electron-file-explorer-with-quasar-and-vue-7bf94f1bbf6
import getFolders from '@/modules/files/services/get-folders'

export default {
  name: 'App',
  
  created () {
    getFolders('/home/mathias')
  }
}
</script>

<style>
</style>
