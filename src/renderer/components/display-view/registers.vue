<template>
  <div>
    <!-- <v-toolbar dark>
      <v-toolbar-side-icon>
        <v-icon>memory</v-icon>
      </v-toolbar-side-icon>
      <v-toolbar-title>Register</v-toolbar-title>
    </v-toolbar> -->

    <v-list>
      <v-subheader>Registers</v-subheader>
      <v-list-tile v-for="reg in regs" :key="reg.id">
        <v-list-tile-avatar>
          {{ reg.name }}
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ reg.value }}</v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-icon color="gray">edit</v-icon>
        </v-list-tile-action>
      </v-list-tile>
      <v-divider></v-divider>
    </v-list>
  </div>
</template>

<script>
  export default {
    name: 'registers',

    data: () => ({
      regs: [
        {
          id: 0,
          name: 'r0',
          value: 1
        },
        {
          id: 1,
          name: 'r1',
          value: 2
        }
      ]
    }),

    methods: {

    }
  }
</script>

<style scoped>

</style>