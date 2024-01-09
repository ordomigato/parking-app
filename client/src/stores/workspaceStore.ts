import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import type { IWorkspace } from '@/types'

export const useWorkspaceStore = defineStore('workspace', () => {
  const workspaces: Ref<IWorkspace[]> = ref([])
  const currentWorkspace: Ref<IWorkspace | null> = ref(null)
  function setWorkspaces(ws: IWorkspace[]) {
    workspaces.value = ws
  }
  function setActiveWorkspace(workspace: IWorkspace) {
    currentWorkspace.value = workspace
  }

  return { workspaces, currentWorkspace, setWorkspaces, setActiveWorkspace }
})
