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
                <c-button @click="downloadCSVModalOpen = true">Download CSV</c-button>
            </header>
            <PermitTable
                :formId="form.form_id"
            />
            <ModalPopup
                v-if="downloadCSVModalOpen"
            >
                <div class="modal-delete-container">
                    <header>
                        Download Permits (CSV)
                    </header>
                    <section>
                        <text-input
                            ref="from"
                            label="from"
                            type="datetime-local"
                        />
                        <text-input
                            ref="to"
                            label="to"
                            type="datetime-local"
                        />
                    </section>
                    <error-display :error="error"></error-display>
                    <footer>
                        <c-button
                            :disabled="busy"
                            @click="downloadCSVModalOpen = false"
                            inverse
                        >
                            Cancel
                        </c-button>
                        <c-button
                            :disabled="busy"
                            @click="onDownloadCSV"
                        >
                            Download
                        </c-button>
                    </footer>
                </div>
            </ModalPopup>
        </div>
        <div class="card">
            <header>
                <h2>Form Settings</h2>
            </header>
            <FormConfig v-if="form" :formInfo="form" />
        </div>
        <div class="card">
            <DeleteForm :form="form" />
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue'
import FormConfig from '@/components/form/form-config.vue';
import DeleteForm from '@/components/form/delete-form.vue';
import PermitTable from '@/components/permit/permit-table.vue'
import { getForm } from '@/services/form.service';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { handleError } from '@/utils/error';
import { computed, onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';
import { type IForm, type IPermit } from '@/types';
import { convertDate } from '@/utils/date'
import { downloadPermits } from '@/services/permit.service';
import { mkConfig, generateCsv, download } from "export-to-csv";
import ModalPopup from '@/components/common/modal-popup.vue';

const csvConfig = mkConfig({ useKeysAsHeaders: true });

const route = useRoute()
const workspaceStore = useWorkspaceStore()

const form: Ref<IForm | null> = ref(null)
const permits: Ref<IPermit[]> = ref([])

const downloadCSVModalOpen = ref(false)
const from = ref<InstanceType<typeof TextInput>>()
const to = ref<InstanceType<typeof TextInput>>()

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
        if (!from.value?.value) {
            throw new Error('Please select a from date')
        }
        if (!to.value?.value) {
            throw new Error('Please select a to date')
        }
        // guard to block potential spamming
        if (permits.value?.length === 0) {
            permits.value = await downloadPermits(
                workspaceStore.currentWorkspace!.workspace_id,
                form.value!.form_id,
                {
                    from: new Date(from.value.value).toISOString(),
                    to: new Date(to.value.value).toISOString(),
                }
            )

            if (permits.value?.length === 0) {
                throw new Error('No permits were found between these dates.')
            }

            const csv = generateCsv(csvConfig)(permits.value)
            download(csvConfig)(csv)
        }
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