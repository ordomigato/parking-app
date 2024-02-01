<template>
    <div class="card" v-if="userStore.user">
        <div>
            <h2>Profile</h2>
            <p><strong>Email: </strong><span class="email">{{ userStore.user.username }} <CircleCheckmark v-if="userStore.user.verified" /></span></p>
            <p><strong>User ID: </strong>{{ userStore.user.client_id }}</p>
            <p><strong>Last Login: </strong>{{ convertDateTime(userStore.user.last_login) }}</p>
            <p><strong>Updated At: </strong>{{ convertDateTime(userStore.user.updated_at) }}</p>
            <p><strong>Created At: </strong>{{ convertDateTime(userStore.user.created_at) }}</p>
        </div>
    </div>
    <div class="card">
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
import CircleCheckmark from "@/components/icons/circle-checkmark.vue"
import { routeNames } from '@/router/routeNames';
import { logout } from '@/services/account.service';
import { useUserStore } from '@/stores/userStore';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';
import { convertDateTime } from '@/utils/date'

const router = useRouter()
const userStore = useUserStore()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleLogout = async () => {
    error.value = null
    busy.value = true
    try {
        await logout()
        router.push({ name: routeNames.auth })
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}
</script>
<style lang="scss" scoped>
.email {
    display: inline-flex;
    align-items: center;
    svg {
        margin-left: 0.25rem;
        height: 0.875rem;
        fill: var(--success-color);
    }
}
</style>
