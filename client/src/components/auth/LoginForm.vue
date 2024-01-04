<template>
    <div>
        <header class="pb-6 text-center">
            <h1 class="text-lg font-semibold">Welcome to SecurePark</h1>
            <p>Create and manage parking spaces with ease</p>
        </header>
        <text-input ref="email" label="Email" type="email" autocomplete="email" />
        <text-input ref="pass" label="Password" type="password" autocomplete="password" />
        <error-display :error="error"></error-display>
        <c-button @click="handleLogin">Login</c-button>
        <footer class="pt-12">
            <p class="text-center">Don't have an account? <c-button @click="handleChangeView" isLink>Register</c-button></p>
        </footer>
    </div>
</template>
<script setup lang="ts">
import { loginUser, getStatus } from '@/services/account.service';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import TextInput from '../global/TextInput.vue';

const emit = defineEmits(['changeView'])

const email = ref<InstanceType<typeof TextInput>>()
const pass = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)

const handleChangeView = () => {
    emit('changeView')
}

const handleLogin = async () => {
    error.value = null
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
        await getStatus()
    } catch (e) {
        error.value = handleError(e);
    }
}
</script>
