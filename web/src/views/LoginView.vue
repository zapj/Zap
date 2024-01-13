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
// import { useJWTStore } from '../stores/jwt';
import serviceRequest from '../httpclient/client';
// const jwtStore = useJWTStore();
const router = useRouter();
const form = reactive({
  name: '',
  password:'',
  
})

const onSubmit = () => {
    serviceRequest({
        url:"/login",
        method:"post",
        data:{
            username:form.name,
            password:form.password
        }
    }).then((data)=>{
        if(data.code === 0 && data.access_token){
            sessionStorage.setItem("access_token",data.access_token)
            router.push("/dashboard");
        }else{
            alert("login error")
        }
        console.log(data);
    })

    
    
}

</script>