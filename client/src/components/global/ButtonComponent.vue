<template>
    <button
        :class="`
            btn
            ${props.variant}
            ${props.danger ? 'danger' : ''}
            ${props.inverse ? 'inverse' : ''}
        `"
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
    danger: {
        type: Boolean,
        default: false,
    },
    inverse: {
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
</script>
<style lang="scss">
.btn {
    &.block {
        white-space: nowrap;
        padding: 0.25rem 0.5rem;
        background-color: var(--main-color);
        color: white;
        border: 1px solid var(--main-color);
        border-radius: 4px;
        transition: all 0.1s linear;
        &:hover:not(:disabled),
        &:focus:not(:disabled) {
            background-color: white;
            color: var(--main-color);
        }
        &:disabled {
            opacity: 0.5;
        }
        &.inverse {
            background-color: white;
            border-color: var(--main-color);
            color: var(--main-color);
            &:hover:not(:disabled),
            &:focus:not(:disabled) {
                border-color: var(--main-color);
                background-color: var(--main-color);
                color: white;
            }
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
        background-color: var(--main-color);
        transition: all 0.1s linear;
        svg {
            fill: white;
        }
        &:hover:not(:disabled),
        &:focus:not(:disabled) {
            background-color: var(--main-color-light);
            svg {
                fill: var(--main-color);
            }
        }
    }
    &.danger {
        border-color: var(--danger-color);
        background-color: var(--danger-color);
        &:hover:not(:disabled),
        &:focus:not(:disabled) {
            background-color: white;
            color: var(--danger-color);
        }
        &.inverse {
            background-color: white;
            border-color: var(--danger-color);
            color: var(--danger-color);
            &:hover:not(:disabled),
            &:focus:not(:disabled) {
                border-color: var(--danger-color);
                background-color: var(--danger-color);
                color: white;
            }
        }
    }
}
</style>