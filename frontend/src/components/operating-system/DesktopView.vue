<script setup lang="ts">
import { ref, h, type VNode } from 'vue'
import DockView from '@/components/operating-system/DockView.vue'
import StatusBarView from '@/components/operating-system/StatusBarView.vue'
import DesktopMain from '@/components/operating-system/components/DesktopMain.vue'
import { useDesktopStore } from '@/stores/desktop'

import ContextMenu from '@imengyu/vue3-context-menu'
import { IconRefresh } from '@arco-design/web-vue/es/icon'

const onContextMenu = (e : MouseEvent) => {
  //prevent the browser's default menu
  e.preventDefault();
  //show your menu
  ContextMenu.showContextMenu({
    theme: 'mac',
    x: e.x,
    y: e.y,
    items: [
      {
        label: "刷新",
        icon: ():VNode => {
          return h(IconRefresh)
        },
        onClick: () => {
          location.reload()
        }
      },
      {
        label: "A submenu",
        children: [
          { label: "Item1" },
          { label: "Item2" },
          { label: "Item3" },
        ]
      },
    ]
  });
}

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
    <DesktopMain @contextmenu="onContextMenu" />
    <!--窗口区-->
    <div class="absolute inset-0 top-8 pointer-events-auto">
      <slot></slot>
    </div>

  </div>
</template>

<style scoped>
.popup-video {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover; /* 保持宽高比，裁剪内容以适应父元素 */
  z-index: -1; /* 如果需要将其放在背景层 */
}
</style>
