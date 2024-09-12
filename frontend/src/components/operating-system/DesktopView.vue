<script setup lang="ts">
import { ref } from 'vue'
import desktopBackgroundImg from '@/assets/background.jpeg'
import DockView from '@/components/operating-system/DockView.vue'
import StatusBarView from '@/components/operating-system/StatusBarView.vue'
import DesktopMain from '@/components/operating-system/components/DesktopMain.vue'

const preview = ref(false)

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
    <DesktopMain />
    <!--窗口区-->
    <div class="absolute inset-0 top-8 pointer-events-auto">
      <slot></slot>
    </div>

  </div>
</template>

<style scoped>

</style>
