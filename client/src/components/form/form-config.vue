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
                <label class="d-flex align-items-center">
                    <span class="text-sm leading-6 font-medium me-2">Enable cycle constraints:</span>
                    <input
                        type="checkbox"
                        v-model="enableCycle"
                    />
                </label>
                <div v-if="enableCycle" class="p-2">
                    <p class="label">Cycle Interval</p>
                    <div class="d-flex align-items-end">
                        <text-input
                            class="me-4"
                            ref="durationIntervalLimit"
                            type="number"
                            :min="1"
                            :disabled="busy || !enableCycle"
                            :defaultValue="form?.cycle_data.reset_interval.value.toString() || '1'"
                            @keyup.enter="onHandleSubmit"
                        ></text-input>
                        <div>
                            <ConstraintTypeDropdown
                                ref="durationIntervalMeasurementUnit"
                                :disabledOptions="[
                                    IFormDurationMeasurementUnits.none,
                                    IFormDurationMeasurementUnits.minutes,
                                    IFormDurationMeasurementUnits.hours,
                                ]"
                                :default="form?.cycle_data.reset_interval.unit || IFormDurationMeasurementUnits.months"
                                :disabled="busy || !enableCycle"
                            />
                        </div>
                    </div>
                    <div v-if="durationIntervalMeasurementUnit?.selected?.value === IFormDurationMeasurementUnits.months">
                        <p class="label">Cycle Reset Day (of month) + Time</p>
                        <div class="d-flex">
                            <text-input
                                ref="durationResetDay"
                                type="number"
                                :disabled="busy || !enableCycle"
                                :min="1"
                                :max="31"
                                defaultValue="1"
                                @keyup.enter="onHandleSubmit"
                            ></text-input>
                            <div class="at-container"><span>@</span></div>
                            <text-input
                                ref="durationResetTime"
                                type="time"
                                :disabled="busy || !enableCycle"
                                defaultValue="00:00"
                                @keyup.enter="onHandleSubmit"
                            ></text-input>
                        </div>
                    </div>
                    <div class="d-flex align-items-end">
                        <text-input
                            class="me-4"
                            ref="durationLimit"
                            type="number"
                            :min="1"
                            :disabled="busy || !enableCycle"
                            label="Duration Limit"
                            :defaultValue="form?.cycle_data.duration_limit.value.toString() || '1'"
                            @keyup.enter="onHandleSubmit"
                        ></text-input>
                        <div>
                            <ConstraintTypeDropdown
                                ref="durationMeasurementUnit"
                                :disabledOptions="durationLimitDisableOptions"
                                :default="form?.cycle_data.duration_limit.unit || IFormDurationMeasurementUnits.hours"
                                :disabled="busy || !enableCycle"
                            />
                        </div>
                    </div>
                </div>
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
import ConstraintTypeDropdown from '@/components/form/constraint-type-dropdown.vue'
import { createForm, updateForm, updateFormPath } from '@/services/form.service';
import { ref, watch, type Ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { IFormDurationMeasurementUnits, type IForm, type IFormCreateRequest, type IFormUpdateRequest } from '@/types';
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
const enableCycle = ref(false)
const durationIntervalMeasurementUnit = ref<InstanceType<typeof ConstraintTypeDropdown>>()
const durationIntervalLimit = ref<InstanceType<typeof TextInput>>()
const durationMeasurementUnit = ref<InstanceType<typeof ConstraintTypeDropdown>>()
const durationLimit = ref<InstanceType<typeof TextInput>>()
const durationResetDay = ref<InstanceType<typeof TextInput>>()
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

const durationLimitDisableOptions = computed(() => {
    let disabledOptions = [ IFormDurationMeasurementUnits.none ]
    if (durationIntervalMeasurementUnit.value?.selected?.value === IFormDurationMeasurementUnits.months) {
        disabledOptions.push(IFormDurationMeasurementUnits.months)
    }
    if (durationIntervalMeasurementUnit.value?.selected?.value === IFormDurationMeasurementUnits.days) {
        disabledOptions.push(IFormDurationMeasurementUnits.months)
        disabledOptions.push(IFormDurationMeasurementUnits.days)
    }
    return disabledOptions
})

const onHandleSubmit = async () => {
    busy.value = true
    error.value = null
    try {
        const name = formName.value?.value
        const fPath = formPath.value?.value
        const dimu = durationIntervalMeasurementUnit.value?.selected?.value
        const dil = durationIntervalLimit.value?.value
        const du = durationMeasurementUnit.value?.selected?.value
        const dl = durationLimit.value?.value
        const drd = durationResetDay.value?.value
        const drt = durationResetTime.value?.value
    
        // configure reference reset date + time
        const refDate = new Date()
        const day = drd ? parseInt(drd) : 1
        const [ hour, min ] = drt ? drt.split(":").map(v => parseInt(v)) : [0, 0]
        refDate.setHours(hour)
        refDate.setMinutes(min)
        refDate.setSeconds(0)
        refDate.setMilliseconds(0)
        refDate.setDate(day)

        if (!name) {
            throw new Error('name cannot be blank')
        }
        if (!fPath) {
            throw new Error('path cannot be blank')
        }
        if (!validatePath(fPath)) {
            throw new Error('Path is incorrectly formatted')
        }
        if (dl && isNaN(parseInt(dl))) {
            throw new Error('constraint limit must be a number')
        }
        const payload: IFormUpdateRequest | IFormCreateRequest = {
            name,
            path: `${workspaceStore.currentWorkspace?.path}${fPath}`,
            cycle_data: {
                enable_cycle: enableCycle.value,
                duration_limit: {
                    value: enableCycle.value ? parseInt(dl || '0') : 0,
                    unit: du || IFormDurationMeasurementUnits.none,
                },
                reset_interval: {
                    value: dil ? parseInt(dil) : 0,
                    unit: dimu || IFormDurationMeasurementUnits.none,
                    reference_date: refDate.toISOString()
                }
            }
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

onMounted(() => {
    if (props.formInfo?.cycle_data.enable_cycle) {
        enableCycle.value = true
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