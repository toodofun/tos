<script setup lang="ts">

defineProps({
  src: {
    type: String,
    required: true
  },
  customClass: {
    type: String,
    required: false
  }
})
</script>

<template>
  <div v-if="src.startsWith('internal://')"
       class="w-full h-full flex items-center justify-center iconfont"
       :class="src.replace('internal://', '') + ' ' + customClass"
  ></div>
  <a-image v-else width="100%" height="100%" :src="src" :preview="false" alt="" fit="cover" />
</template>

<style scoped>

</style>
