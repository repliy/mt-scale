<template>
  <div style="margin-right: 12px;height: 100%;">
    <v-progress-linear
      indeterminate
      :active="specLoading"
      height="2"
    >
    </v-progress-linear>
    <v-row no-gutters class="mt-3">
      <v-col>
        <div class="title-font">{{$tc('ssr.selectSpecies')}}: {{specTag}}</div>
        <p class="subtitle-font">{{$tc('ssr.selectSpeciesType')}}</p>
      </v-col>
    </v-row>
    <!--species select-->
    <v-item-group
      v-model="specSeleIndex"
      @change="changeSpec"
      mandatory
    >
      <v-container>
        <v-row no-gutters>
          <template v-for="(item,i) in speciesItems">
            <v-col
              :key="item.id"
              cols="6"
            >
              <v-item v-slot:default="{ active, toggle }">
                <v-card
                  style="with: 107px;height: 37px;border-radius: 6px;"
                  class="ma-1 pa-0"
                  :color="active ? 'primary lighten-1' : ''"
                  @click="toggle"
                >
                  <v-row
                    no-gutters
                    align="center"
                  >
                    <v-col cols="3">
                      <v-icon
                        class="d-inline"
                        :color="item.color"
                      >mdi-label</v-icon>
                    </v-col>
                    <v-col>
                      <p
                        class="mb-0"
                        style="line-height: 37px;word-break:break-all;"
                      >{{item.name}}</p>
                    </v-col>
                  </v-row>
                </v-card>
              </v-item>
            </v-col>
            <v-responsive
              v-if="(i + 1) % 2 === 0"
              :key="`width-${i}`"
              width="100%"
            ></v-responsive>
          </template>
        </v-row>
      </v-container>
    </v-item-group>
    <v-divider darkr></v-divider>
    <v-row
      no-gutters
      class="mt-5"
    >
      <v-col adjust="center">
        <v-progress-linear
          indeterminate
          :active="boxLoading"
          height="2"
        >
        </v-progress-linear>
        <div class="title-font d-inline">{{$tc('ssr.selectBox')}}:</div>
        <v-btn
          class="ml-5 d-inline"
          small
          color="indigo"
          @click="boxDialog = true"
        >
          <v-icon dark>mdi-pencil</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row
      no-gutters
      align="end"
    >
      <v-col>
        <p class="subtitle-font">{{$tc('ssr.selectBoxType')}}</p>
      </v-col>
    </v-row>
    <!--box select-->
    <v-item-group
      v-model="boxSeleIndex"
      @change="changeBox"
      mandatory
    >
      <v-container>
        <v-row no-gutters>
          <template v-for="(box, n) in boxItems">
            <v-col
              :key="n"
              cols="6"
            >
              <v-item v-slot:default="{ active, toggle }">
                <v-card
                  style="with: 107px;height: 52px;border-radius: 6px;"
                  class="ma-1 pa-0"
                  :color="active ? 'primary lighten-1' : ''"
                  @click="toggle"
                  :disabled="box.num == ''"
                >
                  <p
                    class="mb-0 pl-2"
                    style="font-size: 20px;word-break:break-all;"
                  >{{box.name}}</p>
                  <p
                    class="mb-0 pr-2"
                    style="font-size: 16px;word-break:break-all;text-align: right;"
                  >{{box.num}}</p>
                </v-card>
              </v-item>
            </v-col>
            <v-responsive
              v-if="(n + 1) % 2 === 0"
              :key="`width-${n}`"
              width="100%"
            ></v-responsive>
          </template>
        </v-row>
      </v-container>
    </v-item-group>
    <v-row justify="center">
      <v-dialog
        v-model="tagDialog"
        persistent
        max-width="350"
      >
        <v-card>
          <v-card-title class="headline">{{$tc('ssr.tagInfo')}}</v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="specTag"
                  clearable
                  :label="$tc('ssr.enterTagNo')"
                  type="text"
                >
                </v-text-field>
              </v-col>
              <v-col></v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="cancelTagDialog"
            >{{$tc('ssr.disagree')}}</v-btn>
            <v-btn
              color="green darken-1"
              text
              @click="okTagDialog"
            >{{$tc('ssr.agree')}}</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
    <v-row
      no-gutters
      class="mt-2"
    >
      <v-col align="center">
        <v-btn
          height="40"
          width="100"
          color="primary"
          text-color="white"
          @click="record"
          v-if="!editMode"
        >{{$tc('ssr.recordBtn')}}</v-btn>
        <v-row v-else>
          <v-col>
            <v-btn
              height="40"
              width="100"
              color="primary"
              text-color="white"
              @click="update"
            >{{$tc('ssr.updBtn')}}</v-btn>
          </v-col>
          <v-col>
            <v-btn
              height="40"
              width="100"
              color="primary"
              text-color="white"
              @click="cancelUpdate"
            >{{$tc('ssr.cancelBtn')}}</v-btn>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <!--add box tag dialog-->
    <v-row>
      <v-dialog
        v-model="boxDialog"
        persistent
        max-width="350"
        class="d-flex"
      >
        <v-card :loading="'primary' && boxTagLoading">
          <v-alert
            class="bindBoxAlert"
            :value="bindAlert"
            type="error"
          >{{bindBoxTagError}}</v-alert>
          <v-card-title>
            <span class="headline">{{$tc('ssr.bindBoxNo')}}</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-text-field
                :label="$tc('ssr.smallBox')"
                required
                v-model="boxItems[0].num"
              ></v-text-field>
              <v-text-field
                :label="$tc('ssr.middleBox')"
                required
                v-model="boxItems[1].num"
              ></v-text-field>
              <v-text-field
                :label="$tc('ssr.bigBox')"
                v-model="boxItems[2].num"
                required
              ></v-text-field>
              <v-text-field
                :label="$tc('ssr.extraLargeBox')"
                required
                v-model="boxItems[3].num"
              ></v-text-field>
              <v-text-field
                :label="$tc('ssr.specialBox')"
                required
                v-model="boxItems[4].num"
              ></v-text-field>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="boxDialog = false"
            >{{$tc('ssr.closeBtn')}}</v-btn>
            <v-btn
              color="blue darken-1"
              text
              :loading="boxTagSaveLoading"
              @click="bindBoxNum"
            >{{$tc('ssr.saveBtn')}}</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </div>
