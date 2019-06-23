import Vue from 'vue'

import VFileLine from '@/modules/files/components/v-file-line.vue'
import VFile from '@/modules/files/components/v-file.vue'

import VDisplay from '@/modules/vm/components/v-display.vue'
import VRegisters from '@/modules/vm/components/v-registers.vue'

Vue.component('v-file', VFile)
Vue.component('v-file-line', VFileLine)
Vue.component('v-display', VDisplay)
Vue.component('v-registers', VRegisters)
