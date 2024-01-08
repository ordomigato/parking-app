<template>
  <div>
    <SideNav />
    <error-display :error="error"></error-display>
    <c-button @click="handleLogout">Logout</c-button>
  </div>
</template>
<script setup lang="ts">
import SideNav from '@/components/nav/side-nav.vue'
import { routeNames } from '@/router/routeNames';
import { logout } from '@/services/account.service';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

const error: Ref<Error | null> = ref(null)

const handleLogout = async () => {
  try {
    await logout()
    router.push({ name: routeNames.auth })
  } catch(e) {
    error.value = handleError(e)
  }
}
</script>
