<template>
  <v-data-table
    class="pl-2"
    hide-default-header
    hide-default-footer
    height=480
    @click:row="itemExpand"
    :headers="headers"
    :items="desserts"
    single-expand
    :expanded.sync="expanded"
    item-key="name"
    calculate-widths
  >
    <template v-slot:expanded-item="{ item }">
      <td :colspan="100">
        <v-row>
          <v-col>
            <v-btn tile block class="d-inline pa-2 deep-purple white--text" @click="itemEdit(item)">编辑</v-btn>
          </v-col>
          <v-col>
            <v-btn tile block class="d-inline pa-2 red white--text" @click="itemDel(item)">删除</v-btn>
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
              :color="getColor(item.name)"
            >mdi-label</v-icon>
          </v-col>
          <v-col>
            <div style="height: 37px;">
              <p class="ma-0 species-color-tag">Sea Cucumber</p>
              <p class="ma-0 species-color-tag">537bl</p>
            </div>
          </v-col>
        </v-row>
      </div>
    </template>
    <template v-slot:item.weight="{ item }">
      <div class="species-color-tag">{{item.weight}}bl</div>
    </template>
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
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
      desserts: [
        {
          index: 1,
          name: 'Sea Cucumber',
          weight: '342'
        },
        {
          index: 2,
          name: 'Snow crab',
          weight: '195'
        },
        {
          index: 3,
          name: 'Bluefin tuna',
          weight: '238'
        }
      ]
    }
  },
  methods: {
    itemExpand(item, value) {
      value.expand(!value.isExpanded)
    },
    getColor(name) {
      return 'red'
    },
    itemEdit(item) {
      console.log('item del index', item.index)
    },
    itemDel(item) {

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
