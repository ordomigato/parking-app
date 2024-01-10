<template>
    <div class="card">
        <header>
            <c-button class="mb-2 ms-auto" @click="$router.push({ name: routeNames.createForm })">+ Create Form</c-button>
        </header>
        <table class="text-left" v-if="forms.length">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Created At</th>
                    <th>Updated At</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="form in forms" :key="form.form_id">
                    <td>
                        <c-button @click="$router.push({ name: routeNames.form, params: { id: form.form_id }})" isLink>{{ form.name }}</c-button>
                    </td>
                    <td>{{ convertDate(form.created_at) }}</td>
                    <td>{{ convertDate(form.updated_at) }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import type { IForm } from '@/types';
import { onMounted, ref, type Ref } from 'vue';
import { convertDate } from '@/utils/date'
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { handleError } from '@/utils/error';
import { getForms } from '@/services/form.service';

const workspaceStore = useWorkspaceStore()

const forms: Ref<IForm[]> = ref([])

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onGetForms = async () => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error("Something went wrong")
        }
        const resp = await getForms(workspaceStore.currentWorkspace?.workspace_id)
        forms.value = resp
        console.log(resp)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async() => {
    await onGetForms()
})

</script>
