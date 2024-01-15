<template>
    <button
        :class="`btn ${props.variant}`"
        :disabled="disabled"
    >
        <slot></slot>
    </button>
</template>
<script setup lang="ts">
const props = defineProps({
    variant: {
        type: String,
        validator: (val: string) => ['block', 'link', 'icon'].includes(val),
        default: 'block'
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
</script>
<style lang="scss">
.btn {
    &.block {
        padding: 0.25rem 0.5rem;
        background-color: var(--main-color);
        color: white;
        border: 1px solid var(--main-color);
        border-radius: 4px;
        transition: all 0.1s linear;
        &:hover,
        &:focus {
            background-color: white;
            color: var(--main-color);
        }
        &:disabled {
            opacity: 0.5;
        }
    }

    &.link {
        color: var(--link-color);
    }

    &.icon {
        box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
        // padding: 0.5rem;
        height: 2rem;
        max-width: 2rem;
        min-width: 2rem;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 20px;
        background-color: var(--main-color-light);
        transition: all 0.1s linear;
        svg {
            fill: var(--main-color);
        }
        &:hover,
        &:focus {
            background-color: white;
        }
    }
}
</style>