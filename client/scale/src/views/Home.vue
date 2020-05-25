<template>
  <div class="home">
    <v-btn @click="goRecording" target="_blank" text>
      <span class="mr-2">go recording</span>
      <v-icon>mdi-open-in-new</v-icon>
    </v-btn>
    <img alt="Vue logo" src="../assets/logo.png" />
    <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
  </div>
</template>

<script>
// import AiScale from '@/components/AiScale.vue'

export default {
  name: 'Home',
  components: {
    // AiScale
  },
  methods: {
    goRecording() {
      // this.$router.push({ name: 'Recording' })
      this.$router.push({ name: 'Login' })
    }
  }
}
</script>
