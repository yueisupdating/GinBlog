import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

import Index from '../components/admin/index.vue'
import UserList from '../components/admin/user/UserList.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    component: Login
  },
  {
    path: '/admin',
    component: Admin,
    children: [
      {
        path: '/admin/index',
        component: Index,
        meta: {
          title: 'GinBlog 后台管理页面'
        }
      },
      {
        path: '/admin/userList',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      }

    ]
  }
]

const router = new VueRouter({
  routes
})
export default router

router.beforeEach((to, from, next) => {
  if (to.path === '/login') {
    return next()
  }
  if (!window.sessionStorage.getItem('token')) {
    next('/login')
  } else {
    next()
  }
})
