<template>
  <q-layout view="hHh lpR fFf">
    <q-header>
        <!-- <q-bar>
          <div>EPIC VIrtual Machine</div>

          <q-space />

          <q-btn dense flat icon="minimize" @click="minimize" />
          <q-btn dense flat icon="crop_square" @click="maximize" />
          <q-btn dense flat icon="close" @click="closeApp" />
        </q-bar> -->

      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="left = !left" />

        <q-toolbar-title>
          EPIC VIrtual Machine
        </q-toolbar-title>

        <!-- <q-btn dense flat round icon="menu" @click="right = !right" /> -->
      </q-toolbar>
    </q-header>

<!--
  @mouseover="miniState = false"
  @mouseout="miniState = true"
-->
    <q-drawer
      v-model="left"
      :mini="miniState"
      bordered>
      <q-list>
        <q-item clickable tag="a" target="_blank" href="http://v1.quasar-framework.org">
          <q-item-section avatar>
            <q-icon name="folder" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Files</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <!-- <q-drawer v-model="right" side="right" bordered>
      Right
    </q-drawer> -->

    <q-page-container>
      <q-splitter v-model="splitter">

        <template v-slot:before>
          <router-view />
        </template>

        <template v-slot:after>
          <q-tabs
            v-model="tab"
            dense
            class="text-grey"
            active-color="primary"
            indicator-color="primary"
            align="justify"
            narrow-indicator
          >
            <q-tab name="registers" icon="alarm" label="Registers" />
            <q-tab name="others" icon="movie" label="Others" />
          </q-tabs>

          <q-separator />

          <q-tab-panels v-model="tab" animated>
            <q-tab-panel name="registers">
              Registers
            </q-tab-panel>
            <q-tab-panel name="others">
              Others
            </q-tab-panel>
          </q-tab-panels>
        </template>
      </q-splitter>
    </q-page-container>

    <q-footer bordered class="bg-grey-8 text-white">
      Footer
    </q-footer>
  </q-layout>
</template>

<script>
export default {
  name: 'MyLayout',
  data () {
    return {
      left: this.$q.platform.is.desktop,
      miniState: true,
      // right: true
      splitter: 75,
      tab: 'registers'
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
