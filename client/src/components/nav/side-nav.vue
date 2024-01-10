<template>
    <nav>
        <div>
            <p class="text-sm">MENU</p>
            <ul v-if="workspaceStore.currentWorkspace">
                <li class="nav-item" v-for="i in navItems" :key="i.to">
                    <router-link class="text-sm" :to="i.to">{{ i.name }}</router-link>
                </li>
            </ul>
        </div>
    </nav>
</template>
<script setup lang="ts">
import { useWorkspaceStore } from '@/stores/workspaceStore';
import { onMounted, ref, type Ref } from 'vue';

const workspaceStore = useWorkspaceStore()

class NavItem {
    name: string;
    to: string;

    constructor(name: string, to: string) {
        this.name = name
        this.to = to
    }
}

const navItems: Ref<NavItem[]> = ref([])

const setNavItems = () => {
    navItems.value = [
        new NavItem('Overview', '/dashboard/overview'),
        new NavItem('Forms', '/dashboard/form'),
        new NavItem('Settings', `/dashboard/workspace/${workspaceStore.currentWorkspace?.workspace_id}`),
    ]
}

workspaceStore.$subscribe((mutation, state) => {
    if (state.currentWorkspace) {
        setNavItems()
    }
})

onMounted(() => {
    if (workspaceStore.currentWorkspace) {
        setNavItems()
    }
})
</script>
<style lang="scss" scoped>
nav {
    min-width: 200px;
    background-color: #083344;
    color: white;
    padding: 1rem;
    height: 100vh;
    position: fixed;
    .nav-item a {
        display: block;
        padding: 0.25rem 0.5rem;
        border-radius: 4px;
        margin-bottom: 0.25rem;
        &:hover {
            background-color: #164e63;
        }
    }
    p {
        color: white;
        padding: 0.75rem 0.5rem;
        font-weight: 400;
        opacity: 0.5;
    }
}
</style>
