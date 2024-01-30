<template>
    <div>
        <text-input
            ref="workspaceName"
            label="Workspace Name"
            :default-value="isUpdate ? workspaceStore.currentWorkspace?.name : ''"
            :disabled="busy"
            @keyup.enter="handleSubmit"
        />
        <text-input
            ref="workspacePath"
            label="Base Path"
            :default-value="isUpdate ? workspaceStore.currentWorkspace?.path : ''"
            :disabled="busy || isUpdate"
            @keyup.enter="handleSubmit"
        />
        <error-display :error="error"></error-display>
        <c-button
            @click="handleSubmit"
            @keyup.enter="handleSubmit"
            :disabled="busy || !hasChanged"
        >
            {{ isUpdate ? "Save" : "Create" }}
        </c-button>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
import { handleError } from '@/utils/error';
import { ref, watch, type Ref, computed, type ComputedRef } from 'vue';
import { updateWorkspace, createWorkspace } from '@/services/workspace.service'
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { useRouter } from 'vue-router';
import { routeNames } from '@/router/routeNames';
import { formatPath, validatePath } from '@/utils/string';
import type { IWorkspaceCreateRequest, IWorkspaceUpdateRequest } from '@/types'
import { useToastStore } from '@/stores/toastStore';

const workspaceStore = useWorkspaceStore()
const toastStore = useToastStore()
const router = useRouter()

const workspaceName = ref<InstanceType<typeof TextInput>>()
const workspacePath = ref<InstanceType<typeof TextInput>>()

const props = defineProps({
    isUpdate: {
        type: Boolean,
        default: false
    }
})

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const hasChanged: ComputedRef<boolean> = computed(() => {
    return workspaceName.value?.value !== workspaceStore.currentWorkspace?.name
})

const handleSubmit = async () => {
    if (!hasChanged.value) {
        return
    }
    error.value = null
    busy.value = true
    try {
        const name = workspaceName.value?.value
        const path = workspacePath.value?.value
        if (!name) {
            throw new Error("Name cannot be blank")
        }
        if (!path) {
            throw new Error('Path cannot be blank')
        }
        if (!validatePath(path)) {
            throw new Error('Path is incorrectly formatted')
        }
        const payload = { path, name }
        if (props.isUpdate) {
            await onUpdateWorkspace(payload)
        } else {
            await onCreateWorkspace(payload)
        }
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const onUpdateWorkspace = async (payload: IWorkspaceUpdateRequest) => {
    if (!workspaceStore.currentWorkspace) {
        throw new Error("Something went wrong")
    }
    await updateWorkspace(workspaceStore.currentWorkspace.workspace_id, payload)
    workspaceStore.setActiveWorkspace({
        ...workspaceStore.currentWorkspace,
        ...payload,
    })
    toastStore.updateState("Successfully Updated Workspace", "success")
}

const onCreateWorkspace = async (payload: IWorkspaceCreateRequest) => {
    const resp = await createWorkspace(payload)
    workspaceStore.setActiveWorkspace(resp)
    router.push({ name: routeNames.dashboard })
    toastStore.updateState("Successfully Created Workspace", "success")
}

watch(() => workspaceName.value?.value, (newVal) => {
    if (workspacePath.value && !props.isUpdate) {
        workspacePath.value.value = formatPath(newVal || '')
    }
})
</script>