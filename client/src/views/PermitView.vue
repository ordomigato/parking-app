<template>
    <div v-if="permit">
        <div class="card">
            <div class="permit-data">
                <p><strong>Permit:</strong> {{ permit.permit_id }}</p>
                <h3>General</h3>
                <p v-if="permit.start_date"><strong>Start Date:</strong> {{ permit.start_date }}</p>
                <p v-if="permit.end_date"><strong>End Date:</strong> {{ permit.end_date }}</p>
                <p><strong>Submitted:</strong> {{ permit.created_at }}</p>
                <p><strong>Updated:</strong> {{ permit.updated_at }}</p>
                <h3>Vehicle</h3>
                <p><strong>Plate:</strong> {{ permit.v_plate }}</p>
                <p><strong>Make:</strong> {{ permit.v_make }}</p>
                <p><strong>Model:</strong> {{ permit.v_model }}</p>
                <p><strong>Color:</strong> {{ permit.v_color }}</p>
                <h3>Person</h3>
                <p><strong>First Name:</strong> {{ permit.first_name }}</p>
                <p><strong>Last Name:</strong> {{ permit.last_name }}</p>
                <p><strong>Email:</strong> {{ permit.email }}</p>
                <p><strong>Phone:</strong> {{ permit.primary_phone }}</p>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { getPermit } from '@/services/permit.service';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import type { IFormattedPermit } from '@/types';
import { handleError } from '@/utils/error';
import { formatPermits } from '@/utils/format';
import { onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';

const workspaceStore = useWorkspaceStore()
const route = useRoute()

const permit: Ref<IFormattedPermit | null> = ref(null)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleGetPermit = async () => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error('Something went wrong')
        }
        const resp = await getPermit(
            workspaceStore.currentWorkspace.workspace_id,
            route.params.id as string,
            route.params.permitId as string,
        )

        permit.value = formatPermits([resp])[0]
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async() => {
    await handleGetPermit()
})
</script>
<style lang="scss" scoped>
.permit-data {
    display: flex;
    flex-direction: column;
    h3 {
        font-weight: bold;
        font-size: 1rem;
        margin: 0.5rem 0;
    }
}
</style>