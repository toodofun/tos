<script setup lang="ts">
import TerminalView from '@/components/TerminalView.vue'
import { useWindowsStore } from '@/stores/windows'

const windowsStore = useWindowsStore()

defineProps({
  id: {
    type: String,
    required: true
  },
})
</script>

<template>
  <div class="absolute inset-0">
    <TerminalView class="w-full h-full" :uri="'/api/v1/terminal'" :exit="() => windowsStore.closeWindow(id)" />
  </div>
</template>

<style scoped>

</style>
