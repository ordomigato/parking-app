<template>
    <nav>
        <p @click="$router.push({ name: routeNames.workspaces })" class="active-workspace">
            <strong>{{ workspaceStore.currentWorkspace?.name || 'Select Workspace' }}</strong>
        </p>
    </nav>
</template>
<script setup lang="ts">
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { routeNames } from '@/router/routeNames'

const workspaceStore = useWorkspaceStore()
</script>
<style lang="scss" scoped>
nav {
    height: 48px;
    background-color: #0284c7;
    color: white;
    display: flex;
    align-items: center;
    padding: 0 1rem;
    .active-workspace {
        color: white;
        padding: 0.25rem 1rem;
        background-color: #0369a1;
        border-radius: 4px;
        min-width: 200px;
        cursor: pointer;
    }
}
</style>