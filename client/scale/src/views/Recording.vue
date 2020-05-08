<template>
  <v-container fluid class=" grey lighten-5">
    <v-row class="red">
      <RecordStatistics />
    </v-row>
    <v-row class="pink">
      <v-col class="blue" cols="3">
        <RecordTab />
      </v-col>
      <v-col cols="9"></v-col>
    </v-row>
  </v-container>
</template>
<script>
// @ is an alias to /src
import RecordStatistics from '@/components/RecordStatistics.vue'
import RecordTab from '@/components/RecordTab.vue'

export default {
  name: 'Recording',
  components: {
    RecordStatistics,
    RecordTab
  },
  methods: {},
}
</script>
