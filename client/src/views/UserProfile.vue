<template>
    <div class=card>
        <c-button
            @click="handleLogout"
            @keyup.enter="handleLogout"
        >
            Logout
        </c-button>
        <error-display :error="error"></error-display>
    </div>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { logout } from '@/services/account.service';
import { useUserStore } from '@/stores/userStore';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const userStore = useUserStore()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleLogout = async () => {
    error.value = null
    busy.value = true
    try {
        await logout()
        userStore.setUser(null)
        router.push({ name: routeNames.auth })
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}
</script>
