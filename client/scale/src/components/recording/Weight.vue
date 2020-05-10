<template>
  <v-row style="height: 100%;">
    <v-col
      cols="6"
      class="pa-0"
    >
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
          <input readonly class="weight_input" />
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
    </v-col>
    <v-col cols="6">
      <v-row>
        <v-col class="pt-0 pb-0">
          <div class="tag_font">箱子:</div>
        </v-col>
      </v-row>
      <v-row align="end">
        <v-col class="pt-0 pb-0">
          <div class="tag_font">{{boxType}}-{{boxTag}}</div>
        </v-col>
        <v-col class="pb-0">
          <v-btn
            class="pa-0"
            small
            color="indigo"
          >
            <v-icon dark>mdi-pencil</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-row class="mt-10">
        <v-col class="pt-0 pb-0">
          <div class="tag_font">物种:</div>
          <p style="height: 11px;color: rgba(255, 42, 42, 1);font-size: 8px;">请选择物种类型</p>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <template v-for="n in 7">
          <v-col
            :key="n"
            cols="6"
          >
            <div
              style="with: 107px;height: 37px;border: solid 1px black;border-radius: 6px;"
              color="grey lighten-5"
              class="ma-2"
            >
              <v-row
                no-gutters
                align="center"
              >
                <v-col cols="3">
                  <v-icon
                    class="d-inline"
                    color="red"
                  >mdi-label</v-icon>
                </v-col>
                <v-col>
                  <p
                    class="mb-0"
                    style="line-height: 37px;word-break:break-all;"
                  >hello </p>
                </v-col>
              </v-row>
            </div>
          </v-col>
          <v-responsive
            v-if="n % 2 === 0"
            :key="`width-${n}`"
            width="100%"
          ></v-responsive>
        </template>
      </v-row>
      <v-row class="mt-2">
        <v-col align="center">
          <v-btn
            height="40"
            color="primary"
            text-color="white"
          >记录</v-btn>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>
<script>
import Keyboard from '@/components/recording/Keyboard.vue'

export default {
  name: 'Weight',
  data: () => ({
    boxType: '大号',
    boxTag: '#999',
    index: 69,
    weight: 352
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
