<template>
  <Teleport to="body">
      <ul v-if="toastStore.toasts.length">
          <li
            v-for="toast in toastStore.toasts"
            :key="toast.text"
            :class="toast.status"
          >
            <span>
              {{ toast.text }}
            </span>
          </li>
      </ul>
  </Teleport>
</template>
<script setup lang="ts">
import { useToastStore } from '@/stores/toastStore';

const toastStore = useToastStore()
</script>
<style lang="scss" scoped>
ul {
    position: absolute;
    bottom: 1rem;
    left: 1rem;
    right: 1rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    align-items: center;
    justify-content: center;
    user-select: none;
    li {
        background-color: var(--font-color);
        color: white;
        padding: 0.5rem 1rem;
        border-radius: 100px;
        animation-name: fadeInOut;
        animation-duration: 3s;
        animation-iteration-count: 1;
        opacity: 0;
        &.success {
            background-color: var(--success-color);
        }
    }
}

@keyframes fadeInOut {
  0% {opacity: 0}
  5% { opacity: 1}
  95% { opacity: 1}
  100% {opacity: 0}
}
</style>
