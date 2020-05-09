<template>
  <div class="ma-0">
    <v-row align="center">
      <v-col cols="3">
        <v-card class="mx-auto" max-width="200" tile>
          <v-card-text>
            <div class="weight-tag-font">总重量</div>
            <p class="display-1 weight-font">
              123,456 <span class="wight-unit-font">bl</span>
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="2">
        <v-card class="mx-auto" max-width="150" tile>
          <v-card-text>
            <div class="weight-tag-font">记录次数</div>
            <p class="display-1 weight-font">
              68 <span class="wight-unit-font">次</span>
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="5">
        <v-card @click="switchCard" v-if="speciesShow" class="mx-auto" tile>
          <div class="weight-tag-font">物种</div>
          <v-row no-gutters style="margin-top: 10px;">
            <template v-for="n in 9">
              <v-col :key="n">
                <v-chip class="mb-1" label color="grey lighten-5">
                  <v-icon left :color="getSpeciesTagColor()">mdi-label</v-icon>
                  <div style="height: 24px;">
                    <p class="ma-0 species-color-tag">Sea Cucumber</p>
                    <p class="ma-0 species-color-tag">537bl</p>
                  </div>
                </v-chip>
              </v-col>
              <v-responsive v-if="n === 3" :key="`width-${n}`" width="100%"></v-responsive>
            </template>
          </v-row>
        </v-card>
        <v-card @click="switchCard" v-else class="mx-auto" tile>
          <div class="weight-tag-font">箱子</div>
          <v-row no-gutters style="margin-top: 10px;">
            <template v-for="n in 6">
              <v-col :key="n">
                <v-chip class="mb-1" label color="grey lighten-5">
                  <v-icon left :color="getSpeciesTagColor()">mdi-label</v-icon>
                  <div style="height: 24px;">
                    <p class="ma-0 species-color-tag">Sea Cucumber</p>
                    <p class="ma-0 species-color-tag">537bl</p>
                  </div>
                </v-chip>
              </v-col>
              <v-responsive v-if="n === 2" :key="`width-${n}`" width="100%"></v-responsive>
            </template>
          </v-row>
        </v-card>
      </v-col>
      <v-col cols=2 align="center">
        <v-btn class="complete-btn" height="63">记录完成</v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'RecordStatistics',
  data () {
    return {
      speciesShow: false
    }
  },
  components: {},
  methods: {
    switchCard () {
      this.speciesShow = !this.speciesShow
    },
    getSpeciesTagColor () {
      return "green"
    }
  },
}
</script>
<style scoped>
.weight-tag-font {
  font-size: 20px;
  height: 24px;
  color: rgba(16, 16, 16, 1);
}
.weight-font {
  font-size: 60px;
  color: rgba(16, 16, 16, 1);
  height: 24px;
}
.wight-unit-font {
  font-size: 18px;
  color: rgba(16, 16, 16, 1);
  height: 20px;
}
.species-color-tag {
  height: 12px;
  font-size: 6px;
}
.complete-btn {
  font-size: 20px;
}
</style>