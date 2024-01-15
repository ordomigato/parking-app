<template>
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
            <tr v-for="permit in permits" :key="permit.permit_id">
                <td>{{ permit.permit_id }}</td>
                <td>{{ permit.v_plate }}</td>
                <td>{{ permit.v_make }}</td>
                <td>{{ permit.v_model }}</td>
                <td>{{ permit.v_color }}</td>
                <td>{{ convertDate(permit.expiry) }}</td>
                <td>{{ permit.first_name }}</td>
                <td>{{ permit.last_name }}</td>
                <td>{{ permit.email }}</td>
                <td>{{ permit.primary_phone }}</td>
                <td>{{ convertDate(permit.created_at) }}</td>
                <td>{{ convertDate(permit.updated_at) }}</td>
            </tr>
        </tbody>
    </table>
</template>
<script setup lang="ts">
import { getPermits } from '@/services/permit.service';
import type { IPermit } from '@/types';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import { convertDate } from '@/utils/date'

const props = defineProps({
    formId: {
        type: String,
        required: true
    }
})

const permits: Ref<IPermit[]> = ref([])

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const handleGetPermits = async () => {
    error.value = null
    busy.value = true
    try {
        permits.value = await getPermits(props.formId)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async () => {
    await handleGetPermits()
})
</script>