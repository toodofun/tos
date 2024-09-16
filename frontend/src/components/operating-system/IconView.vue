<script setup lang="ts">

defineProps({
  src: {
    type: String,
    required: true
  },
  customClass: {
    type: String,
    required: false
  },
  preview: {
    type: Boolean,
    default: false
  }
})
</script>

<template>
  <div v-if="src.startsWith('internal://')"
       class="w-full h-full flex items-center justify-center iconfont"
       :class="src.replace('internal://', '') + ' ' + customClass"
  ></div>
  <div v-else class="w-full h-full rounded overflow-hidden" :class="customClass">
    <a-image width="100%" height="100%" :src="src" :preview="preview" alt="" fit="cover" />
  </div>
</template>

<style scoped>

</style>
