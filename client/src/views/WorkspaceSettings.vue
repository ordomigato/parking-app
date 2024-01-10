<template>
    <div v-if="workspaceStore.currentWorkspace">
        <div class="card">
            <text-input
                ref="workspaceName"
                :disabled="busy"
                label="Workspace Name"
                :defaultValue="workspaceStore.currentWorkspace.name"
            ></text-input>
            <c-button
                @click="onUpdateWorkspace"
            >
                Save
            </c-button>
        </div>
        <div class="card">
            <h2 class="danger-text">Danger Zone</h2>
            <c-button
                :disabled="busy"
                @click="onDeleteWorkspace"
            >
                Delete Workspace
            </c-button>
            <error-display :error="error"></error-display>
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { deleteWorkspace, updateWorkspace } from '@/services/workspace.service'
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { useRouter } from 'vue-router';
import { routeNames } from '@/router/routeNames';

const workspaceStore = useWorkspaceStore()
const router = useRouter()

const workspaceName = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onUpdateWorkspace = async () => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error("Missing workspace ID")
        }
        if (!workspaceName.value) {
            throw new Error("Something went wrong")
        }
        await updateWorkspace(workspaceStore.currentWorkspace.workspace_id, workspaceName.value.value)
        workspaceStore.setActiveWorkspace({
            ...workspaceStore.currentWorkspace,
            name: workspaceName.value.value
        })
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const onDeleteWorkspace = async () => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace) {
            throw new Error("Missing workspace ID")
        }
        await deleteWorkspace(workspaceStore.currentWorkspace.workspace_id)
        workspaceStore.setActiveWorkspace(null)
        router.push({ name: routeNames.workspaces })
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}
</script>
