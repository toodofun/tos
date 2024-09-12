<script setup lang="ts">
import { ref } from 'vue'
import numeral from 'numeral'
import axios from '@/utils/request'
import dayjs from 'dayjs'

interface SystemInfo {
  basic: Basic
}

interface Basic {
  version: string;
  commit: string;
  buildTime: string;
  hostname: string;
}

interface InfoItem {
  title: string;
  value: Function;
}

interface InfoGroup {
  title: string;
  items: Array<InfoItem>;
}

const systemInfo = ref<SystemInfo>()

const init = async () => {
  systemInfo.value = (await axios.get<SystemInfo>('/system/info'))
}

const infos: InfoGroup[] = [
  {
    title: '基本信息',
    items: [
      {
        title: '服务器名称',
        value: () => {
          return getValueByPath(systemInfo.value, 'basic.hostname')
        }
      },
      {
        title: 'TOS版本',
        value: () => {
          return getValueByPath(systemInfo.value, 'basic.version')
        }
      },
      {
        title: '运维版本',
        value: () => {
          return getValueByPath(systemInfo.value, 'basic.commit')
        }
      },
      {
        title: '构建时间',
        value: () => {
          return getValueByPath(systemInfo.value, 'basic.buildTime')
        }
      }
    ]
  },
  {
    title: '硬件',
    items: [
      {
        title: 'CPU',
        value: () => {
          let result: Array<string> = []
          const cpus = getValueByPath(systemInfo.value, 'hardware.cpu', null)
          if (cpus) {
            cpus.forEach((cpu: any) => {
              const t = getValueByPath(cpu, 'modelName')
              if (!result.includes(t)) {
                result.push(t)
              }
            })
          }
          return result.join(', ')
        }
      },
      {
        title: 'CPU 时钟频率(MHz)',
        value: () => {
          return getValueByPath(systemInfo.value, 'hardware.cpu.0.mhz')
        }
      },
      {
        title: 'CPU 内核数',
        value: () => {
          let result = 0
          const cpus = getValueByPath(systemInfo.value, 'hardware.cpu', null)
          if (cpus) {
            cpus.forEach((cpu: any) => {
              result += getValueByPath(cpu, 'cores', 0)
            })
          }
          return result
        }
      },
      {
        title: '物理内存大小',
        value: () => {
          return numeral(getValueByPath(systemInfo.value, 'hardware.mem.total')).format('0 ib')
        }
      },
      {
        title: '已用内存大小',
        value: () => {
          return numeral(getValueByPath(systemInfo.value, 'hardware.mem.used')).format('0 ib')
        }
      },
      {
        title: '可用内存大小',
        value: () => {
          return numeral(getValueByPath(systemInfo.value, 'hardware.mem.available')).format('0 ib')
        }
      },
      {
        title: '已用内存比例',
        value: () => {
          return numeral(getValueByPath(systemInfo.value, 'hardware.mem.usedPercent')).format('0.00') + '%'
        }
      }
    ]
  },
  {
    title: '系统信息',
    items: [
      {
        title: '开机时间',
        value: () => {
          return dayjs(getValueByPath(systemInfo.value, 'hardware.host.bootTime') * 1000).format('YYYY-MM-DD HH:mm:ss')
        }
      },
      {
        title: '运行时间',
        value: () => {
          return numeral(getValueByPath(systemInfo.value, 'hardware.host.uptime')).format('00:00:00')
        }
      },
      {
        title: '虚拟化系统',
        value: () => {
          const vm = getValueByPath(systemInfo.value, 'hardware.host.virtualizationSystem')
          return vm ? vm : '非虚拟机'
        }
      }
    ]
  }
]

const getValueByPath = (obj: any, path: string, defaultValue: any = 'UNKNOWN'): any => {
  return path.split('.').reduce((o, key) => (o && o[key] !== undefined ? o[key] : defaultValue), obj)
}


init()

</script>

<template>
  <div class="w-full h-full">
    <a-collapse :default-active-key="infos.map(i => i.title)" :bordered="true">
      <a-collapse-item v-for="info in infos" :key="info.title" :header="info.title">
        <a-list :bordered="false">
          <a-list-item v-for="(item, index) in info.items" :key="index">
            <div class="flex items-center">
              <div class="flex-1 text-xs font-medium truncate text-ellipsis overflow-hidden">{{ item.title }}</div>
              <div class="flex-1 text-xs text-gray-600 tabular-nums truncate text-ellipsis overflow-hidden">
                {{ item.value() }}
              </div>
            </div>
          </a-list-item>
        </a-list>
      </a-collapse-item>
    </a-collapse>
  </div>
</template>

<style scoped>
:deep(.arco-collapse-item-content) {
  padding: 0;
}

:deep(.arco-collapse-item-content-box) {
  padding: 0;
}

:deep(.arco-list-item) {
  padding: 0.5rem 2rem !important;
}
</style>
