<template>
    <div>
        <div>
            <div class="mb-4">
                <text-input
                    ref="formName"
                    :disabled="busy"
                    label="Form Name"
                    :defaultValue="form?.name || ''"
                    @keyup.enter="onHandleSubmit"
                ></text-input>
                <text-input
                    ref="formPath"
                    :label="`Form Path: ${renderedFormPath}`"
                    :disabled="busy"
                    :defaultValue="defaultFormPathValue"
                    @keyup.enter="onHandleSubmit"
                />
                <FormConstraint
                    ref="cycleData"
                    :cycleData="props.formInfo.cycle_data"
                />
            </div>
            <error-display :error="error"></error-display>
            <c-button
                @click="onHandleSubmit"
                @keyup.enter="onHandleSubmit"
            >
                {{ props.formInfo ? "Save" : "Create" }}
            </c-button>
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
import FormConstraint from './form-constraint.vue';
import { createForm, updateForm, updateFormPath } from '@/services/form.service';
import { ref, watch, type Ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { type IForm, type IFormCreateRequest, type IFormUpdateRequest } from '@/types';
import { formatPath, validatePath } from '@/utils/string';
import { handleError } from '@/utils/error';
import { routeNames } from '@/router/routeNames';
import { useToastStore } from '@/stores/toastStore';


const router = useRouter()
const workspaceStore = useWorkspaceStore()
const toastStore = useToastStore()

const props = defineProps({
    formInfo: {
        type: Object as () => IForm,
        default: null,
    },
})

const formName = ref<InstanceType<typeof TextInput>>()
const formPath = ref<InstanceType<typeof TextInput>>()
const cycleData = ref<InstanceType<typeof FormConstraint>>()

const form: Ref<IForm | null> = ref(props.formInfo || null)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const renderedFormPath = computed(() => {
    if (workspaceStore.currentWorkspace) {
        return window.origin + workspaceStore.currentWorkspace?.path + formPath.value?.value
    } else {
        return ''
    }
})

const defaultFormPathValue = computed(() => {
    if (workspaceStore.currentWorkspace && props.formInfo) {
        return props.formInfo.path.path.replace(workspaceStore.currentWorkspace.path, "")
    } else {
        return ''
    }
})

const onHandleSubmit = async () => {
    busy.value = true
    error.value = null
    try {
        const name = formName.value?.value
        const fPath = formPath.value?.value
        const cd = cycleData.value?.getData()
    
        if (!name) {
            throw new Error('name cannot be blank')
        }
        if (!fPath) {
            throw new Error('path cannot be blank')
        }
        if (!validatePath(fPath)) {
            throw new Error('path is incorrectly formatted')
        }
        const payload: IFormUpdateRequest | IFormCreateRequest = {
            name,
            path: `${workspaceStore.currentWorkspace?.path}${fPath}`,
            cycle_data: cd ? cd : null,
        }
        if (props.formInfo) {
            await onUpdateForm(payload)
            if (props.formInfo.path.path !== payload.path) {
                await onUpdateFormPath(payload.path)
            }
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
    await updateForm(workspaceStore.currentWorkspace?.workspace_id, form.value.form_id, payload)
    toastStore.updateState('Form Successfully Updated', 'success')
    form.value = {
        ...form.value,
        ...payload
    }
}

const onUpdateFormPath = async (path: string) => {
    if (!workspaceStore.currentWorkspace!) {
        throw new Error('Something went wrong')
    }
    if (!form.value) {
        throw new Error('Something went wrong')
    }
    await updateFormPath(workspaceStore.currentWorkspace?.workspace_id, form.value.form_id, { path: path })
}

const onCreateForm = async (payload: IFormCreateRequest) => {
    if (!workspaceStore.currentWorkspace!) {
        throw new Error('Something went wrong')
    }
    const form = await createForm(workspaceStore.currentWorkspace?.workspace_id, payload)
    toastStore.updateState('Form Successfully Created', 'success')
    router.push({ name: routeNames.form, params: { id: form.form_id } })
}

watch(() => formName.value?.value, (newVal) => {
    if (formPath.value && !props.formInfo) {
        formPath.value.value = formatPath(newVal || '')
    }
})
</script>
<style>
.workspace-path {
    width: fit-content;
}
.at-container {
    padding: 0.25rem 1rem;
}
</style>