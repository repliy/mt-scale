<template>
  <div>
    <v-container
      class="pg-height"
      fluid
    >
      <AppBar></AppBar>
      <v-row
        align="center"
        justify="center"
        style="height: 100%;"
      >
        <v-col
          cols="12"
          md="5"
        >
          <v-card>
            <v-toolbar
              color="primary"
              dark
              flat
            >
              <v-toolbar-title>{{$tc('ssr.loginTitle')}}</v-toolbar-title>
              <v-spacer></v-spacer>
            </v-toolbar>
            <v-card-text>
              <v-form>
                <v-text-field
                  v-model="username"
                  label="Name"
                  name="name"
                  type="text"
                  clearable
                ></v-text-field>
                <v-text-field
                  v-model="password"
                  label="Password"
                  name="password"
                  type="password"
                  clearable
                ></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="primary"
                :loading="btnLoading"
                @click="userLogin"
              >{{$tc('ssr.loginBtn')}}</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
    <!--error message-->
    <v-alert
      :type="alertType"
      class="pg-alert"
      transition="slide-y-transition"
      :value="showAlert"
    >
      {{alertMessage}}
    </v-alert>
  </div>
</template>
<script>
import AppBar from '@/components/AppBar.vue'
import { login } from '@/core/api/user.js'
export default {
  name: 'Login',
  data: () => ({
    btnLoading: false,
    showAlert: false,
    alertMessage: '',
    alertType: 'error',
    username: '',
    password: ''
  }),
  components: {
    AppBar
  },
  props: {
    source: String
  },
  mounted() { },
  methods: {
    userLogin() {
      this.btnLoading = true
      login({
        username: this.username,
        password: this.password
      }).then((response) => {
        this.btnLoading = false
        this.$store.commit('updateAuthToken', response.data.token)
        this.$store.commit('SET_USERNAME', response.data.username)
        this.$router.push('/recording')
      }).catch((err) => {
        this.btnLoading = false
        this.tipErrorMessage(err.message)
      })
    },
    tipErrorMessage(msg) {
      const _this = this
      this.alertMessage = msg
      this.showAlert = true
      setTimeout(() => {
        _this.showAlert = false
      }, 2000)
    }
  }
}
</script>
<style scoped>
.pg-height {
  height: calc(100vh - 64px);
}
.pg-alert {
  width: 100vw;
  position: absolute;
  top: 0;
}
</style>
