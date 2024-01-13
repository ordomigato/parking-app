<template>
    <div class="dashboard-container">
        <SideNav />
        <div class="main-container">
            <TopNav />
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
<style lang="scss">
.dashboard-container {
    display: flex;
    .main-container {
        position: fixed;
        right: 0;
        top: 0;
        bottom: 0;
        width: calc(100vw - 200px);
        background-color: var(--off-white-color);
        overflow: hidden;
        overflow-y: auto;
        main {
            padding: 1rem;
        }
    }
}
</style>