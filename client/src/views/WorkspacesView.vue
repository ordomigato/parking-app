<template>
    <div>
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
                        <c-button @click="() => setActiveWorkspace(ws)" isLink>{{ ws.name }}</c-button>
                    </td>
                    <td>{{ convertDate(ws.created_at) }}</td>
                    <td>{{ convertDate(ws.updated_at) }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import type { IWorkspace } from '@/types';
import { useRouter } from 'vue-router';

const router = useRouter()

const workspaceStore = useWorkspaceStore()

const convertDate = (date: Date) => {
    const d = new Date(date)
    return d.toLocaleString()
}

const setActiveWorkspace = (ws: IWorkspace) => {
    workspaceStore.setActiveWorkspace(ws)
    router.push({ name: routeNames.dashboard })
}
</script>
