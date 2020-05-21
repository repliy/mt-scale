<template>
  <div>
    <v-data-table
      class="pl-2"
      :headers="headers"
      :items="tableData"
      hide-default-header
      hide-default-footer
      height=480
      @click:row="itemExpand"
      single-expand
      :expanded.sync="expanded"
      item-key="index"
      calculate-widths
      :loading="loading"
      :loading-text="loadingText"
    >
      <template v-slot:expanded-item="{ item }">
        <td :colspan="100">
          <v-row>
            <v-col>
              <v-btn
                tile
                block
                class="d-inline pa-2 deep-purple white--text"
                @click="itemEdit(item)"
              >编辑</v-btn>
            </v-col>
            <v-col>
              <v-btn
                tile
                block
                class="d-inline pa-2 red white--text"
                @click="itemDel(item)"
              >删除</v-btn>
            </v-col>
          </v-row>
        </td>
      </template>
      <template v-slot:item.name="{ item }">
        <div
          style="with: 107px;height: 37px;"
          color="grey lighten-5"
          class="ma-1"
        >
          <v-row
            no-gutters
            align="center"
            style="height: 100%;"
          >
            <v-col cols="3">
              <v-icon
                class="d-inline"
                :color="getColor(item.species)"
              >mdi-label</v-icon>
            </v-col>
            <v-col>
              <div style="height: 37px;">
                <p class="ma-0 species-color-tag">{{item.species}}</p>
                <p class="ma-0 species-color-tag">{{item.tags}}</p>
              </div>
            </v-col>
          </v-row>
        </div>
      </template>
      <template v-slot:item.weight="{ item }">
        <div class="species-color-tag">{{item.weight}}bl</div>
      </template>
    </v-data-table>
    <v-dialog
      v-model="dialog"
      max-width="290"
    >
      <v-card>
        <v-card-title class="headline">Notice</v-card-title>
        <v-card-text>
          Delete Record?
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="green darken-1"
            text
            @click="dialog = false"
          >
            Disagree
          </v-btn>
          <v-btn
            color="green darken-1"
            text
            @click="deleteRecord"
          >
            Agree
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import { FetchWeightRecord, DeleteWeightRecord } from '@/core/api/record.js'
export default {
  data() {
    return {
      dialog: false,
      loading: false,
      loadingText: 'Loading...Please wait...',
      delRecord: null,
      editRecord: null,
      expanded: [],
      headers: [
        {
          text: 'index',
          align: 'start',
          sortable: false,
          value: 'index'
        },
        { text: 'name', value: 'name' },
        { text: 'weight', value: 'weight' }
      ],
      tableData: []
    }
  },
  mounted() {
    this.$on('taskReady', () => {
      this.getWeightRecord()
    })
    this.$on('refreshData', () => {
      this.getWeightRecord()
    })
  },
  methods: {
    itemExpand(item, value) {
      value.expand(!value.isExpanded)
    },
    getColor(name) {
      return 'red'
    },
    itemEdit(item) {
      this.editRecord = item
      this.$emit('editRecordTabItem', item)
    },
    itemDel(item) {
      this.delRecord = item
      this.dialog = true
    },
    deleteRecord() {
      if (this.delRecord == null) {
        return
      }
      DeleteWeightRecord({
        id: this.delRecord.id
      }).then((response) => {
        this.dialog = false
        const taskId = this.$store.getters.taskId
        this.getWeightRecord({ task_id: taskId })
      }).catch((error) => {
        this.dialog = false
        console.log(error)
      })
    },
    getWeightRecord() {
      this.loading = true
      FetchWeightRecord({
        task_id: this.$store.getters.taskId
      }).then((response) => {
        this.loading = false
        this.tableData = response.data
        if (this.tableData[0]) {
          this.$store.commit('SET_RECORD_INDEX', this.tableData[0].index)
          this.$emit('recordTabChange', {})
        }
      }).catch((err) => {
        this.loading = false
        console.log(err)
      })
    }
  }
}
</script>
<style scoped>
tb {
  background-color: red;
}
.tab-font {
  size: 5px;
}
.species-color-tag {
  height: 16px;
  font-size: 6px;
}
.wight-unit-font {
  font-size: 18px;
  color: rgba(16, 16, 16, 1);
  height: 20px;
}
.text-start {
  padding: 0px;
}
</style>
