import { ref } from 'vue'
import { defineStore } from 'pinia'
import bgDefault1 from '@/assets/bg-default-1-mini.jpg'
import bgDefault2 from '@/assets/bg-default-2-mini.jpg'
import bgDefault3 from '@/assets/bg-default-3-mini.jpg'

interface Wallpaper {
  src: string
  type: 'img' | 'video'
}

export const useDesktopStore = defineStore('desktop', () => {
    // StatusBar
    const statusBarShowSeconds = ref(false)
    const dateFormat = ref('M月D日 dddd')
    const systemName = ref('Toodo Cloud OS')
    const wallpapers: Wallpaper[] = [
      {
        src: bgDefault1,
        type: 'img'
      },
      {
        src: bgDefault2,
        type: 'img'
      },
      {
        src: bgDefault3,
        type: 'img'
      },
      {
        src: 'https://img-baofun.zhhainiao.com/market/5/381773.mp4',
        type: 'video'
      },
      {
        src: 'https://img-baofun.zhhainiao.com/market/5/409819.mp4',
        type: 'video'
      }
    ]

    // 壁纸
    const wallpaper = ref<Wallpaper>({
      src: bgDefault1,
      type: 'img'
    })

    function setStatusBarShowSeconds(value: boolean = false) {
      statusBarShowSeconds.value = value
    }

    function setDateFormat(value: string = 'M月D日 dddd') {
      dateFormat.value = value
    }

    function setWallpaper(value: Wallpaper) {
      wallpaper.value = value
    }

    return {
      systemName,
      wallpaper,
      wallpapers,
      statusBarShowSeconds,
      dateFormat,
      setStatusBarShowSeconds,
      setDateFormat,
      setWallpaper
    }
  },
  {
    persist: true
  }
)
