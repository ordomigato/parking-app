<template>
  <div>
    <h1>Dashboard</h1>
    <error-display :error="error"></error-display>
    <button @click="() => handleLogout">Logout</button>
  </div>
</template>
<script setup lang="ts">
import { logout, getStatus } from '@/services/account.service';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';

const error: Ref<Error | null> = ref(null)

const handleLogout = async () => {
  try {
    await logout()
  } catch(e) {
    error.value = handleError(e)
  }
}

onMounted(async() => {
  await getStatus()
})
</script>
