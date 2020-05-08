<template>
  <v-container fluid class=" grey lighten-5">
    <v-row class="red">
      <RecordStatistics />
    </v-row>
    <v-row class="pink">
      <v-col class="blue" cols="3">
        <RecordList />
      </v-col>
      <v-col cols="9"></v-col>
    </v-row>
  </v-container>
</template>
<script>
// @ is an alias to /src
import RecordStatistics from '@/components/RecordStatistics.vue'
import RecordList from '@/components/RecordList.vue'

export default {
  name: 'Recording',
  components: {
    RecordStatistics,
    RecordList
  },
  methods: {},
}
</script>
