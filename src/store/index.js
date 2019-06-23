import Vue from 'vue'
import Vuex from 'vuex'
import Epic from '@/modules/vm/engine/epic'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    vm: new Epic(),
    file: []
  },
  // modules: {
  //   // example
  // },
  getters: {
    lines: state => {
      return state.file
    }
  },
  mutations: {
    setFile(state, file) {
      state.file = file
      console.log(file)
    }
  },
  // enable strict mode (adds overhead!)
  // for dev mode only
  strict: process.env.DEV
})
