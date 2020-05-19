<template>
  <div style="height: 100%;">
    <v-row no-gutters>
      <v-col class="pb-0">
        <div class="tag_font">序号: {{index}}</div>
      </v-col>
    </v-row>
    <v-row no-gutters align="end">
      <v-col
        cols="4"
      >
        <p class="d-inline tag_font">当前重量: </p>
      </v-col>
      <v-col>
        <input
          readonly
          class="weight_input"
          :value="weight"
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
      align="end"
      no-gutters
    >
      <Keyboard v-on:kbClick="kbClick"></Keyboard>
    </v-row>
  </div>
</template>
<script>
import Keyboard from '@/components/recording/Keyboard.vue'

export default {
  name: 'Weight',
  data: () => ({
    index: 0,
    weight: '0'
  }),
  components: {
    Keyboard
  },
  methods: {
    kbClick(data) {
      let tmpNum
      if (data.num >= 0) {
        if (data.num === 0 && this.weight === '0') {
          return
        }
        if (this.weight.substr(0, 1) === '0') {
          tmpNum = this.weight.substr(1, this.weight.length - 1)
        } else {
          tmpNum = this.weight
        }
        tmpNum += data.num.toString()
      } else {
        tmpNum = this.weight.substr(0, this.weight.length - 1)
      }
      if (tmpNum.length === 0) {
        tmpNum = '0'
      }
      this.weight = tmpNum
      this.$emit('realWeight', this.weight)
    }
  }
}
</script>
<style scoped>
.tag_font {
  padding: 0px 12px 0px 12px;
  font-size: 20px;
  height: 28px;
  color: rgba(16, 16, 16, 1);
}
.weight_input {
  font-size: 72px;
  color: rgba(51, 48, 48, 1);
  width: 183px;
  height: 100px;
  padding: 0px;
  text-align: right;
}
.keyboard-layout {
  height: 39vh;
}
.species-color-tag {
  height: 24px;
  font-size: 12px;
}
</style>
