<template>
    <button
        :class="btnStyle()"
    >
        <slot></slot>
    </button>
</template>
<script setup lang="ts">
const props = defineProps({
    isLink: {
        type: Boolean,
        default: false,
    }
})

const btnStyle = (): string => {
    if (props.isLink) {
        return "text-indigo-600"
    } else {
        return "flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
    }
}
</script>