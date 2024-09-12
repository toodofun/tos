<script setup lang="ts">
import IconView from '@/components/operating-system/IconView.vue'

interface UserViewProps {
  avatar: string
  username: string
  id: string
}

defineProps({
  user: {
    type: Object as () => UserViewProps,
    required: true
  }
})

</script>

<template>
  <div class="flex items-center gap-2">
    <div class="min-w-9 min-h-9 w-9 h-9 max-w-9 max-h-9 aspect-square rounded-full overflow-hidden">
      <IconView :src="user.avatar" />
    </div>
    <div class="w-full flex flex-col gap-0.5 items-start">
      <div class="font-medium capitalize truncate text-ellipsis overflow-hidden">{{ user.username }}</div>
      <div class="text-xs text-gray-500 max-w-20 truncate text-ellipsis overflow-hidden">{{ user.id }}</div>
    </div>
  </div>
</template>

<style scoped>

</style>
