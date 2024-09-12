<script setup lang="ts">
import { ref } from 'vue'
import { useWindowsStore } from '@/stores/windows'
import IconView from '@/components/operating-system/IconView.vue'
import DynamicBackground from '@/components/operating-system/components/DynamicBackground.vue'

const windowsStore = useWindowsStore()

const activeIndex = ref(-1)

const getScale = (index: number) => {
  const distance = Math.abs(index - activeIndex.value)
  if (activeIndex.value === -1) return 1
  if (distance === 0) return 1.5
  if (distance === 1) return 1.2
  return 1
}

</script>

<template>
  <div
    class="absolute z-[50] bottom-2 left-0 right-0 mx-auto w-fit bg-white/20 backdrop-blur-sm rounded-xl p-2 shadow-lg">
    <div class="flex items-end gap-4">
      <div
        v-for="(item, index) in windowsStore.fixedApps"
        :key="index"
        class="group relative flex flex-col items-center justify-center transition-all duration-300 ease-in-out"
        @mousemove="activeIndex = index"
        @mouseleave="activeIndex = -1"
        @click="() => {windowsStore.openWindow(item); activeIndex = -1}"
        :style="{zIndex: getScale(index) + 1}"
      >
        <DynamicBackground
          class="relative w-10 h-10 flex items-center justify-center rounded-lg bg-white/100 overflow-hidden backdrop-blur-none text-3xl cursor-pointer will-change-transform iconfont"
          :style="{
            transform: `scale(${getScale(index)})`,
            transformOrigin: 'bottom',
            transition: 'all 0.3s cubic-bezier(0.25, 0.1, 0.25, 1)',
          }"
          :background="item.background"
        >
          <IconView :src="item.icon" customClass="text-3xl" />
        </DynamicBackground>
        <span
          class="absolute -bottom-6 px-2 py-1 bg-gray-800 text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-nowrap">
          {{ item.title }}
        </span>
      </div>
      <div class="h-10 w-[0.2rem]" v-if="windowsStore.minimizedApps.length > 0"></div>
      <div
        v-for="(item, index) in windowsStore.minimizedApps"
        :key="index+windowsStore.fixedApps.length"
        class="group relative flex flex-col items-center justify-center transition-all duration-300 ease-in-out rounded-lg shadow shadow-cyan-400"
        @mousemove="activeIndex = index+windowsStore.fixedApps.length"
        @mouseleave="activeIndex = -1"
        @click="() => {windowsStore.restoreWindow(item.id); activeIndex = -1}"
      >
        <div
          class="relative w-10 h-10 flex items-center justify-center rounded-lg bg-white/100 overflow-hidden backdrop-blur-none text-3xl cursor-pointer will-change-transform iconfont"
          :style="{
            transform: `scale(${getScale(index+windowsStore.fixedApps.length)})`,
            transformOrigin: 'bottom',
            transition: 'all 0.3s cubic-bezier(0.25, 0.1, 0.25, 1)',
          }"
        >
          <IconView :src="item.icon" customClass="text-3xl" />
          <!--启动应用标识-->
          <div class="absolute bottom-0 left-0 right-0 flex justify-center" v-if="item.active">
            <div class="w-full h-[0.1rem] bg-sky-500 rounded-full"></div>
          </div>
        </div>
        <span
          class="absolute -bottom-6 px-2 py-1 bg-gray-800 text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-nowrap">
          {{ item.title }}
        </span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.group {
  transform: translateY(0);
  transition: transform 0.3s cubic-bezier(0.25, 0.1, 0.25, 1);
}

.group:hover {
  transform: translateY(-10px);
}

/* 优化性能 */
.will-change-transform {
  will-change: transform;
}

/* 创建平滑的Dock效果 */
.flex:hover .group {
  transform: translateY(0);
}

.flex:hover .group:hover {
  transform: translateY(-10px);
}

.flex:hover .group:hover ~ .group {
  transform: translateY(-5px);
}
</style>
