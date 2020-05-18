<template>
  <div style="margin-right: 12px;height: 100%;">
    <v-row no-gutters>
      <v-col>
        <div class="tag_font">物种选择: {{specTag}}</div>
        <p style="height: 11px;color: rgba(255, 42, 42, 1);font-size: 8px;">请选择物种类型</p>
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
                        color="red"
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
        <div class="tag_font d-inline">箱子选择:</div>
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
        <p style="height: 11px;color: rgba(255, 42, 42, 1);font-size: 8px;">请选择箱子类型</p>
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
          <template v-for="(box, n) in BoxItems">
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
          <v-card-title class="headline">Tag 信息</v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="specTag"
                  clearable
                  label="请输入Tag编号"
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
            >Disagree</v-btn>
            <v-btn
              color="green darken-1"
              text
              @click="okTagDialog"
            >Agree</v-btn>
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
          color="primary"
          text-color="white"
          @click="record"
        >记录</v-btn>
      </v-col>
    </v-row>
    <!--add box tag dialog-->
    <v-row>
      <v-dialog
        v-model="boxDialog"
        persistent
        max-width="350"
      >
        <v-card>
          <v-alert class="bindBoxAlert" :value="bindAlert" type="error">{{bindBoxTagError}}</v-alert>
          <v-card-title>
            <span class="headline">绑定箱子号</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-text-field
                label="小号*"
                required
                v-model="BoxItems[0].num"
              ></v-text-field>
              <v-text-field
                label="中号*"
                required
                v-model="BoxItems[1].num"
              ></v-text-field>
              <v-text-field
                label="大号*"
                v-model="BoxItems[2].num"
                required
              ></v-text-field>
              <v-text-field
                label="特大*"
                required
                v-model="BoxItems[3].num"
              ></v-text-field>
              <v-text-field
                label="特殊*"
                required
                v-model="BoxItems[4].num"
              ></v-text-field>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="boxDialog = false"
            >Close</v-btn>
            <v-btn
              color="blue darken-1"
              text
              @click="bindBoxNum"
            >Save</v-btn>
          </v-card-actions>
        </v-card>
        <v-overlay :value="loading">
          <v-progress-circular
            indeterminate
            size="64"
          ></v-progress-circular>
        </v-overlay>
      </v-dialog>
    </v-row>
  </div>
</template>

<script>
import { getSpecies } from '@/core/api/species.js'
import { createBoxList } from '@/core/api/box.js'

export default {
  data: () => ({
    loading: false,
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
    BoxItems: [
      {
        id: '',
        type: 's',
        num: '',
        name: '小号'
      },
      {
        id: '',
        type: 'm',
        num: '',
        name: '中号'
      },
      {
        id: '',
        type: 'l',
        num: '',
        name: '大号'
      },
      {
        id: '',
        type: 'xl',
        num: '',
        name: '超大号'
      },
      {
        id: '',
        type: 'special',
        num: '',
        name: '特殊'
      }
    ]
  }),
  mounted() {
    this.getSpeciesInfo()
  },
  methods: {
    record() {
      const speciesId = this.speciesItems[this.specSeleIndex].id
      const boxId = this.BoxItems[this.boxSeleIndex].id
      this.$emit('weightRecord', {
        species_id: speciesId,
        box_id: boxId
      })
    },
    getSpeciesInfo() {
      getSpecies().then((response) => {
        this.speciesItems = response.data
      }).catch((err) => { console.log(err) })
    },
    editSpeciesTag(index) {
      console.log(index)
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
        console.log('timeout functioin')
        _this.bindAlert = false
        _this.bindBoxTagError = ''
      }, 1500)
      this.bindBoxTagError = msg
      this.bindAlert = true
    },
    bindBoxNum() {
      this.loading = true
      const params = []
      const boxNums = []
      for (const item of this.BoxItems) {
        if (item.num) {
          if (boxNums.indexOf(item.num) !== -1) {
            // find same box num
            this.showBindBoxAlert(item.type + ' type box num ' + item.num + 'duplication')
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
        this.loading = false
        this.boxDialog = false
        const binds = response.data
        for (const bind of binds) {
          for (const box of this.BoxItems) {
            if (bind.box_num === box.num && bind.box_type === box.type) {
              // same box
              box.id = bind.box_id
            }
          }
        }
      }).catch((error) => {
        this.showBindBoxAlert(error.msg)
      })
    },
    changeBox(val) {
      this.boxSeleIndex = val
      console.log(val)
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
</style>
