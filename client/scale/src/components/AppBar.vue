<template>
      <v-app-bar
      app
      color="primary"
      dark
    >
      <v-btn
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
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
  mounted() {},
  methods: {
    clickRightBtn() {
      this.$emit('clickRightBtn')
    }
  }
}
</script>
