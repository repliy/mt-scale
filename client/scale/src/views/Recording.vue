<template>
  <v-row
    no-gutters
    class="pg-height"
  >
    <v-col cols="8">
      <div class="top-height pa-3">
        <RecordStatistics class="layout-border"></RecordStatistics>
      </div>
      <div class="bottom-height">
        <v-row
          align="stretch"
          justify="space-between"
          style="height: 100%;"
          no-gutters
        >
          <v-col cols="5">
            <RecordTab class="layout-border ml-3 mr-3"></RecordTab>
          </v-col>
          <v-col cols="7">
            <Weight class="layout-border mr-3"></Weight>
          </v-col>
        </v-row>
      </div>
    </v-col>
    <v-col cols="4">
      <div class="right-height">
        <Option class="layout-border pt-3" v-on:weightRecord="weightRecord"></Option>
      </div>
    </v-col>
  </v-row>
</template>
<script>
import RecordStatistics from '@/components/recording/RecordStatistics.vue'
import RecordTab from '@/components/recording/RecordTab.vue'
import Weight from '@/components/recording/Weight.vue'
import Option from '@/components/recording/Option.vue'
import { getLatestTask } from '@/core/api/task.js'

export default {
  name: 'Recording',
  data: () => ({

  }),
  components: {
    RecordStatistics,
    RecordTab,
    Weight,
    Option
  },
  mounted() {
    this.getTask()
  },
  methods: {
    getTask() {
      getLatestTask().then((response) => {
        this.$store.commit('SET_TASK_ID', response.data)
      }).catch((err) => {
        console.log(err)
      })
    },
    weightRecord(data) {
      console.log(data)
    }
  }
}
</script>
<style scoped>
.pg-height {
  height: calc(100vh - 64px);
}
.top-height {
  height: 28vh;
}
.bottom-height {
  height: calc(71vh - 64px);
}
.right-height {
  height: calc(99vh - 64px);
}
.layout-border {
  border: 1px solid rgba(187, 187, 187, 1);
}
</style>
