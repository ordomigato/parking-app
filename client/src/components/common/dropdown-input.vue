<template>
    <div class="dropdown-container">
        <div class="dropdown-btn" @click="toggleOpen">
            {{ props.selected ? props.selected.name : "Select" }}
        </div>
        <div class="dropdown" v-if="open">
            <ul>
                <li
                    v-for="item in props.items"
                    :key="item.value"
                    @click="() => onSelect(item)"
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
    .dropdown-btn {
        padding: 0.25rem 1rem;
        border: 1px solid black;
    }
    .dropdown {
        position: absolute;
        background-color: white;
        padding: 0.25rem 1rem;
        box-shadow: 0 3px 6px rgba(0, 0, 0, 0.1);
        width: 100%;
    }
}
</style>
