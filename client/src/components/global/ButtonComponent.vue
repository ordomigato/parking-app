<template>
    <button
        :class="btnStyle()"
        :disabled="disabled"
    >
        <slot></slot>
    </button>
</template>
<script setup lang="ts">
const props = defineProps({
    isLink: {
        type: Boolean,
        default: false,
    },
    fullWidth: {
        type: Boolean,
        default: false,
    },
    disabled: {
        type: Boolean,
        default: false,
    }
})

const btnStyle = (): string => {
    if (props.isLink) {
        return "text-sky-700"
    } else {
        return `flex ${props.fullWidth && 'w-full'} justify-center rounded-md bg-sky-700 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-700 disabled:opacity-50 disabled:bg-sky-700`
    }
}
</script>