<template>
<el-row :gutter="20">
    <el-col :span="24" class="mb-2">
        <el-card :body-style="{ padding: '0px' }"  shadow="hover">
          <el-table :data="tbDataRef" stripe style="width: 100%">

            <el-table-column prop="ID" label="ID" width="180" />
            <el-table-column prop="title" label="任务名称" width="180" />
            <el-table-column prop="start_time" label="开始时间" width="180" />
            <el-table-column prop="end_time" label="完成时间" width="180" />
            <el-table-column prop="Cmd" label="Cmd" />
            <el-table-column prop="Status" label="Status" />
          </el-table>
          
       </el-card>
    </el-col>

  
   
  </el-row>
</template>
<script setup>
import { onMounted } from 'vue';
import apiRequest from '../../httpclient/client';
const tbDataRef = ref([]) 

onMounted(()=>{
  getAppList()
})

const getAppList = ()=>{
  apiRequest({
    url:'/v1/task/appinstall/tasklist'
  }).then((resp)=>{
    tbDataRef.value = resp.data
  })
}


function installApp(Id,actionName){
  console.log(Id,actionName);
}
</script>