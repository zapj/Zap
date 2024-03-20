<template>
  <el-row :gutter="20">
    <el-col
      :xs="24"
      :sm="24"
      :md="8"
      :lg="8"
      :xl="6"
      class="mb-2"
      v-for="(app, i) in appListRef"
      :key="i"
    >
      <el-card :body-style="{ padding: '0px' }" shadow="hover">
        <div style="padding: 14px">
          <div>
            <h3 v-text="app.title" class="font-size-4"></h3>
            <p v-text="app.description" class="font-size-3"></p>
          </div>
          <div class="bottom mt-2">
            <el-tag
              class="tag"
              type="info"
              effect="dark"
              size="small"
              v-text="app.category"
            ></el-tag>

            <time class="time font-size-3 ms-2 me-2">latest:{{ app.version[0] }}</time>
            <template v-if="app.actions && app.allow_installation === 'yes'">
              <el-button
                type="success"
                v-for="(actTitle, actionName) in app.actions"
                @click.prevent="confirmInstallApp(i, actionName)"
                >{{ actTitle }}
              </el-button>
            </template>
            <template v-else>
              <el-button
                type="success"
                v-if="app.allow_installation === 'yes'"
                @click.prevent="confirmInstallApp(i, 'install')"
                >安装</el-button
              >
            </template>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>

  <InstallConfirm ref="installConfirmRef" v-model="installConfirmVisible" />
</template>
<script setup>
import { onMounted } from 'vue'
import apiRequest from '../../httpclient/client'
import { ElMessage } from 'element-plus'
import InstallConfirm from '@/views/appstore/InstallConfirm.vue'
const appListRef = ref()
const confirmData = ref()
const installConfirmRef = ref()
const installConfirmVisible = ref(false)
// const  InstallConfirm = defineAsyncComponent(()=> import('@/views/appstore/InstallConfirm.vue'))
onMounted(() => {
  getAppList()
})

const getAppList = () => {
  apiRequest({
    url: '/v1/app/appstore'
  }).then((resp) => {
    appListRef.value = resp.data
  })
}
function confirmInstallApp(i, actionName) {
  confirmData.value = appListRef.value[i]
  confirmData.value.action = actionName
  installConfirmRef.value.setAppInfo(confirmData.value)
  installConfirmVisible.value = true
}
</script>
