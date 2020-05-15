<template>
  <div style="margin-right: 12px;height: 100%;">
    <v-row no-gutters>
      <v-col>
        <div class="tag_font">物种选择:</div>
        <p style="height: 11px;color: rgba(255, 42, 42, 1);font-size: 8px;">请选择物种类型</p>
      </v-col>
    </v-row>
    <!--species select-->
    <v-item-group
      v-model="speciesSelected"
      @change="changeSpeciesSelect"
    >
      <v-container>
        <v-row no-gutters>
          <template v-for="n in 7">
            <v-col
              :key="n"
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
                      >hello </p>
                    </v-col>
                  </v-row>
                </v-card>
              </v-item>
            </v-col>
            <v-responsive
              v-if="n % 2 === 0"
              :key="`width-${n}`"
              width="100%"
            ></v-responsive>
          </template>
        </v-row>
      </v-container>
    </v-item-group>
    <v-divider darkr></v-divider>
    <v-row no-gutters class="mt-5">
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
      v-model="boxSelected"
      @change="changeBoxSelect"
    >
      <v-container>
        <v-row no-gutters>
          <template v-for="n in 7">
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
                >
                  <p
                    class="mb-0 pl-2"
                    style="font-size: 20px;word-break:break-all;"
                  >hello </p>
                  <p
                    class="mb-0 pr-2"
                    style="font-size: 16px;word-break:break-all;text-align: right;"
                  >#100</p>
                </v-card>
              </v-item>
            </v-col>
            <v-responsive
              v-if="n % 2 === 0"
              :key="`width-${n}`"
              width="100%"
            ></v-responsive>
          </template>
        </v-row>
      </v-container>
    </v-item-group>
    <v-row justify="center">
      <v-dialog
        v-model="dialog"
        persistent
        max-width="350"
      >
        <v-card>
          <v-card-title class="headline">Tag 信息</v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="speceiesTag"
                  clearable
                  label="请输入Tag编号"
                  type="text"
                >
                  <template v-slot:append-outer>
                    <v-menu
                      style="top: -12px"
                      offset-y
                    >
                      <template v-slot:activator="{ on }">
                        <v-btn
                          color="primary"
                          v-on="on"
                        >
                          <v-icon left>mdi-menu</v-icon>
                          Menu
                        </v-btn>
                      </template>
                      <v-list>
                        <v-list-item
                          v-for="(item, index) in items"
                          :key="index"
                          @click="selectSpeciesTag"
                        >
                          <v-list-item-title>{{ item.title }}</v-list-item-title>
                        </v-list-item>
                      </v-list>
                    </v-menu>
                  </template>
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
              @click="dialog = false"
            >Disagree</v-btn>
            <v-btn
              color="green darken-1"
              text
              @click="dialog = false"
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
    <v-row>
      <v-dialog
        v-model="boxDialog"
        persistent
        max-width="350"
      >
        <v-card>
          <v-card-title>
            <span class="headline">绑定箱子号</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-text-field
                label="小号*"
                required
              ></v-text-field>
              <v-text-field
                label="中号*"
                required
              ></v-text-field>
              <v-text-field
                label="大号*"
                required
              ></v-text-field>
              <v-text-field
                label="特大*"
                required
              ></v-text-field>
              <v-text-field
                label="特殊*"
                required
              ></v-text-field>
            </v-container>
            <small>*indicates required field</small>
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
              @click="dialog = false"
            >Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </div>
</template>

<script>
export default {
  data: () => ({
    boxSelected: null,
    speciesSelected: null,
    items: [{
      title: '#tag1'
    }],
    dialog: false,
    boxDialog: false,
    boxType: '大号',
    boxTag: '#999',
    speceiesTag: '',
    weight: 352
  }),
  methods: {
    record() {
      this.boxSelected = 1
    },
    editSpeciesTag(index) {
      console.log(index)
      this.dialog = true
    },
    selectSpeciesTag() { },
    changeSpeciesSelect(val) {

    },
    changeBoxSelect(val) {
      console.log(val)
    }
  }
}
</script>
<style scoped>
</style>
