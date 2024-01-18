<template>
    <el-row class="mb-3 text-center" :gutter="10" >
      
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statData.user_stats.websiteCount">
        <template #title>
          <div style="display: inline-flex; align-items: center">
            <Icon icon="material-symbols-light:globe" style="font-size: 18px;" :inline="true"  /> 网站数量 
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
            <Icon icon="material-symbols:database-outline" style="font-size: 18px;" :inline="true"  /> 数据库
          </div>
        </template>
        
      </el-statistic>
      </el-card>
      
    </el-col>
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statData.user_stats.cronjobCount">
        <template #title>
          <div >
            <Icon icon="ic:outline-task-alt" style="font-size: 18px;" :inline="true"  /> 计划任务
          </div>
        </template>
       
      </el-statistic>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <div class="el-statistic">
        <div class="el-statistic__head v-mid">
            <Icon icon="ri:time-line" style="font-size: 18px;" :inline="true"  /> 运行时间
          </div>
          <div class="el-statistic__content">
            <span class="el-statistic__number">{{ statData.host.uptime }}</span>
          </div>
       
      </div>
   
      </el-card>
    </el-col>

    
  </el-row>

  <el-card shadow="hover" class="mb-3">
    <div class="card-header font-size-4 mb-2">
        系统运行状态
          </div>
  <el-row class="mb-3 text-center" :gutter="10">
    <el-col :md="6" :sm="12">

      <el-progress type="dashboard" :percentage="cpuPercent" status="success" class="cpu">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">CPU 使用率</span>
                
      </template>
    </el-progress>
    <div class="font-size-3">物理/逻辑核心:{{ statData.cpu.physical_count }}/{{ statData.cpu.physical_count }}</div>
 
    </el-col>
    <el-col :md="6" :sm="12">
      <el-progress type="dashboard" :percentage="memoryPercent" status="success">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">内存使用率</span>
        
      </template>
    </el-progress>
    <div class="font-size-3">总大小:{{ statData.memory.total }}/可用:{{ statData.memory.available }}</div>
      
    </el-col>
    <el-col :md="6" :sm="12">
      <el-progress type="dashboard" :percentage="swapmemPercent" status="success">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">Swap内存使用率</span>
        
      </template>
    </el-progress>
    <div class="font-size-3">总大小:{{ statData.swapmem.total }}/可用:{{ statData.swapmem.free }}</div>
      
    </el-col>
    <el-col :md="6" :sm="12">
      <el-progress type="dashboard" :percentage="diskPercent" status="success">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">挂载点：/</span>
      </template>
    </el-progress>
    <div class="font-size-3">总大小:{{ statData.disk.total }}/可用:{{ statData.disk.free }}</div>
    </el-col>
  </el-row>
</el-card>


<el-row :gutter="20">
    <el-col :md="10" :sm="24">
      <el-card class="box-card mb-3">
        <div class="card-header font-size-4  mb-2">
            服务器信息
          </div>
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
              <td class="color-zinc-400">{{ statData.host.Platform }} {{ statData.host.PlatformVersion }}</td>
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
        <div class="card-header font-size-4  mb-2">
            监控
            <el-radio-group v-model="switchTrafficDisk">
      <el-radio-button label="流量"  />
      <el-radio-button label="硬盘" />
    </el-radio-group>
          </div>
        <div id="zapcharts" style="height: 400px;"></div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import * as echarts from 'echarts/core';
import {
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  DataZoomComponent
} from 'echarts/components';
import { LineChart } from 'echarts/charts';
import { UniversalTransition } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
echarts.use([
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  DataZoomComponent,
  LineChart,
  CanvasRenderer,
  UniversalTransition
]);


import { computed, onMounted, onUnmounted, reactive } from 'vue';
import { Icon } from '@iconify/vue';
import serviceRequest from '../../httpclient/client'
const reqDone = ref(true)
const switchTrafficDisk = ref('流量')
const statData = reactive({
  user_stats:{},
  cpu:{},
  memory:{},
  host:{},
  swapmem:{},
  disk:{},
})
const bootTime = computed(() => {
  var unixTimestamp = new Date(statData.host.BootTime*1000);
  return unixTimestamp.toLocaleString()
})
const cpuPercent = computed(()=> statData.cpu.used_percent ? parseFloat(statData.cpu.used_percent) : 0.00 )
const memoryPercent = computed(()=> statData.memory.used_percent ? parseFloat(statData.memory.used_percent) : 0.00 )
const swapmemPercent = computed(()=> statData.swapmem.used_percent ? parseFloat(statData.swapmem.used_percent) : 0.00 )
const diskPercent = computed(()=> statData.disk.used_percent ? parseFloat(statData.disk.used_percent) : 0.00 )

