<script setup lang="ts">
import TerminalView from '@/components/TerminalView.vue'

</script>

<template>
  <div class="absolute inset-0">
    <TerminalView class="w-full h-full" :uri="'/api/v1/terminal'" />
  </div>
</template>

<style scoped>

</style>
