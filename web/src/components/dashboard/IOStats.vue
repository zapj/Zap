<template>
  <el-card class="box-card" style="width: 100%" id="drawContent">
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
    <div id="zapcharts" style="height: 300px; "></div>
  </el-card>
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
import { inject, reactive, ref } from 'vue'
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
import { computed, onMounted, onUnmounted } from 'vue'
import { Icon } from '@iconify/vue'
const statsData = inject('statsData')
const switchTrafficDisk = ref('流量')
const monitorTags = reactive({
  tag1: '',
  tag2: '',
  tag3: '',
  tag4: ''
})

const resizeChart = () => {
  window.myChart.resize()
}
onMounted(() => {
  drawCharts()
  renderCharts()
  window.addEventListener('resize', resizeChart)
})
onUnmounted(()=>{
  window.removeEventListener('resize', resizeChart)
})
watch(
  () => statsData.net,
  (newVal, oldVal) => {
    renderCharts()
  }
)

const renderCharts = () => {
  monitorTags.tag1 = '已发送:' + formatBytes(statsData.net.send_bytes, 2)
  monitorTags.tag2 = '已接收:' + formatBytes(statsData.net.recv_bytes, 2)
  let [sendBytes, sendUnit] = fmtBytes(statsData.net.send_bytes_per_sec, 2)
  let [recvBytes, recvUnit] = fmtBytes(statsData.net.recv_bytes_per_sec, 2)
  monitorTags.tag3 = `下行：${recvBytes}${recvUnit}/秒`
  monitorTags.tag4 = `上行：${sendBytes}${sendUnit}/秒`

  var option = window.myChart.getOption()
  option.tooltip = {
    trigger: 'axis',
    formatter: function (params) {
      let relVal = params[0].axisValueLabel
      for (let i = 0; i < params.length; i++) {
        relVal += '<br/>' + params[i].marker + params[i].seriesName + ':  ' + params[i].data

        if (params[i].seriesName == '下行') {
          relVal += `${recvUnit}`
        } else {
          relVal += `${sendUnit}`
        }
      }
      return relVal
    }
  }
  option.yAxis = [
    {
      boundaryGap: [0, '100%'],
      alignTicks: true,
      axisLabel: {
        formatter: `{value} ${recvUnit}`
      }
    }
  ]
  let xAxisData = option.xAxis[0].data
  let seriesData = option.series[0].data
  let series1Data = option.series[1].data
  if (seriesData.length > 200) {
    xAxisData.shift()
    seriesData.shift()
    series1Data.shift()
  }
  xAxisData.push([new Date().toLocaleTimeString()])

  seriesData.push(sendBytes)
  series1Data.push(recvBytes)

  window.myChart.setOption(option)
}

function drawCharts() {
  var chartDom = document.getElementById('zapcharts')
  var myChart = echarts.init(chartDom)
  window.myChart = myChart
  var option = {
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
    yAxis: [{}],
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
