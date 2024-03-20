<template>
  <el-row class="mb-3 text-center" :gutter="10">
    <el-col :md="6" :sm="12">
      <el-card shadow="hover">
        <el-statistic :value="statsData.user_stats.websiteCount">
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
        <el-statistic :value="statsData.user_stats.databaseCount">
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
        <el-statistic :value="statsData.user_stats.cronjobCount">
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
            <span class="el-statistic__number">{{ statsData.host.uptime }}</span>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>
<script setup>
import { inject } from 'vue'
import { Icon } from '@iconify/vue'
const statsData = inject('statsData')
</script>
