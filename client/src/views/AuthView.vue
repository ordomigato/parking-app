<template>
  <main class="flex items-center justify-center bg-sky-50 h-screen">
    <CardWrapper class="m-2 max-w-md">
      <RegisterForm v-if="isRegister" @changeView="() => isRegister = false" />
      <LoginForm v-else @changeView="() => isRegister = true" />
    </CardWrapper>
  </main>
</template>
<script setup lang="ts">
import CardWrapper from "../components/CardWrapper.vue"
import RegisterForm from "../components/auth/RegisterForm.vue"
import LoginForm from "../components/auth/LoginForm.vue"
import { ref } from "vue";

const isRegister = ref(false)
</script>
