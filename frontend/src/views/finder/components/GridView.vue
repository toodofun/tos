<script setup lang="ts">
import { type FileInfo, getFileIcon } from '@/views/finder/FinderView'
import type { PropType } from 'vue'
import IconView from '@/components/operating-system/IconView.vue'
import dayjs from 'dayjs'

defineProps({
  items: {
    type: Array as PropType<FileInfo[]>,
    required: true
  },
  onDoubleClick: {
    type: Function as PropType<(file: FileInfo) => void>,
    required: true
  },
  onClick: {
    type: Function as PropType<(file: FileInfo) => void>,
    required: true
  },
  selectedInfo: {
    type: Array as PropType<FileInfo[]>,
    required: false,
    default: () => []
  }
})
</script>

<template>
  <div class="w-full h-full grid gap-4 p-4" style="grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));">
    <div v-for="(item, index) in items" :key="index"
         class="w-full h-full cursor-pointer rounded-lg p-2 flex flex-col items-center justify-between gap-2"
         :class="selectedInfo.includes(item) ? 'bg-slate-100' : 'hover:bg-slate-100'"
         @dblclick="onDoubleClick(item)"
         @click="onDoubleClick(item)"
         @mouseenter="onClick(item)"
    >
      <div class="aspect-square">
        <IconView :src="getFileIcon(item)" custom-class="text-5xl" />
      </div>
      <div class="flex flex-col items-center gap-1">
        <div class="text-xs text-wrap text-center max-w-20 truncate text-ellipsis break-words line-clamp-2">
          {{ item.name }}
        </div>
        <div class="text-xs text-gray-500 max-w-20 truncate text-ellipsis overflow-hidden">
          {{ dayjs(item.modTime).format('MM/DD HH:mm') }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
