<template>
    <div class="table-container">
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>V. Plate</th>
                    <th>V. Make</th>
                    <th>V. Model</th>
                    <th>V. Color</th>
                    <th>Expiry</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Email</th>
                    <th>Phone</th>
                    <th>Created At</th>
                    <th>Updated At</th>
                </tr>
            </thead>
            <tbody>
                <tr
                    @click="() => handleDeletePermits(permit.permit_id)"
                    v-for="permit in permits"
                    :key="permit.permit_id"
                    :class="`permit ${isExpired(permit.expiry) ? 'expired' : ''}`"
                >
                    <td>{{ permit.permit_id }}</td>
                    <td>{{ permit.v_plate }}</td>
                    <td>{{ permit.v_make }}</td>
                    <td>{{ permit.v_model }}</td>
                    <td>{{ permit.v_color }}</td>
                    <td  v-if="permit.expiry">{{ convertDate(permit.expiry) }}, {{ convertTime(permit.expiry) }}</td>
                    <td v-else></td>
                    <td>{{ permit.first_name }}</td>
                    <td>{{ permit.last_name }}</td>
                    <td>{{ permit.email }}</td>
                    <td>{{ permit.primary_phone }}</td>
                    <td>{{ convertDate(permit.created_at) }}, {{ convertTime(permit.created_at) }}</td>
                    <td>{{ convertDate(permit.updated_at) }}, {{ convertTime(permit.updated_at) }}</td>
                </tr>
            </tbody>
        </table>
        <error-display :error="error"></error-display>
    </div>
    <PaginationInput @queryPage="handleGetPermits" :count="count" />
</template>
<script setup lang="ts">
import { getPermits, deletePermit } from '@/services/permit.service';
import type { IPagination, IPermit } from '@/types';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import { convertDate, convertTime } from '@/utils/date'
import { useWorkspaceStore } from '@/stores/workspaceStore';
import PaginationInput from '../common/pagination-input.vue';
import { PaginationQuery } from '@/utils/pagination';
import { useToastStore } from '@/stores/toastStore';

const workspaceStore = useWorkspaceStore()
const toastStore = useToastStore()

const props = defineProps({
    formId: {
        type: String,
        required: true
    }
})

const permits: Ref<IPermit[]> = ref([])
const count: Ref<number> = ref(0)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const isExpired = (exp: Date): boolean => {
    if (!exp) {
        return false
    }
    const expiry = new Date(exp)
    const now = new Date(Date.now())
    return expiry < now
}

const handleGetPermits = async (query: IPagination) => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error('something went wrong')
        }

        const { data, count: c } = await getPermits(workspaceStore.currentWorkspace.workspace_id, props.formId, query)
        permits.value = data
        count.value = c
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const handleDeletePermits = async (permitId: string) => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error('something went wrong')
        }
        await deletePermit(workspaceStore.currentWorkspace.workspace_id, props.formId, permitId)
        permits.value = permits.value.filter(p => p.permit_id !== permitId)
        toastStore.updateState("Successfully Deleted Permit", "success")
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async () => {
    await handleGetPermits(new PaginationQuery())
})
</script>
<style lang="scss" scoped>
.permit {
    cursor: pointer;
    &.expired {
        background-color: #fecaca;
    }
}
</style>