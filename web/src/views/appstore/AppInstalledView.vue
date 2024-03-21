<template>
  <el-table :data="appListRef" stripe style="width: 100%">
    <el-table-column prop="id" label="Id" width="50" />
    <el-table-column prop="name" label="Name" width="150" />
    <el-table-column prop="title" label="Title" />
    <el-table-column prop="version" label="版本" />
    <el-table-column prop="status" label="状态" width="100" />
    <el-table-column prop="app_status" label="运行状态" >
      <template #default="scope">
        <el-tag v-if="scope.row.app_status === 'running' && scope.row.pid_file" type="success">运行中</el-tag>
        <el-tag v-else-if="scope.row.app_status === 'stop' && scope.row.pid_file" type="info">已停止</el-tag>
        <template v-else>
        </template>
       </template>
    </el-table-column>
    <el-table-column prop="install_date" label="安装时间" />
    <el-table-column fixed="right" label="操作" width="120">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="appSettingPanel(scope.row)"><Icon icon="mdi:settings-outline" /> 设置</el-button>
        <el-button link type="primary" size="small" @click="confirmUninstall(scope.row)">卸载</el-button>
      </template>
    </el-table-column>
  </el-table>
  <AppSettingView v-model="appSettingVisable" ref="appSettingRef" />
</template>
<script setup>
import { onMounted } from 'vue'
import apiRequest from '../../httpclient/client'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Icon } from '@iconify/vue';
import AppSettingView from './AppSettingView.vue'
const appListRef = ref([])
const appSettingRef = ref()
const appSettingVisable = ref(false)
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
    url: '/v1/app/appstore/already_install',
    loading:false
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

const appSettingPanel = (row) => {
  appSettingVisable.value = true
  appSettingRef.value.setAppInfo(row)
}
</script>
