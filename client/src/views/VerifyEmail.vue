<template>
    <main class="flex items-center justify-center bg-sky-50 h-screen">
        <div class="card text-center">
            <div v-if="!busy && !error" class="flex flex-col justify-center">
                <h2>Your email has been verified!</h2>
                <p></p>
                <footer>
                    <c-button
                        class="mx-auto"
                        @click="() => $router.push({ name: routeNames.dashboard })"
                    >
                        {{ userStore.user ? 'Access Dashboard' : 'Login' }}
                    </c-button>
                </footer>
            </div>
            <div v-else>
                <h2>Email Verification Failed</h2>
                <error-display :error="error"></error-display>
            </div>
        </div>
    </main>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { verifyEmail } from '@/services/account.service';
import { useUserStore } from '@/stores/userStore';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const userStore = useUserStore()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleVerify = async () => {
    error.value = null
    busy.value = true
    try {
        const email = route.query.email as string
        const otp = route.query.otp as string
        if (!email) {
            throw new Error("email is missing")
        }
        if (!otp) {
            throw new Error("otp is missing")
        }
        await verifyEmail(email, otp)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async() => {
    await handleVerify()
})
</script>