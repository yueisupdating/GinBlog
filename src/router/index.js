import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ArticleList from '../components/Article/ArticleList.vue'
import Details from '../components/Article/Details.vue'
import Search from '../components/Article/Search.vue'
import CateArt from '../components/Article/CateArt.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    children: [
      {
        path: '/',
        component: ArticleList
      },
      {
        path: '/detail/:id',
        component: Details
      },
      {
        path: '/search/:title',
        component: Search
      },
      {
        path: '/category/:id',
        component: CateArt
      }
    ]
  }

]

const router = new VueRouter({
  routes
})

export default router
