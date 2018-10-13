<template>
  <v-layout row wrap justify-center>
    <!-- <div v-for="(row, r) in display" :key="r" class="display__row">
      <div v-for="(cell, c) in row" :key="r * cols + c">
        X
      </div>
      h
    </div> -->
    <div class="display-view__container">
      <table class="display">
        <tr v-for="(row, r) in display" :key="r" class="display__row">
          <td v-for="(cell, c) in row" :key="r * cols + c" class="display__cell">
            {{ cell }}
          </td>
        </tr>
      </table>
    </div>
  </v-layout>
</template>

<script>
  export default {
    name: 'display-view',

    data: function () {
      return {
        display: [],
        rows: 30, // 25,
        cols: 80
      }
    },

    methods: {

      initDisplay: function () {
        this.display = []
        for (let r = 0; r < this.rows; r++) {
          let row = []
          for (let c = 0; c < this.cols; c++) {
            row.push(String.fromCharCode(64))
          }
          this.display.push(row)
        }
      }
    },

    mounted: function () {
      this.initDisplay()
    }
  }
</script>

<style lang="stylus" scoped>
  .display-view__container
    padding 2px
    border 2px solid #282828

  .display
    background-color #282828
    font-family monospace
    font-size 16px

  .display__cell
    line-height 16px
</style>
