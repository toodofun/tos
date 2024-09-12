import { ref, h, defineAsyncComponent, type VNode } from 'vue'
import { defineStore } from 'pinia'
import { v4 as uuidv4 } from 'uuid'

export const defaultBackground = 'white'

export interface WindowInfo {
  id: string
  title: string
  icon: string
  page: string
  width: number
  height: number
  x: number
  y: number
  z?: number
  active?: boolean
  theme: 'light' | 'dark'
  background: string
}

interface CreateNewWindowInfo {
  title: string
  icon: string
  page: string
  width?: number
  height?: number
  x?: number
  y?: number
  theme?: 'light' | 'dark'
  background?: string
}

const internalApps: { [key: string]: VNode } = {
  'internal://setting': h(defineAsyncComponent(() => import('@/views/setting/SettingView.vue'))),
  'internal://terminal': h(defineAsyncComponent(() => import('@/views/terminal/TerminalView.vue'))),
  'internal://finder': h(defineAsyncComponent(() => import('@/views/finder/FinderView.vue')))
}

export const getWindow = (src: string): VNode => {
  if (src.startsWith('internal')) {
    return internalApps[src] || h('div', 'Not Found')
  }
  if (src.startsWith('http')) {
    return h('iframe', {
      src,
      class: 'w-full h-full will-change-auto',
      allow: 'camera;microphone;clipboard-write;clipboard-read;',
      sandbox: 'allow-same-origin allow-scripts allow-popups allow-forms'
    })
  }
  return h('div', 'Not Found')
}

export const useWindowsStore = defineStore('windows', () => {
    // Windows Manager
    const windowList = ref<Array<WindowInfo>>([])
    const maxZ = ref(1)

    const restoreFunctionMap: { [key: string]: () => void } = {}

    // Dock
    const fixedApps = ref<Array<CreateNewWindowInfo>>([
      { icon: 'internal://icon-oss', title: '文件', page: 'internal://finder' },
      {
        icon: 'internal://icon-app',
        title: '启动台',
        page: '//docker.ac.cn',
        background: 'linear-gradient(to right, #4e54c8, #8f94fb)'
      },
      {
        icon: 'internal://icon-app-store',
        title: '应用商店',
        page: 'internal://finder',
        background: 'linear-gradient(to right, #06beb6, #48b1bf)'
      },
      {
        icon: 'internal://icon-setting',
        title: '设置',
        page: 'internal://setting',
        background: 'linear-gradient(to right, #536976, #292e49)'
      },
      { icon: 'internal://icon-task', title: '任务管理', page: '//toodo.fun' },
      { icon: 'internal://icon-terminal', title: '终端', page: 'internal://terminal', theme: 'dark' }
    ])
    const minimizedApps = ref<Array<WindowInfo>>([])

    const desktopApps = ref<Array<CreateNewWindowInfo>>([
      {
        icon: 'https://i04piccdn.sogoucdn.com/a72804451f0e9825',
        title: '原神',
        width: 1280,
        height: 720,
        page: 'https://genshin.titlecan.cn/'
      },
      // {
      //   icon: 'https://hexgl.bkcore.com/play/css/title.png',
      //   title: 'HexGL赛车',
      //   width: 1280,
      //   height: 720,
      //   page: 'https://hexgl.bkcore.com/play/'
      // },
      // {
      //   icon: 'https://th.bing.com/th/id/OIP.nfFu7l8TPI6fnX5Fb8bJ_QHaHa?rs=1&pid=ImgDetMain',
      //   title: '亲戚称呼计算器',
      //   width: 338,
      //   height: 600,
      //   page: 'https://passer-by.com/relationship/vue/#/',
      //   theme: 'dark'
      // },
      {
        icon: 'internal://icon-app-store',
        title: '应用商店',
        page: 'internal://finder',
        background: 'linear-gradient(to right, #06beb6, #48b1bf)'
      },
      {
        icon: 'internal://icon-terminal',
        title: '终端',
        page: 'internal://terminal',
        theme: 'dark'
      },
      {
        icon: 'internal://icon-setting',
        title: '设置',
        page: 'internal://setting',
        background: 'linear-gradient(to right, #536976, #292e49)'
      }
      // { icon: 'internal://icon-container', title: '容器管理', page: '//toodo.fun' },
      // { icon: 'internal://icon-cluster', title: '集群管理', page: '//toodo.fun', theme: 'dark' },
    ])

    function clickWindow(w: WindowInfo) {
      maxZ.value += 1
      w.z = maxZ.value
    }

    function openWindow(w: CreateNewWindowInfo) {
      maxZ.value += 1
      windowList.value.push({
        id: uuidv4(),
        title: w.title,
        icon: w.icon,
        page: w.page,
        width: w.width || 800,
        height: w.height || 480,
        x: w.x || window.innerWidth / 2 - (w.width || 800) / 2,
        y: w.y || window.innerHeight / 2 - (w.height || 480) / 2,
        z: maxZ.value,
        active: true,
        theme: w.theme || 'light',
        background: w.background || defaultBackground
      })
    }

    function closeWindow(id: string) {
      windowList.value = windowList.value.filter((item) => {
        return item.id !== id
      })
    }

    function minimizeWindow(id: string, restoreFunction: () => void) {
      const w = windowList.value.find((item) => {
        return item.id === id
      })
      if (w) {
        minimizedApps.value.push(w)
        restoreFunctionMap[id] = restoreFunction
      }
    }

    function restoreWindow(id: string) {
      restoreFunctionMap[id]()
      minimizedApps.value = minimizedApps.value.filter((item) => {
        return item.id !== id
      })
    }

    return {
      windowList,
      maxZ,
      fixedApps,
      minimizedApps,
      clickWindow,
      openWindow,
      closeWindow,
      minimizeWindow,
      restoreWindow,
      desktopApps,
      getWindow
    }
  },
  {
    persist: {
      pick: []
    }
  }
)
