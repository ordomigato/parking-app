<template>
    <p class="block text-sm leading-6 font-medium">Constraint Type</p>
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
import { IFormSubmissionConstraintTypes, type IDropdownItem } from '@/types'

const props = defineProps({
    default: {
        type: String as () => IFormSubmissionConstraintTypes,
        default: IFormSubmissionConstraintTypes.none
    },
    disabled: {
        type: Boolean,
        default: false
    }
})

const selected: Ref<IDropdownItem<IFormSubmissionConstraintTypes, string> | null> = ref(null)
const submissionTypes: Ref<IDropdownItem<IFormSubmissionConstraintTypes, string>[]> = ref([
    {
        value: IFormSubmissionConstraintTypes.none,
        name: 'None'
    },
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

onMounted(() => {
    if (props.default) {
        selected.value = submissionTypes.value.find(x => x.value === props.default) || submissionTypes.value[0]
    } else {
        selected.value = submissionTypes.value[0]
    }
})
    
</script>