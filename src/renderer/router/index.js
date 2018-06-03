import Vue from 'vue'
import Router from 'vue-router'

// import DisplayView from 'components/display-view'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'display-view',
      component: require('@/components/display-view').default
    }
  ]
})
