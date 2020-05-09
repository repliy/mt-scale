<template>
  <v-container class="bk-width">
    <v-row no-gutters>
      <template v-for="n in 3">
        <v-col :key="n">
          <v-btn class="pa-2 row-height" outlined x-large block tile>
            {{n}}
          </v-btn>
        </v-col>
      </template>
    </v-row>
    <v-row no-gutters>
      <template v-for="n in 3">
        <v-col :key="n">
          <v-btn class="pa-2 row-height" outlined x-large block tile>
            {{n+3}}
          </v-btn>
        </v-col>
      </template>
    </v-row>
    <v-row no-gutters>
      <template v-for="n in 3">
        <v-col :key="n">
          <v-btn class="pa-2 row-height" outlined x-large block tile>
            {{n+6}}
          </v-btn>
        </v-col>
      </template>
    </v-row>
    <v-row no-gutters>
      <v-col cols=8>
        <v-btn class="pa-2 row-height" outlined x-large block tile>
          0
        </v-btn>
      </v-col>
      <v-col>
        <v-btn class="pa-2 row-height" outlined x-large block tile>
          Del
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>

export default {
  name: 'Keyboard',
  data: () => ({
    index: 69,
    weight: 352
  }),
  components: {
  },
  methods: {},
}
</script>
<style scoped>
.row-height {
  height: 80px;
}
.bk-width {
  width: 244px;
}
</style>