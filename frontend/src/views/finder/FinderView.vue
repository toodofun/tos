<script setup lang="ts">
import { type FileInfo, listDirectory, onUploadFile, getSP } from '@/views/finder/FinderView'
import {
  ref,
  computed,
  toRaw
} from 'vue'
import IconView from '@/components/operating-system/IconView.vue'
import numeral from 'numeral'
import {
  IconLeft,
  IconRight,
  IconRefresh,
  IconCaretDown,
  IconApps,
  IconList,
  IconFolderAdd,
  IconHome
} from '@arco-design/web-vue/es/icon'
import { useQueueStore } from '@/stores/queue'
import type { RequestOption, UploadRequest } from '@arco-design/web-vue'
import eventBus from '@/plugins/eventBus'
import GridView from '@/views/finder/components/GridView.vue'
import ListView from '@/views/finder/components/ListView.vue'

const queueStore = useQueueStore()

eventBus.on('finder-refresh', (path) => {
  const p = path as string
  if (p === currentPath.value) {
    getData(false)
  }
})

const showInList = ref(false)
const sp = ref<FileInfo[]>([])
const files = ref<FileInfo[]>([])
const loading = ref(false)
const currentPath = ref('/')
const selectedInfo = ref<FileInfo[]>([])
const filepathRef = ref()
const pathHistory = ref<string[]>([currentPath.value])
const currentIndex = ref(0)

const init = async () => {
  loading.value = true
  await getSPData()
  await getData()
  loading.value = false
}

const getSPData = async () => {
  sp.value = await getSP()
}

const getData = async (useLoading: boolean = true) => {
  if (useLoading) {
    loading.value = true
  }
  await listDirectory(currentPath.value)
    .then((res) => {
      files.value = res
      selectedInfo.value = []
    })
    .catch(() => {
      files.value = []
    })
    .finally(() => {
      if (useLoading) {
        loading.value = false
      }
      filepathRef.value.blur()
    })
}

const onChangeFilepath = (path: string): void => {
  currentPath.value = path
}

const goToPath = (path: string) => {
  if (path === currentPath.value) return
  currentPath.value = path
  pathHistory.value = pathHistory.value.slice(0, currentIndex.value + 1)
  pathHistory.value.push(currentPath.value)
  currentIndex.value = pathHistory.value.length - 1
  getData()
}

const onDoubleClick = (item: FileInfo) => {
  if (item.isDir) {
    goToPath(item.path)
  }
}

const onClick = (item: FileInfo) => {
  // 暂时不用多选
  // if (selectedInfo.value.includes(item)) {
  //   selectedInfo.value = selectedInfo.value.filter((i) => item !== i)
  // } else {
  //   selectedInfo.value.push(item)
  // }

  selectedInfo.value = [item]
}

// 前往上一个路径
const onClickPrePath = () => {
  if (currentIndex.value > 0) {
    currentIndex.value--
    currentPath.value = pathHistory.value[currentIndex.value]
    getData()
  }
}

// 前往下一个路径
const onClickNextPath = () => {
  if (currentIndex.value < pathHistory.value.length - 1) {
    currentIndex.value++
    currentPath.value = pathHistory.value[currentIndex.value]
    getData()
  }
}

// 判断是否可以返回上一个路径
const canGoBack = computed(() => {
  return currentIndex.value > 0
})

// 判断是否可以前往下一个路径
const canGoForward = computed(() => {
  return currentIndex.value < pathHistory.value.length - 1
})

const switchShowInList = () => {
  showInList.value = !showInList.value
}

const handleUploadModeIgnore = (option: RequestOption): UploadRequest => {
  const { fileItem, name } = option
  const title = name || fileItem.name

  const initialPath = toRaw(currentPath.value)

  queueStore.submit(`Upload ${title}`, async (onProgress) => {
    option.onProgress = onProgress
    await onUploadFile(option, initialPath, 'ignore')
    eventBus.emit('finder-refresh', initialPath)
  })
  return {}
}

const handleUploadModeOverwrite = (option: RequestOption): UploadRequest => {
  const { fileItem, name } = option
  const title = name || fileItem.name

  const initialPath = toRaw(currentPath.value)

  queueStore.submit(`Upload ${title}`, async (onProgress) => {
    option.onProgress = onProgress
    await onUploadFile(option, initialPath, 'overwrite')
    eventBus.emit('finder-refresh', initialPath)
  })
  return {}
}

init()
</script>

