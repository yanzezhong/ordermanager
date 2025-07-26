import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/login', component: LoginView },
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/dept',
      name: 'dept',
      component: () => import('../views/dept/DeptView.vue')
    },
    {
      path: '/menu',
      name: 'menu',
      component: () => import('../views/menu/MenuView.vue')
    },
    {
      path: '/order',
      name: 'order',
      component: () => import('../views/order/OrderView.vue')
    },
    {
      path: '/product',
      name: 'product',
      component: () => import('../views/product/ProductView.vue')
    }
  ]
})

// 全局前置守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  
  // 如果访问登录页，直接放行
  if (to.path === '/login') {
    // 如果已经登录，跳转到首页
    if (token) {
      next('/')
    } else {
      next()
    }
    return
  }
  
  // 检查是否已登录
  if (!token) {
    // 未登录，跳转到登录页，并记录原本要访问的页面
    next({
      path: '/login',
      query: { redirect: to.fullPath }
    })
  } else {
    // 已登录，放行
    next()
  }
})

export default router
