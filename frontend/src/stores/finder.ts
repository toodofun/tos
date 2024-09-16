import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useFinderStore = defineStore('finder', () => {
  const showInList = ref(false)

  const setShowInList = (value: boolean) => {
    showInList.value = value
  }
  return {
    showInList,
    setShowInList
  }
}, {
  persist: true
})
