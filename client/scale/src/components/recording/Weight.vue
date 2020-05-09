<template>
  <v-container>
    <div>统计记录</div>
    <v-row class="red">
      <v-col cols="6">
        <div>序号: {{index}}</div>
        <div>当前重量: {{weight}}</div>
        <Keyboard />
      </v-col>
      <v-col cols="6">
        <div>物种</div>
        <div>请选择物种类型</div>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import Keyboard from '@/components/recording/Keyboard.vue'

export default {
  name: 'Weight',
  data: () => ({
    index: 69,
    weight: 352
  }),
  components: {
    Keyboard
  },
  methods: {},
}
</script>
