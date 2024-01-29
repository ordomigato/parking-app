<template>
    <div v-if="form">
        <div class="card">
            <header>
                <h2>{{ form.name }}</h2>
            </header>
            <div class="text-sm">
                <p><strong>ID:</strong> {{ form.form_id }}</p>
                <p><strong>Created At:</strong> {{ convertDate(form.created_at) }}</p>
                <p><strong>Updated At:</strong> {{ convertDate(form.updated_at) }}</p>
                <p><strong>Path:</strong> <a class="link" :href="formPath" target="_blank">{{ formPath }}</a></p>
            </div>
        </div>
        <div class="card">
            <header class="title-btn-combo">
                <h2>Parking Permits</h2>
                <c-button @click="onDownloadCSV">Download CSV</c-button>
            </header>
            <PermitTable
                :formId="form.form_id"
            />
        </div>
        <div class="card">
            <header>
                <h2>Form Settings</h2>
            </header>
            <FormConfig v-if="form" :formInfo="form" />
        </div>
        <div class="card">
            <header>
                <h2 class="danger-text">Danger Zone</h2>
            </header>
            <c-button
                :disabled="busy"
                @click="onDeleteForm"
                danger
            >
                Delete Form
            </c-button>
            <error-display :error="error"></error-display>
        </div>
    </div>
</template>
<script setup lang="ts">
import FormConfig from '@/components/form/form-config.vue';
import PermitTable from '@/components/permit/permit-table.vue'
import { getForm, deleteForm } from '@/services/form.service';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { handleError } from '@/utils/error';
import { computed, onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';
import { type IForm, type IPermit } from '@/types';
import router from '@/router';
import { routeNames } from '@/router/routeNames';
import { convertDate } from '@/utils/date'
import { downloadPermits } from '@/services/permit.service';
import { mkConfig, generateCsv, download } from "export-to-csv";

const csvConfig = mkConfig({ useKeysAsHeaders: true });

const route = useRoute()
const workspaceStore = useWorkspaceStore()

const form: Ref<IForm | null> = ref(null)
const permits: Ref<IPermit[] | null> = ref(null)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const formPath = computed(() => window.location.origin + form.value?.path.path)

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

const onDownloadCSV = async () => {
    busy.value = true
    error.value = null
    try {
        if (!workspaceStore.currentWorkspace!) {
            throw new Error('Something went wrong')
        }
        if (!form.value) {
            throw new Error('Something went wrong')
        }
        // guard to block potential spamming
        if (!permits.value) {
            permits.value = await downloadPermits(
                workspaceStore.currentWorkspace!.workspace_id,
                form.value!.form_id,
                {
                    from: new Date(form.value.created_at).toISOString(),
                    to: new Date().toISOString(),
                }
            )
        }

        const csv = generateCsv(csvConfig)(permits.value)
        download(csvConfig)(csv)
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