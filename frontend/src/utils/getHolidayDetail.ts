import axios from '@/utils/request'
import dayjs from 'dayjs'
import isBetween from 'dayjs/plugin/isBetween'

dayjs.extend(isBetween)

interface HolidayAPI {
  Name: string
  Version: string
  Generated: string
  Timezone: string
  Author: string
  URL: string
  Years: { [key: string]: Year[] }
}

interface Year {
  Name: string;
  StartDate: string;
  EndDate: string;
  Duration: number;
  CompDays: string[];
  URL: string;
  Memo: string;
}

interface Tag {
  name: string
  color: string
}

export interface HolidayDetail {
  isWorkDay: boolean
  isHoliday: boolean
  isInLieu: boolean
  offline: boolean
  date: dayjs.Dayjs
  tags: Tag[]
}

const getFestivals = (date: dayjs.Dayjs = dayjs()): Tag[] => {
  switch (date.format('MM-DD')) {
    case '02-14':
      return [{ name: '情人节', color: 'bg-red-500' }]
    case '03-08':
      return [{ name: '妇女节', color: 'bg-red-500' }]
    case '05-04':
      return [{ name: '青年节', color: 'bg-red-500' }]
    case '06-01':
      return [{ name: '儿童节', color: 'bg-red-500' }]
    case '07-01':
      return [{ name: '建党节', color: 'bg-red-500' }]
    case '08-01':
      return [{ name: '建军节', color: 'bg-red-500' }]
    case '09-10':
      return [{ name: '教师节', color: 'bg-red-500' }]
    case '10-31':
      return [{ name: '万圣节', color: 'bg-red-500' }]
    case '12-25':
      return [{ name: '圣诞节', color: 'bg-red-500' }]
    case '12-24':
      return [{ name: '平安夜', color: 'bg-red-500' }]
    default:
      return []
  }
}

let holidayAPI: HolidayAPI | null = null

export const getHolidayDetail = async (date: string = dayjs().format('YYYY-MM-DD')): Promise<HolidayDetail> => {
  const dateDayjs = dayjs(date)
  const holidayDetail = {
    isWorkDay: false,
    isHoliday: false,
    isInLieu: false,
    offline: true,
    date: dateDayjs,
    tags: getFestivals(dateDayjs)
  } as HolidayDetail


  if (!holidayAPI) {
    try {
      holidayAPI = await axios.get<HolidayAPI>('/system/holiday')
      holidayDetail.offline = false
      setTimeout(() => {
        holidayAPI = null
      }, 1000 * 60 * 60)
    } catch (error) {
      return holidayDetail
    }
  }
  const year = date.split('-')[0]
  const yearlyHolidays = holidayAPI.Years[year]
  if (!yearlyHolidays) {
    return holidayDetail
  }
  yearlyHolidays.forEach((item: Year) => {
    if (dateDayjs.isBetween(item.StartDate, item.EndDate) || dateDayjs.isSame(item.StartDate) || dateDayjs.isSame(item.EndDate)) {
      holidayDetail.tags.push({ name: item.Name, color: 'bg-sky-500' })
      holidayDetail.isHoliday = true
    } else if (item.CompDays.includes(date)) {
      holidayDetail.tags.push({ name: '调休工作日', color: 'bg-red-500' })
      holidayDetail.isInLieu = true
    }
  })
  holidayDetail.isWorkDay = !holidayDetail.isHoliday

  if ((dateDayjs.format('d') === '6' || dateDayjs.format('d') === '0') && !holidayDetail.isInLieu && !holidayDetail.isHoliday) {
    holidayDetail.tags.push({ name: '休息日', color: 'bg-sky-500' })
    holidayDetail.isWorkDay = false
  }

  return holidayDetail
}
