<template>
    <div class="search-bar-container">
        <TextInput
            :disabled="disabled"
            placeholder="Search"
            ref="search"
            @keyup.enter="handleSearch"
        />
        <div class="search-button-container">
            <c-button
                @click="handleSearch"
                @keyup.enter="handleSearch"
            >
                Search
            </c-button> 
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import TextInput from '../global/TextInput.vue';

const props = defineProps({
    disabled: {
        type: Boolean,
        default: false,
    }
})

const emit = defineEmits(['onSearch'])

const search = ref<InstanceType<typeof TextInput>>()

const handleSearch = () => {
    if (props.disabled || !search.value) {
        return
    }
    emit('onSearch', search.value.value)
}

</script>
<style lang="scss" scoped>
.search-bar-container {
    display: flex;
    .search-button-container {
        margin-left: 0.5rem;
    }
}
</style>