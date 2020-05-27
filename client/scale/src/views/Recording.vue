<template>
  <div>
    <AppBar
      :rightBtnTitle="$tc('ssr.complete')"
      :leftBtnTitle="$tc('ssr.logout')"
      :leftTitle="username"
      v-on:clickLeftBtn="logoutAction"
      v-on:clickRightBtn="recordCompleteAction"
    ></AppBar>
    <v-row
      no-gutters
      class="pg-height"
    >
      <v-col cols="8">
        <div class="top-height pa-3">
          <RecordStatistics
            ref="stat"
            class="layout-border"
            v-on:alertMessage="receiveAlertMessage"
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
                v-on:recordIndexChange="getCurrentIndex"
                v-on:alertMessage="receiveAlertMessage"
              ></RecordTab>
            </v-col>
            <v-col cols="7">
              <Weight
                ref="weight"
                v-model="weightNum"
                :index="weightIndex"
                class="layout-border mr-3"
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
            :editMode="optionEditMode"
            v-on:addRecord="addRecord"
            v-on:updateRecord="updateRecord"
            v-on:cancelUpdate="cancelUpdate"
            v-on:alertMessage="receiveAlertMessage"
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
        :type="alertType"
        class="pg-alert"
        transition="slide-y-transition"
        :value="showAlert"
      >
        {{alertMessage}}
      </v-alert>
      <v-dialog
        v-model="dialog"
        max-width="290"
      >
        <v-card>
          <v-card-title class="headline">Notice</v-card-title>
          <v-card-text>
            {{dialogText}}
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="dialog = false"
            >
              {{$tc('ssr.disagree')}}
            </v-btn>
            <v-btn
              color="green darken-1"
              text
              @click="dialogAgree"
            >
              {{$tc('ssr.agree')}}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
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

const DIALOG = {
  TYPE: {
    COMPLETE: 'Record-Complete',
    LOGOUT: 'Logout'
  }
}

export default {
  name: 'Recording',
  data: () => ({
    optionEditMode: false,
    dialog: false,
    dialogText: '',
    dialogType: DIALOG.TYPE.COMPLETE,
    username: '',
    showAlert: false,
    alertMessage: '',
    alertType: 'error',
    loading: false,
    weightIndex: 0,
    weightNum: '0',
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
    this.username = this.$store.getters.getUsername
    console.log('111', this.username)
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
        this.tipAlertMessage('error', err.message)
      })
    },
    // add new record
    addRecord(data) {
      const params = Object.assign(data)
      params.task_id = this.$store.getters.taskId
      params.index = this.weightIndex
      params.weight = Number(this.weightNum)
      if (params.weight <= 0) {
        this.tipAlertMessage('warning', '重量不能为0')
        return
      }
      this.loading = true
      AddWeightRecord(params).then((response) => {
        this.loading = false
        this.refreshData()
      }).catch((err) => {
        this.loading = false
        this.tipAlertMessage('error', err.message)
      })
    },
    // edit record
    updateRecord(data) {
      this.loading = true
      const params = Object.assign(data)
      params.index = this.weightIndex
      params.weight = Number(this.weightNum)
      params.id = this.editItem.id
      console.log(params)
      UpdWeightRecord(params).then((response) => {
        this.loading = false
        this.optionEditMode = false
        this.refreshData()
      }).catch((err) => {
        this.loading = false
        this.tipAlertMessage('error', err.message)
      })
    },
    cancelUpdate() {
      this.optionEditMode = false
      this.weightNum = '0'
      this.$refs.recordTab.closeExpandItem()
      this.getCurrentIndex()
    },
    refreshData() {
      this.$refs.recordTab.$emit('refreshData')
      this.$refs.stat.$emit('refreshData')
    },
    recordCompleteAction() {
      this.dialogText = this.$tc('ssr.completeTip') + '?'
      this.dialogType = DIALOG.TYPE.COMPLETE
      this.dialog = true
    },
    logoutAction() {
      this.dialogText = this.$tc('ssr.logout') + '?'
      this.dialogType = DIALOG.TYPE.LOGOUT
      this.dialog = true
    },
    dialogAgree() {
      switch (this.dialogType) {
        case DIALOG.TYPE.COMPLETE:
          this.loading = true
          updateTaskStatus({
            id: this.$store.getters.taskId,
            status: 'complete'
          }).then((response) => {
            this.loading = false
            const elemIF = document.createElement('iframe')
            const timestamp = (new Date()).valueOf()
            elemIF.src = '/api/test/excel?snapshotTime=' + timestamp
            elemIF.style.display = 'none'
            document.body.appendChild(elemIF)
          }).catch((err) => {
            this.loading = false
            this.tipAlertMessage('error', err.message)
          })
          break
        case DIALOG.TYPE.LOGOUT:
          this.$store.dispatch('FedLogOut')
          break
        default :
      }
    },
    // record table data change
    recordTabChange() {
      this.refreshData()
    },
    getCurrentIndex() {
      const curIndex = this.$store.getters.recordIndex
      this.weightIndex = curIndex + 1
    },
    // click record table item edit button
    editRecordTabItem(item) {
      this.editItem = item
      this.weightIndex = item.index
      this.weightNum = item.weight.toString()
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
      this.optionEditMode = true
    },
    receiveAlertMessage(data) {
      this.tipAlertMessage(data.type, data.msg)
    },
    tipAlertMessage(type, msg) {
      const _this = this
      this.alertMessage = msg
      this.alertType = type
      this.showAlert = true
      setTimeout(() => {
        _this.showAlert = false
      }, 2000)
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
