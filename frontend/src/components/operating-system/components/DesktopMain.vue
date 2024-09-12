<script setup lang="ts">

import DynamicBackground from '@/components/operating-system/components/DynamicBackground.vue'
import IconView from '@/components/operating-system/IconView.vue'
import { useWindowsStore } from '@/stores/windows'

const windowsStore = useWindowsStore()
</script>

<template>
  <div class="absolute top-8 bottom-0 left-0 right-0 z-[1] flex">
    <div
      class="font-bold flex flex-col flex-wrap gap-6 p-6 select-none">
      <div class="aspect-square h-[78px] rounded-3xl flex items-center justify-center cursor-pointer"
           @click="windowsStore.openWindow(app)" v-for="(app, index) in windowsStore.desktopApps" :key="index">
        <div class="flex flex-col items-center gap-1">
          <DynamicBackground :background="app.background"
                             class="w-16 h-16 aspect-square rounded-2xl bg-white/100 overflow-hidden">
            <IconView :src="app.icon" custom-class="text-5xl" />
          </DynamicBackground>
          <div class="text-white text-xs truncate text-ellipsis overflow-hidden">{{ app.title }}</div>
        </div>
      </div>
    </div>
    <div class="flex-grow"></div>
  </div>
</template>

<style scoped>

</style>
