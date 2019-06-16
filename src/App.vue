<template>
  <div id="q-app">
    <router-view />
  </div>
</template>

<script>
import openAssemblyFile from '@/modules/files/services/file-dialog'

export default {
  name: 'App',

  created () {
    openAssemblyFile(filePath => {
      console.log(filePath)
    })
  }
}
</script>

<style>
</style>
