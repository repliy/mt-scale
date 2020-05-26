<template>
  <v-app-bar app color="primary" dark>
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
    <div class="d-flex align-center">
      <p v-if="username != ''" class="ma-0">{{ username }}</p>
      <v-divider class="mx-4" vertical></v-divider>
      <v-btn v-if="leftBtnTitle != ''" class="pa-0" text @click="clickLeftBtn">{{ leftBtnTitle }}</v-btn>
    </div>
    <v-spacer></v-spacer>
    <div class="d-flex align-center">
      <v-btn v-if="rightBtnTitle != ''" text @click="clickRightBtn">{{ rightBtnTitle }}</v-btn>
    </div>
  </v-app-bar>
</template>

<script>
export default {
  name: 'AppBar',
  data: () => ({
    username: ''
  }),
  props: {
    rightBtnTitle: {
      type: String,
      default: ''
    },
    leftBtnTitle: {
      type: String,
      default: ''
    }
  },
  mounted() {
    this.username = this.$store.getters.getUsername
  },
  methods: {
    clickRightBtn() {
      this.$emit('clickRightBtn')
    },
    clickLeftBtn() {
      this.$emit('clickLeftBtn')
    }
  }
}
</script>
