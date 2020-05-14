<template>
  <v-row>
    <v-col cols="8" class="red">
      <div
        class="top-height green"
      >
        <RecordStatistics />
      </div>
      <div class="mt-2 bottom-height">
        <v-row
          align="stretch"
          justify="space-between"
          style="height: 100%;"
          class="ma-0"
        >
          <v-col cols="5">
            <RecordTab />
          </v-col>
          <v-col cols="6" class="pr-0 pl-0">
            <Weight />
          </v-col>
        </v-row>
      </div>
    </v-col>
    <v-col cols="4" class="">
      <Option />
    </v-col>
  </v-row>
</template>
<script>
import RecordStatistics from '@/components/recording/RecordStatistics.vue'
import RecordTab from '@/components/recording/RecordTab.vue'
import Weight from '@/components/recording/Weight.vue'
import Option from '@/components/recording/Option.vue'

export default {
  name: 'Recording',
  components: {
    RecordStatistics,
    RecordTab,
    Weight,
    Option
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
  margin-left: 12px;
}
.bottom-height {
  height: calc(74vh - 64px);
}
</style>
