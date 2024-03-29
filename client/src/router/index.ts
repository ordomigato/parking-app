import { createRouter, createWebHistory } from 'vue-router'
import AuthView from '../views/AuthView.vue'
import { routeNames } from './routeNames'
import { getStatus } from '@/services/account.service'
import { getWorkspaces } from '@/services/workspace.service'
import { useWorkspaceStore } from '@/stores/workspaceStore'
import { getForm } from '@/services/form.service'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/admin',
      name: routeNames.auth,
      component: AuthView
    },
    {
      path: '/verify/email',
      name: routeNames.verifyEmail,
      component: () => import('../views/VerifyEmail.vue')
    },
    {
      path: '/:workspacePath/:formPath',
      name: routeNames.formLanding,
      component: () => import('../views/FormLanding.vue')
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
          path: '/dashboard/profile',
          name: routeNames.profile,
          component: () => import('../views/UserProfile.vue'),
        },
        {
          path: '/dashboard/workspace',
          name: routeNames.workspaces,
          component: () => import('../views/WorkspacesView.vue')
        },
        {
          path: '/dashboard/workspace/:id',
          name: routeNames.workspaceSettings,
          component: () => import('../views/WorkspaceSettings.vue')
        },
        {
          path: '/dashboard/workspace/create',
          name: routeNames.createWorkspace,
          component: () => import('../views/CreateWorkspace.vue')
        },
        {
          path: '/dashboard/form',
          name: routeNames.forms,
          component: () => import('../views/FormsView.vue')
        },
        {
          path: '/dashboard/form/:id',
          name: routeNames.form,
          component: () => import('../views/FormView.vue')
        },
        {
          path: '/dashboard/form/create',
          name: routeNames.createForm,
          component: () => import('../views/CreateForm.vue')
        },
        {
          path: '/dashboard/form/:id/permit/:permitId',
          name: routeNames.permit,
          component: () => import('../views/PermitView.vue')
        }
      ]
    }
  ]
})

router.beforeEach(async(to, from, next) => {
  if (to.name === routeNames.formLanding) {
    return next()
  }
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

    // Reset current form
    workspaceStore.setCurrentForm(null)
    // get form if on form specific path
    if (to.name === routeNames.form) {
      if (!workspaceStore.currentWorkspace?.workspace_id) {
        router.push({ name: routeNames.workspaces})
      } else {
        const form = await getForm(workspaceStore.currentWorkspace.workspace_id, to.params.id as string)
        workspaceStore.setCurrentForm(form)
      }
    }
  }
  next()
})

export default router
