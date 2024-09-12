declare module '@/components/TerminalView.vue' {
  import { DefineComponent } from 'vue'

  const component: DefineComponent<{
    uri: string
    exit: Function
    isActive: boolean
  }, {}, any>

  export default component
}
