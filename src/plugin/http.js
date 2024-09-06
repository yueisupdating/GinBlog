import Vue from 'vue'
import axios from 'axios'

const baseURL = 'http://localhost:3000/api/v1'
axios.defaults.baseURL = baseURL

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

export { baseURL }
