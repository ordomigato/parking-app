<template>
    <p v-if="props.error" class="text-sm text-center text-red-600 mb-4">
        {{ props.error.message}}
    </p>
</template>
<script setup lang="ts">
const props = defineProps({
    error: {
        type: Error,
        default: null,
    }
})
</script>