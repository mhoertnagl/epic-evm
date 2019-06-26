import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    vm: null,
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
