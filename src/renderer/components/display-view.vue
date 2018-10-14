<template>
  <v-layout row wrap justify-center>
    <div class="display-view__container">
      <table class="display">
        <tr v-for="(row, r) in displayMemory" :key="r" class="display__row">
          <td v-for="(cell, c) in row" :key="r * 80 + c" class="display__cell">
            {{ cell }}
          </td>
        </tr>
      </table>
    </div>
  </v-layout>
</template>

<script>
  // import Vue from 'vue'
  import VM from '../../main/vm/vm'

  // export default class DisplayView extends Vue {
  //   get displayMemory () {
  //     return VM.displayMemory
  //   }
  // }

  export default {
    name: 'display-view',

    computed: {
      displayMemory: function () {
        return VM.displayMemory
      }
    }
  }
</script>

<style lang="stylus" scoped>
  .display-view__container
    padding 2px
    border 2px solid #282828

  .display
    padding 2px
    background-color #282828
    font-family monospace
    font-size 16px

  .display__cell
    line-height 16px
    height 16px
    width 10px
</style>
