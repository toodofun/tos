<template>
  <Transition :css="false">
    <!--有bug，先屏蔽-->
    <!--      <Transition :css="false" @before-leave="beforeLeave" @leave="leave" @enter="enter" @after-enter="afterEnter">-->
    <vue-drag-resize
      v-show="!isMinimized"
      class="rounded-lg flex flex-col overflow-hidden shadow-2xl shadow-cyan-500/30"
      :w="width"
      :h="height"
      :x="x"
      :y="y"
      :z="z"
      :minw="minWidth"
      :minh="minHeight"
      :isActive="true"
      :isResizable="true"
      :isDraggable="isDraggable"
      :parentLimitation="true"
      :stickSize="4"
      :sticks="['br', 'tr', 'bl', 'tl']"
      @resizing="onResizing"
    >
      <div class="w-full h-full rounded flex flex-col overflow-hidden" ref="targetDiv"
           :class="theme === 'light' ? 'bg-white text-black' : 'bg-slate-800 text-white'">
        <div class="flex justify-between items-center pl-2 select-none cursor-move border-b"
             :class="theme === 'light' ? 'border-slate-100' : 'border-slate-900'"
             @mouseenter="enableDraggable"
             @mouseleave="disableDraggable" @dblclick="toggleFullscreen">
          <div class="flex gap-1 items-center text-sm">
            <DynamicBackground :background="background" class="w-4 h-4 rounded aspect-square">
              <IconView :src="icon as string" custom-class="text-xs" />
            </DynamicBackground>
            <div>{{ title }}</div>
          </div>
          <div class="flex items-center gap-0" @mouseenter="disableDraggable">
            <a-button type="text" size="small" status="success" @click="toggleMinimize">
              <div class="text-sm">—</div>
            </a-button>
            <a-button type="text" size="small" status="warning" @click="toggleFullscreen">
              <div class="text-sm">□</div>
            </a-button>
            <a-button type="text" size="small" status="danger" @click="windowsStore.closeWindow(id)">
              <div class="text-sm">X</div>
            </a-button>
          </div>
        </div>
        <div class="w-full h-full relative bg-[#f8f8f8]" @mousedown="(e) => {e.stopPropagation()}">
          <component :is="windowsStore.getWindow(page)" />
          <!--          <iframe v-if="typeof page === 'string'" :src="page" class="w-full h-full will-change-auto"-->
          <!--                  allow="camera;microphone;clipboard-write;clipboard-read;"-->
          <!--                  sandbox="allow-same-origin allow-scripts allow-popups allow-forms"-->
          <!--          ></iframe>-->
          <!--          <component v-else :is="page" />-->
        </div>
      </div>
    </vue-drag-resize>
  </Transition>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { useWindowsStore } from '@/stores/windows'

import gsap from 'gsap'
import VueDragResize from '@/components/VueDragResize.vue'
import IconView from '@/components/operating-system/IconView.vue'
import DynamicBackground from '@/components/operating-system/components/DynamicBackground.vue'
import eventBus from '@/plugins/eventBus'
import { useDockStore } from '@/stores/dock'

const windowsStore = useWindowsStore()
const dockStore = useDockStore()

const props = defineProps({
  id: {
    type: String,
    required: true
  },
  title: String,
  icon: String,
  page: {
    type: String,
    required: true
  },
  theme: {
    type: String,
    required: true,
    validator: (val: string) => ['light', 'dark'].includes(val)
  },
  background: {
    type: String,
    required: true
  }
})

const width = defineModel('width', { type: Number, required: true })
const height = defineModel('height', { type: Number, required: true })
// const x = ref<string | number>(window.innerWidth / 2 - width.value / 2)
// const y = ref<string | number>(window.innerHeight / 2 - height.value / 2)
const x = defineModel('x', { type: Number, required: true })
const y = defineModel('y', { type: Number, required: true })
const z = ref(10)
const minWidth = ref<string | number>(280)
const minHeight = ref<string | number>(280)

const isDraggable = ref(false)
const isFullscreen = ref(false)
const targetDiv = ref<HTMLElement | null>(null)
const isMinimized = ref(false)
const preState = ref<{ w: number, h: number, x: number, y: number }>({
  w: 0,
  h: 0,
  x: 0,
  y: 0
})

const onResizing = () => {
  eventBus.emit('terminal-resize')
}

