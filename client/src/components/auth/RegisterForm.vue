<template>
    <div>
        <header class="text-center pb-6">
            <h1 class="text-lg font-semibold">Create an Account</h1>
            <p>Create and manage parking spaces with ease</p>
        </header>
        <text-input ref="email" label="Email" type="email" autocomplete="email" />
        <text-input ref="pass" label="Password" type="password" />
        <text-input ref="confPass" label="Confirm Password" type="password" />
        <error-display :error="error"></error-display>
        <c-button @click="handleRegister">Register</c-button>
        <footer class="pt-12">
            <p class="text-center">Already have an account? <c-button @click="handleChangeView" isLink>Login</c-button></p>
        </footer>
    </div>
</template>
<script setup lang="ts">
import { ref, type Ref } from 'vue';
import TextInput from '../global/TextInput.vue';
import { handleError } from '../../utils/error';
import { registerUser } from '../../services/account.service'

const emit = defineEmits(['changeView'])

const email = ref<InstanceType<typeof TextInput>>()
const pass = ref<InstanceType<typeof TextInput>>()
const confPass = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)

const handleChangeView = () => {
    emit('changeView')
}

const handleRegister = async () => {
    error.value = null
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
    } catch (e) {
        error.value = handleError(e);
    }
}
</script>
