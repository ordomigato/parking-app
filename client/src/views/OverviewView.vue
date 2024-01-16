<template>
    <div class="card">
        <c-button
            class="mb-4"
            fullWidth
            @click="$router.push({ name: routeNames.forms })"
            @keyup.enter="$router.push({ name: routeNames.forms })"
        >
            View Parking Forms
        </c-button>
        <c-button
            fullWidth
            @click="$router.push({ name: routeNames.workspaceSettings, params: { id: workspaceStore.currentWorkspace?.workspace_id } })"
            @keyup.enter="$router.push({ name: routeNames.workspaceSettings, params: { id: workspaceStore.currentWorkspace?.workspace_id } })"
        >
            Manage Workspace Settings
        </c-button>
    </div>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { useWorkspaceStore } from '@/stores/workspaceStore';

const workspaceStore = useWorkspaceStore()
</script>