<template>
  <div class="">
    <v-progress-linear
      indeterminate
      :active="loading"
      height="2"
    >
    </v-progress-linear>
    <v-row
      align="center"
      justify="space-around"
      class=""
    >
      <v-card
        tile
        max-width=200
      >
        <v-card-text>
          <div class="weight-tag-font">{{$tc('ssr.totalWeight')}}</div>
          <p class="display-1 weight-font">
            {{displayWeight(weight)}} <span class="wight-unit-font">{{$tc('ssr.pound')}}</span>
          </p>
        </v-card-text>
      </v-card>
      <v-card
        @click="switchCard"
        v-if="speciesShow"
        class="top-height pl-2 pr-2"
        tile
        min-width=400
      >
        <div class="weight-tag-font mt-3">{{$tc('ssr.species')}}</div>
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
                      :color="item.color"
                    >mdi-label</v-icon>
                  </v-col>
                  <v-col>
                    <div style="height: 37px;">
                      <p class="ma-0 species-color-tag">{{item.name}}</p>
                      <p class="ma-0 species-color-tag">{{displayWeight(item.weight)}} {{$tc('ssr.pound')}}</p>
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
        class="top-height pl-2 pr-2"
        tile
        min-width=400
      >
        <div class="weight-tag-font mt-3">{{$tc('ssr.box')}}</div>
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
                >{{item.type}}:{{displayWeight(item.weight)}} {{$tc('ssr.pound')}}</p>
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
import { toThousands } from '@/utils/func.js'
export default {
  name: 'RecordStatistics',
  data() {
    return {
      loading: false,
      speciesShow: true,
      weight: 0,
      speciesData: [],
      boxData: [
        {
          type: 's',
          weight: 0
        },
        {
          type: 'm',
          weight: 0
        },
        {
          type: 'l',
          weight: 0
        },
        {
          type: 'xl',
          weight: 0
        },
        {
          type: 'special',
          weight: 0
        }
      ]
    }
  },
  components: {},
  mounted() {
    this.$on('taskReady', () => {
      this.getStatInfo()
    })
    this.$on('refreshData', () => {
      this.getStatInfo()
    })
  },
  methods: {
    displayWeight(weight) {
      return toThousands(weight)
    },
    switchCard() {
      this.speciesShow = !this.speciesShow
    },
    getStatInfo() {
      this.loading = true
      StatWeight({
        task_id: this.$store.getters.taskId
      }).then((response) => {
        this.loading = false
        const retData = response.data
        this.speciesData = retData.species
        if (retData.box) {
          this.boxData = retData.box
        }
        this.weight = retData.total.weight
      }).catch((err) => {
        this.loading = false
        this.tipErrorMessage(err.message)
      })
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
  height: 18px;
  font-size: 6px;
}
.top-height {
  height: 25vh;
}
</style>
