<script setup lang="ts">
import DynamicBackground from '@/components/operating-system/components/DynamicBackground.vue'
import type { PropType } from 'vue'
import type { CreateNewWindowInfo } from '@/stores/windows'
import IconView from '@/components/operating-system/IconView.vue'

defineProps({
  app: {
    type: Object as PropType<CreateNewWindowInfo>,
    required: true
  }
})
</script>

<template>
  <div class="w-[100px] h-[100px] flex justify-center items-center cursor-pointer">
    <div class="flex flex-col items-center gap-1">
      <DynamicBackground
        :background="app.background"
        class="w-16 h-16 aspect-square rounded-2xl bg-white/100 overflow-hidden"
      >
        <IconView :src="app.icon" custom-class="text-5xl" />
      </DynamicBackground>
      <a-popover trigger="hover" :content="app.title">
        <div
          class="text-white ltr text-center text-sm font-bold max-w-[88px] text-wrap line-clamp-1 break-words truncate text-ellipsis w-full">
          {{ app.title }}
        </div>
      </a-popover>
    </div>
  </div>
</template>

<style scoped>
.ltr {
  direction: ltr;
}
</style>
