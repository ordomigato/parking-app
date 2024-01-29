<template>
    <div>
        <label class="d-flex align-items-center">
            <span class="text-sm leading-6 font-medium me-2">Enable cycle constraint:</span>
            <input
                type="checkbox"
                v-model="enableCycle"
            />
        </label>
        <div v-if="enableCycle" class="section">
            <p class="label">Max Cycle Duration</p>
            <div class="options-container">
                <text-input
                    ref="selectedDurationLimit"
                    :defaultValue="props.cycleData?.duration_limit.value.toString() || '1'"
                    @keyup.enter="onHandleContinue"
                />
                <DropdownInput
                    :selected="selectedDurationLimitUnit"
                    :items="durationLimitUnits"
                    :disabled="props.disabled"
                    @onSelect="onSelectedDurationLimitUnit"
                />
                <p class="text-center">per</p>
                <DropdownInput
                    :selected="selectedResetIntervalUnit"
                    :items="resetIntervalUnits"
                    :disabled="props.disabled"
                    @onSelect="onSelectedResetIntervalUnit"
                />
            </div>
            <p class="label">Reset Time</p>
            <text-input
                ref="time"
                type="time"
                :defaultValue="getDefaultTime(
                    props.cycleData
                    ? new Date(props.cycleData.reset_interval.reference_date)
                    : null
                )"
            />
            <text-input
                ref="dayOfMonth"
                v-if="enableDayOfMonth"
                type="number"
                label="Reset Day"
                :min="1"
                :max="31"
                :defaultValue="getDefaultDay(
                    props.cycleData
                    ? new Date(props.cycleData.reset_interval.reference_date)
                    : null
                )"
            ></text-input>
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
import DropdownInput from '@/components/common/dropdown-input.vue'
import { IFormDurationMeasurementUnits, type IDropdownItem, type IDurationUnitMeasurement, type ICycleData, type IResetInterval } from '@/types'
import { computed, ref, type ComputedRef, type Ref, onMounted } from 'vue';

type IUnitDropdownItem = IDropdownItem<IFormDurationMeasurementUnits, string>

const durationUnitOptions: IUnitDropdownItem[] = [
    {
        name: 'Minute(s)',
        value: IFormDurationMeasurementUnits.minutes,
    },
    {
        name: 'Hour(s)',
        value: IFormDurationMeasurementUnits.hours,
    },
    {
        name: 'Day(s)',
        value: IFormDurationMeasurementUnits.days,
    }
]

const resetIntervalUnitOptions: IUnitDropdownItem[] = [
    {
        name: 'Day',
        value: IFormDurationMeasurementUnits.days,
    },
    {
        name: 'Month',
        value: IFormDurationMeasurementUnits.months,
    }
]

const props = defineProps({
    disabled: {
        type: Boolean,
        default: false,
    },
    cycleData: {
        type: Object as () => ICycleData | null,
        default: null
    }
})

const emit = defineEmits(['next'])

const enableCycle = ref(false)
const selectedResetIntervalUnit: Ref<IUnitDropdownItem | null> = ref(null)
const selectedDurationLimitUnit: Ref<IUnitDropdownItem | null> = ref(null)
const selectedDurationLimit = ref<InstanceType<typeof TextInput>>()
const time = ref<InstanceType<typeof TextInput>>()
const dayOfMonth = ref<InstanceType<typeof TextInput>>()

const enableDayOfMonth: ComputedRef<boolean> = computed(() => {
    return !!(selectedResetIntervalUnit.value?.value === IFormDurationMeasurementUnits.months)
})

const durationLimitUnits = computed(() => {
    if (selectedResetIntervalUnit.value?.value === IFormDurationMeasurementUnits.days) {
        return durationUnitOptions.filter(o => o.value !== IFormDurationMeasurementUnits.days)
    }
    return durationUnitOptions
})
const resetIntervalUnits = computed(() => {
    return resetIntervalUnitOptions
})

const computedDurationLimit: ComputedRef<IDurationUnitMeasurement | null> = computed(() => {
    if (!selectedDurationLimitUnit.value?.value || !selectedDurationLimit.value?.value) {
        return null
    }
    return {
        unit: selectedDurationLimitUnit.value.value,
        value: parseInt(selectedDurationLimit.value.value)
    }
})

const computedResetInterval: ComputedRef<IResetInterval | null> = computed(() => {

    const day = dayOfMonth.value?.value || '1'
    const t = time.value?.value || '00:00'

    // configure reference reset date + time
    const refDate = new Date()
    const [ hour, min ] = t ? t.split(":").map(v => parseInt(v)) : [0, 0]
    refDate.setHours(hour)
    refDate.setMinutes(min)
    refDate.setSeconds(0)
    refDate.setMilliseconds(0)
    refDate.setDate(parseInt(day))

    switch (selectedResetIntervalUnit.value?.value) {
    case IFormDurationMeasurementUnits.days:
        return {
            unit: IFormDurationMeasurementUnits.days,
            value: 1,
            reference_date: refDate,
        }
    case IFormDurationMeasurementUnits.months:
        return {
            unit: IFormDurationMeasurementUnits.months,
            value: 1,
            reference_date: refDate,
        }
    default:
        return null
    }
})

const onSelectedDurationLimitUnit = (payload: IUnitDropdownItem) => {
    selectedDurationLimitUnit.value = payload
}

const onSelectedResetIntervalUnit = (payload: IUnitDropdownItem) => {
    selectedResetIntervalUnit.value = payload
    if (
        payload.value === IFormDurationMeasurementUnits.days
        && selectedDurationLimitUnit.value?.value === IFormDurationMeasurementUnits.days
    ) {
        selectedDurationLimitUnit.value = null
    }
}

const getDefaultDay = (date: Date | null) => {
    if (!date) {
        return '1'
    }
    return date.getDate().toString()
}

const getDefaultTime = (date: Date | null) => {
    if (!date) {
        return '00:00'
    }
    const hour = ("0" + date.getHours()).slice(-2)
    const min = ("0" + date.getMinutes()).slice(-2)
    return `${hour}:${min}`
}

const onHandleContinue = () => {
    emit('next')
}

const getData = (): ICycleData => {
    if (!computedResetInterval.value) {
        throw new Error('reset interval data is not properly configured')
    }
    if (!computedDurationLimit.value) {
        throw new Error('duration limit data is not properly configured')
    }
    return {
        reset_interval: computedResetInterval.value,
        duration_limit: computedDurationLimit.value,
    }
}

onMounted(() => {
    // set default values
    selectedResetIntervalUnit.value =
        resetIntervalUnitOptions.find(o => o.value === props.cycleData?.reset_interval.unit)
        || resetIntervalUnitOptions[1]
    selectedDurationLimitUnit.value =
        durationUnitOptions.find(o => o.value === props.cycleData?.duration_limit.unit)
        || durationUnitOptions[2]

    if (props.cycleData) {
        enableCycle.value = true
    }
})

defineExpose({
    getData,
})

</script>
<style lang="scss" scoped>
.options-container {
    display: flex;
    flex-direction: column;
    p {
        margin-bottom: 1rem;
    }
}
</style>
