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
import { useRouter } from 'vue-router';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { routeNames } from '@/router/routeNames';

const workspaceStore = useWorkspaceStore()

const workspaceName = ref<InstanceType<typeof TextInput>>()
const router = useRouter()

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
        workspaceStore.setActiveWorkspace(resp)
        router.push({ name: routeNames.dashboard })

    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

</script>