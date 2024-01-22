<template>
  <el-row class="mb-3 text-center" :gutter="10">
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statData.user_stats.websiteCount">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              <Icon icon="material-symbols-light:globe" style="font-size: 18px" :inline="true" />
              网站数量
            </div>
          </template>
        </el-statistic>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statData.user_stats.databaseCount">
          <template #title>
            <div style="display: inline-flex; align-items: center">
              <Icon
                icon="material-symbols:database-outline"
                style="font-size: 18px"
                :inline="true"
              />
              数据库
            </div>
          </template>
        </el-statistic>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statData.user_stats.cronjobCount">
          <template #title>
            <div>
              <Icon icon="ic:outline-task-alt" style="font-size: 18px" :inline="true" /> 计划任务
            </div>
          </template>
        </el-statistic>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <div class="el-statistic">
          <div class="el-statistic__head v-mid">
            <Icon icon="ri:time-line" style="font-size: 18px" :inline="true" /> 运行时间
          </div>
          <div class="el-statistic__content">
            <span class="el-statistic__number">{{ statData.host.uptime }}</span>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>

  <el-card shadow="hover" class="mb-3">
    <div class="card-header font-size-4 mb-2">系统运行状态</div>
    <el-row class="mb-3 text-center" :gutter="10">
      <el-col :md="6" :sm="12">
        <el-progress type="dashboard" :percentage="cpuPercent" status="success" class="cpu">
          <template #default="{ percentage }">
            <span class="percentage-value">{{ percentage }}%</span>
            <span class="percentage-label">CPU 使用率</span>
          </template>
        </el-progress>
        <div class="font-size-3">CPU核心数:{{ statData.cpu.physical_count }}</div>
      </el-col>
      <el-col :md="6" :sm="12">
        <el-progress type="dashboard" :percentage="memoryPercent" status="success">
          <template #default="{ percentage }">
            <span class="percentage-value">{{ percentage }}%</span>
            <span class="percentage-label">内存使用率</span>
          </template>
        </el-progress>
        <div class="font-size-3">
          总大小:{{ statData.memory.total }}/可用:{{ statData.memory.available }}
        </div>
      </el-col>
      <el-col :md="6" :sm="12">
        <el-progress type="dashboard" :percentage="swapmemPercent" status="success">
          <template #default="{ percentage }">
            <span class="percentage-value">{{ percentage }}%</span>
            <span class="percentage-label">Swap内存使用率</span>
          </template>
        </el-progress>
        <div class="font-size-3">
          总大小:{{ statData.swapmem.total }}/可用:{{ statData.swapmem.free }}
        </div>
      </el-col>
      <el-col :md="6" :sm="12">
        <el-progress type="dashboard" :percentage="diskPercent" status="success">
          <template #default="{ percentage }">
            <span class="percentage-value">{{ percentage }}%</span>
            <span class="percentage-label">挂载点：/</span>
          </template>
        </el-progress>
        <div class="font-size-3">
          总大小:{{ statData.disk.total }}/可用:{{ statData.disk.free }}
        </div>
      </el-col>
    </el-row>
  </el-card>

  <el-row :gutter="20">
    <el-col :md="10" :sm="24">
      <el-card class="box-card mb-3">
        <div class="card-header font-size-4 mb-2">服务器信息</div>
        <div class="text item">
          <table style="width: 100%" class="border-1">
            <tr>
              <td>主机名称</td>
              <td class="color-zinc-400">{{ statData.host.Hostname }}</td>
            </tr>
            <tr>
              <td>操作系统</td>
              <td class="color-zinc-400">{{ statData.host.OS }}</td>
            </tr>
            <tr>
              <td>发行版</td>
              <td class="color-zinc-400">
                {{ statData.host.Platform }} {{ statData.host.PlatformVersion }}
              </td>
            </tr>
            <tr>
              <td>CPU 架构</td>
              <td class="color-zinc-400">{{ statData.host.KernelArch }}</td>
            </tr>
            <tr>
              <td>内核版本</td>
              <td class="color-zinc-400">{{ statData.host.KernelVersion }}</td>
            </tr>
            <tr>
              <td>服务器启动时间</td>
              <td class="color-zinc-400">{{ bootTime }}</td>
            </tr>
            <tr>
              <td>服务器已经运行</td>
              <td class="color-zinc-400">{{ statData.host.uptime }}</td>
            </tr>
          </table>
        </div>
      </el-card>
    </el-col>
    <el-col :md="14" :sm="24">
      <el-card class="box-card">
        <div class="card-header font-size-4 mb-2">
          监控
          <!-- <el-radio-group v-model="switchTrafficDisk">
              <el-radio-button label="流量"  />
              <el-radio-button label="硬盘" />
            </el-radio-group> -->
        </div>
        <div style="position: absolute">
          <el-tag class="ml-2" type="success">{{ monitorTags.tag1 }}</el-tag>
          <el-tag class="ml-2" type="success">{{ monitorTags.tag2 }}</el-tag>
          <el-tag class="ml-2" type="success">{{ monitorTags.tag3 }}</el-tag>
          <el-tag class="ml-2" type="success">{{ monitorTags.tag4 }}</el-tag>
        </div>
        <div id="zapcharts" style="height: 300px"></div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import * as echarts from 'echarts/core'
