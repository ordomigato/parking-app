<template>
    <div class="table-components">
        <SearchBar
            :disabled="busy"
            @onSearch="onSearch"
        />
        <div v-if="tableColumns.length">
            <TableFilter
                :columns="tableColumns"
                :settingsId="LocalSettings.PermitTableFilter"
            />
        </div>
    </div>
    <div class="table-container">
        <table>
            <thead>
                <tr>
                    <th
                        v-for="col in filterTableCols(tableColumns)" :key="col.value"
                    >
                        {{ col.name }}
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr
                    @click="() => handleDeletePermits(permit.permit_id)"
                    v-for="permit in permits"
                    :key="permit.permit_id"
                    :class="`permit ${isExpired(new Date(permit.expiry)) ? 'expired' : ''}`"
                >
                    <td
                        v-for="(data, i) in filteredPermit(permit)"
                        :key="i"
                    >{{ data }}</td>
                </tr>
            </tbody>
        </table>
        <error-display :error="error"></error-display>
    </div>
    <PaginationInput @queryPage="handleGetPermits" :count="count" />
</template>
<script setup lang="ts">
import { getPermits, deletePermit } from '@/services/permit.service';
import type { IColumn, IFormattedPermit, IPagination } from '@/types';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import PaginationInput from '../common/pagination-input.vue';
import { PaginationQuery } from '@/utils/pagination';
import { useToastStore } from '@/stores/toastStore';
import SearchBar from '../common/search-bar.vue';
import TableFilter from '../common/table-filter.vue';
import { formatPermits } from '@/utils/format';
import { filterTableCols, filterTableData } from '@/utils/table';
import { LocalSettings, loadLocal } from '@/utils/storage';

const workspaceStore = useWorkspaceStore()
const toastStore = useToastStore()

const props = defineProps({
    formId: {
        type: String,
        required: true
    }
})

const permits: Ref<IFormattedPermit[]> = ref([])
const count: Ref<number> = ref(0)

const tableColumns: Ref<IColumn[]> = ref([])

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const filteredPermit = (permit: IFormattedPermit): string[] => {
    const cols: string[] = filterTableCols(tableColumns.value).map(c => c.value)
    return filterTableData(permit, cols)
}

const onSearch = (val: string) => {
    handleGetPermits(new PaginationQuery(), val)
}

const isExpired = (exp: Date): boolean => {
    if (!exp) {
        return false
    }
    const expiry = new Date(exp)
    const now = new Date(Date.now())
    return expiry < now
}

const handleGetPermits = async (query: IPagination, search?: string) => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error('something went wrong')
        }

        const { data, count: c } = await getPermits(workspaceStore.currentWorkspace.workspace_id, props.formId, query, search)
        permits.value = formatPermits(data)
        count.value = c
        const tableColumnsLocalSave = loadLocal(LocalSettings.PermitTableFilter)
        if (tableColumnsLocalSave) {
            tableColumns.value = JSON.parse(tableColumnsLocalSave)
        } else {
            tableColumns.value = [
                "permit_id",
                "v_plate",
                "expiry",
                "v_make",
                "v_model",
                "v_color",
                "first_name",
                "last_name",
                "email",
                "primary_phone",
                "created_at",
                "updated_at"
            ].map(c => ({
                name: c,
                value: c,
                visible: true,
            }))
        }
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

.table-components {
    display: flex;
    justify-content: space-between;
}
.permit {
    cursor: pointer;
    &.expired {
        background-color: #fecaca;
    }
}
</style>