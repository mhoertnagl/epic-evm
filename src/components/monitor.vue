<template>
  <div class="monitor">
    <div v-for="r in display.rows" :key="'r' + r" class="monitor-row">
      <div v-for="c in display.cols" :key="'r' + r + 'c' + c" class="monitor-col">
        {{ read(r-1, c-1) }}
      </div>
    </div>
  </div>
</template>

<script>
import Display from '../vm/display.js'

export default {
  name: 'Monitor',

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
.monitor {
  font-style: "Courier New", Courier, monospace;
  font-size: 12px;
  color: #eee;
  background: #111;
  padding: 10px;
}

.monitor-row {
  height: 12px;
}

.monitor-col {
  display: inline-block;
  width: 9px;
}
</style>
