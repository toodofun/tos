<script setup lang="ts">

import IconView from '@/components/operating-system/IconView.vue'
</script>

<template>
  <a-popover trigger="click"
             position="tl"
             :content-style="{background: 'rgba(255,255,255,0)', padding: '0.5rem', border: 'none', marginRight: '0.6rem', boxShadow: 'none'}"
             :arrow-style="{display: 'none'}">
    <div class="w-6 h-6 max-w-6 max-h-6 min-w-6 min-h-6 cursor-pointer select-none">
      <IconView src="internal://icon-control" />
    </div>
    <template #content>
      <div class="select-none bg-white/40 backdrop-blur-sm p-2 min-w-64 rounded-lg">
        <div class="bg-white p-2 rounded-lg">on development...</div>
      </div>
    </template>
  </a-popover>
</template>

<style scoped>

</style>
