<template>
  <!-- <el-skeleton :animated="true">

    <template  #default>
      <SmallStatics></SmallStatics>
    </template>
  </el-skeleton> -->
  <Suspense>
    <template v-slot:default>
      <SmallStatics></SmallStatics>
    </template>

    <template v-slot:fallback>
      <el-skeleton :rows="4" animated translate="yes" />
    </template>
</Suspense>


  <el-row :gutter="20">
    <el-col :md="12" :sm="24">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>Card name</span>
            <el-button class="button" @click="testCon">Operation button</el-button>
          </div>
        </template>
        <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
      </el-card>
    </el-col>
    <el-col :md="12" :sm="24">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>Card name</span>
            <el-button class="button" text>Operation button</el-button>
          </div>
        </template>
        <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
      </el-card>
    </el-col>
  </el-row>
</template>


<script setup lang="ts">
import { ChatLineRound, Male,Monitor } from '@element-plus/icons-vue'
import serviceRequest from '../httpclient/client'
import {defineAsyncComponent} from 'vue'
function testCon() {
  serviceRequest({
    url: '/v1/loadavg'
  })
}


const SmallStatics = defineAsyncComponent(() => import('../components/dashboard/SmallStatistics.vue'))


</script>
<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}


</style>