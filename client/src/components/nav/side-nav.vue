<template>
    <nav :class="`side-nav ${props.mobileView ? 'mobile' : ''}`">
        <div>
            <div v-if="props.mobileView" class="cross-container">
                <c-button
                    variant="icon"
                    @click="() => emit('close')"
                >
                    <CrossIcon />
                </c-button>
            </div>
            <p class="text-sm">General</p>
            <ul v-if="workspaceStore.currentWorkspace">
                <li class="nav-item" v-for="i in navItems" :key="i.to">
                    <router-link class="text-sm" :to="i.to" @click="() => emit('close')">{{ i.name }}</router-link>
                </li>
            </ul>
        </div>
    </nav>
</template>
<script setup lang="ts">
import CrossIcon from '@/components/icons/cross-icon.vue'
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

const props = defineProps({
    mobileView: {
        type: Boolean,
        default: false,
    }
})

const emit = defineEmits(['close'])

const navItems: Ref<NavItem[]> = ref([])

const setNavItems = () => {
    navItems.value = [
        new NavItem('Overview', '/dashboard/overview'),
        new NavItem('Parking Forms', '/dashboard/form'),
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
.side-nav {
    min-width: 200px;
    background-color: var(--main-color-dark);
    color: white;
    padding: 1rem;
    height: 100vh;
    height: 100dvh;
    position: fixed;
    z-index: 10;
    &.mobile {
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        z-index: 1;
        padding: 0.5rem 1rem;
    }
    .cross-container {
        display: flex;
        justify-content: end;
    }
    .nav-item a {
        display: block;
        padding: 0.25rem 0.5rem;
        border-radius: 4px;
        margin-bottom: 0.25rem;
        &:hover:not(.router-link-active) {
            background-color: #164e63;
        }
        &.router-link-active {
            opacity: 0.25;
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
