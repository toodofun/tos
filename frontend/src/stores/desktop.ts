import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useDesktopStore = defineStore('desktop', () => {
    // StatusBar
    const statusBarShowSeconds = ref(false)
    const dateFormat = ref('M月D日 dddd')
    const systemName = ref('Toodo Cloud OS')

    function setStatusBarShowSeconds(value: boolean = false) {
      statusBarShowSeconds.value = value
    }

    function setDateFormat(value: string = 'M月D日 dddd') {
      dateFormat.value = value
    }

    return {
      systemName,
      statusBarShowSeconds,
      dateFormat,
      setStatusBarShowSeconds,
      setDateFormat
    }
  },
  {
    persist: true
  }
)
