<script setup lang="ts">
import WindowView from '@/components/operating-system/WindowView.vue'
import { useWindowsStore } from '@/stores/windows'

const windowsStore = useWindowsStore()

</script>

<template>
  <WindowView
    @click="windowsStore.clickWindow(w)"
    v-for="w in windowsStore.windowList"
    :key="w.id"
    :title="w.title"
    :icon="w.icon"
    :page="w.page"
    :id="w.id"
    :width="w.width"
    :height="w.height"
    :x="w.x"
    :y="w.y"
    :theme="w.theme"
    :style="{zIndex: w.z}"
  />
</template>

<style scoped>

</style>
