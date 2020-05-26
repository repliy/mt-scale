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
      <p v-if="leftTitle != ''" class="ma-0">{{ leftTitle }}</p>
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
  }),
  props: {
    rightBtnTitle: {
      type: String,
      default: ''
    },
    leftBtnTitle: {
      type: String,
      default: ''
    },
    leftTitle: {
      type: String,
      default: ''
    }
  },
  mounted() {},
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
