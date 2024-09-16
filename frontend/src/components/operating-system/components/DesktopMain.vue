<script setup lang="ts">
import { ref, watch, onMounted, onBeforeUnmount, h, type VNode } from 'vue'
import { useWindowsStore } from '@/stores/windows'
import DesktopIcon from '@/components/operating-system/components/DesktopIcon.vue'
import { GridLayout, GridItem, type Layout, type LayoutItem } from 'vue3-draggable-grid'
import 'vue3-draggable-grid/dist/style.css'
import CalendarWidget from '@/components/operating-system/components/CalendarWidget.vue'
import { storeToRefs } from 'pinia'

const windowsStore = useWindowsStore()

const { desktopApps } = storeToRefs(windowsStore)

const gridContainer = ref()

const padding = ref(16)
const col = ref(12)
const row = ref(0)
const gutter = ref(20)
const rowH = ref(100)

const layout = ref<Layout>([])

const getItem = (id: string): VNode => {
  if (id.startsWith('widget://')) {
    return h(CalendarWidget)
  }
  return h(DesktopIcon, {
    app: windowsStore.desktopApps[Number(id)],
    onClick: () => windowsStore.openWindow(windowsStore.desktopApps[Number(id)])
  })
}

const calculateGridSize = () => {
  col.value = Math.floor((gridContainer.value.clientWidth - padding.value * 2 + gutter.value) / (rowH.value + gutter.value))
  row.value = Math.floor((gridContainer.value.clientHeight - padding.value * 2 + gutter.value) / (rowH.value + gutter.value))
  layout.value = []

  layout.value.push({
    x: 1,
    y: 1,
    h: 2,
    w: 3,
    id: 'widget://test'
  })

  windowsStore.desktopApps.forEach((item, index) => {
    // 计算当前元素的列号（从右往左）
    const currentCol = col.value - Math.floor(index / row.value)
    // 计算当前元素的行号（从上到下）
    const currentRow = (index % row.value) + 1

    layout.value.push({
      id: index.toString(),  // 使用 item 作为 id
      x: currentCol,        // 列号，从右往左
      y: currentRow,        // 行号，从上到下
      h: 1,  // 每个元素的高度
      w: 1   // 每个元素的宽度
    })
  })
}

onMounted(() => {
  calculateGridSize()
  window.addEventListener('resize', calculateGridSize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', calculateGridSize)
})

// 验证更新数据是否正确
watch(desktopApps, () => {
  calculateGridSize()
}, { deep: true })

const draggableStart = (id: string) => {
  // console.log('拖拽开始', id)
}

const draggableHandle = (id: string, data: LayoutItem) => {
  // console.log('拖拽中', id, data)
}

const draggableEnd = (data: Layout) => {
  // console.log('拖拽结束', data)
}

const remove = (id: string) => {
  // console.log('删除', id)
}
</script>

<template>
  <div ref="gridContainer" class="absolute top-8 bottom-0 left-0 right-0 select-none p-4">
    <grid-layout
      v-model:data="layout"
      :drage="false"
      :draggable="false"
      :remove="false"
      :resize="false"
      :is-draw-grid-lines="false"
      :col="col"
      :gutter="gutter"
      :row-h="rowH"
    >
      <grid-item v-for="item in layout" :key="item.id" :id="item.id">
        <component :is="getItem(item.id)" />
      </grid-item>
    </grid-layout>
  </div>
</template>

<style scoped>
:deep(.grid-item) {
  background-color: transparent !important;
  border: none;
}
</style>
