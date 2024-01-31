<template>
    <header class="title-btn-combo">
        <h2>Blacklist</h2>
        <c-button @click="blacklistModalOpen = true">Add Plate</c-button>
    </header>
    <div class="table-container">
        <table>
            <thead>
                <tr>
                    <th>Plate</th>
                    <th class="text-end">Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr
                    v-for="plate in blacklist"
                    :key="plate"
                >
                    <td>{{ plate }}</td>
                    <td>
                        <c-button
                            class="ms-auto"
                            @click="() => onDeleteFromBlacklist(plate)"
                            inverse
                        >
                            Remove
                        </c-button>
                    </td>
                </tr>
            </tbody>
        </table>
        <error-display :error="error"></error-display>
    </div>
    <PaginationInput @queryPage="onGetBlacklist" :count="count" />
    <ModalPopup v-if="blacklistModalOpen">
        <div class="modal-delete-container">
            <header>
                <h2>Add To Blacklist</h2>
            </header>
            <section>
                <text-input
                    ref="plate"
                    label="Plate"
                    uppercase
                    @keyup.enter="onAddToBlacklist"
                />
            </section>
            <error-display :error="error"></error-display>
            <footer>
                <c-button
                    :disabled="busy"
                    @click="blacklistModalOpen = false"
                    inverse
                >
                    Cancel
                </c-button>
                <c-button
                    :disabled="busy"
                    @click="onAddToBlacklist"
                >
                    Add Plate
                </c-button>
            </footer>
        </div>
    </ModalPopup>
</template>
<script lang="ts" setup>
import { addToBlacklist, deleteFromBlacklist, getBlacklist } from '@/services/blacklist.service';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import ModalPopup from '../common/modal-popup.vue';
import TextInput from '@/components/global/TextInput.vue'
import PaginationInput from '@/components/common/pagination-input.vue'
import type { IPagination } from '@/types';
import { PaginationQuery } from '@/utils/pagination';

const props = defineProps({
    workspaceId: {
        type: String,
        required: true,
    },
    formId: {
        type: String,
        required: true,
    }
})

const blacklist: Ref<string[]> = ref([])
const blacklistModalOpen = ref(false)
const plate = ref<InstanceType<typeof TextInput>>()
const count: Ref<number> = ref(0)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onGetBlacklist = async (query: IPagination) => {
    busy.value = true
    error.value = null
    try {
        const response = await getBlacklist(props.workspaceId, props.formId, query)
        blacklist.value = response.data || []
        count.value = response.count
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const onAddToBlacklist = async () => {
    busy.value = true
    error.value = null
    try {
        if (!plate.value?.value) {
            throw new Error('Please enter the plate you wish to blacklist')
        }
        await addToBlacklist(props.workspaceId, props.formId, plate.value.value)
        blacklist.value = blacklist.value.concat(plate.value.value)
        count.value = count.value + 1
        blacklistModalOpen.value = false
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const onDeleteFromBlacklist = async (plate: string) => {
    busy.value = true
    error.value = null
    try {
        await deleteFromBlacklist(props.workspaceId, props.formId, plate)
        blacklist.value = blacklist.value.filter(p => p !== plate)
        count.value = count.value - 1
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async () => {
    await onGetBlacklist(new PaginationQuery())
})
</script>