<template>
  <el-card class="box-card">
    <el-radio-group v-model="activeName" type="success" @change="handleSwitch">
      <el-radio-button label="应用市场" />
      <el-radio-button label="已安装" />
     
    </el-radio-group>
  </el-card>
  <div class="mt-3">
    <component :is="currentItem"></component>
  </div>
</template>
<script lang="ts" setup>
import { defineAsyncComponent, onMounted, ref, shallowRef } from 'vue'
const AppListView = defineAsyncComponent(() => import('../appstore/AppListView.vue'))
const AppInstalledView = defineAsyncComponent(() => import('../appstore/AppInstalledView.vue'))
 

const activeName = ref('应用市场')
const currentItem = shallowRef(AppListView)

const handleSwitch = () => {
  if (activeName.value === '应用市场') {
    currentItem.value = AppListView
  } else if (activeName.value === '已安装') {
    currentItem.value = AppInstalledView
  }
}

onMounted(() => {})
</script>
<style></style>
