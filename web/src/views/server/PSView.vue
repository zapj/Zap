<template>
  <el-table
    :data="tableDataRef"
    border
    style="width: 100%"
    :default-sort="{ prop: 'pid', order: 'descending' }"
    :height="tableHeight"
    v-loading="tbLoading"
  >
    <el-table-column prop="pid" label="PID" width="120" sortable fixed />
    <el-table-column prop="name" label="进程名称" sortable />
    <el-table-column prop="username" label="用户" width="150" sortable />
    <el-table-column prop="cpu_percent" label="CPU%" width="100" sortable>
      <template #default="scope"> {{ scope.row.cpu_percent }} % </template>
    </el-table-column>
    <el-table-column prop="mem_percent" label="MEM%" width="100" sortable>
      <template #default="scope"> {{ scope.row.mem_percent }} % </template>
    </el-table-column>

    <el-table-column prop="create_time" label="启动时间" width="180" />
    <el-table-column prop="cmdline" label="COMMAND" />
  </el-table>
  <p class="text-center font-size-3">总进程数: {{ totalRows }}</p>
</template>
<script setup>
import { computed, onMounted, ref } from 'vue'
import apiRequest from '../../httpclient/client'

const tableDataRef = ref([])
const tbLoading = ref(false)
const tableHeight = ref(300)
const resizeTable = () => {
  let tbHeight = document.body.clientHeight - 215 - 140
  if (tbHeight < 300) {
    tableHeight.value = 300
  } else {
    tableHeight.value = tbHeight
  }
  console.log(tbHeight)
}
const totalRows = computed(() => {
  return tableDataRef.value.length
})

onMounted(() => {
  getServerProcessList()
  window.onresize = resizeTable
})

const getServerProcessList = () => {
  tbLoading.value = true
  apiRequest({
    url: '/v1/server/processlist',
    loading: false
  })
    .then((resp) => {
      tableDataRef.value = resp.data
      resizeTable()
      tbLoading.value = false
    })
    .finally(() => (tbLoading.value = false))
}
</script>
