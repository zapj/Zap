<template>
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
        <div class="font-size-3">CPU核心数:{{ statsData.cpu.physical_count }}</div>
      </el-col>
      <el-col :md="6" :sm="12">
        <el-progress type="dashboard" :percentage="memoryPercent" status="success">
          <template #default="{ percentage }">
            <span class="percentage-value">{{ percentage }}%</span>
            <span class="percentage-label">内存使用率</span>
          </template>
        </el-progress>
        <div class="font-size-3">
          总大小:{{ statsData.memory.total }}/可用:{{ statsData.memory.available }}
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
          总大小:{{ statsData.memory.swap_total }}/可用:{{ statsData.memory.swap_free }}
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
          总大小:{{ statsData.disk.total }}/可用:{{ statsData.disk.free }}
        </div>
      </el-col>
    </el-row>
  </el-card>
</template>
<script setup>
import { inject, computed } from 'vue'
import { Icon } from '@iconify/vue'
const statsData = inject('statsData')

const cpuPercent = computed(() =>
  statsData.cpu.used_percent ? parseFloat(statsData.cpu.used_percent) : 0.0
)
const memoryPercent = computed(() =>
  statsData.memory.used_percent ? parseFloat(statsData.memory.used_percent) : 0.0
)
const swapmemPercent = computed(() =>
  statsData.memory.swap_used_percent ? parseFloat(statsData.memory.swap_used_percent) : 0.0
)
const diskPercent = computed(() =>
  statsData.disk.used_percent ? parseFloat(statsData.disk.used_percent) : 0.0
)
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
