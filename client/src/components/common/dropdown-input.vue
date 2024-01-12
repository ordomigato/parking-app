<template>
    <div class="dropdown-container">
        <div
            tabindex="0"
            :class="`dropdown-btn ${open ? 'open' : ''}`"
            @click="toggleOpen"
            @keyup.enter="toggleOpen"
        >
            {{ props.selected ? props.selected.name : "Select" }}
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
    font-size: 0.85rem;
    line-height: 1.5rem;
    .dropdown-btn {
        padding: 0.375rem 0.5rem;
        border-radius: 0.375rem;
        border: 1px solid rgb(209 213 219);
        &.open {
            border-bottom: none;
            border-radius: 0.375rem 0.375rem 0 0;
        }
    }
    .dropdown {
        // position: absolute;
        background-color: white;
        padding: 0.375rem 0;
        width: 100%;
        border: 1px solid rgb(209 213 219);
        border-top: none;
        border-radius: 0 0 0.375rem 0.375rem;
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