<template>
  <a-layout class="absolute inset-0 select-none">
    <a-layout-sider :resize-directions="['right']" style="min-width: 5rem;">
      <div class="relative h-full">
        <div class="flex flex-col gap-2 p-2">
          <div class="cursor-pointer hover:bg-slate-200 rounded p-0.5" v-for="(item, index) in sp" :key="index"
               @click="goToPath(item.path)">
            <div class="flex items-center justify-start gap-1 text-gray-800">
              <div class="w-6 h-6 min-w-6 max-w-6 min-h-6 max-h-6">
                <IconView src="internal://icon-finder-folder" />
              </div>
              <div>{{ item.name }}</div>
            </div>
          </div>
        </div>
        <!--显示右侧文件选中信息-->
        <div
          v-if="selectedInfo.length > 0"
          class="absolute bottom-0 left-0 right-0 w-full text-right bg-white p-1 border-t border-slate-500/10 text-xs text-gray-800">
          <!--          {{ files.length }}&nbsp;个项目&emsp;选中&nbsp;{{ selectedInfo.length }}&nbsp;个-->
          {{ numeral(selectedInfo.reduce((sum, item) => sum + item.size, 0)).format('0.[00] ib') }}
        </div>
      </div>
    </a-layout-sider>
    <a-layout class="overflow-hidden">
      <a-layout-header class="p-2">
        <div class="px-4 flex items-center justify-start gap-1 z-1 box-border">
          <a-button-group size="small" rounded>
            <a-button type="secondary" :disabled="!canGoBack" @click="onClickPrePath">
              <icon-left />
            </a-button>
            <a-button type="secondary" :disabled="!canGoForward" @click="onClickNextPath">
              <icon-right />
            </a-button>
          </a-button-group>
          <a-button-group size="small" rounded>
            <a-button @click="() => {getData()}">
              <icon-refresh size="mini" :class="loading ? 'animate-spin' : ''" />
            </a-button>
          </a-button-group>
          <a-breadcrumb :style="{width:'100%', background: '#f0f1f3'}" :max-count="3">
            <a-breadcrumb-item @click="goToPath('/')">
              <icon-home />
            </a-breadcrumb-item>
            <a-breadcrumb-item
              v-for="item in currentPath.split('/').filter(item => item.trim() !== '')"
              class="cursor-pointer"
              :key="item"
              @click="() => {
                goToPath(currentPath.split('/').slice(0, currentPath.split('/').indexOf(item) + 1).join('/'))
              }"
            >{{ item }}
            </a-breadcrumb-item>
          </a-breadcrumb>
          <a-button-group size="small" rounded>
            <a-dropdown :popup-max-height="false" position="bl">
              <a-button>
                <div class="flex gap-1 items-center">
                  <div class="text-sm">新增</div>
                  <icon-caret-down size="mini" />
                </div>
              </a-button>
              <template #content>
                <a-doption>
                  <div class="flex items-center gap-1">
                    <icon-folder-add />
                    <div>新建文件</div>
                  </div>
                </a-doption>
                <a-doption>
                  <div class="flex items-center gap-1">
                    <icon-folder-add />
                    <div>新建文件夹</div>
                  </div>
                </a-doption>
              </template>
            </a-dropdown>
            <a-dropdown :popup-max-height="false" position="bl">
              <a-button>
                <div class="flex gap-1 items-center">
                  <div class="text-sm">上传</div>
                  <icon-caret-down size="mini" />
                </div>
              </a-button>
              <template #content>
                <a-doption>
                  <div class="flex items-center gap-1">
                    <div class="w-4 h-fit flex items-center">
                      <icon-folder-add />
                    </div>
                    <a-upload
                      multiple
                      :custom-request="handleUploadModeIgnore"
                    >
                      <template #upload-button>
                        <div type="secondary">上传文件-忽略</div>
                      </template>
                    </a-upload>
                  </div>
                </a-doption>
                <a-doption>
                  <div class="flex items-center gap-1">
                    <div class="w-4 h-fit flex items-center">
                      <icon-folder-add />
                    </div>
                    <a-upload
                      multiple
                      :custom-request="handleUploadModeOverwrite"
                    >
                      <template #upload-button>
                        <div type="secondary">上传文件-覆盖</div>
                      </template>
                    </a-upload>
                  </div>
                </a-doption>
                <a-doption>
                  <div class="flex items-center gap-1">
                    <div class="w-4 h-fit flex items-center">
                      <icon-folder-add />
                    </div>
                    <a-upload
                      directory
                      :custom-request="handleUploadModeIgnore"
                    >
                      <template #upload-button>
                        <div type="secondary">上传文件夹-忽略</div>
                      </template>
                    </a-upload>
                  </div>
                </a-doption>
                <a-doption>
                  <div class="flex items-center gap-1">
                    <div class="w-4 h-fit flex items-center">
                      <icon-folder-add />
                    </div>
                    <a-upload
                      directory
                      :custom-request="handleUploadModeOverwrite"
                    >
                      <template #upload-button>
                        <div type="secondary">上传文件夹-覆盖</div>
                      </template>
                    </a-upload>
                  </div>
                </a-doption>
              </template>
            </a-dropdown>
          </a-button-group>
          <a-button-group size="small" rounded>
            <a-button v-if="showInList" @click="switchShowInList">
              <icon-list size="mini" />
            </a-button>
            <a-button v-else @click="switchShowInList">
              <icon-apps size="mini" />
            </a-button>
          </a-button-group>
        </div>
      </a-layout-header>
      <a-layout-content class="w-full relative overflow-y-auto">
        <div v-if="files.length > 0">
          <ListView v-if="showInList" :on-click="onClick" :on-double-click="onDoubleClick" :items="files" />
          <GridView v-else :on-click="onClick" :on-double-click="onDoubleClick" :selected-info="selectedInfo"
                    :items="files" />
        </div>
        <div v-else class="absolute inset-0 flex items-center justify-center z-0">
          <a-empty description="暂无文件" />
        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<style scoped>

</style>
