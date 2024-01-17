<template>
    <footer>
        <div class="controls">
            <span>Page: </span>
            <text-input
                class="page-input"
                :disabled="busy"
                ref="pageNumber"
                type="number"
                :defaultValue="props.defaultPage"
                @keyup.enter="queryPage"
            />
            <span>of {{ maxPageNumber }}</span>
        </div>
        <div class="results">
            <p>Results: {{ count }}</p>
        </div>
    </footer>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue';
import TextInput from '../global/TextInput.vue';
import { PaginationQuery } from '@/utils/pagination';

const props = defineProps({
    defaultPage: {
        type: String,
        default: "1",
    },
    count: {
        type: Number,
        required: true
    },
    limit: {
        type: Number,
        default: new PaginationQuery().limit
    }
})

const emit = defineEmits(['queryPage'])

const pageNumber = ref<InstanceType<typeof TextInput>>()

const busy = ref(false)

const maxPageNumber = computed(() => Math.ceil(props.count / props.limit))

const queryPage = () => {
    if (pageNumber.value) {
        let page = parseInt(pageNumber.value.value)
        if (page > maxPageNumber.value) {
            const maxNum = maxPageNumber.value
            pageNumber.value.value = maxNum.toString()
            page = maxNum
        }
        if (page < 1) {
            pageNumber.value.value = "1"
            page = 1
        }
        const payload = new PaginationQuery(props.limit, page)
        emit('queryPage', payload)
    }
}

</script>
<style lang="scss" scoped>
footer {
    padding: 1rem;
    position: sticky;
    left: 0;
    display: flex;
    align-items: center;
    // justify-content: center;
    white-space: nowrap;
    .controls {
        display: flex;
        align-items: center;
        white-space: nowrap;
    }
    .results {
        margin-left: auto;
    }
    label {
        margin: 0 1rem;
    }
    .page-input {
        width: 100px;
    }
}
</style>
