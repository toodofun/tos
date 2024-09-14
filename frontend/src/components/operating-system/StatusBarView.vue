<script setup lang="ts">
import DateTime from './components/DateTime.vue'
import { useDesktopStore } from '@/stores/desktop'
import ControlCenter from '@/components/operating-system/components/ControlCenter.vue'
import TaskCenter from '@/components/operating-system/components/TaskCenter.vue'

const desktopStore = useDesktopStore()
</script>

<template>
  <div
    class="absolute z-[50] top-0 left-0 right-0 h-8 bg-white/20 p-1 px-4 backdrop-blur text-white flex items-center justify-between text-sm select-none">
    <div class="font-bold">{{ desktopStore.systemName }}</div>
    <div class="flex items-center flex-nowrap gap-2">
      <!--网速插件-->
      <!--      <div class="flex flex-col items-end tabular-nums text-[0.7rem] leading-3">-->
      <!--        <div>↑ 29.82KB/s</div>-->
      <!--        <div>↓ 20.12KB/s</div>-->
      <!--      </div>-->
      <!--任务中心-->
      <TaskCenter />
      <!--控制中心-->
      <!--<ControlCenter />-->
      <!--时间插件-->
      <DateTime />
    </div>
  </div>
</template>

<style scoped>

</style>
