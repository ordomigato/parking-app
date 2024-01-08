<template>
    <div>
        <text-input ref="workspaceName" />
        <c-button @click="onWorkspaceSubmit">Submit</c-button>
    </div>
</template>
<script lang="ts" setup>
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
        const resp = await createWorkspace(workspaceName.value.value)
        console.log(resp)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

</script>