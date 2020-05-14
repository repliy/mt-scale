<template>
  <div>
    <v-row>
      <v-col class="pt-0 pb-0">
        <div class="tag_font">序号: {{index}}</div>
      </v-col>
    </v-row>
    <v-row align="end">
      <v-col
        class="pr-0"
        cols="4"
      >
        <p class="d-inline tag_font">当前重量: </p>
      </v-col>
      <v-col class="pl-0">
        <input
          readonly
          class="weight_input"
        />
        <v-divider darkr></v-divider>
      </v-col>
      <v-col>
        <p class="d-inline">lb</p>
      </v-col>
    </v-row>
    <v-row
      class="keyboard-layout mt-5"
      justify="center"
    >
      <Keyboard />
    </v-row>
  </div>
</template>
<script>
import Keyboard from '@/components/recording/Keyboard.vue'

export default {
  name: 'Weight',
  data: () => ({
    index: 69
  }),
  components: {
    Keyboard
  },
  methods: {}
}
</script>
<style scoped>
.tag_font {
  font-size: 20px;
  height: 28px;
  color: rgba(16, 16, 16, 1);
}
.weight_input {
  font-size: 72px;
  color: rgba(51, 48, 48, 1);
  width: 183px;
  height: 100px;
  text-align: center;
  padding-top: 30px;
}
.keyboard-layout {
  height: 39vh;
}
.species-color-tag {
  height: 24px;
  font-size: 12px;
}
</style>
