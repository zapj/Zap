<template>
  <el-container>
    <el-main class="login">
      <el-form :model="form" label-width="100px" :rules="formRules" ref="ruleFormRef" status-icon>
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>ZAP</span>
              <!-- <el-button class="button" text>Operation button</el-button> -->
            </div>
          </template>
          <div>
            <el-form-item label="用户名" prop="username">
              <el-input v-model="form.username" placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input
                type="password"
                v-model="form.password"
                placeholder="请输入密码"
                show-password
              />
            </el-form-item>
          </div>
          <template #footer>
            <el-form-item>
              <el-button type="primary" @click="onSubmit(ruleFormRef)">登陆</el-button>
            </el-form-item>
          </template>
        </el-card>
      </el-form>
    </el-main>
  </el-container>
</template>
<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import apiRequest from '../httpclient/client'
import { ElMessageBox, ElMessage } from 'element-plus'
const router = useRouter()
const form = reactive({
  username: '',
  password: ''
})
const ruleFormRef = ref(null)
const formRules = reactive({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
})

const onSubmit = (formEl) => {
  if (!formEl) return
  formEl.validate((valid, fields) => {
    if (valid) {
      login()
    }
  })
}

const login = () => {
  apiRequest({
    url: '/login',
    method: 'post',
    data: {
      username: form.username,
      password: form.password
    }
  }).then((data) => {
    if (data.code === 0 && data.access_token) {
      sessionStorage.setItem('access_token', data.access_token)
      ElMessage({ message: '登陆成功', type: 'success' })
      router.push('/dashboard')
    }
    // } else {
    //   ElMessageBox.alert(data.msg)
    // }
  })
}
</script>

<style scoped>
.card-header {
  text-align: center;
  color: rgb(3, 86, 3);
  font-weight: bold;
}
.el-container {
  height: 100vh;
  align-items: center;
  justify-content: center;
}

@media (min-width: 768px) {
  .box-card {
    top: 200px;
    width: 480px;
    margin: 0 auto;
  }
}
</style>
