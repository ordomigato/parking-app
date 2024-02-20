<template>
    <div class="breadcrumb-container">
        <span>
            <router-link class="link" :to="{ name: routeNames.workspaces }">Workspaces</router-link>
        </span>
        <span v-if="workspaceStore.currentWorkspace?.workspace_id">
            / <router-link class="link" :to="{ name: routeNames.overview }">
                {{ workspaceStore.currentWorkspace?.name || 'Select Workspace' }}
            </router-link>
        </span>
        <span v-if="workspaceStore.currentWorkspace && workspaceStore.currentForm">
            / <router-link class="link" :to="{ name: routeNames.forms }">Forms</router-link>
        </span>
        <span v-if="workspaceStore.currentForm">
            / {{ workspaceStore.currentForm.name }}
        </span>
    </div>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { useWorkspaceStore } from '@/stores/workspaceStore';

const workspaceStore = useWorkspaceStore()
</script>
<style lang="scss" scoped>
.breadcrumb-container {
    color: var(--main-color);
    font-size: 0.875rem;
}
</style>