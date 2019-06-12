<template>
  <div class="v-display">
    <div v-for="r in display.rows" :key="'r' + r" class="v-display-row">
      <div v-for="c in display.cols" :key="'r' + r + 'c' + c" class="v-display-col">
        {{ read(r-1, c-1) }}
      </div>
    </div>
  </div>
</template>

<script>
import Display from '../engine/display.js'

export default {
  name: 'Display',

  data () {
    return {
      display: null
    }
  },

  methods: {
    read (r, c) {
      let addr = r * this.display.cols + c
      let val = this.display.read(addr)
      return String.fromCharCode(val)
    }
  },

  created () {
    this.display = new Display()
  }
}
</script>

<style scoped>
.v-display {
  font-style: "Courier New", Courier, monospace;
  font-size: 12px;
  color: #eee;
  background: #111;
  padding: 10px;
}

.v-display-row {
  height: 12px;
}

.v-display-col {
  display: inline-block;
  width: 9px;
}
</style>
