<template>
  <el-container>
    <el-main class="login">
      <el-form :model="form" label-width="120px">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>ZAP</span>
              <!-- <el-button class="button" text>Operation button</el-button> -->
            </div>
          </template>
          <div>
            <el-form-item label="用户名">
              <el-input v-model="form.name" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="form.password" />
            </el-form-item>
          </div>
          <template #footer>
            <el-form-item>
              <el-button type="primary" @click="onSubmit">登陆</el-button>
            </el-form-item>
          </template>
        </el-card>
      </el-form>
    </el-main>
  </el-container>
</template>
<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import serviceRequest from '../httpclient/client'
import { ElMessageBox, ElMessage } from 'element-plus'
const router = useRouter()
const form = reactive({
  name: '',
  password: ''
})

const onSubmit = () => {
  serviceRequest({
    url: '/login',
    method: 'post',
    data: {
      username: form.name,
      password: form.password
    }
  }).then((data) => {
    if (data.code === 0 && data.access_token) {
      sessionStorage.setItem('access_token', data.access_token)

      ElMessage({
        message: '登陆成功',
        type: 'success'
      })
      router.push('/dashboard')
    } else {
      ElMessageBox.alert(data.msg)
    }
  })
}
</script>

<style scoped>
.card-header {
  text-align: center;
  color: green;
  font-weight: bold;
}
.el-container {
  height: 100vh;
  align-items: center;
  justify-content: center;
}
.box-card {
  top: 200px;
  width: 480px;
  margin: 0 auto;
}
</style>
