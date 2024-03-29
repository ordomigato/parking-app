<template>
    <div class="dropdown-container">
        <div
            tabindex="0"
            :class="`dropdown-btn ${open ? 'open' : ''}`"
            @click="toggleOpen"
            @keyup.enter="toggleOpen"
        >
            {{ props.selected ? props.selected.name : "Select" }}<span class="arrow"></span>
        </div>
        <div class="dropdown" v-if="open">
            <ul>
                <li
                    tabindex="0"
                    v-for="item in props.items"
                    :key="item.value"
                    @click="() => onSelect(item)"
                    @keyup.enter="() => onSelect(item)"
                >
                    {{ item.name }}
                </li>
            </ul>
        </div>
    </div>
</template>
<script setup lang="ts">
import type { IDropdownItem } from "@/types";
import { ref } from "vue";

const props = defineProps({
    selected: {
        type: Object as () => IDropdownItem<any, any> | null,
        default: null,
    },
    items: {
        type: Array as () => IDropdownItem<any, any>[],
        required: true,
    },
    disabled: {
        type: Boolean,
        default: false,
    }
})

const emit = defineEmits(['onSelect'])

const open = ref(false)

const toggleOpen = () => {
    open.value = !open.value
}

const onSelect = (payload: IDropdownItem<any, any>) => {
    emit('onSelect', payload)
    toggleOpen()
}

</script>
<style lang="scss" scoped>
.dropdown-container {
    position: relative;
    user-select: none;
    margin-bottom: 1rem;
    line-height: 1.5rem;
    min-width: 100px;
    .dropdown-btn {
        padding: 0.375rem 0.5rem;
        border-radius: 0.375rem;
        border: 1px solid rgb(209 213 219);
        &.open {
            border-bottom: 1px solid transparent;
            border-radius: 0.375rem 0.375rem 0 0;
            .arrow {
                transform: translate(0, -50%) rotate(180deg);
            }
        }
        .arrow {
            width: 4px;
            height: 4px;
            border-left: 4px solid transparent;
            border-right: 4px solid transparent;
            border-top: 5px solid rgb(209, 213, 219);

            position: absolute;
            right: 1rem;
            top: 50%;
            transform: translate(0, -50%);
        }
    }
    .dropdown {
        position: absolute;
        background-color: white;
        padding: 0.375rem 0;
        width: 100%;
        border: 1px solid rgb(209 213 219);
        border-top: none;
        border-radius: 0 0 0.375rem 0.375rem;
        max-height: 6rem;
        overflow: hidden;
        overflow-y: auto;
        z-index: 1;
        li {
            padding: 0 0.5rem;
            cursor: pointer;
            &:hover {
                background-color: rgb(209 213 219);
            }
        }
    }
}
</style>
