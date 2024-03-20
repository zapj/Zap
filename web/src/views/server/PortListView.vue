<template>
  <el-table :data="tableData" border style="width: 100%" v-loading="tbLoadding">
    <el-table-column prop="protocol" label="协议" width="180" />
    <el-table-column prop="local_addr" label="Local Addr" />
    <el-table-column prop="remote_addr" label="Remote Addr" />
    <el-table-column prop="status" label="Status" width="180" />
    <el-table-column prop="proc_name" label="进程名称" />
    <el-table-column prop="pid" label="进程PID" width="180" />
  </el-table>
</template>
<script setup>
// /server/port_list
import { ref, onMounted } from 'vue'
import apiRequest from '@/httpclient/client'
const tableData = ref([])
const tbLoadding = ref(false)
onMounted(() => {
  loadData()
})

const loadData = () => {
  tbLoadding.value = true
  apiRequest({
    url: '/v1/server/port_list',
    method: 'get',
    loading: false
  })
    .then((res) => {
      tableData.value = res.data
    })
    .finally(() => {
      tbLoadding.value = false
    })
}
</script>
