<script setup lang="ts">
import { ref } from 'vue'
import desktopBackgroundImg from '@/assets/background.jpeg'
import DockView from '@/components/operating-system/DockView.vue'
import StatusBarView from '@/components/operating-system/StatusBarView.vue'
import { useWindowsStore } from '@/stores/windows'
import IconView from '@/components/operating-system/IconView.vue'
import DynamicBackground from '@/components/operating-system/components/DynamicBackground.vue'

const preview = ref(false)
const windowsStore = useWindowsStore()

</script>

<template>
  <div class="fixed inset-0 z-0">
    <!--背景图片-->
    <div class="absolute inset-0">
      <a-image width="100%" height="100%" :src="desktopBackgroundImg" alt="" fit="cover" :preview="preview" />
    </div>
    <!--状态栏-->
    <StatusBarView />
    <!--Dock栏-->
    <DockView />
    <!--桌面内容-->
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
    <!--窗口区-->
    <div class="absolute inset-0 pointer-events-auto">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>

</style>
