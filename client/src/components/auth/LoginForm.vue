<template>
    <div>
        <header class="pb-6 text-center">
            <h1 class="text-lg font-semibold">Welcome to SecurePark</h1>
            <p>Create and manage parking spaces with ease</p>
        </header>
        <text-input
            :disabled="busy"
            ref="email"
            label="Email"
            type="email"
            autocomplete="email"
            :defaultValue="userStore.prevUserEmail"
            @keyup.enter="handleLogin"
        />
        <text-input
            :disabled="busy"
            ref="pass"
            label="Password"
            type="password"
            autocomplete="password"
            @keyup.enter="handleLogin"
        />
        <error-display :error="error"></error-display>
        <c-button
            class="w-100"
            :disabled="busy"
            fullWidth
            @click="handleLogin"
            @keyup.enter="handleLogin"
        >
            Login
        </c-button>
        <footer class="pt-12">
            <p
                class="text-center"
            >
                Don't have an account? <c-button
                    :disabled="busy"
                    @click="handleChangeView"
                    @keyup.enter="handleChangeView"
                    variant="link"
                >
                    Register
                </c-button>
            </p>
        </footer>
    </div>
</template>
<script setup lang="ts">
import { loginUser } from '@/services/account.service';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import TextInput from '../global/TextInput.vue';
import { useRouter } from 'vue-router';
import { routeNames } from '@/router/routeNames'
import { useUserStore } from '@/stores/userStore';

const router = useRouter()

const userStore = useUserStore()

const emit = defineEmits(['changeView'])

const email = ref<InstanceType<typeof TextInput>>()
const pass = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleChangeView = () => {
    emit('changeView')
}

const handleLogin = async () => {
    error.value = null
    busy.value = true
    try {
        const userEmail = email.value?.value
        const userPass = pass.value?.value

        if (!userEmail) {
            throw new Error("Email cannot be blank")
        }

        if (!userPass) {
            throw new Error("Password cannot be blank")
        }

        const { user } = await loginUser(userEmail, userPass)
        userStore.setUser(user)
        router.push({
            name: routeNames.overview,
        })

    } catch (e) {
        error.value = handleError(e);
    } finally {
        busy.value = false
    }
}
</script>
