import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useDockStore = defineStore('dock', () => {
  const showLaunchpad = ref(false)
  const showDock = ref(true)
  const dockSize = ref(20)

  const setShowLaunchpad = (value: boolean = false) => {
    showLaunchpad.value = value
  }

  const setShowDock = (value: boolean = true) => {
    showDock.value = value
  }

  const setDockSize = (value: number = 20) => {
    dockSize.value = value
  }

  return {
    showLaunchpad,
    setShowLaunchpad,
    showDock,
    setShowDock,
    dockSize,
    setDockSize
  }
}, {
  persist: {
    pick: [
      'dockSize'
    ]
  }
})
