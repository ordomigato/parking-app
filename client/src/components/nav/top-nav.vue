<template>
    <nav>
        <c-button
            @click="$router.push({ name: routeNames.workspaces })"
            @keyup.enter="$router.push({ name: routeNames.workspaces })"
        >
            <strong>{{ workspaceStore.currentWorkspace?.name || 'Select Workspace' }}</strong>
        </c-button>
        <div class="profile-container">
            <p class="username">{{ userStore.user?.username }}</p>
            <c-button
                variant="icon"
                @click="$router.push({ name: routeNames.profile })"
                @keyup.enter="$router.push({ name: routeNames.profile })"
            >
                <UserIcon color="white" />
            </c-button>
        </div>
    </nav>
</template>
<script setup lang="ts">
import UserIcon from "@/components/icons/user-icon.vue"
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { routeNames } from '@/router/routeNames'
import { useUserStore } from "@/stores/userStore";

const workspaceStore = useWorkspaceStore()
const userStore = useUserStore()
</script>
<style lang="scss" scoped>
nav {
    height: 48px;
    background-color: var(--main-color-light);
    color: white;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 1rem;
    .profile-container {
        display: flex;
        align-items: center;
        .username {
            margin-right: 1rem;
            color: white;
        }
    }
}
</style>