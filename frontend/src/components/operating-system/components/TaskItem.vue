<script setup lang="ts">
import { type TaskItem, useQueueStore } from '@/stores/queue'
import { type PropType, h, type VNode } from 'vue'
import {
  IconLoading,
  IconClockCircle,
  IconRefresh,
  IconCheck
} from '@arco-design/web-vue/es/icon'

const queueStore = useQueueStore()

const ps = defineProps({
  task: {
    type: Object as PropType<TaskItem>,
    required: true
  }
})


const success = (props?: Object): VNode => h('div', { class: 'cursor-pointer' }, [h(IconCheck, { style: { color: 'rgb(34 197 94)' }, ...props })])

const running = (props?: Object): VNode => h('div', {
  class: 'cursor-pointer',
  onClick: () => {
    queueStore.pauseTask(ps.task.id)
  }
}, [h(IconLoading, { style: { color: 'rgb(14 165 233)' }, ...props })])

const pending = (props?: Object): VNode => h('div', { class: 'cursor-pointer' }, [h(IconClockCircle, { style: { color: 'rgb(234 179 8)' }, ...props })])

const error = (props?: Object): VNode => h('div', {
  class: 'cursor-pointer',
  onClick: () => {
    queueStore.retryTask(ps.task.id)
  }
}, [h(IconRefresh, { style: { color: 'rgb(239 68 68)' }, ...props })])

const paused = (props?: Object): VNode => h('div', {
  class: 'cursor-pointer',
  onClick: () => {
    queueStore.resumeTask(ps.task.id)
  }
}, [h(IconRefresh, { style: { color: 'rgb(100 116 139)' }, ...props })])

const getOperations = (status: string): VNode[] => {
  switch (status) {
    case 'success':
      return [success()]
    case 'running':
      return [running()]
    case 'pending':
      return [pending()]
    case 'error':
      return [error()]
    case 'paused':
      return [paused()]
  }
  return []
}

const getTitle = (title: string) => {
  if (title.length > 25) {
    return title.slice(0, 15) + '...' + title.slice(-10)
  }
  return title
}

</script>

<template>
  <a-popover class="overflow-hidden" trigger="hover" :content="task.title">
    <div class="flex flex-col bg-slate-100 px-2 rounded py-2">
      <div class="text-xs flex justify-between items-center">
        <div class="max-w-52 truncate text-ellipsis overflow-hidden" :title="task.title">{{ getTitle(task.title) }}
        </div>
      </div>
      <div class="w-full flex gap-2 items-center">
        <a-progress
          :show-text="false"
          trackColor="rgb(226 232 240)"
          status='success'
          :percent="task.percent / 100"
        />
        <div class="flex gap-1">
          <component v-for="operation in getOperations(task.status)" :key="operation" :is="operation" />
        </div>
      </div>
    </div>
  </a-popover>
</template>

<style scoped>

</style>
