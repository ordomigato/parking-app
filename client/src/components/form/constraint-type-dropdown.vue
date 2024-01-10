<template>
    <DropdownInput
        :selected="selected"
        :items="submissionTypes"
        @onSelect="onSelect"
    />
</template>
<script setup lang="ts">
import DropdownInput from '@/components/common/dropdown-input.vue'
import { type Ref, ref } from 'vue';
import { IFormSubmissionConstraintTypes, type IDropdownItem } from '@/types'

const selected: Ref<IDropdownItem<IFormSubmissionConstraintTypes, string> | null> = ref(null)
const submissionTypes: Ref<IDropdownItem<IFormSubmissionConstraintTypes, string>[]> = ref([
    {
        value: IFormSubmissionConstraintTypes.minutes,
        name: 'Minutes'
    },
    {
        value: IFormSubmissionConstraintTypes.hours,
        name: 'Hours'
    },
    {
        value: IFormSubmissionConstraintTypes.days,
        name: 'Days'
    },
    {
        value: IFormSubmissionConstraintTypes.months,
        name: 'Months'
    }
])

const onSelect = (payload: IDropdownItem<IFormSubmissionConstraintTypes, string>) => {
    selected.value = payload
}

defineExpose({
    selected,
})
    
</script>