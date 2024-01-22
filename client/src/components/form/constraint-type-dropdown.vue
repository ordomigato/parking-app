<template>
    <DropdownInput
        :selected="selected"
        :items="computedOptions"
        :disabled="disabled"
        @onSelect="onSelect"
    />
</template>
<script setup lang="ts">
import DropdownInput from '@/components/common/dropdown-input.vue'
import { type Ref, ref, onMounted, computed } from 'vue';
import { IFormDurationMeasurementUnits, type IDropdownItem } from '@/types'

const props = defineProps({
    default: {
        type: String as () => IFormDurationMeasurementUnits,
        default: IFormDurationMeasurementUnits.none
    },
    disabled: {
        type: Boolean,
        default: false
    },
    disabledOptions: {
        type: Array as () => IFormDurationMeasurementUnits[],
        default: () => []
    }
})

const selected: Ref<IDropdownItem<IFormDurationMeasurementUnits, string> | null> = ref(null)
const submissionTypes: Ref<IDropdownItem<IFormDurationMeasurementUnits, string>[]> = ref([
    {
        value: IFormDurationMeasurementUnits.none,
        name: 'None',
    },
    {
        value: IFormDurationMeasurementUnits.minutes,
        name: 'Minute(s)'
    },
    {
        value: IFormDurationMeasurementUnits.hours,
        name: 'Hour(s)'
    },
    {
        value: IFormDurationMeasurementUnits.days,
        name: 'Day(s)'
    },
    {
        value: IFormDurationMeasurementUnits.months,
        name: 'Month(s)'
    }
])

const computedOptions = computed(() => {
    return submissionTypes.value.filter(t => !props.disabledOptions.includes(t.value))
})

const onSelect = (payload: IDropdownItem<IFormDurationMeasurementUnits, string>) => {
    selected.value = payload
}

defineExpose({
    selected,
})

onMounted(() => {
    if (props.default) {
        selected.value = submissionTypes.value.find(x => x.value === props.default) || submissionTypes.value[0]
    } else {
        selected.value = submissionTypes.value[0]
    }
})
    
</script>