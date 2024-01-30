<template>
    <div v-if="workspaceStore.currentWorkspace">
        <div class="card">
            <h2>Workspace Settings</h2>
            <WorkspaceConfig isUpdate />
        </div>
        <div class="card">
            <DeleteWorkspace
                :workspace="workspaceStore.currentWorkspace"
            />
        </div>
    </div>
</template>
<script setup lang="ts">
import WorkspaceConfig from '@/components/workspace/workspace-config.vue';
import DeleteWorkspace from '@/components/workspace/delete-workspace.vue'
import { useWorkspaceStore } from '@/stores/workspaceStore';

const workspaceStore = useWorkspaceStore()
</script>
