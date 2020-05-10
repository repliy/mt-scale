<template>
  <div class="ma-0">
    <v-row align="center">
      <v-col cols="3">
        <v-card
          class="mx-auto"
          max-width="200"
          tile
        >
          <v-card-text>
            <div class="weight-tag-font">总重量</div>
            <p class="display-1 weight-font">
              123,456 <span class="wight-unit-font">bl</span>
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="2">
        <v-card
          class="mx-auto"
          max-width="150"
          tile
        >
          <v-card-text>
            <div class="weight-tag-font">记录次数</div>
            <p class="display-1 weight-font">
              68 <span class="wight-unit-font">次</span>
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="5">
        <v-card
          @click="switchCard"
          v-if="speciesShow"
          class="mx-auto top-height"
          tile
        >
          <div class="weight-tag-font">物种</div>
          <v-row
            no-gutters
            class="mt-5"
          >
            <template v-for="n in 7">
              <v-col
                :key="n"
                cols="4"
              >
                <div
                  style="with: 107px;height: 37px;border: solid 1px black;border-radius: 6px;"
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
                        :color="getSpeciesTagColor()"
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
              </v-col>
              <v-responsive
                v-if="n % 3 === 0"
                :key="`width-${n}`"
                width="100%"
              ></v-responsive>
            </template>
          </v-row>
        </v-card>
        <v-card
          @click="switchCard"
          v-else
          class="mx-auto top-height"
          tile
        >
          <div class="weight-tag-font">箱子</div>
          <v-row
            no-gutters
            class="mt-5"
          >
            <template v-for="n in 5">
              <v-col
                :key="n"
                cols="4"
              >
                <div
                  style="with: 107px;height: 37px;border: solid 1px black;border-radius: 6px;"
                  color="grey lighten-5"
                  class="ma-1"
                >
                  <p
                    style="line-height: 37px;"
                    class="ml-2"
                  >{{boxName}}:{{weight}}lb</p>
                </div>
              </v-col>
              <v-responsive
                v-if="n % 3 === 0"
                :key="`width-${n}`"
                width="100%"
              ></v-responsive>
            </template>
          </v-row>
        </v-card>
      </v-col>
      <v-col
        cols=2
        align="center"
      >
        <v-btn
          height="63"
          color="primary"
          text-color="white"
        >记录完成</v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'RecordStatistics',
  data() {
    return {
      boxName: '大号',
      weight: 2108,
      speciesShow: false
    }
  },
  components: {},
  methods: {
    switchCard() {
      this.speciesShow = !this.speciesShow
    },
    getSpeciesTagColor() {
      return 'green'
    }
  }
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
.top-height {
  height: 25vh;
}
</style>