import {
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  DataZoomComponent,
  LegendComponent
} from 'echarts/components'
import { LineChart } from 'echarts/charts'
import { UniversalTransition } from 'echarts/features'
import { CanvasRenderer } from 'echarts/renderers'

echarts.use([
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  DataZoomComponent,
  LineChart,
  CanvasRenderer,
  UniversalTransition,
  LegendComponent
])

import { formatBytes, fmtBytes } from '@/commons/commons'
import { computed, onMounted, onUnmounted, reactive } from 'vue'
import { Icon } from '@iconify/vue'
import apiRequest from '@/httpclient/client'
const reqDone = ref(true)
const switchTrafficDisk = ref('流量')
const monitorTags = reactive({
  tag1: '',
  tag2: '',
  tag3: '',
  tag4: ''
})
const statData = reactive({
  user_stats: {},
  cpu: {},
  memory: {},
  host: {},
  swapmem: {},
  disk: {}
})
const bootTime = computed(() => {
  var unixTimestamp = new Date(statData.host.BootTime * 1000)
  return unixTimestamp.toLocaleString()
})
const cpuPercent = computed(() =>
  statData.cpu.used_percent ? parseFloat(statData.cpu.used_percent) : 0.0
)
const memoryPercent = computed(() =>
  statData.memory.used_percent ? parseFloat(statData.memory.used_percent) : 0.0
)
const swapmemPercent = computed(() =>
  statData.swapmem.used_percent ? parseFloat(statData.swapmem.used_percent) : 0.0
)
const diskPercent = computed(() =>
  statData.disk.used_percent ? parseFloat(statData.disk.used_percent) : 0.0
)

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
  drawCharts()
  window.addEventListener('resize', resizeChart)
})

onUnmounted(() => {
  clearInterval(refreshInterval)
  window.removeEventListener('resize', resizeChart)
})

const resizeChart = () => {
  window.myChart.resize()
}

//刷新统计数据
function refreshData() {
  reqDone.value = false
  apiRequest({
    url: '/v1/statistics/dashboard'
  })
    .then((data) => {
      statData.user_stats = data.user_stats
      statData.cpu = data.cpu
      statData.memory = data.memory
      statData.host = data.host
      statData.swapmem = data.swapmem
      statData.disk = data.disk
      reqDone.value = true
      
      //监控Tag
      monitorTags.tag1 = '已发送:' + formatBytes(data.netBytesSent, 2)
      monitorTags.tag2 = '已接收:' + formatBytes(data.netBytesRecv, 2)
      let [sendBytes, sendUnit] = fmtBytes(data.netBytesSentPerSec, 2)
      let [recvBytes, recvUnit] = fmtBytes(data.netBytesRecvPerSec, 2)
      monitorTags.tag3 = `下行：${sendBytes}${sendUnit}/秒`
      monitorTags.tag4 = `上行：${recvBytes}${recvUnit}/秒`
    
      var option = window.myChart.getOption()
      option.yAxis = [
      {
        boundaryGap: [0, '100%'],
        alignTicks: true,
        axisLabel: {
          formatter: '{value} '+ sendUnit
        }
      },{
        boundaryGap: [0, '100%'],
        alignTicks: true,
        axisLabel: {
          formatter: '{value} '+ recvUnit
        }
      }
      ]
      let xAxisData = option.xAxis[0].data
      let seriesData = option.series[0].data
      let series1Data = option.series[1].data
      if (seriesData.length > 100) {
        xAxisData.shift()
        seriesData.shift()
        series1Data.shift()
      }
      xAxisData.push([new Date().toLocaleTimeString()])

      seriesData.push(sendBytes)
      series1Data.push(recvBytes)

      window.myChart.setOption(option)
    })
    .catch(() => {
      reqDone.value = false
    })
}

function drawCharts() {
  var chartDom = document.getElementById('zapcharts')
  var myChart = echarts.init(chartDom)
  window.myChart = myChart
  var option

  option = {
    legend: {
      data: ['上行', '下行'],
      right: 10
    },
    dataZoom: [
      {
        type: 'inside'
      },
      {}
    ],
    tooltip: {
      trigger: 'axis'
    },
    xAxis: [
      {
        data: [],
        boundaryGap: false,
        showSymbol: false,
        splitLine: {
          show: false
        }
      }
    ],
    yAxis: [
      {}
    ],
    series: [
      {
        type: 'line',
        name: '上行',
        showSymbol: false,
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgba(0, 94, 235, .3)'
            },
            {
              offset: 1,
              color: 'rgba(0, 94, 235, 0)'
            }
          ])
        },
        data: []
      },
      {
        type: 'line',
        name: '下行',
        showSymbol: false,
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgba(27, 143, 60, .3)'
            },
            {
              offset: 1,
              color: 'rgba(27, 143, 60, 0)'
            }
          ])
        },
        data: []
      }
    ]
  }

  option && myChart.setOption(option)
}
</script>
<style scoped>
.el-col {
  margin-bottom: 8px;
}
.percentage-value {
  display: block;
  margin-top: 10px;
  font-size: 28px;
}
.percentage-label {
  display: block;
  margin-top: 10px;
  font-size: 12px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  /* padding: 10px !important;
  font-size: 0.75rem; */
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>
