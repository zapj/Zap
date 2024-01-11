import { createRouter, createWebHistory } from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
const Layout = () => import('@/layout/Layout.vue')
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:'/login',
      name:'login',
      component:()=> import('../views/LoginView.vue')
    },
    {
      path: '/',
      name: 'home',
      component: Layout,
      meta:{
        "requiresAuth":true
      },
      children:[
        {
          path: '/dashboard',
          name: 'dashboard',
          component: ()=> import('../views/DashboardView.vue')
        },
        {
          path: '/about',
          name: 'about',
          component: () => import('../views/AboutView.vue')
        },
        {
          path: '/table',
          name: 'table',
          component: () => import('../views/TableView.vue')
        },
        {
          path: '/console',
          name: 'console',
          component: () => import('../views/ConsoleView.vue')
        }
      ]
    }
  ]
})
router.beforeEach((to, from, next) => {
  const isLoggedIn = localStorage.getItem('zap-token')
  if (to.meta.requiresAuth && !isLoggedIn) {
    next('/login')
  } else {
    next()
  }
  NProgress.start()
})

// 页面加载成功之后，关闭进度条
router.afterEach(() => {
  NProgress.done()
})
export default router
