import { createRouter, createWebHistory } from 'vue-router'
import AuthView from '../views/AuthView.vue'
import { routeNames } from './routeNames'
import { getStatus } from '@/services/account.service'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: routeNames.auth,
      component: AuthView
    },
    {
      path: '/dashboard',
      name: routeNames.dashboard,
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../components/layouts/dashboard-layout.vue'),
      children: [
        {
          path: '/dashboard',
          name: routeNames.dashboard,
          component: () => import('../views/OverviewView.vue'),
        },
        {
          path: '/dashboard/overview',
          name: routeNames.overview,
          component: () => import('../views/OverviewView.vue'),
        },
        {
          path: '/dashboard/workspace/create',
          name: routeNames.createWorkspace,
          component: () => import('../views/CreateWorkspace.vue')
        },
      ]
    }
  ]
})

router.beforeEach(async(to, from, next) => {
  // add basic auth guarding
  const isLoggedIn = await getStatus()
  if (to.name === routeNames.auth && isLoggedIn) {
    router.push({ name: routeNames.dashboard })
  }
  if (to.name !== routeNames.auth && !isLoggedIn) {
    router.push({ name: routeNames.auth })
  }
  next()
})

export default router
