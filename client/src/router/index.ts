import { createRouter, createWebHistory } from 'vue-router'
import AuthView from '../views/AuthView.vue'
import { routeNames } from './routeNames'
import { getStatus } from '@/services/account.service'
import { getWorkspaces } from '@/services/workspace.service'
import { useWorkspaceStore } from '@/stores/workspaceStore'

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
          path: '/dashboard/workspace',
          name: routeNames.workspaces,
          component: () => import('../views/WorkspacesView.vue')
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
  if (isLoggedIn) {
    // get workspaces
    const workspaceStore = useWorkspaceStore()
    const wp = await getWorkspaces()
    workspaceStore.setWorkspaces(wp)
  }
  next()
})

export default router
