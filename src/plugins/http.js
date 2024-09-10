import Vue from 'vue'
import axios from 'axios'

const baseURL = 'http://localhost:3000/api/v1'
axios.defaults.baseURL = baseURL

Vue.prototype.$http = axios
