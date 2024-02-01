<template>
    <div>
        <header class="text-center pb-6">
            <h1 class="text-lg font-semibold">Create an Account</h1>
            <p>Create and manage parking spaces with ease</p>
        </header>
        <text-input
            :disabled="busy"
            ref="email"
            label="Email"
            type="email"
            autocomplete="email"
            @keyup.enter="handleRegister"
        />
        <text-input
            :disabled="busy"
            ref="pass"
            label="Password"
            type="password"
            @keyup.enter="handleRegister"
        />
        <text-input
            :disabled="busy"
            ref="confPass"
            label="Confirm Password"
            type="password"
            @keyup.enter="handleRegister"
        />
        <error-display :error="error"></error-display>
        <c-button
            class="w-100"
            :disabled="busy"
            fullWidth
            @click="handleRegister"
            @keyup.enter="handleRegister"
        >
            Register
        </c-button>
        <footer class="pt-12">
            <p class="text-center">Already have an account? <c-button :disabled="busy" @click="handleChangeView" variant="link">Login</c-button></p>
        </footer>
    </div>
</template>
<script setup lang="ts">
import { ref, type Ref } from 'vue';
import TextInput from '../global/TextInput.vue';
import { handleError } from '../../utils/error';
import { loginUser, registerUser } from '../../services/account.service'
import { useRouter } from 'vue-router';
import { routeNames } from '@/router/routeNames';
import { useUserStore } from '@/stores/userStore';

const router = useRouter()
const userStore = useUserStore()

const emit = defineEmits(['changeView'])

const email = ref<InstanceType<typeof TextInput>>()
const pass = ref<InstanceType<typeof TextInput>>()
const confPass = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleChangeView = () => {
    emit('changeView')
}

const handleRegister = async () => {
    error.value = null
    busy.value = true
    try {
        const userEmail = email.value?.value
        const userPass = pass.value?.value
        const userConfPass = confPass.value?.value

        if (!userEmail) {
            throw new Error("Email cannot be blank")
        }

        if (!userPass) {
            throw new Error("Password cannot be blank")
        }

        if (userPass !== userConfPass) {
            throw new Error("Passwords do not match")
        }

        await registerUser(userEmail, userPass, userConfPass)
        const { user } = await loginUser(userEmail, userPass)
        userStore.setUser(user)
        router.push({
            name: routeNames.workspaces,
        })
    } catch (e) {
        error.value = handleError(e);
    } finally {
        busy.value = false
    }
}
</script>
