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
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/app',
      name: 'app',
      children: [
        {
          path: 'terminal',
          name: 'terminal',
          component: () => import('../views/terminal/TerminalView.vue')
        },
        {
          path: 'setting',
          name: 'setting',
          component: () => import('../views/setting/SettingView.vue')
        },
        {
          path: 'finder',
          name: 'finder',
          component: () => import('../views/finder/FinderView.vue')
        },
        {
          path: 'app-store',
          name: 'app-store',
          component: () => import('../views/app-store/AppStoreView.vue')
        }
      ]
    }
  ]
})

export default router
