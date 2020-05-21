<template>
  <v-row
    no-gutters
    class="pg-height"
  >
    <v-col cols="8">
      <div class="top-height pa-3">
        <RecordStatistics ref="stat" class="layout-border"></RecordStatistics>
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
              v-on:editRecordTabItem="editRecordTabItem"
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
          v-on:addRecord="addRecord"
          v-on:updateRecord="updateRecord"
        ></Option>
      </div>
    </v-col>
    <v-overlay
      light
      :value="loading"
      opacity="0.23"
    >
      <v-progress-circular
        indeterminate
        size="48"
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
import { AddWeightRecord, UpdWeightRecord } from '@/core/api/record.js'

export default {
  name: 'Recording',
  data: () => ({
    loading: true,
    weightNum: '',
    editItem: null
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
        // notify stat
        this.$refs.stat.$emit('taskReady', {
          task_id: response.data
        })
      }).catch((err) => {
        this.loading = false
        console.log(err)
      })
    },
    // add new record
    addRecord(data) {
      this.loading = true
      const params = Object.assign(data)
      params.task_id = this.$store.getters.taskId
      params.index = this.$refs.weight.index
      params.weight = Number(this.weightNum)
      AddWeightRecord(params).then((response) => {
        console.log(response.data)
        this.loading = false
      }).catch((error) => {
        console.log(error)
        this.loading = false
      })
    },
    // edit record
    updateRecord(data) {
      console.log('updateRecord')
      this.loading = true
      const params = Object.assign(data)
      params.index = this.$refs.weight.index
      params.weight = Number(this.weightNum)
      params.id = this.editItem.id
      console.log(params)
      UpdWeightRecord(params).then((response) => {
        this.loading = false
        this.switchEditMode()
      }).catch((error) => {
        console.log(error)
        this.loading = false
        this.switchEditMode()
      })
    },
    // realtime weight output
    realWeight(val) {
      this.weightNum = val
      console.log('recording:', this.weightNum)
    },
    // record table data change
    recordTabChange() {
      const curIndex = this.$store.getters.recordIndex
      this.$refs.weight.index = curIndex + 1
    },
    // click record table item
    editRecordTabItem(item) {
      this.editItem = item
      this.$refs.weight.index = item.index
      this.$refs.weight.weight = item.weight.toString()
      this.$refs.option.specTag = item.tag
      // species
      const specItems = this.$refs.option.speciesItems
      for (let i = 0; i < specItems.length; i++) {
        if (specItems[i].name === item.species) {
          this.$refs.option.specSeleIndex = i
          break
        }
      }
      // box
      const boxItems = this.$refs.option.boxItems
      for (let i = 0; i < boxItems.length; i++) {
        if (boxItems[i].type === item.box_type) {
          this.$refs.option.boxSeleIndex = i
          boxItems[i].id = item.box_id
          boxItems[i].num = item.box_num
          break
        }
      }
      this.switchEditMode()
    },
    // option edit mode
    switchEditMode() {
      console.log(this.$refs.recordTab)
      this.$refs.option.editMode = !this.$refs.option.editMode
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
