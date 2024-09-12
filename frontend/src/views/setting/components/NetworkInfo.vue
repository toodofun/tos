<script setup lang="ts">
import { ref } from 'vue'
import axios from '@/utils/request'

const networks = ref<NetworkItem[]>([])
const activeKeys = ref<Array<string | number>>([])

export interface NetworkItem {
  index: number;
  mtu: number;
  name: string;
  hardwareAddr: string;
  flags: string[] | null;
  addrs: Addr[];
}

export interface Addr {
  addr: string;
}

const init = async () => {
  const temp = await axios.get<NetworkItem[]>('/system/network')
  networks.value = temp.filter((item: NetworkItem) => {
    // item.addrs = item.addrs.filter((addr: Addr) => !addr.addr.startsWith('fe80') && !addr.addr.startsWith('127.0.0.1'))
    return item.hardwareAddr && item.addrs && item.addrs.length > 0
  })
  activeKeys.value = networks.value.map((item: NetworkItem) => item.name)
}

const onChange = (key: (string | number)[], ev: Event): any => {
  activeKeys.value = key
}
init()
</script>

<template>
  <div class="w-full h-full">
    <a-collapse :active-key="activeKeys" expand-icon-position="right" :bordered="false" @change="onChange">
      <a-collapse-item v-for="network in networks" :key="network.name">
        <template #header>
          <div class="font-bold">{{ network.name }}</div>
        </template>
        <div class="flex flex-row w-full items-start text-sm">
          <div class="basis-1/3">MAC地址:</div>
          <div class="basis-2/3">{{ network.hardwareAddr }}</div>
        </div>
        <div class="flex flex-row w-full items-start text-sm">
          <div class="basis-1/3">IP地址:</div>
          <div class="basis-2/3 select-text">
            <div v-for="addr in network.addrs" :key="addr.addr">
              {{ addr.addr }}
            </div>
          </div>
        </div>
        <div class="flex flex-row w-full items-start text-sm">
          <div class="basis-1/3">网络状态:</div>
          <div class="basis-2/3">MTU {{ network.mtu }}</div>
        </div>
      </a-collapse-item>
    </a-collapse>
  </div>
</template>

<style scoped>

</style>
