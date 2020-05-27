<template>
  <v-app>
    <v-content >
      <router-view class="pg-height"></router-view>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  components: {},
  created() {
    this.$i18n.locale = 'zh-CN'
  },
  data: () => ({
  }),
  methods: {}
}
</script>
<style scoped>
.pg-height {
  height: calc(100vh - 4px);
}
</style>
