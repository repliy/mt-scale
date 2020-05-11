<template>
  <div
    fluid
    class="grey lighten-4 pg-height"
  >
    <div fluid class="top-height">
      <RecordStatistics />
    </div>
    <div class="mt-2 bottom-height">
      <v-row
        align="stretch"
        justify="space-between"
        style="height: 100%;"
        class="ma-0"
      >
        <v-col cols="4">
          <RecordTab />
        </v-col>
        <v-col cols="8">
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
  methods: {}
}
</script>
<style scoped>
.pg-height {
  height: calc(100vh - 64px);
}
.top-height {
  height: 25vh;
}
.bottom-height {
  height: calc(74vh - 64px);
}
</style>
