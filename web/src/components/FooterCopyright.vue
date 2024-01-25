<template>
  <el-backtop :right="100" :bottom="100" />

  <el-footer class="text-gray:70 font-size-3 border-1 bg-light-color:700">
    <el-row justify="space-between">
      <el-col :span="12">@ Copyright zap.cn</el-col>
      <el-col :span="12" class="text-right">
        <el-text> 当前版本 {{ globalStore.settings.version }} <el-link @click="upgrade">更新</el-link> </el-text>
      </el-col>
    </el-row>
  </el-footer>
</template>
<script setup>
import { ElMessage } from 'element-plus'
import apiRequest from '../httpclient/client'
import { useGlobalStore } from '../stores/global';
const globalStore = useGlobalStore()
function upgrade(){
  apiRequest({
    url:"/v1/upgrade/check"
  }).then((resp)=>{
    ElMessage({message:"提交成功"})
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
