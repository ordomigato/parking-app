<template>
    <div v-if="form">
        <div class="card">
            <h2>Form Info</h2>
            <div class="text-sm">
                <p><strong>Name:</strong> {{ form.name }}</p>
                <p><strong>ID:</strong> {{ form.form_id }}</p>
                <p><strong>Create At:</strong> {{ convertDate(form.created_at) }}</p>
                <p><strong>Updated At:</strong> {{ convertDate(form.updated_at) }}</p>
                <p><strong>Path:</strong> <a class="link" :href="formPath" target="_blank">{{ formPath }}</a></p>
            </div>
        </div>
        <div class="card">
            <h2>Form Settings</h2>
            <FormConfig v-if="form" :formInfo="form" />
        </div>
        <div class="card">
            <h2 class="danger-text">Danger Zone</h2>
            <c-button
                :disabled="busy"
                @click="onDeleteForm"
            >
                Delete Form
            </c-button>
            <error-display :error="error"></error-display>
        </div>
    </div>
</template>
<script setup lang="ts">
import FormConfig from '@/components/form/form-config.vue';
import { getForm, deleteForm } from '@/services/form.service';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { handleError } from '@/utils/error';
import { computed, onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';
import { type IForm } from '@/types';
import router from '@/router';
import { routeNames } from '@/router/routeNames';
import { convertDate } from '@/utils/date'

const route = useRoute()
const workspaceStore = useWorkspaceStore()

const form: Ref<IForm | null> = ref(null)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const formPath = computed(() => window.location.origin + form.value?.path)

const onGetForm = async () => {
    busy.value = true
    error.value = null
    try {
        if (!workspaceStore.currentWorkspace!) {
            throw new Error('Something went wrong')
        }
        form.value = await getForm(workspaceStore.currentWorkspace?.workspace_id, route.params.id as string)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const onDeleteForm = async () => {
    busy.value = true
    error.value = null
    try {
        if (!workspaceStore.currentWorkspace!) {
            throw new Error('Something went wrong')
        }
        if (!form.value) {
            throw new Error('Something went wrong')
        }
        await deleteForm(workspaceStore.currentWorkspace?.workspace_id, form.value.form_id)
        router.push({ name: routeNames.forms })
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

onMounted(async() => {
    await onGetForm()
})

</script>