const toggleFullscreen = () => {
  isFullscreen.value = !isFullscreen.value
  if (isFullscreen.value && targetDiv.value) {
    dockStore.setShowDock(false)
    z.value = 51
    document.body.style.overflow = 'hidden' // 禁用滚动
    const td = targetDiv.value
    td.style.position = 'fixed'
    td.style.top = '32px'
    td.style.left = '0'
    td.style.width = '100vw'
    td.style.height = 'calc(100vh - 32px)'
    td.style.zIndex = '9999' // 确保全屏时在最上层
  } else if (targetDiv.value) {
    dockStore.setShowDock(true)
    z.value = 10
    document.body.style.overflow = '' // 恢复滚动
    const td = targetDiv.value
    td.style.position = ''
    td.style.top = ''
    td.style.left = ''
    td.style.width = ''
    td.style.height = ''
    td.style.zIndex = '10'
  }
  nextTick(() => {
    eventBus.emit('terminal-resize')
  })
}

const enableDraggable = () => {
  isDraggable.value = true
}

const disableDraggable = () => {
  isDraggable.value = false
}

const toggleMinimize = () => {
  isMinimized.value = !isMinimized.value
  if (isMinimized.value) {
    preState.value = {
      w: width.value,
      h: height.value,
      x: x.value,
      y: y.value
    }
    windowsStore.minimizeWindow(props.id, () => {
      toggleMinimize()
    })
  } else {
    nextTick(() => {
      width.value = preState.value.w as number
      height.value = preState.value.h as number
      x.value = preState.value.x as number
      y.value = preState.value.y as number
    })
    nextTick(() => {
      eventBus.emit('terminal-resize')
    })
  }
}

const beforeLeave = (el: Element) => {
  const rect = el.getBoundingClientRect()
  gsap.set(el, isFullscreen.value ? {
    position: 'fixed',
    top: `${32}px`,
    left: `${0}px`,
    width: `${window.innerWidth}px`,
    height: `${window.innerHeight - 132}px`,
    transformOrigin: 'center center',
    zIndex: 9999,
    boxShadow: '0 4px 6px rgba(0,0,0,0.1)'
  } : {
    position: 'fixed',
    top: `${rect.top}px`,
    left: `${rect.left}px`,
    width: `${rect.width}px`,
    height: `${rect.height}px`,
    transformOrigin: 'center center',
    zIndex: 9999,
    boxShadow: '0 4px 6px rgba(0,0,0,0.1)'
  })
}

const leave = (el: Element, done: () => void) => {
  const screenBottom = window.innerHeight
  const screenCenter = window.innerWidth / 2
  gsap.to(el, {
    duration: 0.5,
    x: screenCenter - (el as HTMLElement).offsetWidth / 2,
    y: screenBottom + 20, // 稍微超出屏幕底部
    scale: 0.12,
    opacity: 0,
    ease: 'back.in(1.7)',
    boxShadow: '0 0 0 rgba(0,0,0,0)',
    onComplete: done
  })
}

const enter = (el: Element, done: () => void) => {
  const screenBottom = window.innerHeight
  const screenCenter = window.innerWidth / 2
  gsap.set(el, isFullscreen.value ? {
    x: screenCenter - (preState.value.w as number) / 2,
    y: screenBottom + 20, // 从屏幕稍微下方开始
    width: window.innerWidth,
    height: window.innerHeight - 132,
    scale: 0.12,
    opacity: 0,
    position: 'fixed',
    zIndex: 9999,
    boxShadow: '0 0 0 rgba(0,0,0,0)'
  } : {
    x: screenCenter - (preState.value.w as number) / 2,
    y: screenBottom + 20, // 从屏幕稍微下方开始
    width: preState.value.w,
    height: preState.value.h,
    scale: 0.12,
    opacity: 0,
    position: 'fixed',
    zIndex: 9999,
    boxShadow: '0 0 0 rgba(0,0,0,0)'
  })
  gsap.to(el, isFullscreen.value ? {
    duration: 0.5,
    x: 0,
    y: 0,
    scale: 1,
    opacity: 1,
    ease: 'back.out(1.7)',
    boxShadow: '0 4px 6px rgba(0,0,0,0.1)',
    onComplete: done
  } : {
    duration: 0.5,
    x: preState.value.x,
    y: preState.value.y,
    scale: 1,
    opacity: 1,
    ease: 'back.out(1.7)',
    boxShadow: '0 4px 6px rgba(0,0,0,0.1)',
    onComplete: done
  })
}

const afterEnter = (el: Element) => {
  const style = (el as HTMLElement).style
  nextTick(() => {
    width.value = preState.value.w as number
    height.value = preState.value.h as number
    x.value = preState.value.x
    y.value = preState.value.y
    style.removeProperty('transform')
    style.removeProperty('opacity')
    style.removeProperty('position')
    style.removeProperty('z-index')
    style.removeProperty('box-shadow')
    style.position = 'absolute'
  })
}

</script>

<style scoped>
</style>
