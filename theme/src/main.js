import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import marked from "./utils/marked"
import api from "./api"

import "normalize.css"
import "@primer/css"


Vue.use(marked)
Vue.use(api)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
