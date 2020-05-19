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
            <RecordTab
              ref="recordTab"
              class="layout-border ml-3 mr-3"
              v-on:recordTabChange="recordTabChange"
            ></RecordTab>
          </v-col>
          <v-col cols="7">
            <Weight
              ref="weight"
              class="layout-border mr-3"
              v-on:realWeight="realWeight"
            ></Weight>
          </v-col>
        </v-row>
      </div>
    </v-col>
    <v-col cols="4">
      <div class="right-height">
        <Option
          ref="option"
          class="layout-border pt-3"
          v-on:recordAction="recordAction"
        ></Option>
      </div>
    </v-col>
    <v-overlay :value="loading">
      <v-progress-circular
        indeterminate
        size="64"
      ></v-progress-circular>
    </v-overlay>
  </v-row>
</template>
<script>
import RecordStatistics from '@/components/recording/RecordStatistics.vue'
import RecordTab from '@/components/recording/RecordTab.vue'
import Weight from '@/components/recording/Weight.vue'
import Option from '@/components/recording/Option.vue'
import { getLatestTask } from '@/core/api/task.js'
import { AddWeightRecord } from '@/core/api/record.js'

export default {
  name: 'Recording',
  data: () => ({
    loading: false,
    weightNum: ''
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
      this.loading = true
      getLatestTask().then((response) => {
        this.loading = false
        this.$store.commit('SET_TASK_ID', response.data)
        // notify option
        this.$refs.option.$emit('taskReady', {
          task_id: response.data
        })
        // notify record table
        this.$refs.recordTab.$emit('taskReady', {
          task_id: response.data
        })
      }).catch((err) => {
        this.loading = false
        console.log(err)
      })
    },
    // tap record button, take species_id, box_id
    recordAction(data) {
      const params = Object.assign(data)
      params.task_id = this.$store.getters.getTaskId
      params.index = this.$refs.weight.index
      params.weight = Number(this.weightNum)
      AddWeightRecord(params).then((response) => {
        console.log(response.data)
      }).catch((error) => {
        console.log(error)
      })
    },
    // realtime weight output
    realWeight(val) {
      this.weightNum = val
      console.log('recording:', this.weightNum)
    },
    // record table data change
    recordTabChange() {
      const curIndex = this.$store.getters.getRecordIndex
      this.$refs.weight.index = curIndex + 1
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
