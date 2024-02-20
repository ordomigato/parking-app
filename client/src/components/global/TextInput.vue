<template>
    <label class="block mb-4">
        <span class="label">{{ props.label }}</span>
        <input
            :type="props.type"
            :placeholder="placeholder"
            :autocomplete="props.autocomplete"
            :disabled="disabled"
            :min="min?.toString()"
            :max="max?.toString()"
            v-model="value"
        />
    </label>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps({
    label: {
        type: String,
        default: '',
    },
    type: {
        type: String,
        default: 'text',
    },
    placeholder: {
        type: String,
        default: '',
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
        type: Number || undefined,
        default: undefined,
    },
    max: {
        type: Number || undefined,
        default: undefined,
    },
    uppercase: {
        type: Boolean,
        default: false,
    },
})

const value = ref(props.defaultValue)

watch(value, (val) => {
    if (props.uppercase) {
        value.value = val.toUpperCase()
    }
})

defineExpose({
    value,
})
</script>
<style lang="scss" scoped>
input {
    width: 100%;
    min-width: 100px;
    padding:  0.25rem 0.5rem;
    border: 1px solid rgb(209, 213, 219);
    border-radius: 0.375rem;

}
</style>
