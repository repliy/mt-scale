<template>
  <div class="">
    <v-row align="center" justify="space-around" class="">
        <v-card
          tile
          max-width=200
        >
          <v-card-text>
            <div class="weight-tag-font">总重量</div>
            <p class="display-1 weight-font">
              {{weight}} <span class="wight-unit-font">bl</span>
            </p>
          </v-card-text>
        </v-card>
        <v-card
          @click="switchCard"
          v-if="speciesShow"
          class="top-height pl-2 pr-2"
          tile
          max-width="400"
        >
          <div class="weight-tag-font">物种</div>
          <v-row
            no-gutters
            class="mt-5"
          >
            <template v-for="(item, n) in speciesData">
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
                        <p class="ma-0 species-color-tag">{{item.name}}</p>
                        <p class="ma-0 species-color-tag">{{item.weight}}bl</p>
                      </div>
                    </v-col>
                  </v-row>
                </div>
              </v-col>
              <v-responsive
                v-if="(n + 1) % 3 === 0"
                :key="`width-${n}`"
                width="100%"
              ></v-responsive>
            </template>
          </v-row>
        </v-card>
        <v-card
          @click="switchCard"
          v-else
          class="top-height  pl-2 pr-2"
          tile
          max-width="400"
        >
          <div class="weight-tag-font">箱子</div>
          <v-row
            no-gutters
            class="mt-5"
          >
            <template v-for="(item, n) in boxData">
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
                  >{{item.type}}:{{item.weight}} lb</p>
                </div>
              </v-col>
              <v-responsive
                v-if="(n + 1) % 3 === 0"
                :key="`width-${n}`"
                width="100%"
              ></v-responsive>
            </template>
          </v-row>
        </v-card>
    </v-row>
  </div>
</template>

<script>
import { StatWeight } from '@/core/api/record.js'
export default {
  name: 'RecordStatistics',
  data() {
    return {
      boxName: '大号',
      speciesShow: true,
      weight: 0,
      speciesData: [],
      boxData: []
    }
  },
  components: {},
  mounted() {
    this.$on('taskReady', (data) => {
      this.getStatInfo(data)
    })
  },
  methods: {
    switchCard() {
      this.speciesShow = !this.speciesShow
    },
    getSpeciesTagColor() {
      return 'green'
    },
    getStatInfo(params) {
      StatWeight(params).then((response) => {
        console.log(response.data)
        const retData = response.data
        this.speciesData = retData.species
        this.boxData = retData.box
        this.weight = retData.total.weight
      }).catch((error) => {
        console.log(error)
      })
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
