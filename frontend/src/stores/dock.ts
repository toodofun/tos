import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useDockStore = defineStore('dock', () => {
  const showLaunchpad = ref(false)
  const showDock = ref(true)

  const setShowLaunchpad = (value: boolean = false) => {
    showLaunchpad.value = value
  }

  const setShowDock = (value: boolean = true) => {
    showDock.value = value
  }

  return { showLaunchpad, setShowLaunchpad, showDock, setShowDock }
},{
  persist: {
    pick: []
  }
})