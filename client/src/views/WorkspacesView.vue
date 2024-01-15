<template>
    <div class="card">
        <header class="title-btn-combo">
            <h2>Workspaces</h2>
            <c-button class="ms-auto" @click="$router.push({ name: routeNames.createWorkspace })">+ Create Workspace</c-button>
        </header>
        <div class="table-container" v-if="workspaceStore.workspaces.length">
            <table class="text-left">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Created At</th>
                        <th>Updated At</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="ws in workspaceStore.workspaces" :key="ws.workspace_id">
                        <td>
                            <c-button @click="() => setActiveWorkspace(ws)" variant="link">{{ ws.name }}</c-button>
                        </td>
                        <td>{{ convertDate(ws.created_at) }}</td>
                        <td>{{ convertDate(ws.updated_at) }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div v-else>
            <p>No workspaces found. Create one to get started!</p>
        </div>
    </div>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import type { IWorkspace } from '@/types';
import { useRouter } from 'vue-router';
import { convertDate } from '@/utils/date'

const router = useRouter()

const workspaceStore = useWorkspaceStore()

const setActiveWorkspace = (ws: IWorkspace) => {
    workspaceStore.setActiveWorkspace(ws)
    router.push({ name: routeNames.dashboard })
}
</script>
