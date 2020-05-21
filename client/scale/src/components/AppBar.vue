<template>
  <v-app-bar
    app
    color="primary"
    dark
  >
    <div class="d-flex align-center">
      <v-img
        alt="Vuetify Logo"
        class="shrink mr-2"
        contain
        src="../assets/logo.png"
        transition="scale-transition"
        width="150"
      />
    </div>
    <v-spacer></v-spacer>
    <div class="d-flex align-center">
      <v-btn
        text
        @click="clickRightBtn"
      >{{rightBtnTitle}}</v-btn>
    </div>
  </v-app-bar>
</template>

<script>
export default {
  name: 'AppBar',
  data: () => ({
  }),
  props: {
    rightBtnTitle: {
      type: String,
      default: ''
    }
  },
  mounted() { },
  methods: {
    clickRightBtn() {
      this.$emit('clickRightBtn')
    }
  }
}
</script>
