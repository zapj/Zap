<template>
  <el-drawer title="安装确认" :size="drawerSize" :append-to-body="true" direction="rtl">
    <template #header>
      <span>安装{{ app.title }}</span>
    </template>
    <template #default>
      <el-descriptions title="应用信息" :column="1">
        <el-descriptions-item label="应用名称">{{ app.title }}</el-descriptions-item>
        <el-descriptions-item label="应用Name">{{ app.name }}</el-descriptions-item>

        <el-descriptions-item label="Tag">
          <el-tag size="small" class="me-2" v-for="(v, i) in app.tags" :key="i">{{ v }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="描述">{{ app.description }}</el-descriptions-item>
        <el-descriptions-item label="版本">
          <el-tag size="small" v-for="(v, i) in app.version" :key="i" class="me-2" >{{ v }}</el-tag>
        </el-descriptions-item>
      </el-descriptions>
      <el-form ref="formRef" :model="installAppForm" label-width="auto">
        <el-form-item label="选择安装的版本">
          <el-radio-group v-model="installAppForm.version">
            <el-radio border :value="v"  :label="v" 
            v-for="(v, i) in app.version" :disabled="app.already_installed_version.indexOf(v) !== -1" :key="i" >
            {{ v }}
            </el-radio>
          </el-radio-group>
        </el-form-item>
       
        <el-form-item>
          <template v-if="app.actions && app.allow_installation === 'yes'">
              <el-button
                type="success"
                v-for="(actTitle, actionName) in app.actions"
                @click.prevent="installApp(actionName)"
                >{{ actTitle }}
              </el-button>
            </template>
            <template v-else>
              <el-button
                type="success"
                v-if="app.allow_installation === 'yes'"
                @click.prevent="installApp('install')">安装</el-button>
            </template>
        </el-form-item>
      </el-form>
    </template>
  </el-drawer>
</template>
<script setup>
import {
  onMounted,
  onUnmounted,
  ref,
  reactive
} from 'vue'
import { ElMessage } from 'element-plus'
import apiRequest from '@/httpclient/client'

let drawerSize = '50%'
const installAppForm = reactive({
  version: '',
  id: 0
})
 
const formRef = ref()
const app = ref({})

 
onMounted(() => {
  window.addEventListener('resize', resizeDrawer)
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


function installApp(actionName) {
   apiRequest({
     url: '/v1/app/appstore/install',
     method: 'post',
     data: {
       id: app.value.app_id,
       action: actionName,
       version: installAppForm.version
     },
     dataType: 'form'
   }).then((resp) => {
     console.log(resp)
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
    installAppForm.id = app.value.id
    let version = ''
    for (let index = 0; index < app.value.version.length; index++) {
        if (app.value.already_installed_version.indexOf(app.value.version[index]) !== -1){
            version = app.value.version[index]
            break
        }
        
    }
    // installAppForm.version = app.value.version[0]
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
