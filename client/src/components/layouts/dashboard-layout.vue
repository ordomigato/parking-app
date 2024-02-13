<template>
    <div class="dashboard-container">
        <SideNav class="side-nav" />
        <div ref="mainContainer" class="main-container">
            <div :class="`top-nav ${isTop ? '' : 'scrolling'}`">
                <TopNav />
            </div>
            <main>
                <RouterView :key="$route.fullPath" />
            </main>
        </div>
    </div>
</template>
<script setup lang="ts">
import { onMounted, onUnmounted, ref, type Ref } from 'vue';
import SideNav from '../nav/side-nav.vue'
import TopNav from '../nav/top-nav.vue'
import { RouterView } from 'vue-router'

const isTop = ref(true)

const mainContainer: Ref<HTMLElement | undefined> = ref()

const handleScroll = () => {
    if (mainContainer.value) {
        isTop.value = mainContainer.value?.scrollTop === 0
    }
}

onMounted(() => {
    mainContainer.value?.addEventListener('scroll', handleScroll, false)
})

onUnmounted(() => {
    mainContainer.value?.removeEventListener('scroll', handleScroll, false)
})

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
            z-index: 999;
            transition: all 0.1s linear;
            &.scrolling {
                box-shadow: 0 3px 12px rgba(0,0,0,0.1);
                border-bottom: 1px solid var(--main-color-light);
            }
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