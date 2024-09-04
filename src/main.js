import Vue from 'vue'
import App from './App.vue'
import router from './router'
import '../src/plugin/antui'
import axios from 'axios'
import './assets/css/style.css'

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

Vue.config.productionTip = false

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
