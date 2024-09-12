declare module '@/components/VueDragResize.vue' {
  import { DefineComponent } from 'vue'

  const component: DefineComponent<{
    w: number
    h: number
    x: number
    y: number
    z: number
    minw: number
    minh: number
    isActive: boolean
    isResizable: boolean
    isDraggable: boolean
    stickSize: number
    sticks: string[]
  }, {}, any>

  export default component
}