//定时刷新
let refreshInterval;
onMounted(()=>{
  refreshData()
  refreshInterval = setInterval(() => {
        if(reqDone.value === false){
          return;
        }
        refreshData()
      }, 4000); 
  drawCharts()
})

onUnmounted(()=>{
  clearInterval(refreshInterval)
})


//刷新统计数据
function refreshData(){
  reqDone.value = false
  serviceRequest({
    url: '/v1/statistics/dashboard'
  }).then((data)=>{
    statData.user_stats = data.user_stats
    statData.cpu = data.cpu
    statData.memory = data.memory
    statData.host = data.host
    statData.swapmem = data.swapmem
    statData.disk = data.disk
    reqDone.value = true
  }).catch(()=>{reqDone.value = false})
}



function drawCharts(){
  var chartDom = document.getElementById('zapcharts');
var myChart = echarts.init(chartDom);

// prettier-ignore
const data = [["2000-06-05 11:00", 116], ["2000-06-06 11:05", 129], ["2000-06-07", 135], ["2000-06-08", 86], ["2000-06-09", 73], ["2000-06-10", 85], ["2000-06-11", 73], ["2000-06-12", 68], ["2000-06-13", 92], ["2000-06-14", 130], ["2000-06-15", 245], ["2000-06-16", 139], ["2000-06-17", 115], ["2000-06-18", 111], ["2000-06-19", 309], ["2000-06-20", 206], ["2000-06-21", 137], ["2000-06-22", 128], ["2000-06-23", 85], ["2000-06-24", 94], ["2000-06-25", 71], ["2000-06-26", 106], ["2000-06-27", 84], ["2000-06-28", 93], ["2000-06-29", 85], ["2000-06-30", 73], ["2000-07-01", 83], ["2000-07-02", 125], ["2000-07-03", 107], ["2000-07-04", 82], ["2000-07-05", 44], ["2000-07-06", 72], ["2000-07-07", 106], ["2000-07-08", 107], ["2000-07-09", 66], ["2000-07-10", 91], ["2000-07-11", 92], ["2000-07-12", 113], ["2000-07-13", 107], ["2000-07-14", 131], ["2000-07-15", 111], ["2000-07-16", 64], ["2000-07-17", 69], ["2000-07-18", 88], ["2000-07-19", 77], ["2000-07-20", 83], ["2000-07-21", 111], ["2000-07-22", 57], ["2000-07-23", 55], ["2000-07-24", 60]];
const dateList = data.map(function (item) {
  return item[0];
});
const valueList = data.map(function (item) {
  return item[1];
});
const valueList1 = data.map(function (item) {
  return 23;
});
var option = {
  // Make gradient line here
  visualMap: [
    {
      show: false,
      type: 'continuous',
      seriesIndex: 0,
      min: 0,
      max: 400
    },
    {
      show: false,
      type: 'continuous',
      seriesIndex: 1,
      dimension: 0,
      min: 0,
      max: dateList.length - 1
    }
  ],
  title: [
    {
      left: 'center',
      text: 'Gradient along the y axis'
    },
    {
      top: '55%',
      left: 'center',
      text: 'Gradient along the x axis'
    }
  ],
  tooltip: {
    trigger: 'axis'
  },
  xAxis: [
    {
      data: dateList
    },
    {
      data: dateList,
      gridIndex: 1
    }
  ],
  yAxis: [
    {},
    {
      gridIndex: 1
    }
  ],
  grid: [
    {
      bottom: '60%'
    },
    {
      top: '60%'
    }
  ],
  series: [
    {
      type: 'line',
      showSymbol: false,
      data: valueList
    },
    {
      type: 'line',
      showSymbol: false,
      data: valueList1
    },
    {
      type: 'line',
      showSymbol: false,
      data: valueList
    },
    {
      type: 'line',
      showSymbol: false,
      data: valueList,
      xAxisIndex: 1,
      yAxisIndex: 1
    }
  ]
};

option && myChart.setOption(option);
}

</script>
<style scoped>
.el-col{
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