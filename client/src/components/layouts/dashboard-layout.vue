<template>
    <div class="dashboard-container">
        <SideNav />
        <main>
            <RouterView :key="$route.fullPath" />
        </main>
    </div>
</template>
<script setup lang="ts">
import SideNav from '../nav/side-nav.vue'
import { RouterView } from 'vue-router'

</script>
<style lang="scss">
.dashboard-container {
    display: flex;
    main {
        position: fixed;
        right: 0;
        top: 0;
        bottom: 0;
        width: calc(100vw - 200px);
        padding: 1rem;
        background-color: #f0fdfa;
    }
}
</style>