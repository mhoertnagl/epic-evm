<template>
  <div>
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
        <q-list dense bordered separator class="rounded-borders">
          <q-item v-for="(regValue, rid) in cpu.gp_regs" :key="rid">
            <q-item-section>
              {{ regNames[rid] }}
            </q-item-section>
            
            <q-item-section>
              <!-- <q-input filled dense v-model="regValue" /> -->
              {{ regValue }}
            </q-item-section>
          </q-item>
        </q-list>        
      </q-tab-panel>
      <q-tab-panel name="others">
        <q-list dense bordered separator class="rounded-borders">
          <q-item>
            <q-item-section>
              csr
            </q-item-section>
            
            <q-item-section>
              {{ cpu.cs_reg }}
            </q-item-section>
          </q-item>
        </q-list>  
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script>
import Cpu from '../engine/cpu'

export default {
  name: 'Registers',

  data () {
    return {
      tab: 'registers',
      regNames: [
        'r0', 'r1', 'r2', 'r3', 'r4', 'r5', 'r6', 'r7',
        'r8', 'r9', 'r10', 'r11', 'r12', 'fp', 'rp', 'ip'
      ],
      cpu: null
    }
  },

  methods: {

  },

  created () {
    this.cpu = new Cpu()
  }
}
</script>

<style scoped>

</style>
