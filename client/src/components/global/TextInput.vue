<template>
    <label class="block mb-4">
        <span class="block text-sm leading-6 font-medium">{{ props.label }}</span>
        <input
            class="block w-full rounded-md border-0 px-2 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            :type="props.type"
            :autocomplete="props.autocomplete"
            :disabled="disabled"
            v-model="value"
        />
    </label>
</template>
<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps({
    label: {
        type: String,
        default: '',
    },
    type: {
        type: String,
        default: 'text',
    },
    autocomplete: {
        type: String,
        default: 'off',
    },
    disabled: {
        type: Boolean,
        default: false,
    },
    defaultValue: {
        type: String,
        default: ''
    }
})

const value = ref(props.defaultValue)

defineExpose({
    value,
})
</script>
