<script setup lang="ts">
import { ref } from 'vue'
import DockView from '@/components/operating-system/DockView.vue'
import StatusBarView from '@/components/operating-system/StatusBarView.vue'
import DesktopMain from '@/components/operating-system/components/DesktopMain.vue'
import { useDesktopStore } from '@/stores/desktop'

const preview = ref(false)
const desktopStore = useDesktopStore()

</script>

<template>
  <div class="fixed inset-0 z-0">
    <!--背景图片-->
    <div class="absolute inset-0">
      <a-image
        v-if="desktopStore.wallpaper.type === 'img'"
        width="100%"
        height="100%"
        :src="desktopStore.wallpaper.src"
        alt=""
        fit="cover"
        :preview="preview"
      />
      <video
        v-if="desktopStore.wallpaper.type === 'video'"
        :src="desktopStore.wallpaper.src"
        width="100%"
        height="100%"
        :muted="true"
        :loop="true"
        :autoplay="true"
        x5-video-player-type="h5"
        class="popup-video"
      />
    </div>
    <!--状态栏-->
    <StatusBarView />
    <!--Dock栏-->
    <DockView />
    <!--桌面内容-->
    <DesktopMain />
    <!--窗口区-->
    <div class="absolute inset-0 top-8 pointer-events-auto">
      <slot></slot>
    </div>

  </div>
</template>

<style scoped>

</style>
