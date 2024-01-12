<template>
    <div class="card">
        <h2 class="text-center">Create Form</h2>
        <text-input
            ref="formName"
            label="Form Name"
            :disabled="busy"
        />
        <ConstraintTypeDropdown
            ref="constraintType"
            :disabled="busy"
        />
        <text-input
            ref="constraintLimit"
            label="Constraint Limit"
            :disabled="busy"
        ></text-input>
        <error-display :error="error"></error-display>
        <c-button @click="onFormSubmit">Submit</c-button>
    </div>
</template>
<script lang="ts" setup>
import TextInput from '@/components/global/TextInput.vue';
import ConstraintTypeDropdown from '@/components/form/constraint-type-dropdown.vue'
import { type Ref, ref } from 'vue';
import { handleError } from '../utils/error';
import { useRouter } from 'vue-router';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { routeNames } from '@/router/routeNames';
import { createForm } from '@/services/form.service';
import type { IFormCreateRequest } from '@/types';

const workspaceStore = useWorkspaceStore()

const router = useRouter()

const formName = ref<InstanceType<typeof TextInput>>()
const constraintType = ref<InstanceType<typeof ConstraintTypeDropdown>>()
const constraintLimit = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onFormSubmit = async () => {
    error.value = null
    busy.value = true
    try {
        const name = formName.value?.value
        const ct = constraintType.value?.selected?.value
        if (!workspaceStore.currentWorkspace) {
            throw new Error('Something went wrong')
        }
        if ((!ct && ct !== '') || ct === null) {
            throw new Error('Constraint type cannot be blank')
        }
        if (!name){
            throw new Error('Workspace name cannot be blank')
        }
        if (!constraintLimit.value) {
            throw new Error('Something went wrong')
        }
        const payload: IFormCreateRequest = {
            name: name,
            submission_constraint_type: ct,
            submission_constraint_limit: parseInt(constraintLimit.value.value)
        }
        await createForm(workspaceStore.currentWorkspace?.workspace_id, payload)
        router.push({ name: routeNames.forms })
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

</script>