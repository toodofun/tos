<script setup lang="ts">
import { ref } from 'vue'
import useTimer from '@/hooks/useTimer'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import axios from '@/utils/request'
import { getHolidayDetail, type HolidayDetail } from '@/utils/getHolidayDetail'

dayjs.locale('zh-cn')

const props = defineProps({
  showSeconds: {
    type: Boolean,
    default: false
  }
})

const currentTime = ref('')
const next7DaysEvent = ref<Array<HolidayDetail>>([])


useTimer(async () => {
  const time = await axios.get<number>('/system/timestamp')
  currentTime.value = dayjs(time).format(props.showSeconds ? 'M月D日 dddd HH:mm:ss' : 'M月D日 dddd HH:mm')
}, { interval: props.showSeconds ? 1000 : 1000 * 60 })

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
  <div
    class="w-full h-full select-none bg-white p-4 rounded-lg flex gap-0 flex-nowrap min-w-64 shadow-lg overflow-hidden">
    <div class="flex flex-col flex-1 gap-4">
      <div class="flex items-center gap-0">
        <div class="text-5xl font-bold">{{ dayjs().format('D') }}</div>
        <div class="w-[1px] bg-slate-500 h-6 ml-2 mr-1"></div>
        <div class="writing-vertical-lr select-none font-bold">{{ dayjs().format('ddd') }}</div>
      </div>
      <div class="h-full w-full flex items-center justify-start text-xs text-gray-500">无事项</div>
    </div>
    <div class="flex flex-col gap-2 flex-1">
      <template v-if="next7DaysEvent.length > 0">
        <div class="flex flex-col gap-1" v-for="(item, index) in next7DaysEvent" :key="index">
          <div class="text-xs font-bold text-nowrap">{{ item.date.format('M月D日 dddd') }}</div>
          <div class="text-xs font-bold text-nowrap flex gap-1 items-center" v-for="(tag, index) in item.tags"
               :key="index">
            <div class="w-[0.3rem] h-4 rounded-full" :class="tag.color"></div>
            <div>{{ tag.name }}</div>
          </div>
        </div>
      </template>
      <div v-else class="w-full h-full flex items-center justify-center">
        <a-empty />
      </div>
    </div>
  </div>
</template>

<style scoped>
.writing-vertical-lr {
  writing-mode: vertical-lr;
  text-orientation: upright;
}
</style>
