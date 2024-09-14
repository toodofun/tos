<script setup lang="ts">
import { ref, h, defineAsyncComponent } from 'vue'
import IconView from '@/components/operating-system/IconView.vue'
import CommonInfo from '@/views/setting/components/CommonInfo.vue'

const active = ref<SettingItem>()

interface SettingItem {
  id: string
  icon?: string
  title?: string
  left?: ReturnType<typeof h>
  right: ReturnType<typeof h>
}

interface SettingGroup {
  name?: string
  customClass?: string
  items: Array<SettingItem>
}

const settings: Array<SettingGroup> = [
  {
    customClass: 'sticky top-0 bg-white pt-2',
    items: [
      {
        id: 'user',
        left: h(defineAsyncComponent(() => import('@/components/UserView.vue')), {
          user: {
            avatar: 'https://www.dmoe.cc/random.php',
            username: 'TOS账户',
            id: '10001'
          }
        }),
        right: h('div', 'on development...')
      }
    ]
  },
  {
    name: '',
    items: [
      {
        id: 'common',
        icon: 'internal://icon-setting-common',
        title: '通用',
        right: h(CommonInfo)
      },
      {
        id: 'appearance',
        icon: 'internal://icon-setting-appearance',
        title: '外观',
        right: h('div', 'on development...')
      }
    ]
  },
  {
    name: '服务',
    items: [
      {
        id: 'user-manager',
        icon: 'internal://icon-setting-user',
        title: '用户',
        right: h('div', 'on development...')
      },
      {
        id: 'notify',
        icon: 'internal://icon-setting-notify',
        title: '通知',
        right: h('div', 'on development...')
      },
      {
        id: 'network',
        icon: 'internal://icon-setting-network',
        title: '网络',
        right: h(defineAsyncComponent(() => import('@/views/setting/components/NetworkInfo.vue')))
      },
      {
        id: 'secure',
        icon: 'internal://icon-setting-secure',
        title: '安全',
        right: h('div', 'on development...')
      }
    ]
  },
  {
    name: '系统',
    items: [
      {
        id: 'info',
        icon: 'internal://icon-setting-info',
        title: '信息',
        right: h(defineAsyncComponent(() => import('./components/SystemInfo.vue')))
      },
      {
        id: 'update',
        icon: 'internal://icon-setting-update',
        title: '更新',
        right: h('div', 'on development...')
      }
    ]
  },
]

active.value = settings[1].items[0]

const onClick = (item: SettingItem) => {
  active.value = item
}
</script>

<template>
  <a-layout class="absolute inset-0 select-none">
    <a-layout-sider :resize-directions="['right']" style="min-width: 5rem;">
      <div class="px-2 pb-2 flex flex-col gap-2 box-border">
        <div v-for="(group, index) in settings" :key="index" class="flex flex-col gap-0 items-start w-full" :class="group.customClass">
          <div v-if="group.name" class="text-xs text-gray-500 font-bold py-1 px-2 max-w-20 truncate text-ellipsis overflow-hidden">{{ group.name }}</div>
          <div v-for="item in group.items" :key="item.id"
               class="p-2 py-1.5 rounded cursor-pointer w-full"
               :class="active?.id === item.id ? 'bg-slate-300' : 'hover:bg-slate-200'"
               @click="onClick(item)"
          >
            <component v-if="item.left" :is="item.left" />
            <div v-else class="flex items-center flex-nowrap gap-2">
              <div class="w-6 h-6 min-w-6 max-w-6 min-h-6 max-h-6 aspect-square bg-sky-100 rounded-lg">
                <IconView :src="item.icon as string" />
              </div>
              <div class="font-normal text-slate-700 max-w-20 truncate text-ellipsis overflow-hidden">{{ item.title }}</div>
            </div>
          </div>
        </div>
      </div>
    </a-layout-sider>
    <a-layout-content class="scrollbar-hidden">
      <component :is="active?.right" />
    </a-layout-content>
  </a-layout>
</template>

<style scoped>

</style>
