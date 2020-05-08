<template>
  <v-container style="height: 100%;" class=" green lighten-5 ma-0">
    <v-label class="">统计记录</v-label>
    <v-row class="red">
      <v-col cols="3">
        <v-card class="mx-auto" max-width="344" outlined>
          <v-card-text>
            <div>重量</div>
            <p class="display-1 text--primary">
              123,456 bl
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="2">
        <v-card class="mx-auto" max-width="344" outlined>
          <v-card-text>
            <div>记录次数</div>
            <p class="display-1 text--primary">
              68 次
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card class="mx-auto" max-width="344" outlined>
          <v-card-text>
            <div>物种</div>
            <v-row no-gutters>
              <template v-for="n in 6">
                <v-col :key="n" cols="6">
                  <v-card class="pa-1" outlined tile>
                    Column
                  </v-card>
                </v-col>
                <v-responsive v-if="n === 2" :key="`width-${n}`" width="100%"></v-responsive>
              </template>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="2">
        <v-card class="mx-auto" max-width="344" outlined>
          <v-card-text>
            <div>箱子</div>
            <template v-for="n in 4">
              <v-row :key="n">
                <v-card class="pa-1" outlined tile>
                  Column
                </v-card>
              </v-row>
            </template>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols=2>
        <v-card>
          <v-btn>记录完成</v-btn>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'RecordStatistics',
  components: {},
  methods: {},
}
</script>