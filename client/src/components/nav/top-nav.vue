<template>
    <nav class="top-nav">
        <c-button
            class="workspace-btn"
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
                <UserIcon />
            </c-button>
        </div>
        <c-button
            variant="icon"
            class="hamburger-btn"
            @click="toggleMobileNav"
        >
            <HamburgerIcon />
        </c-button>
    </nav>
    <SideNav
        v-if="openMobileNav"
        class="mobile-nav"
        mobileView
        @close="toggleMobileNav"
    />
</template>
<script setup lang="ts">
import UserIcon from "@/components/icons/user-icon.vue"
import HamburgerIcon from "@/components/icons/hamburger-icon.vue"
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { routeNames } from '@/router/routeNames'
import { useUserStore } from "@/stores/userStore";
import SideNav from "./side-nav.vue";
import { ref } from "vue";

const workspaceStore = useWorkspaceStore()
const userStore = useUserStore()

const openMobileNav = ref(false)

const toggleMobileNav = () => {
    openMobileNav.value = !openMobileNav.value
}
</script>
<style lang="scss" scoped>
.top-nav {
    height: 48px;
    background-color: var(--main-color-light);
    color: white;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 1rem;
    .workspace-btn {
        white-space: nowrap;
        max-width: 300px;
        overflow: hidden;
        text-overflow: ellipsis;
        margin-right: 1rem;
    }
    .profile-container {
        display: flex;
        align-items: center;
        margin-left: auto;
        .username {
            margin-right: 1rem;
            color: var(--main-color);
            max-width: 300px;
            overflow: hidden;
            text-overflow: ellipsis;
        }
    }
    .hamburger-btn {
        margin-left: 1rem;
        flex: 1;
    }
}

@media (min-width: 992px) {
    .top-nav {
        .hamburger-btn {
            display: none;
        }
    }
    .mobile-nav {
        display: none;
    }
}
</style>