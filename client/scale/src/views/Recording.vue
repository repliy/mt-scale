<template>
  <div style="height: 100%;" fluid class="grey lighten-4">
    <div>
      <RecordStatistics />
    </div>
    <div>
      <v-row class="pink">
        <v-col class="blue" cols="3">
          <RecordTab />
        </v-col>
        <v-col cols="9">
          <Weight />
        </v-col>
      </v-row>
    </div>
  </div>
</template>
<script>
import RecordStatistics from '@/components/recording/RecordStatistics.vue'
import RecordTab from '@/components/recording/RecordTab.vue'
import Weight from '@/components/recording/Weight.vue'

export default {
  name: 'Recording',
  components: {
    RecordStatistics,
    RecordTab,
    Weight
  },
  methods: {},
}
</script>
<style scoped>
</style>
