<template>
 <el-table :data="appListRef" stripe style="width: 100%">
    <el-table-column prop="id" label="Id" width="180" />
    <el-table-column prop="name" label="Name" width="180" />
    <el-table-column prop="title" label="Title" width="180" />
    <el-table-column prop="version" label="版本" />
    <el-table-column prop="status" label="状态" />
    <el-table-column prop="install_date" label="安装时间" />
    <el-table-column fixed="right" label="操作" width="120">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="uninstallApp(scope.row)">卸载</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>
<script setup>
import { onMounted } from 'vue';
import apiRequest from '../../httpclient/client';
import { ElMessage } from 'element-plus';
const appListRef = ref([]) 

onMounted(()=>{
  getAppList()
})

const uninstallApp = (row)=>{
  apiRequest({
    url:'/v1/app/appstore/uninstall',
    method:'post',
    data:{
      id:row.id
    },
    dataType:'form'
  }).then((resp)=>{
    ElMessage.success('卸载成功')
    getAppList()
  })
}

const getAppList = ()=>{
  apiRequest({
    url:'/v1/app/appstore/already_install'
  }).then((resp)=>{
    appListRef.value = resp.data
  })
}
</script>