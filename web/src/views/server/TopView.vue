<template>
  <pre class="border-1 border-rd color-gray-500 font-size-3">{{ header }}</pre>

  <el-table :data="tableDataRef" style="width: 100%" :height="tableHeight">
    <el-table-column prop="0" label="PID" width="180" sortable />
    <el-table-column prop="1" label="USER" width="180" sortable />
    <el-table-column prop="2" label="PR" />
    <el-table-column prop="3" label="NI" />
    <el-table-column prop="4" label="VIRT" />
    <el-table-column prop="5" label="RES" sortable />
    <el-table-column prop="6" label="SHR" sortable />
    <el-table-column prop="7" label="S" />
    <el-table-column prop="8" label="%CPU" sortable />
    <el-table-column prop="9" label="%MEM" sortable />
    <el-table-column prop="10" label="TIME+" />
    <el-table-column prop="11" label="COMMAND" />
  </el-table>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import apiRequest from '../../httpclient/client'
const header = ref('')
const tableDataRef = ref([])
const tableHeight = ref(300)

const resizeTable = () => {
  let tbHeight = document.body.clientHeight - 260 - 140
  if (tbHeight < 300) {
    tableHeight.value = 300
  } else {
    tableHeight.value = tbHeight
  }
  console.log(tbHeight)
}
onMounted(() => {
  getServerTop()
  window.onresize = resizeTable
})

const getServerTop = () => {
  apiRequest({
    url: '/v1/server/top_info',
    loading: false
  }).then((resp) => {
    header.value = resp.data
    tableDataRef.value = resp.rows
    resizeTable()
  })
}
</script>
