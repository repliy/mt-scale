<template>
<v-container>
  <v-label>统计记录</v-label>
  <v-row class="red">
    <v-col cols="6">
      <v-label>序号: {{index}}</v-label>
      <v-label>当前重量: {{weight}}</v-label>
      <v-card></v-card>
    </v-col>
    <v-col cols="6"></v-col>
  </v-row>
</v-container>
</template>
<script>

export default {
  name: 'Weight',
  data: () => ({
    index: 69,
    weight: 352
  }),
  components: {
  },
  methods: {},
}
</script>
