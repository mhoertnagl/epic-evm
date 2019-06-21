<template>
  <q-layout view="hHh lpR fFf">
    <q-header>
      <q-toolbar>
        <q-toolbar-title>
          EPIC Virtual Machine
        </q-toolbar-title>

        <q-btn dense flat round icon="play_arrow" />
        <q-btn dense flat round icon="pause" />
        <q-btn dense flat round icon="stop" />
      </q-toolbar>
    </q-header>

    <q-drawer v-model="left" mini bordered>
      <q-list>
        <q-item clickable v-ripple @click="$router.push('/')">
          <q-item-section avatar>
            <q-icon name="desktop_windows" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Virtual Machine</q-item-label>
          </q-item-section>
        </q-item>
        <q-item clickable v-ripple @click="$router.push('/debug')">
          <q-item-section avatar>
            <q-icon name="bug_report" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Debug</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container>
      <q-splitter v-model="splitter">

        <template v-slot:before>
          <router-view />
        </template>

        <template v-slot:after>
          <v-registers></v-registers>
        </template>
      </q-splitter>
    </q-page-container>

    <q-footer bordered class="bg-grey-8 text-white q-pl-md">
      Footer
    </q-footer>
  </q-layout>
</template>

<script>
export default {
  name: 'MyLayout',
  data () {
    return {
      left: true,
      splitter: 75,
    }
  },

  methods: {
    // minimize () {
    //   if (process.env.MODE === 'electron') {
    //     this.$q.electron.remote.BrowserWindow.getFocusedWindow().minimize()
    //   }
    // },

    // maximize () {
    //   if (process.env.MODE === 'electron') {
    //     const win = this.$q.electron.remote.BrowserWindow.getFocusedWindow()

    //     if (win.isMaximized()) {
    //       win.unmaximize()
    //     } else {
    //       win.maximize()
    //     }
    //   }
    // },

    // closeApp () {
    //   if (process.env.MODE === 'electron') {
    //     this.$q.electron.remote.BrowserWindow.getFocusedWindow().close()
    //   }
    // }
  }
}
</script>

<style>
</style>
