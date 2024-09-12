<script setup lang="ts">
import { computed } from 'vue'
import { defaultBackground } from '@/stores/windows'

const props = defineProps({
  background: {
    type: String,
    required: false,
    default: defaultBackground
  }
})

const dynamicBackgroundStyle = computed(() => {
  if (props.background?.startsWith('linear-gradient')) {
    return { backgroundImage: props.background }
  } else {
    return { backgroundColor: props.background }
  }
})

</script>

<template>
  <div :style="dynamicBackgroundStyle">
    <slot />
  </div>
</template>

<style scoped>

</style>
