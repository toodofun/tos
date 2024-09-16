<script setup lang="ts">
import type { PropType } from 'vue'
import DesktopIcon from '@/components/operating-system/components/DesktopIcon.vue'
import { type CreateNewWindowInfo, useWindowsStore } from '@/stores/windows'

const windowStore = useWindowsStore()

defineProps({
  onOpenWindow: {
    type: Function as PropType<((app: CreateNewWindowInfo) => void)>,
    required: true
  },
  onClose: {
    type: Function as PropType<() => void>,
    required: true
  }
})
</script>

<template>
  <div class="absolute inset-0 px-[15%] py-[10%] bg-slate-800/60 backdrop-blur-lg" @click="onClose">
    <div class="w-full h-full grid gap-4" style="grid-template-columns: repeat(auto-fill, minmax(136px, 1fr));">
      <DesktopIcon :size="136" :app="app" @click="onOpenWindow(app)" v-for="app in windowStore.allApps" :key="app.id" />
    </div>
  </div>
</template>

<style scoped>

</style>
