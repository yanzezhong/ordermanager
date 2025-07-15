import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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

export default router
