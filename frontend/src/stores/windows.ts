import { ref, h, defineAsyncComponent, type VNode } from 'vue'
import { defineStore } from 'pinia'
import { v4 as uuidv4 } from 'uuid'
import { useDockStore } from '@/stores/dock'
import axios from '@/utils/request'

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

export interface CreateNewWindowInfo {
  id?: string
  title: string
  icon: string
  page: string
  width?: number
  height?: number
  x?: number
  y?: number
  theme?: 'light' | 'dark'
  background?: string
  singleton?: boolean
  fixOnDock?: boolean
  fixOnDesk?: boolean
}

const listApplications = async (): Promise<CreateNewWindowInfo[]> => {
  return axios.get('/application/list')
}

const internalApps: { [key: string]: VNode } = {
  'internal://setting': h(defineAsyncComponent(() => import('@/views/setting/SettingView.vue'))),
  'internal://terminal': h(defineAsyncComponent(() => import('@/views/terminal/TerminalView.vue'))),
  'internal://finder': h(defineAsyncComponent(() => import('@/views/finder/FinderView.vue'))),
  'internal://app-store': h(defineAsyncComponent(() => import('@/views/app-store/AppStoreView.vue')))
}

export const getWindow = (src: string, id: string): VNode => {
  if (src.startsWith('internal')) {
    const app = internalApps[src]
    return app ? h(app, { id }) : h('div', { id }, 'Not Found')
  }

  if (src.startsWith('http') || src.startsWith('//')) {
    return h('iframe', {
      id,
      src,
      class: 'w-full h-full will-change-auto',
      allow: 'camera;microphone;clipboard-write;clipboard-read;',
      sandbox: 'allow-same-origin allow-scripts allow-popups allow-forms'
    })
  }

  return h('div', { id }, 'Not Found')
}

export const useWindowsStore = defineStore('windows', () => {
    // Windows Manager
    const windowList = ref<Array<WindowInfo>>([])
    const maxZ = ref(1)

    const restoreFunctionMap: { [key: string]: () => void } = {}
    const dockStore = useDockStore()

    const allApps = ref<Array<CreateNewWindowInfo>>([])

    // Dock
    const fixedApps = ref<Array<CreateNewWindowInfo>>([
      // { icon: 'internal://icon-task', title: '任务管理', page: '//toodo.fun' },
    ])
    const minimizedApps = ref<Array<WindowInfo>>([])

    const desktopApps = ref<Array<CreateNewWindowInfo>>([
      // {
      //   icon: 'https://i04piccdn.sogoucdn.com/a72804451f0e9825',
      //   title: '原神',
      //   width: 1280,
      //   height: 720,
      //   page: 'https://genshin.titlecan.cn/'
      // },
      // // {
      // //   icon: 'https://hexgl.bkcore.com/play/css/title.png',
      // //   title: 'HexGL赛车',
      // //   width: 1280,
      // //   height: 720,
      // //   page: 'https://hexgl.bkcore.com/play/'
      // // },
      // // {
      // //   icon: 'https://th.bing.com/th/id/OIP.nfFu7l8TPI6fnX5Fb8bJ_QHaHa?rs=1&pid=ImgDetMain',
      // //   title: '亲戚称呼计算器',
      // //   width: 338,
      // //   height: 600,
      // //   page: 'https://passer-by.com/relationship/vue/#/',
      // //   theme: 'dark'
      // // },
      // { icon: 'internal://icon-container', title: '容器管理', page: '//toodo.fun' },
      // { icon: 'internal://icon-cluster', title: '集群管理', page: '//toodo.fun', theme: 'dark' },
    ])

    listApplications()
      .then(res => {
        allApps.value = res
        fixedApps.value = res.filter(i => i.fixOnDock)
        desktopApps.value = res.filter(i => i.fixOnDesk)
      })
      .catch(e => {
        console.log(e)
      })

    function clickWindow(w: WindowInfo) {
      maxZ.value += 1
      w.z = maxZ.value
    }

    function openWindow(w: CreateNewWindowInfo) {
      if (w.page === 'system://launchpad') {
        dockStore.setShowLaunchpad(true)
        return
      }
      dockStore.setShowLaunchpad(false)

      if (w.singleton) {
        const i = minimizedApps.value.findIndex((item) => {
          return item.page === w.page
        })
        if (i > -1) {
          restoreWindow(minimizedApps.value[i].id)
          return
        }
        const index = windowList.value.findIndex((item) => {
          return item.page === w.page
        })
        if (index > -1) {
          clickWindow(windowList.value[index])
          return
        }
      }
      maxZ.value += 1
      windowList.value.push({
        id: uuidv4(),
        title: w.title,
        icon: w.icon,
        page: w.page,
        width: w.width || 800,
        height: w.height || 525,
        x: w.x || window.innerWidth / 2 - (w.width || 800) / 2,
        y: w.y || window.innerHeight / 2 - (w.height || 525) / 2,
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

      const window = minimizedApps.value.find((item) => {
        return item.id === id
      })
      if (window) {
        clickWindow(window)
        minimizedApps.value = minimizedApps.value.filter((item) => {
          return item.id !== id
        })
      }
    }

    return {
      allApps,
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
