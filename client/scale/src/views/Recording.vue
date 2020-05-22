<template>
  <div>
    <AppBar rightBtnTitle="complete" v-on:clickRightBtn="recordComplete"></AppBar>
    <v-row
      no-gutters
      class="pg-height"
    >
      <v-col cols="8">
        <div class="top-height pa-3">
          <RecordStatistics
            ref="stat"
            class="layout-border"
          ></RecordStatistics>
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
            class="layout-border mt-3"
            v-on:addRecord="addRecord"
            v-on:updateRecord="updateRecord"
          ></Option>
        </div>
      </v-col>
      <!--loading-->
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
      <!--error message-->
      <v-alert
        type="error"
        class="pg-alert"
        transition="slide-y-transition"
        :value="showError"
      >
        {{errorMessage}}
      </v-alert>
    </v-row>
  </div>
</template>
<script>
import RecordStatistics from '@/components/recording/RecordStatistics.vue'
import RecordTab from '@/components/recording/RecordTab.vue'
import Weight from '@/components/recording/Weight.vue'
import Option from '@/components/recording/Option.vue'
import AppBar from '@/components/AppBar.vue'
import { getLatestTask, updateTaskStatus } from '@/core/api/task.js'
import { AddWeightRecord, UpdWeightRecord } from '@/core/api/record.js'

export default {
  name: 'Recording',
  data: () => ({
    showError: false,
    errorMessage: '',
    loading: false,
    weightNum: '',
    editItem: null
  }),
  components: {
    RecordStatistics,
    RecordTab,
    Weight,
    Option,
    AppBar
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
        this.$refs.option.$emit('taskReady')
        // notify record table
        this.$refs.recordTab.$emit('taskReady')
        // notify stat
        this.$refs.stat.$emit('taskReady')
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
        this.$refs.recordTab.$emit('refreshData')
        this.loading = false
      }).catch((error) => {
        console.log(error)
        this.loading = false
      })
    },
    // edit record
    updateRecord(data) {
      this.loading = true
      const params = Object.assign(data)
      params.index = this.$refs.weight.index
      params.weight = Number(this.weightNum)
      params.id = this.editItem.id
      console.log(params)
      UpdWeightRecord(params).then((response) => {
        this.loading = false
        this.switchEditMode()
        this.$refs.recordTab.$emit('refreshData')
      }).catch((error) => {
        console.log(error)
        this.loading = false
        this.switchEditMode()
      })
    },
    recordComplete() {
      const elemIF = document.createElement('iframe')
      const timestamp = (new Date()).valueOf()
      
      elemIF.src = '/api/test/excel?snapshotTime=' + timestamp

      console.log(elemIF.src)
      elemIF.style.display = 'none'
      document.body.appendChild(elemIF)
      // this.loading = true
      // updateTaskStatus({
      //   id: this.$store.getters.taskId,
      //   status: 'complete'
      // }).then((response) => {
      //   this.loading = false
      // }).catch((error) => {
      //   this.loading = false
      //   console.log(error)
      // })
    },
    // realtime weight output
    realWeight(val) {
      this.weightNum = val
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
      this.$refs.option.editMode = !this.$refs.option.editMode
    },
    tipErrorMessage(msg) {
      const _this = this
      this.errorMessage = msg
      this.showError = true
      setTimeout(() => {
        _this.showError = false
      }, 3000)
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
  height: calc(97.5vh - 64px);
}
.layout-border {
  border: 1px solid rgba(187, 187, 187, 1);
}
.pg-alert {
  width: 100vw;
  position: absolute;
  top: 0;
}
</style>
