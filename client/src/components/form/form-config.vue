<template>
    <div>
        <div>
            <text-input
                ref="formName"
                :disabled="busy"
                label="Form Name"
                :defaultValue="form?.name || ''"
            ></text-input>
            <text-input
                ref="formPath"
                label="Path"
                :disabled="busy"
                :defaultValue="form?.path || ''"
            />
            <ConstraintTypeDropdown
                ref="constraintType"
                :default="form?.submission_constraint_type || IFormSubmissionConstraintTypes.none"
            />
            <text-input
                ref="constraintLimit"
                :disabled="busy"
                label="Constraint Limit"
                :defaultValue="form?.submission_constraint_limit.toString() || ''"
            ></text-input>
            <error-display :error="error"></error-display>
            <c-button
                @click="onHandleSubmit"
            >
                {{ props.formInfo ? "Save" : "Create" }}
            </c-button>
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
import ConstraintTypeDropdown from '@/components/form/constraint-type-dropdown.vue'
import { createForm, updateForm } from '@/services/form.service';
import { ref, watch, type Ref } from 'vue';
import { useRouter } from 'vue-router';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { IFormSubmissionConstraintTypes, type IForm, type IFormCreateRequest, type IFormUpdateRequest } from '@/types';
import { formatPath, validatePath } from '@/utils/string';
import { handleError } from '@/utils/error';
import { routeNames } from '@/router/routeNames';

const router = useRouter()
const workspaceStore = useWorkspaceStore()

const props = defineProps({
    formInfo: {
        type: Object as () => IForm,
        default: null,
    },
})

const formName = ref<InstanceType<typeof TextInput>>()
const formPath = ref<InstanceType<typeof TextInput>>()
const constraintType = ref<InstanceType<typeof ConstraintTypeDropdown>>()
const constraintLimit = ref<InstanceType<typeof TextInput>>()

const form: Ref<IForm | null> = ref(props.formInfo || null)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onHandleSubmit = async () => {
    busy.value = true
    error.value = null
    try {
        const name = formName.value?.value
        const path = formPath.value?.value
        const ct = constraintType.value?.selected?.value
        const cl = constraintLimit.value?.value
        if (!name) {
            throw new Error('name cannot be blank')
        }
        if (!path || !validatePath(path)) {
            throw new Error('path cannot be blank')
        }
        if (!ct && ct !== '') {
            throw new Error('something went wrong')
        }
        if (cl && isNaN(parseInt(cl))) {
            throw new Error('constraint limit must be a number')
        }
        const payload: IFormUpdateRequest | IFormCreateRequest = {
            name,
            path,
            submission_constraint_type: ct,
            submission_constraint_limit: parseInt(cl || '0'),
        }
        if (props.formInfo) {
            await onUpdateForm(payload)
        } else {
            await onCreateForm(payload)
        }
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const onUpdateForm = async (payload: IFormUpdateRequest) => {
    if (!workspaceStore.currentWorkspace!) {
        throw new Error('Something went wrong')
    }
    if (!form.value) {
        throw new Error('Something went wrong')
    }
    updateForm(workspaceStore.currentWorkspace?.workspace_id, form.value.form_id, payload)
    form.value = {
        ...form.value,
        ...payload
    }
}

const onCreateForm = async (payload: IFormCreateRequest) => {
    if (!workspaceStore.currentWorkspace!) {
        throw new Error('Something went wrong')
    }
    await createForm(workspaceStore.currentWorkspace?.workspace_id, payload)
    router.push({ name: routeNames.forms })
}

watch(() => formName.value?.value, (newVal) => {
    if (formPath.value) {
        formPath.value.value = formatPath(newVal || '')
    }
})
</script>