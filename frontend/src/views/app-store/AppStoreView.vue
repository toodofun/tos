<script setup lang="ts">
import { IconApps } from '@arco-design/web-vue/es/icon'

const images = [
  'https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/cd7a1aaea8e1c5e3d26fe2591e561798.png~tplv-uwbnlip3yd-webp.webp',
  'https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/6480dbc69be1b5de95010289787d64f1.png~tplv-uwbnlip3yd-webp.webp',
  'https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/0265a04fddbd77a19602a15d9d55d797.png~tplv-uwbnlip3yd-webp.webp',
  'https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/24e0dd27418d2291b65db1b21aa62254.png~tplv-uwbnlip3yd-webp.webp'
]
</script>

<template>
  <a-layout class="absolute inset-0 select-none">
    <a-layout-sider class="p-2 box-border">
      <a-input :style="{width:'320px'}" placeholder="Please enter something" allow-clear />
      <a-input-search placeholder="搜索应用"></a-input-search>
      <a-menu
      >
        <a-menu-item key="explore">探索
          <template #icon>
            <icon-apps></icon-apps>
          </template>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout-content class="p-4">
      <a-carousel
        :autoPlay="true"
        animation-name="card"
        indicator-position="outer"
        :style="{
      width: '100%',
      height: '240px',
    }"
      >
        <a-carousel-item :key="image" v-for="image in images" :style="{ width: '88%' }">
          <img
            :src="image"
            :style="{
          width: '100%',
        }"
            alt="" />
        </a-carousel-item>
      </a-carousel>
    </a-layout-content>
  </a-layout>
</template>

<style scoped>

</style>
