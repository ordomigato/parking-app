import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import type { IToast, ToastStatus } from '@/types'

const STORE_NAME = 'toast'
const defaultTimeout: number = 3000;

export const useToastStore = defineStore(STORE_NAME, () => {
  const toasts: Ref<IToast[]> = ref([])

  const createToast = (text: string, status: ToastStatus): IToast => ({
    text,
    status,
    id: Math.random() * 1000,
  });

  const updateState = (payload: string, status: ToastStatus) => {
    const toast = createToast(payload, status);

    toasts.value.push(toast);

    setTimeout(() => {
        toasts.value = toasts.value.filter((t) => t.id !== toast.id);
    }, defaultTimeout);
}

  return {
    toasts,
    createToast,
    updateState,
  }
})
