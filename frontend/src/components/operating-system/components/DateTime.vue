<script setup lang="ts">
import { ref, computed } from 'vue'
import useTimer from '@/hooks/useTimer'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import axios from '@/utils/request'
import { getHolidayDetail, type HolidayDetail } from '@/utils/getHolidayDetail'
import CalendarWidget from '@/components/operating-system/components/CalendarWidget.vue'
import { useDesktopStore } from '@/stores/desktop'

dayjs.locale('zh-cn')

const desktopStore = useDesktopStore()

const currentTime = ref('')
const next7DaysEvent = ref<Array<HolidayDetail>>([])


useTimer(async () => {
  const time = await axios.get<number>('/system/timestamp')
  currentTime.value = dayjs(time).format(desktopStore.statusBarShowSeconds ? `${desktopStore.dateFormat} HH:mm:ss` : `${desktopStore.dateFormat} HH:mm`)
}, { interval: 1000 })

useTimer(async () => {
  const temp: Array<HolidayDetail> = []
  for (let i = 0; i < 7; i++) {
    await getHolidayDetail(dayjs().hour(24 * (i + 1)).format('YYYY-MM-DD')).then((res) => {
      if (res.tags && res.tags.length > 0) {
        temp.push(res)
      }
    })
  }
  next7DaysEvent.value = temp

}, { interval: 1000 * 60 })

</script>

<template>
  <a-popover
    class="overflow-hidden"
    trigger="click"
    position="tl"
    :content-style="{background: 'rgba(255,255,255,0)', padding: '0.5rem', border: 'none', marginRight: '0.6rem', boxShadow: 'none'}"
    :arrow-style="{display: 'none'}">
    <div class="tabular-nums cursor-pointer select-none text-nowrap">{{ currentTime }}</div>
    <template #content>
      <div class="flex flex-col gap-4">
        <CalendarWidget />
      </div>
    </template>
  </a-popover>
</template>

<style scoped>
.writing-vertical-lr {
  writing-mode: vertical-lr;
  text-orientation: upright;
}
</style>
