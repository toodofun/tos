<script setup lang="ts">
import { h, type PropType, reactive } from 'vue'
import { type FileInfo, getFileIcon } from '@/views/finder/FinderView'
import IconView from '@/components/operating-system/IconView.vue'
import numeral from 'numeral'
import dayjs from 'dayjs'
import type { TableColumnData, TableData } from '@arco-design/web-vue'

const props = defineProps({
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
  },
  onContextMenu: {
    type: Function as PropType<(e: MouseEvent, file: FileInfo) => void>,
    required: true
  }
})

const columns = reactive<TableColumnData[]>([
  {
    title: '名称',
    dataIndex: 'name',
    render: (data: { record: TableData; column: TableColumnData; rowIndex: number; }) => {
      return h('div', {
        class: 'flex items-center gap-1 cursor-pointer',
        onContextmenu: (e: MouseEvent) => {
          props.onContextMenu(e, data.record as FileInfo)
        }
      }, [
        h('div', { class: 'w-6 h-6' }, [
          h(IconView, { src: getFileIcon(data.record as FileInfo) })
        ]),
        h('div', data.record.name)
      ])
    }
  },
  {
    title: '大小',
    dataIndex: 'size',
    width: 106,
    minWidth: 88,
    render: (data: { record: TableData; column: TableColumnData; rowIndex: number; }) => {
      return h('div', {
        class: 'cursor-pointer',
        onContextmenu: (e: MouseEvent) => {
          props.onContextMenu(e, data.record as FileInfo)
        }
      }, numeral(data.record.size).format('0.[00] ib'))
    }
  },
  {
    title: '修改日期',
    dataIndex: 'modTime',
    width: 180,
    minWidth: 80,
    render: (data: { record: TableData; column: TableColumnData; rowIndex: number; }) => {
      return h('div', {
        class: 'cursor-pointer',
        onContextmenu: (e: MouseEvent) => {
          props.onContextMenu(e, data.record as FileInfo)
        }
      }, dayjs(data.record.modTime).format('YYYY-MM-DD HH:mm:ss'))
    }
  }
])
</script>

<template>
  <div class="w-full h-full p-4">
    <a-table :columns="columns" :data="items" column-resizable
             :sticky-header="true"
             :bordered="false"
             :pagination="false"
             @row-click="onDoubleClick as any"
             @row-dblclick="onDoubleClick as any"
             @cellMouseEnter="onClick as any"
    ></a-table>
  </div>
</template>

<style scoped>

</style>
