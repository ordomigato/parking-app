<template>
    <div class="dashboard-container">
        <SideNav class="side-nav" />
        <div class="main-container">
            <div class="top-nav">
                <TopNav />
            </div>
            <main>
                <RouterView :key="$route.fullPath" />
            </main>
        </div>
    </div>
</template>
<script setup lang="ts">
import SideNav from '../nav/side-nav.vue'
import TopNav from '../nav/top-nav.vue'
import { RouterView } from 'vue-router'

</script>
<style lang="scss" scoped>
.dashboard-container {
    position: relative;
    display: flex;
    .side-nav {
        display: none;
    }
    .main-container {
        position: fixed;
        right: 0;
        top: 0;
        bottom: 0;
        width: 100%;
        background-color: var(--off-white-color);
        overflow: hidden;
        overflow-y: auto;
        .top-nav {
            position: sticky;
            top: 0;
        }
        main {
            padding: 1rem;
        }
    }
}

@media (min-width: 992px) {
    .dashboard-container {
        .main-container {
            width: calc(100vw - 200px);
        }
        .side-nav {
            display: block;
        }
    }
}
</style>