<template>
  <q-scroll-area class="v-file q-pa-md">
    <v-file-line v-for="(line, idx) in lines" :line="line" :key="idx"></v-file-line>
  </q-scroll-area>
</template>

<script>
export default {
  name: 'File',

  data () {
    return {

    }
  },

  computed: {
    lines () {
      return this.$store.getters.lines
    }
  },

  methods: {

  },

  created () {

  }
}
</script>

<style scoped>
.v-file {
  font-family: "Lucida Console", Monaco, monospace;
  width: 100%;
}
</style>
