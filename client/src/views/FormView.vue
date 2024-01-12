<template>
    <div v-if="form">
        <div class="card">
            <h2>Form Info</h2>
            <div class="text-sm">
                <p><strong>Name:</strong> {{ form.name }}</p>
                <p><strong>ID:</strong> {{ form.form_id }}</p>
                <p><strong>Create At:</strong> {{ convertDate(form.created_at) }}</p>
                <p><strong>Updated At:</strong> {{ convertDate(form.updated_at) }}</p>
            </div>
        </div>
        <div class="card">
            <h2>Form Settings</h2>
            <div>
                <text-input
                    ref="formName"
                    :disabled="busy"
                    label="Form Name"
                    :defaultValue="form.name"
                ></text-input>
                <ConstraintTypeDropdown
                    ref="constraintType"
                    :default="form.submission_constraint_type"
                />
                <text-input
                    ref="constraintLimit"
                    :disabled="busy"
                    label="Constraint Limit"
                    :defaultValue="form.submission_constraint_limit.toString()"
                ></text-input>
                <c-button
                    @click="onUpdateForm"
                >
                    Save
                </c-button>
            </div>
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
import TextInput from '@/components/global/TextInput.vue';
import ConstraintTypeDropdown from '@/components/form/constraint-type-dropdown.vue'
import { getForm, updateForm, deleteForm } from '@/services/form.service';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';
import { type IForm, type IFormUpdateRequest } from '@/types';
import router from '@/router';
import { routeNames } from '@/router/routeNames';
import { convertDate } from '@/utils/date'

const route = useRoute()
const workspaceStore = useWorkspaceStore()

const formName = ref<InstanceType<typeof TextInput>>()
const constraintType = ref<InstanceType<typeof ConstraintTypeDropdown>>()
const constraintLimit = ref<InstanceType<typeof TextInput>>()

const form: Ref<IForm | null> = ref(null)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

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

const onUpdateForm = async () => {
    busy.value = true
    error.value = null
    try {
        if (!workspaceStore.currentWorkspace!) {
            throw new Error('Something went wrong')
        }
        if (!form.value) {
            throw new Error('Something went wrong')
        }
        if (!formName.value) {
            throw new Error('Something went wrong')
        }
        if ((!constraintType.value && constraintType.value !== '') || constraintType.value.selected === null) {
            throw new Error('Something went wrong')
        }
        if (!constraintLimit.value) {
            throw new Error('Something went wrong')
        }
        const payload: IFormUpdateRequest = {
            name: formName.value.value,
            submission_constraint_type: constraintType.value.selected.value,
            submission_constraint_limit: parseInt(constraintLimit.value.value),
        }
        updateForm(workspaceStore.currentWorkspace?.workspace_id, form.value.form_id, payload)
        form.value = {
            ...form.value,
            ...payload
        }
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