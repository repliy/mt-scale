<template>
  <v-data-table hide-default-header calculate-widths hide-default-footer height=300 @click:row="itemExpand" :headers="headers" :items="desserts" single-expand :expanded.sync="expanded" item-key="name">
    <template v-slot:expanded-item="{ headers, item }">
      <!-- <td :colspan="headers.length">More info about {{ item.name }}</td> -->
      <v-btn></v-btn>
    </template>
    <template v-slot:item.flag="{ item }">
      <v-chip :color="getColor(item.name)" dark></v-chip>
    </template>
    <template v-slot:item.name="{ item }">
      <div style="font-size: 5px;width: 50px;">{{ item.name }}</div>
    </template>
  </v-data-table>
</template>

<script>
export default {
  data () {
    return {
      expanded: [],
      headers: [
        {
          text: 'index',
          align: 'start',
          sortable: false,
          value: 'index',
        },
        { text: 'flag', value: 'flag' },
        { text: 'name', class: 'tab-font', width: 50, value: 'name' },
        { text: 'weight', value: 'weight' },
        { text: 'unit', value: 'unit' }
      ],
      desserts: [
        {
          index: 1,
          flag: '',
          name: 'Sea Cucumber',
          weight: '342',
          unit: 'lb'
        },
        {
          index: 2,
          flag: '',
          name: 'Snow crab',
          weight: '195',
          unit: 'lb'
        },
        {
          index: 3,
          flag: '',
          name: 'Bluefin tuna',
          weight: '238',
          unit: 'lb'
        },
      ],
    }
  },
  methods: {
    itemExpand(item, value) {
      value.expand(!value.isExpanded)
    },
    getColor (name) {
        return 'red'
    }
  }
}
</script>
<style scoped>
.tab-font {
  size: 5px;
}
</style>