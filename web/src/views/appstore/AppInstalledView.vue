<template>
  <el-table :data="appListRef" stripe style="width: 100%">
    <el-table-column prop="id" label="Id" width="180" />
    <el-table-column prop="name" label="Name" width="180" />
    <el-table-column prop="title" label="Title" width="180" />
    <el-table-column prop="version" label="版本" />
    <el-table-column prop="status" label="状态" />
    <el-table-column prop="app_status" label="运行状态" />
    <el-table-column prop="install_date" label="安装时间" />
    <el-table-column fixed="right" label="操作" width="120">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="confirmUninstall(scope.row)"
          >卸载</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>
<script setup>
import { onMounted } from 'vue'
import apiRequest from '../../httpclient/client'
import { ElMessage, ElMessageBox } from 'element-plus'
const appListRef = ref([])

onMounted(() => {
  getAppList()
})

const confirmUninstall = (row) => {
  ElMessageBox.confirm('确定要卸载' + row.name + '吗？', 'warning', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
    callback: (action) => {
      if (action === 'confirm') {
        uninstallApp(row)
      }
    }
  })
}

const uninstallApp = (row) => {
  apiRequest({
    url: '/v1/app/appstore/uninstall',
    method: 'post',
    data: {
      id: row.id
    },
    dataType: 'form'
  }).then((resp) => {
    ElMessage({
      type: resp.code === 0 ? 'success' : 'error',
      message: resp.msg
    })
    getAppList()
  })
}

const getAppList = () => {
  apiRequest({
    url: '/v1/app/appstore/already_install'
  }).then((resp) => {
    if (resp.code !== 0) {
      ElMessage({
        type: 'error',
        message: resp.msg
      })
      return
    }
    appListRef.value = resp.data
  })
}
</script>
