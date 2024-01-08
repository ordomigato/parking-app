<template>
    <div>
        <header class="pb-6 text-center">
            <h1 class="text-lg font-semibold">Welcome to SecurePark</h1>
            <p>Create and manage parking spaces with ease</p>
        </header>
        <text-input :disabled="busy" ref="email" label="Email" type="email" autocomplete="email" />
        <text-input :disabled="busy" ref="pass" label="Password" type="password" autocomplete="password" />
        <error-display :error="error"></error-display>
        <c-button :disabled="busy" fullWidth @click="handleLogin">Login</c-button>
        <footer class="pt-12">
            <p class="text-center">Don't have an account? <c-button :disabled="busy" @click="handleChangeView" isLink>Register</c-button></p>
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

const router = useRouter()

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

        await loginUser(userEmail, userPass)
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
