<template>
    <el-row class="mb-3 text-center" :gutter="10">
      
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statData.websiteCount">
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
        <el-statistic :value="138">
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
        <el-statistic :value="138">
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
  <template #header>
      <div class="card-header">
        <span>系统运行状态</span>
      </div>
    </template>
  <el-row class="mb-3 text-center" :gutter="10">
    <el-col :md="6" :sm="12">

      <el-progress type="dashboard" :percentage="parseFloat(statData.cpu.used_percent).toFixed(2)" status="success">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">CPU 使用率</span>
                
      </template>
    </el-progress>
    <div class="font-size-3">物理/逻辑核心:{{ statData.cpu.physical_count }}/{{ statData.cpu.physical_count }}</div>
 
    </el-col>
    <el-col :md="6" :sm="12">
      <el-progress type="dashboard" :percentage="parseFloat(statData.memory.used_percent).toFixed(2)">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">内存使用率</span>
        
      </template>
    </el-progress>
    <div class="font-size-3">总大小:{{ statData.memory.total }}/可用:{{ statData.memory.available }}</div>
      
    </el-col>
    <el-col :md="6" :sm="12">
      <el-progress type="dashboard" :percentage="parseFloat(statData.swapmem.used_percent).toFixed(2)">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">Swap内存使用率</span>
        
      </template>
    </el-progress>
    <div class="font-size-3">总大小:{{ statData.swapmem.total }}/可用:{{ statData.swapmem.free }}</div>
      
    </el-col>
    <el-col :md="6" :sm="12">
      <el-progress type="dashboard" :percentage="80">
      <template #default="{ percentage }">
        <span class="percentage-value">{{ percentage }}%</span>
        <span class="percentage-label">Progressing</span>
      </template>
    </el-progress>
    </el-col>
  </el-row>
</el-card>

</template>

<script setup>
import { onMounted, onUnmounted, reactive } from 'vue';
import { Icon } from '@iconify/vue';
import serviceRequest from '../../httpclient/client'
const reqDone = ref(true)
const statData = reactive({
  websiteCount:0,
  cpuPercent:0,
  cpu:{},
  memory:{},
  host:{},
  swapmem:{},
})
let refreshInterval;
onMounted(()=>{
  refreshData()
  refreshInterval = setInterval(() => {
        if(reqDone.value === false){
          return;
        }
        refreshData()
      }, 5000); 
})
onUnmounted(()=>{

  clearInterval(refreshInterval)
})

function refreshData(){
  reqDone.value = false
  serviceRequest({
    url: '/v1/statistics/dashboard'
  }).then((data)=>{
    statData.websiteCount = data.user_stats.websiteCount
    statData.cpu = data.cpu
    statData.memory = data.memory
    statData.host = data.host
    statData.swapmem = data.swapmem
    reqDone.value = true
  }).catch(()=>{reqDone.value = false})
}

</script>
<style scoped>

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
</style>