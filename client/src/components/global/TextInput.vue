<template>
    <label class="block mb-4">
        <span class="label">{{ props.label }}</span>
        <input
            :type="props.type"
            :autocomplete="props.autocomplete"
            :disabled="disabled"
            :min="min"
            :max="max"
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
    },
    min: {
        type: Number,
        default: 0,
    },
    max: {
        type: Number,
        default: 0,
    },
})

const value = ref(props.defaultValue)

defineExpose({
    value,
})
</script>
<style lang="scss" scoped>
input {
    width: 100%;
    min-width: 100px;
    padding: 0.375rem 0.5rem;
    border: 1px solid rgb(209, 213, 219);
    border-radius: 0.375rem;

}
</style>
