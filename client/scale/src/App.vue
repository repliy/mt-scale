<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-btn @click="toggleDrawer" target="_blank" text>
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
      <v-spacer></v-spacer>
      <div class="d-flex align-center">
        <v-img alt="Vuetify Logo" class="shrink mr-2" contain src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png" transition="scale-transition" width="40" />
      </div>
    </v-app-bar>
    <v-content >
      <router-view class="pg-height"></router-view>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'App',

  components: {},

  data: () => ({
    drawerModel: null
  }),
  methods: {
    toggleDrawer() {
      this.drawerModel = !this.drawerModel
    }
  }
}
</script>
<style scoped>
.pg-height {
  height: calc(100vh - 4px);
}
</style>
