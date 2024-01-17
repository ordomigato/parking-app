<template>
    <p class="block text-sm leading-6 font-medium">Duration Type</p>
    <DropdownInput
        :selected="selected"
        :items="submissionTypes"
        :disabled="disabled"
        @onSelect="onSelect"
    />
</template>
<script setup lang="ts">
import DropdownInput from '@/components/common/dropdown-input.vue'
import { type Ref, ref, onMounted } from 'vue';
import { IFormDurationMeasurementUnits, type IDropdownItem } from '@/types'

const props = defineProps({
    default: {
        type: String as () => IFormDurationMeasurementUnits,
        default: IFormDurationMeasurementUnits.none
    },
    disabled: {
        type: Boolean,
        default: false
    }
})

const selected: Ref<IDropdownItem<IFormDurationMeasurementUnits, string> | null> = ref(null)
const submissionTypes: Ref<IDropdownItem<IFormDurationMeasurementUnits, string>[]> = ref([
    {
        value: IFormDurationMeasurementUnits.none,
        name: 'None'
    },
    {
        value: IFormDurationMeasurementUnits.minutes,
        name: 'Minutes'
    },
    {
        value: IFormDurationMeasurementUnits.hours,
        name: 'Hours'
    },
    {
        value: IFormDurationMeasurementUnits.days,
        name: 'Days'
    },
    {
        value: IFormDurationMeasurementUnits.months,
        name: 'Months'
    }
])

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