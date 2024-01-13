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
      redirect:"/dashboard",
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
        ,
        {
          path: '/charts',
          name: 'charts',
          component: () => import('../views/ChartsView.vue')
        }
      ]
    }
  ]
})
router.beforeEach((to, from, next) => {
  const access_token = sessionStorage.getItem('access_token')
  if (to.meta.requiresAuth && !access_token ) {
    next('/login')
  } else {
    next()
  }
  NProgress.start()
})


router.afterEach(() => {
  NProgress.done()
})
export default router
