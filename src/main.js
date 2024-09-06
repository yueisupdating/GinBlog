import Vue from 'vue'
import App from './App.vue'
import router from './router'
import '../src/plugin/antui'
import './assets/css/style.css'
import '../src/plugin/http'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
