<script setup lang="ts">
import { useDesktopStore } from '@/stores/desktop'

const desktopStore = useDesktopStore()
</script>

<template>
  <div class="w-full h-full p-2 flex flex-col gap-2">
    <div class="bg-white rounded-md p-2 flex flex-col gap-2">
      <div class="flex flex-col gap-2">
        <div class="font-bold">桌面壁纸</div>
        <div class="flex flex-nowrap overflow-x-auto items-center gap-4 p-2">
          <div
            class="aspect-video w-72 min-w-72 rounded-md overflow-hidden cursor-pointer"
            v-for="wallpaper in desktopStore.wallpapers"
            :key="wallpaper.src"
            @click="desktopStore.setWallpaper(wallpaper)"
          >
            <a-image
              v-if="wallpaper.type === 'img'"
              width="100%"
              height="100%"
              :src="wallpaper.src"
              alt=""
              fit="cover"
              :preview="false"
            />
            <video
              v-if="wallpaper.type === 'video'"
              :src="wallpaper.src"
              width="100%"
              height="100%"
              :muted="true"
              :loop="true"
              :autoplay="true"
              x5-video-player-type="h5"
              class="popup-video"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
