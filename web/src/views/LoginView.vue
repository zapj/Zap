<template>
    <el-form :model="form" label-width="120px">
    <el-form-item label="Activity name">
      <el-input v-model="form.name" />
    </el-form-item>


    
    <el-form-item label="Activity name">
      <el-input v-model="form.password" />
    </el-form-item>
     
     
    <el-form-item>
      <el-button type="primary" @click="onSubmit">登陆</el-button>
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
</template>
<script setup>
import { reactive } from 'vue'
import { useRouter } from "vue-router";
const router = useRouter();
const form = reactive({
  name: '',
  password:'',
  
})

const onSubmit = () => {
    localStorage.setItem('zap-token',form.name)
    router.push("/dashboard");
}

</script>