import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import Vue18n from 'vue-i18n'

Vue.use(Vue18n)

const i18n = new Vue18n({
  locale: 'zh-CN',
  messages: {
    'zh-CN': require('./core/lang/cn'),
    'en-US': require('./core/lang/en')
  }
})

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  i18n,
  render: h => h(App)
}).$mount('#app')
