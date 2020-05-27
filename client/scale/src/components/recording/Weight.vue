<template>
  <div style="height: 100%;">
    <v-row no-gutters>
      <v-col class="pb-0">
        <div class="tag_font">{{$tc('ssr.index')}}: {{index}}</div>
      </v-col>
    </v-row>
    <v-row no-gutters align="end">
      <v-col
        cols="4"
      >
        <p class="d-inline tag_font">{{$tc('ssr.curWeight')}}: </p>
      </v-col>
      <v-col>
        <input
          ref="weightInput"
          readonly
          class="weight_input"
          :value="displayWeight(weight)"
        />
        <v-divider darkr></v-divider>
      </v-col>
      <v-col>
        <p class="d-inline">{{$tc('ssr.pound')}}</p>
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
import { toThousands } from '@/utils/func.js'

export default {
  name: 'Weight',
  data: () => ({
  }),
  props: {
    index: {
      type: Number,
      default: 1
    },
    weight: {
      type: String,
      default: '0'
    }
  },
  model: {
    prop: 'weight',
    event: 'change'
  },
  components: {
    Keyboard
  },
  methods: {
    displayWeight(weight) {
      return toThousands(weight)
    },
    kbClick(data) {
      let tmpNum
      if (data.num >= 0) {
        // tap 0~9
        if (data.num === 0 && this.weight === '0') {
          // tap 0 and weight is 0
          return
        }
        if (this.weight.substr(0, 1) === '0') {
          // weight is 0
          tmpNum = this.weight.substr(1, this.weight.length - 1)
        } else {
          tmpNum = this.weight
        }
        tmpNum += data.num.toString()
      } else {
        // tap delete
        tmpNum = this.weight.substr(0, this.weight.length - 1)
      }
      if (tmpNum.length === 0) {
        tmpNum = '0'
      }
      this.changeDisplaySize(tmpNum)
      this.$emit('change', tmpNum)
    },
    changeDisplaySize(num) {
      let fontSize = 72
      if (num.length > 4) {
        let tmp = fontSize - (num.length - 4) * 10
        if (tmp <= 32) {
          tmp = 32
        }
        fontSize = tmp
      }
      this.$refs.weightInput.style.fontSize = fontSize + 'px'
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
  text-align: right;
  padding-top:30px;
}
.keyboard-layout {
  height: 39vh;
}
.species-color-tag {
  height: 24px;
  font-size: 12px;
}
</style>
