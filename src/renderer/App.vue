<template>
  <div id="app">
    <v-app dark>
      <v-navigation-drawer app fixed clipped mini-variant>
        <v-list>
          <v-list-tile exact router :to="item.to" v-for="(item, i) in items" :key="i">
            <v-list-tile-action>
              <v-icon v-html="item.icon"></v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title v-text="item.title"></v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-navigation-drawer>

      <v-toolbar app fixed clipped>
        <v-spacer></v-spacer>
        <v-btn icon @click.native.stop="rightDrawer = !rightDrawer">
          <v-icon>menu</v-icon>
        </v-btn>
      </v-toolbar>

      <v-content>
        <v-container fluid fill-height>
          <v-slide-y-transition mode="out-in">
            <router-view></router-view>
          </v-slide-y-transition>
        </v-container>
      </v-content>

      <v-navigation-drawer app fixed clipped right v-model="rightDrawer">
        <registers></registers>
      </v-navigation-drawer>

      <v-footer fixed app>
      </v-footer>
    </v-app>
  </div>
</template>

<script>
  import registers from './components/display-view/registers'

  export default {
    name: 'epic-evm',

    components: {
      registers
    },

    data: () => ({
      rightDrawer: true,
      items: [
        { icon: 'tv', title: 'Display', to: '/' },
        { icon: 'bug_report', title: 'Debug', to: '/inspire' },
        { icon: 'folder_open', title: 'Files', to: '/inspire' }
      ]
    })
  }
</script>

<style>
  @import url('https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons');
  /* Global CSS */
</style>
