<template>
  <table class="font-size-3 serverinfo">
    <thead>
      <tr>
        <td></td>
        <td></td>
        <td></td>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td width="150px" :colspan="colspan">服务器信息</td>
      </tr>

      <tr>
        <td>操作系统版本</td>
        <td>
          {{ state.serverInfo.OS }} / {{ state.serverInfo.Platform }}-{{
            state.serverInfo.PlatformVersion
          }}
        </td>
        <td></td>
      </tr>
      <tr>
        <td>系统内核</td>
        <td>{{ state.serverInfo.KernelVersion }} {{ state.serverInfo.KernelArch }}</td>
        <td></td>
      </tr>
      <tr>
        <td>主机名</td>
        <td>{{ state.serverInfo.Hostname }}</td>
        <td></td>
      </tr>
      <tr>
        <td>服务器启动时间</td>
        <td>{{ state.serverInfo.BootTime }}</td>
        <td></td>
      </tr>
      <tr>
        <td>运行时间</td>
        <td>{{ state.serverInfo.uptime }}</td>
        <td></td>
      </tr>
      <tr>
        <td>系统负载</td>
        <td>
          <el-tag class="ml-2" type="success">{{ state.serverInfo.load1 }}</el-tag>
          <el-tag class="ml-2" type="success">{{ state.serverInfo.load5 }}</el-tag>
          <el-tag class="ml-2" type="success">{{ state.serverInfo.load15 }}</el-tag>
          <el-progress :percentage="state.serverInfo.system_load1" />
        </td>
        <td></td>
      </tr>

      <tr>
        <td :colspan="colspan">&nbsp;</td>
      </tr>
      <tr>
        <td :colspan="colspan">CPU</td>
      </tr>
      <tr>
        <td>CPU 名称</td>
        <td>{{ state.serverInfo.cpu_name }}</td>
        <td></td>
      </tr>
      <tr>
        <td>CPU 频率</td>
        <td>{{ state.serverInfo.cpu_ghz }}GHz</td>
        <td></td>
      </tr>
      <tr>
        <td>CPU 核心数</td>
        <td>{{ state.serverInfo.cpu_cores }}核心</td>
        <td></td>
      </tr>
      <tr>
        <td :colspan="colspan">&nbsp;</td>
      </tr>
      <tr>
        <td :colspan="colspan">内存</td>
      </tr>
      <tr>
        <td>内存总大小</td>
        <td>{{ state.serverInfo.memory_total }}</td>
        <td></td>
      </tr>
      <tr>
        <td>已用内存</td>
        <td>
          {{ state.serverInfo.memory_used }}
          <el-progress :percentage="state.serverInfo.memory_used_percent" />
        </td>
        <td></td>
      </tr>
      <tr>
        <td>剩余可以内存</td>
        <td>{{ state.serverInfo.memory_available }}</td>
        <td></td>
      </tr>
      <tr>
        <td :colspan="colspan">&nbsp;</td>
      </tr>
      <tr>
        <td :colspan="colspan">Swap内存</td>
      </tr>
      <tr>
        <td>Swap内存总大小</td>
        <td>{{ state.serverInfo.swapmem_total }}</td>
        <td></td>
      </tr>
      <tr>
        <td>已用内存</td>
        <td>
          {{ state.serverInfo.swapmem_used }}
          <el-progress :percentage="state.serverInfo.swapmem_used_percent" />
        </td>
        <td></td>
      </tr>
      <tr>
        <td>剩余可以内存</td>
        <td>{{ state.serverInfo.swapmem_free }}</td>
        <td></td>
      </tr>

      <tr>
        <td :colspan="colspan">&nbsp;</td>
      </tr>
      <tr>
        <td :colspan="colspan">硬盘信息</td>
      </tr>
      <tr>
        <td :colspan="colspan">
          <table>
            <tr v-for="(disk, index) in state.serverInfo.disk" :key="index">
              <td>{{ disk.path }}</td>
              <td>{{ disk.fstype }}</td>
              <td>{{ formatBytes(disk.total) }}</td>
              <td>{{ formatBytes(disk.free) }}</td>
              <td>{{ formatBytes(disk.used) }}</td>
              <td>{{ parseFloat(disk.usedPercent).toFixed(2) }}%</td>
            </tr>
          </table>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
import { onMounted, ref, reactive } from 'vue'
import apiRequest from '../../httpclient/client'
import { formatBytes } from '../../commons/commons'
const state = reactive({
  serverInfo: {}
})
const colspan = 3
onMounted(() => {
  getServerInfo()
})

const getServerInfo = () => {
  apiRequest({
    url: '/v1/server/info'
  }).then((resp) => {
    state.serverInfo = resp.data
    console.log(state.serverInfo)
  })
}
</script>
<style scoped>
.serverinfo {
  line-height: 30px;
}
.serverinfo tbody > tr td {
  color: #6c6c6c;
}

.serverinfo tbody > tr:hover {
  background-color: #f5f5f5;
}
</style>
