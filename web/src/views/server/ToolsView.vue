<template>
  <el-card class="box-card">
  

    <el-radio-group v-model="activeName" type="success" @change="handleSwitch">
      <el-radio-button label="服务器信息"  />
      <el-radio-button label="Top 命令" />
      <el-radio-button label="进程管理" />
    </el-radio-group>

    
  </el-card>
  <el-card class="box-card mt-3">
    <component :is="currentItem" ></component>
  </el-card>
  </template>
  <script lang="ts" setup>
import { defineAsyncComponent, onMounted, ref ,shallowRef} from 'vue'
const ServerInfoView = defineAsyncComponent(()=>import("./ServerInfoView.vue"))
const TopView = defineAsyncComponent(()=>import("./TopView.vue"))
const PSView = defineAsyncComponent(()=>import("./PSView.vue"))
 
const activeName = ref('服务器信息')
const currentItem = shallowRef(ServerInfoView)
 
const handleSwitch = () => {
  if(activeName.value === '服务器信息'){
    currentItem.value = ServerInfoView
  }else if(activeName.value === 'Top 命令'){
    currentItem.value = TopView
  }else if(activeName.value === '进程管理'){
    currentItem.value = PSView
  }
}

  onMounted(()=>{
  
    
  })
  </script>
  <style>
  .demo-tabs > .el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
  }
  </style>