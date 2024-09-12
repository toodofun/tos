<script setup lang="ts">
import type { PropType } from 'vue'
import DesktopIcon from '@/components/operating-system/components/DesktopIcon.vue'
import type { CreateNewWindowInfo } from '@/stores/windows'

defineProps({
  onOpenWindow: {
    type: Function as PropType<((app: CreateNewWindowInfo) => void)>,
    required: true
  },
  onClose: {
    type: Function as PropType<() => void>,
    required: true
  }
})
</script>

<template>
  <div class="absolute inset-0 bg-slate-800/60 backdrop-blur-lg p-16" @click="onClose">
    <DesktopIcon :app="{
        icon: 'internal://icon-setting',
        title: '设置',
        page: 'internal://setting',
        background: 'linear-gradient(to right, #536976, #292e49)'
      }" @click="onOpenWindow({
        icon: 'internal://icon-setting',
        title: '设置',
        page: 'internal://setting',
        background: 'linear-gradient(to right, #536976, #292e49)',
        singleton: true
      })"/>
  </div>
</template>

<style scoped>

</style>
