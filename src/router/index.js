import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

import Index from '../components/index.vue'
import UserList from '../components/user/UserList.vue'
import CategoryList from '../components/cate/CategoryList.vue'
import ArticleList from '../components/article/ArticleList.vue'
import AddArticle from '../components/article/AddArticle.vue'
import Profile from '../components/user/Profile.vue'

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
      },
      {
        path: '/admin/CategoryList',
        component: CategoryList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: '/admin/articleList',
        component: ArticleList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: '/admin/article/add',
        component: AddArticle,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: '/admin/article/update/:id',
        component: AddArticle,
        meta: {
          title: '编辑文章'
        }
      },
      {
        path: '/admin/profile',
        component: Profile,
        meta: {
          title: '用户信息管理'
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
