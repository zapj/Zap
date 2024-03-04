<template>
  <el-backtop :right="100" :bottom="100" />

  <el-footer class="text-gray border-1" style="font-size: 12px">
    <el-row justify="space-between">
      <!-- <el-col :span="12"></el-col> -->
      <el-col :span="24" class="text-right">
        @ Copyright zap.cn
        <span>
          当前版本 {{ globalStore.settings.version }}({{ globalStore.settings.build_date }})
          <el-link underline style="font-size: 12px" @click="upgrade">更新</el-link>
        </span>
      </el-col>
    </el-row>
  </el-footer>
</template>
<script setup>
import { ElMessage } from 'element-plus'
import apiRequest from '../httpclient/client'
import { useGlobalStore } from '../stores/global'
const globalStore = useGlobalStore()
function upgrade() {
  apiRequest({
    url: '/v1/upgrade/check'
  }).then((resp) => {
    ElMessage({ message: resp.msg, type: resp.code ===0 ? 'success' : 'error' })
  })
}
</script>
<style scoped>
.el-footer {
  height: 60px;
  line-height: 60px;
  /* background-color: #fff; */
}
</style>
