<template>
  <div class="bk-width">
    <v-row no-gutters>
      <template v-for="n in 9">
        <v-col :key="n">
          <v-btn
            class="pa-2"
            height="70"
            x-large
            block
            tile
            color="grey lighten-2"
            @click="numberClick(n)"
          >
            {{n}}
          </v-btn>
        </v-col>
        <v-responsive
          v-if="(n) % 3 === 0"
          :key="`width-${n}`"
          width="100%"
        ></v-responsive>
      </template>
    </v-row>
    <v-row no-gutters>
      <v-col cols=8>
        <v-btn
          class="pa-2"
          height="70"
          x-large
          block
          tile
          color="grey lighten-2"
          @click="zeroClick()"
        >
          0
        </v-btn>
      </v-col>
      <v-col>
        <v-btn
          class="pa-2"
          height="70"
          x-large
          block
          tile
          color="grey lighten-2"
          @click="deleteClick()"
        >
          Del
        </v-btn>
      </v-col>
    </v-row>
  </div>
</template>
<script>

export default {
  name: 'Keyboard',
  data: () => ({
  }),
  components: {
  },
  methods: {
    numberClick(val) {
      this.emitKeyboardClick(val)
    },
    zeroClick() {
      this.emitKeyboardClick(0)
    },
    deleteClick() {
      this.emitKeyboardClick(-1)
    },
    emitKeyboardClick(val) {
      this.$emit('kbClick', {
        num: val
      })
    }
  }
}
</script>
<style scoped>
.bk-width {
  width: 244px;
}
</style>
