<template>
  <el-drawer title="应用设置" :size="drawerSize" :append-to-body="true" direction="rtl">
    <template #header>
      <span>安装{{ app.title }}</span>
    </template>
    <template #default>
      <el-descriptions title="应用信息" :column="1">
        <el-descriptions-item label="应用名称">{{ app.title }}</el-descriptions-item>
        <!-- <el-descriptions-item label="应用Name">{{ app.name }}</el-descriptions-item> -->

      
        <el-descriptions-item label="描述">{{ app.description }}</el-descriptions-item>
        <el-descriptions-item label="版本">
          <el-tag size="small"  class="me-2">{{ app.version }}</el-tag>
        </el-descriptions-item>
      </el-descriptions>
      <el-form ref="formRef" :model="appForm" label-width="auto">
         
        <el-form-item label="安装目录" >
          <el-input :value="app.install_dir" disabled />
        </el-form-item>
        <el-form-item label="配置文件" >
          <el-input v-model="appForm.config_file"  placeholder="etc: /etc/nginx/nginx.conf"   />
        </el-form-item>
        <el-form-item label="Pid 文件" >
          <el-input v-model="appForm.pid_file"  placeholder="etc: /run/nginx/nginx.pid"   />
        </el-form-item>
        <el-form-item label="Expose" >
          <el-input v-model="appForm.expose"  type="textarea" :autosize="{ minRows: 2 ,maxRows:5}"  />
        </el-form-item>
        <el-form-item label="Settings" >
          <el-input v-model="appForm.settings"  type="textarea" :autosize="{ minRows: 2 ,maxRows:10}"  />
        </el-form-item>
        <el-form-item>
            <el-button type="warning">取消</el-button>
            <el-button type="success" @click="saveAppSetting" >保存</el-button>
        </el-form-item>
      </el-form>
    </template>
  </el-drawer>
</template>
<script setup>
import { onMounted, onUnmounted, ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import apiRequest from '@/httpclient/client'

let drawerSize = '50%'
const appForm = reactive({
  config_file: '',
  pid_file:'',
  id: 0,
  expose:'',
  settings:''
})

const formRef = ref()
const app = ref({})

onMounted(() => {
  window.addEventListener('resize', resizeDrawer)
  console.log("mounted drawer");
})
onUnmounted(() => {
  window.removeEventListener('resize', resizeDrawer)
})

const resizeDrawer = () => {
  if (document.body.clientWidth < 1200 && document.body.clientWidth > 768) {
    drawerSize = '60%'
  } else if (document.body.clientWidth < 768) {
    drawerSize = '100%'
  } else {
    drawerSize = '50%'
  }
}

function saveAppSetting() {
  apiRequest({
    url: '/v1/app/appstore/save/settings',
    method: 'post',
    data: appForm,
    dataType: 'form'
  }).then((resp) => {
    if (resp.code == 0) {
      ElMessage({
        type: 'success',
        message: resp.msg
      })
    } else {
      ElMessage({
        type: 'error',
        message: resp.msg
      })
    }
  })
}

const setAppInfo = function (_app) {
  app.value = _app
  appForm.id = app.value.id
  appForm.expose = _app.expose
  appForm.config_file = _app.config_file
  appForm.pid_file = _app.pid_file
  appForm.settings = _app.settings
}

defineExpose({
  setAppInfo
})
</script>
<style scoped>
.el-drawer .el-drawer__header {
  height: 50px;
  line-height: 50px;
  font-size: 18px;
  margin-bottom: 0px;
}
</style>
