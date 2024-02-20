import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import type { IForm, IWorkspace } from '@/types'

const STORE_NAME = 'workspace'

export const useWorkspaceStore = defineStore(STORE_NAME, () => {
  const workspaces: Ref<IWorkspace[]> = ref([])
  const currentWorkspace: Ref<IWorkspace | null> = ref(null)
  const currentForm: Ref<IForm | null> = ref(null)
  function setWorkspaces(ws: IWorkspace[]) {
    workspaces.value = ws
  }

  function setActiveWorkspace(workspace: IWorkspace | null) {
    currentWorkspace.value = workspace

    updateLocalState()
  }

  function setCurrentForm(form: IForm | null) {
    currentForm.value = form
  }

  function updateLocalState() {
    localStorage.setItem(STORE_NAME, JSON.stringify({
        currentWorkspace: currentWorkspace.value
    }))
  }

  function setState() {
    const localState = localStorage.getItem(STORE_NAME)
    if (localState) {
        currentWorkspace.value = JSON.parse(localState).currentWorkspace
    } else {
        currentWorkspace.value = null
    }
  }

  setState()

  return {
    workspaces,
    currentWorkspace,
    currentForm,
    setWorkspaces,
    setActiveWorkspace,
    setCurrentForm,
  }
})
