<template>
    <div>
        <header class="pb-6 text-center">
            <h1 class="text-lg font-semibold">Welcome to ParkSecure</h1>
            <p>Create and manage parking spaces with ease</p>
        </header>
        <text-input label="Email" type="email" autocomplete="email" />
        <text-input label="Password" type="password" autocomplete="password" />
        <c-button>Login</c-button>
        <footer class="pt-12">
            <p class="text-center">Don't have an account? <c-button @click="handleChangeView" isLink>Register</c-button></p>
        </footer>
    </div>
</template>
<script setup lang="ts">
const emit = defineEmits(['changeView'])

const handleChangeView = () => {
    emit('changeView')
}
</script>
