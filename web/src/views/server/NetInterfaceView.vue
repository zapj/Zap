<template>
  <el-table
    :data="tableDataRef"
    border
    style="width: 100%"
    :height="tableHeight"
    v-loading="tbLoading"
  >
    <el-table-column prop="index" label="ID" width="50" />
    <el-table-column prop="name" label="设备名称" />
    <el-table-column prop="hardwareAddr" label="硬件地址" width="150" />
    <el-table-column prop="mtu" label="MTU" width="150" />
    <el-table-column prop="flags" label="Flags">
      <template #default="scope">
        <el-tag class="ml-2" type="success" v-for="flag in scope.row.flags">{{ flag }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="addrs" label="IP 地址">
      <template #default="scope">
        <el-tag class="ml-2" type="success" v-for="addr in scope.row.addrs">{{ addr.addr }}</el-tag>
      </template>
    </el-table-column>
  </el-table>
  <p class="text-center font-size-3">网络设备数: {{ totalRows }}</p>
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
    url: '/v1/server/netinterface_list',
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
