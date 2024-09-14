<script setup lang="ts">
import { useDesktopStore } from '@/stores/desktop'
import dayjs from 'dayjs'

const desktopStore = useDesktopStore()

const onDateFormatChange = (value: any) => {
  desktopStore.setDateFormat(value)
}

</script>

<template>
  <div class="w-full h-full p-2 flex flex-col gap-2">
    <div class="bg-white rounded-md p-2 flex flex-col gap-2">
      <div class="flex items-center justify-between">
        <div>显示秒级时间</div>
        <a-switch v-model="desktopStore.statusBarShowSeconds" />
      </div>
      <div class="flex items-center justify-between">
        <div>系统时间格式</div>
        <a-select
          :style="{width:'220px'}"
          :bordered="false"
          :model-value="desktopStore.dateFormat"
          @change="onDateFormatChange"
        >
          <a-option value="M月D日 dddd" :label="dayjs().format('M月D日 dddd')"></a-option>
          <a-option value="M/D/YYYY dddd" :label="dayjs().format('M/D/YYYY dddd')"></a-option>
          <a-option value="YYYY-MM-DD dddd" :label="dayjs().format('YYYY-MM-DD dddd')"></a-option>
          <a-option value="YYYY/MM/DD dddd" :label="dayjs().format('YYYY/MM/DD dddd')"></a-option>
          <a-option value="YYYY年M月D日 dddd" :label="dayjs().format('YYYY年M月D日 dddd')"></a-option>
        </a-select>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
