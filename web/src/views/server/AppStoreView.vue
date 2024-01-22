<template>
  <!-- <Suspense>
    <TopView></TopView>
     
      </Suspense> -->
        <button @click="change">aaa</button>
        <keep-alive>
        <component :is="currentItem"/>
        </keep-alive>
  </template>
  <script lang="ts" setup>
  import { defineAsyncComponent, onMounted, ref } from 'vue'
//   import type { TabsPaneContext } from 'element-plus'

//   import ace from 'ace-builds'
  
 
  const currentItem = ref('TopView')
  const TopView = defineAsyncComponent(()=>import("./TopView.vue"))
  const PSView = defineAsyncComponent(()=>import("./PSView.vue"))
  
 const change = () => {
  currentItem.value ='PSView'
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