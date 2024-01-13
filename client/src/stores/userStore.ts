import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import type { IClient } from '@/types'

const STORE_NAME = 'user'

export const useUserStore = defineStore(STORE_NAME, () => {
  const user: Ref<IClient | null> = ref(null)
  const prevUserEmail: Ref<string> = ref('')

  function setUser(u: IClient | null) {
    user.value = u

    if (u) {
        prevUserEmail.value = u.username
    }
    updateLocalState()
  }
  function updateLocalState() {
    localStorage.setItem(STORE_NAME, JSON.stringify({
        user: user.value,
        prevUserEmail: prevUserEmail.value
    }))
  }

  function setState() {
    const localState = localStorage.getItem(STORE_NAME)
    if (localState) {
        user.value = JSON.parse(localState).user
        prevUserEmail.value = JSON.parse(localState).prevUserEmail
    } else {
        user.value = null
        prevUserEmail.value = ''
    }
  }

  setState()

  return {
    user,
    prevUserEmail,
    setUser,
  }
})
