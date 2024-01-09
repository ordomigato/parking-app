<template>
    <div>
        <h2 class="text-center">Create Workspace</h2>
        <text-input ref="workspaceName" label="Workspace Name" />
        <c-button @click="onWorkspaceSubmit">Submit</c-button>
    </div>
</template>
<script lang="ts" setup>
import TextInput from '@/components/global/TextInput.vue';
import { type Ref, ref } from 'vue';
import { handleError } from '../utils/error';
import { createWorkspace } from '../services/workspace.service'

const workspaceName = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onWorkspaceSubmit = async () => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceName.value?.value){
            throw new Error('Workspace name cannot be blank')
        }
        const resp = await createWorkspace(workspaceName.value.value)
        console.log(resp)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

</script>