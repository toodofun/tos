import { ref } from 'vue'
import { defineStore } from 'pinia'
import PQueue from 'p-queue'
import { v4 as uuidv4 } from 'uuid'

export interface TaskItem {
  id: string
  title: string
  percent: number
  status: 'pending' | 'running' | 'paused' | 'success' | 'error'
  message: string
  fn?: (onProgress: (percent: number) => void, abortSignal?: AbortSignal) => Promise<any>
}

export const useQueueStore = defineStore('queue', () => {
  const queue = new PQueue({ concurrency: 5 })
  const tasks = ref<TaskItem[]>([])
  const abortControllers = new Map<string, AbortController>()

  // 添加任务到任务列表
  const addTask = (id: string, title: string, fn: (onProgress: (percent: number) => void, abortSignal?: AbortSignal) => Promise<any>) => {
    const task: TaskItem = {
      id,
      title,
      percent: 0,
      status: 'pending',
      message: '',
      fn
    }
    tasks.value.push(task)
    return task
  }

  // 更新任务状态和进度
  const updateTask = (id: string, updates: Partial<TaskItem>) => {
    const index = tasks.value.findIndex(task => task.id === id)
    if (index !== -1) {
      tasks.value[index] = { ...tasks.value[index], ...updates }
    }
  }

  // 进度更新回调
  const onProgress = (id: string, percent: number) => {
    updateTask(id, { percent })
  }

  // 删除任务，包括取消正在执行的任务
  const removeTask = (id: string) => {
    const taskIndex = tasks.value.findIndex(task => task.id === id)
    if (taskIndex !== -1) {
      tasks.value.splice(taskIndex, 1) // 从任务列表中移除
      abortControllers.get(id)?.abort() // 取消任务
      abortControllers.delete(id)
    }
  }

  // 暂停单个任务
  const pauseTask = (id: string) => {
    const task = tasks.value.find(task => task.id === id)
    if (task && task.status === 'running') {
      updateTask(id, { status: 'paused', message: `${task.title} paused.` })
      abortControllers.get(id)?.abort() // 取消任务
    }
  }

  // 恢复单个任务
  const resumeTask = (id: string) => {
    const task = tasks.value.find(task => task.id === id)
    if (task && task.status === 'paused') {
      updateTask(id, { status: 'running', message: `${task.title} resumed.` })

      const abortController = new AbortController()
      abortControllers.set(id, abortController)

      queue.add(async () => {
        try {
          await task.fn!((percent: number) => onProgress(task.id, percent), abortController.signal)
          updateTask(task.id, { status: 'success', message: `${task.title} finished successfully!`, percent: 100 })
        } catch (error: any) {
          if (error.name === 'AbortError') {
            updateTask(task.id, { status: 'paused', message: `${task.title} paused.` })
          } else {
            updateTask(task.id, { status: 'error', message: `${task.title} failed: ${error.message}` })
          }
        }
      })
    }
  }

  // 提交任务到队列
  const submit = (title: string, fn: (onProgress: (percent: number) => void, abortSignal?: AbortSignal) => Promise<any>) => {
    const id = uuidv4()
    const task = addTask(id, title, fn)

    const abortController = new AbortController()
    abortControllers.set(id, abortController)

    queue.add(async () => {
      // 任务开始时更新状态为 running
      updateTask(task.id, { status: 'running', message: `${title} is running...` })

      try {
        // 传递 onProgress 回调来更新进度
        await fn((percent: number) => onProgress(task.id, percent), abortController.signal)

        // 任务完成，更新状态为 success，进度为 100%
        updateTask(task.id, { status: 'success', message: `${title} finished successfully!`, percent: 100 })
      } catch (error: any) {
        if (error.name === 'AbortError') {
          updateTask(task.id, { status: 'paused', message: `${title} paused.` })
        } else {
          // 任务失败，更新状态为 error
          updateTask(task.id, { status: 'error', message: `${title} failed: ${error.message}` })
        }
      }
    })
  }

  // 重试任务，只传入任务 ID
  const retryTask = (id: string) => {
    console.log(`Retry task: ${id}`)
    const task = tasks.value.find(task => task.id === id)
    if (task && task.status === 'error' && task.fn) {
      updateTask(id, { status: 'pending', percent: 0, message: `${task.title} retrying...` })

      const abortController = new AbortController()
      abortControllers.set(id, abortController)

      queue.add(async () => {
        try {
          await task.fn!((percent: number) => onProgress(task.id, percent), abortController.signal)
          updateTask(task.id, { status: 'success', message: `${task.title} finished successfully!`, percent: 100 })
        } catch (error: any) {
          if (error.name === 'AbortError') {
            updateTask(task.id, { status: 'paused', message: `${task.title} paused.` })
          } else {
            updateTask(task.id, { status: 'error', message: `${task.title} failed again: ${error.message}` })
          }
        }
      })
    }
  }

  const clearSuccessTasks = () => {
    tasks.value = tasks.value.filter(task => task.status !== 'success')
  }

  const retryFailedTasks = () => {
    tasks.value.forEach(task => {
      if (task.status === 'error') {
        retryTask(task.id)
      }
    })
  }

  const getTasks = () => tasks

  return {
    tasks,
    submit,
    removeTask,
    pauseTask,
    resumeTask,
    retryTask,
    getTasks,
    clearSuccessTasks,
    retryFailedTasks
  }
}, {
  persist: {
    pick: [
    ] // 持久化任务列表
  }
})
