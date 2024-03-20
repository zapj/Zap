<template>
  <UserUsage></UserUsage>
  <SysStats></SysStats>
  <el-row :gutter="20">
    <el-col :md="10" :sm="24">
      <el-card class="box-card mb-3">
        <div class="card-header font-size-4 mb-2">服务器信息</div>
        <div class="text item">
          <table style="width: 100%" class="border-1">
            <tr>
              <td>主机名称</td>
              <td class="color-zinc-400">{{ statsData.host.Hostname }}</td>
            </tr>
            <tr>
              <td>操作系统</td>
              <td class="color-zinc-400">{{ statsData.host.OS }}</td>
            </tr>
            <tr>
              <td>发行版</td>
              <td class="color-zinc-400">
                {{ statsData.host.Platform }} {{ statsData.host.PlatformVersion }}
              </td>
            </tr>
            <tr>
              <td>CPU 架构</td>
              <td class="color-zinc-400">{{ statsData.host.KernelArch }}</td>
            </tr>
            <tr>
              <td>内核版本</td>
              <td class="color-zinc-400">{{ statsData.host.KernelVersion }}</td>
            </tr>
            <tr>
              <td>服务器启动时间</td>
              <td class="color-zinc-400">{{ bootTime }}</td>
            </tr>
            <tr>
              <td>服务器已经运行</td>
              <td class="color-zinc-400">{{ statsData.host.uptime }}</td>
            </tr>
            <tr>
              <td>ZAP 版本</td>
              <td class="color-zinc-400">
                {{ globalStore.settings.version }}({{ globalStore.settings.build_date }})
              </td>
            </tr>
          </table>
        </div>
      </el-card>
    </el-col>
    <el-col :md="14" :sm="24">
      <Suspense>
        <template v-slot:default>
          <IOStats></IOStats>
        </template>

        <template v-slot:fallback>
          <el-skeleton :rows="4" animated translate="yes" />
        </template>
      </Suspense>
    </el-col>
  </el-row>
</template>

<script setup>
import { defineAsyncComponent, provide, onMounted, onUnmounted, ref } from 'vue'
import apiRequest from '@/httpclient/client'
import UserUsage from '@/components/dashboard/UserUsage.vue'
import SysStats from '@/components/dashboard/SysStats.vue'
const reqDone = ref(true)
import { useGlobalStore } from '@/stores/global'
// const SmallStatics = defineAsyncComponent(() => import('../components/dashboard/SmallStatistics.vue'))
const IOStats = defineAsyncComponent(() => import('@/components/dashboard/IOStats.vue'))
const globalStore = useGlobalStore()
const statsData = reactive({
  user_stats: {},
  cpu: {},
  memory: {},
  host: {},
  swapmem: {},
  disk: {},
  net: {}
})
provide('statsData', statsData)
const bootTime = computed(() => {
  var unixTimestamp = new Date(statsData.host.BootTime * 1000)
  return unixTimestamp.toLocaleString()
})

//定时刷新
let refreshInterval
onMounted(() => {
  refreshData()
  refreshInterval = setInterval(() => {
    if (reqDone.value === false) {
      return
    }
    refreshData()
  }, 4000)
})

onUnmounted(() => {
  clearInterval(refreshInterval)
})

//刷新统计数据
function refreshData() {
  reqDone.value = false
  apiRequest({
    url: '/v1/statistics/dashboard',
    loading: false
  })
    .then((data) => {
      statsData.user_stats = data.user_stats
      statsData.cpu = data.cpu
      statsData.memory = data.memory
      statsData.host = data.host
      statsData.swapmem = data.swapmem
      statsData.disk = data.disk
      statsData.net = data.net
    })
    .finally(() => {
      reqDone.value = true
    })
}
</script>
<style scoped>
.text {
  font-size: 14px;
}
</style>
