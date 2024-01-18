<template>
    <div>
        <div>
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
            <div class="d-flex">
                <text-input
                    class="me-4"
                    ref="durationLimit"
                    :disabled="busy || !enableDuration"
                    label="Duration Limit"
                    :defaultValue="form?.duration_limit.toString() || ''"
                    @keyup.enter="onHandleSubmit"
                ></text-input>
                <div>
                    <ConstraintTypeDropdown
                        ref="durationMeasurementUnit"
                        :default="form?.duration_measurement_unit || IFormDurationMeasurementUnits.none"
                    />
                </div>
            </div>
            <text-input
                ref="durationResetTime"
                type="time"
                :disabled="busy"
                label="Duration Reset Time"
                defaultValue="00:00"
                @keyup.enter="onHandleSubmit"
            ></text-input>
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
import ConstraintTypeDropdown from '@/components/form/constraint-type-dropdown.vue'
import { createForm, updateForm, updateFormPath } from '@/services/form.service';
import { ref, watch, type Ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { IFormDurationMeasurementUnits, type IForm, type IFormCreateRequest, type IFormUpdateRequest } from '@/types';
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
const durationMeasurementUnit = ref<InstanceType<typeof ConstraintTypeDropdown>>()
const durationLimit = ref<InstanceType<typeof TextInput>>()
const durationResetTime = ref<InstanceType<typeof TextInput>>()

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

const enableDuration = computed(() => {
    if (durationMeasurementUnit.value?.selected?.value !== "") {
        return true
    } else {
        return false
    }
})

const onHandleSubmit = async () => {
    busy.value = true
    error.value = null
    try {
        const name = formName.value?.value
        const fPath = formPath.value?.value
        const ct = durationMeasurementUnit.value?.selected?.value
        const cl = durationLimit.value?.value
        const drt = durationResetTime.value?.value
        if (!name) {
            throw new Error('name cannot be blank')
        }
        if (!fPath) {
            throw new Error('path cannot be blank')
        }
        if (!validatePath(fPath)) {
            throw new Error('Path is incorrectly formatted')
        }
        if (cl && isNaN(parseInt(cl))) {
            throw new Error('constraint limit must be a number')
        }
        const payload: IFormUpdateRequest | IFormCreateRequest = {
            name,
            path: `${workspaceStore.currentWorkspace?.path}${fPath}`,
            duration_measurement_unit: ct || IFormDurationMeasurementUnits.none,
            duration_limit: enableDuration.value ? parseInt(cl || '0') : 0,
            // duration_reset_time: drt || ''
            duration_reset_time: '1 year'
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
</style>