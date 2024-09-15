<script setup lang="ts">
import IconView from '@/components/operating-system/IconView.vue'
import { useQueueStore } from '@/stores/queue'
import TaskItem from '@/components/operating-system/components/TaskItem.vue'
import { IconDelete, IconRefresh } from '@arco-design/web-vue/es/icon'
import { computed } from 'vue'

const queueStore = useQueueStore()

const sortedTasks = computed(() => {
  return queueStore.tasks.slice().sort((a, b) => {
    const statusOrder = ['running', 'pending', 'error', 'success']

    // 根据 statusOrder 数组中的顺序进行排序
    return statusOrder.indexOf(a.status) - statusOrder.indexOf(b.status)
  })
})


</script>

<template>
  <a-popover
    class="overflow-hidden"
    trigger="click"
    position="tl"
    :content-style="{
        background: 'rgba(255,255,255,0)',
        padding: '0.5rem',
        border: 'none',
        marginRight: '0.6rem',
        boxShadow: 'none'
    }"
    :arrow-style="{display: 'none'}">
    <div
      class="w-6 h-6 max-w-6 max-h-6 min-w-6 min-h-6 cursor-pointer select-none"
      :class="sortedTasks.length > 0 ? 'animate-pulse':''"
    >
      <IconView src="internal://icon-task-center" />
    </div>
    <template #content>
      <div
        class="select-none bg-white/100 backdrop-blur-sm px-2 pb-2 w-72 max-w-72 min-w-72 max-h-[30rem] overflow-y-auto rounded-lg flex flex-col gap-2">
        <div class="flex justify-between items-center sticky bg-white z-20 top-0 py-2">
          <div class="font-bold">任务队列</div>
          <div class="flex items-center gap-1">
            <a-popover class="overflow-hidden" trigger="hover" content="重试失败的任务">
              <a-button size="mini" rounded @click="queueStore.retryFailedTasks">
                <template #icon>
                  <icon-refresh />
                </template>
              </a-button>
            </a-popover>
            <a-popover class="overflow-hidden" trigger="hover" content="清除成功的任务">
              <a-button size="mini" rounded @click="queueStore.clearSuccessTasks">
                <template #icon>
                  <icon-delete />
                </template>
              </a-button>
            </a-popover>
          </div>
        </div>

        <div v-if="queueStore.tasks.length > 0">
          <div class="flex flex-col gap-2">
            <div v-for="task in sortedTasks" :key="task.id">
              <TaskItem :task="task" />
            </div>
          </div>
        </div>
        <div v-else>
          <a-empty description="暂无任务" />
        </div>
      </div>
    </template>
  </a-popover>
</template>

<style scoped>

</style>
