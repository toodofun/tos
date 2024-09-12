import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useDesktopStore = defineStore('desktop', () => {
    // StatusBar
    const statusBarShowSeconds = ref(false)

    return { statusBarShowSeconds }
  },
  {
    persist: true
  }
)