</template>

<script>
import { getSpecies } from '@/core/api/species.js'
import { createBoxList, getLatestBoxes } from '@/core/api/box.js'

export default {
  data() {
    return {
      specLoading: false,
      boxLoading: false,
      boxTagLoading: false,
      boxTagSaveLoading: false,
      // species
      tagDialog: false,
      speciesItems: [],
      specSeleIndex: null,
      preSpecSele: null,
      specTag: null,
      // box
      boxSeleIndex: null,
      boxDialog: false,
      weight: 352,
      bindAlert: false,
      bindBoxTagError: '',
      boxItems: [
        {
          id: '',
          type: 's',
          num: '',
          name: this.$t('ssr.smallBox')
        },
        {
          id: '',
          type: 'm',
          num: '',
          name: this.$tc('ssr.middleBox')
        },
        {
          id: '',
          type: 'l',
          num: '',
          name: this.$tc('ssr.bigBox')
        },
        {
          id: '',
          type: 'xl',
          num: '',
          name: this.$tc('ssr.extraLargeBox')
        },
        {
          id: '',
          type: 'special',
          num: '',
          name: this.$tc('ssr.specialBox')
        }
      ]
    }
  },
  props: {
    editMode: {
      type: Boolean,
      default: false
    }
  },
  mounted() {
    this.getSpeciesInfo()
    this.$on('taskReady', () => {
      this.getBoxInfo()
    })
  },
  methods: {
    record() {
      this.$emit('addRecord', {
        species_id: this.speciesItems[this.specSeleIndex].id,
        box_id: this.boxItems[this.boxSeleIndex].id,
        tag_name: this.specTag
      })
    },
    update() {
      this.$emit('updateRecord', {
        species_id: this.speciesItems[this.specSeleIndex].id,
        box_id: this.boxItems[this.boxSeleIndex].id,
        tag_name: this.specTag
      })
    },
    cancelUpdate() {
      this.$emit('cancelUpdate')
    },
    getSpeciesInfo() {
      this.specLoading = true
      getSpecies().then((response) => {
        this.specLoading = false
        this.speciesItems = response.data
      }).catch((err) => {
        this.specLoading = false
        this.tipErrorMessage(err.message)
      })
    },
    getBoxInfo(data) {
      this.boxLoading = true
      getLatestBoxes({
        task_id: this.$store.getters.taskId
      }).then((response) => {
        this.boxLoading = false
        const recordList = response.data
        if (Array.isArray(recordList)) {
          for (const record of recordList) {
            for (const local of this.boxItems) {
              if (local.type === record.box_type) {
                local.id = record.box_id
                local.num = record.box_num
              }
            }
          }
        }
      }).catch((err) => {
        this.boxLoading = false
        this.tipErrorMessage(err.message)
      })
    },
    editSpeciesTag(index) {
      this.dialog = true
    },
    okTagDialog() {
      if (this.specTag.lenght <= 0) {
        return
      }
      this.tagDialog = false
      this.preSpecSele = this.specSeleIndex
    },
    cancelTagDialog() {
      this.tagDialog = false
      if (this.preSpecSele != null) {
        this.specSeleIndex = this.preSpecSele
        this.specTag = null
      }
    },
    changeSpec(val) {
      const species = this.speciesItems[val]
      if (species.has_tag) {
        this.tagDialog = true
      } else {
        this.preSpecSele = val
        this.specTag = null
      }
    },
    showBindBoxAlert(msg) {
      const _this = this
      this.loading = false
      setTimeout(() => {
        _this.bindAlert = false
        _this.bindBoxTagError = ''
      }, 1500)
      this.bindBoxTagError = msg
      this.bindAlert = true
    },
    showBindBoxNumLoading(show) {
      this.boxTagLoading = show
      this.boxTagSaveLoading = show
    },
    bindBoxNum() {
      this.showBindBoxNumLoading(true)
      const params = []
      const boxNums = []
      for (const item of this.boxItems) {
        if (item.num) {
          if (boxNums.indexOf(item.num) !== -1) {
            // find same box num
            this.showBindBoxAlert(item.type + ' type box num ' + item.num + 'duplication')
            this.showBindBoxNumLoading(false)
            return
          }
          boxNums.push(item.num)
          item.task_id = this.$store.getters.taskId
          params.push(item)
        }
      }
      createBoxList({
        box_list: params
      }).then((response) => {
        this.showBindBoxNumLoading(false)
        this.boxDialog = false
        const binds = response.data
        for (const bind of binds) {
          for (const box of this.boxItems) {
            if (bind.box_num === box.num && bind.box_type === box.type) {
              // same box
              box.id = bind.box_id
            }
          }
        }
      }).catch((error) => {
        this.showBindBoxNumLoading(false)
        this.showBindBoxAlert(error.message)
      })
    },
    changeBox(val) {
      this.boxSeleIndex = val
    },
    tipErrorMessage(msg) {
      this.$emit('alertMessage', {
        type: 'error',
        msg: msg
      })
    }
  }
}
</script>
<style scoped>
.bindBoxAlert {
  position: absolute;
  top: 0;
  width: 100%;
}
.title-font {
  font-size: 20px;
  margin-left: 12px;
  color: rgba(16, 16, 16, 1);
}
.subtitle-font {
  height: 11px;
  color: rgba(255, 42, 42, 1);
  font-size: 8px;
  margin-left: 12px;
}
</style>
