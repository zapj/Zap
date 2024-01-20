<template>
  <div class="zap-overlay"></div>
  <el-aside width="200px">
    <div class="font-bold color-green h-60px lh-60px text-center">
      <span>ZAP</span>
    </div>

    <el-menu :default-active="route.path" background-color="#363636" text-color="#fff" router>
      <el-menu-item index="/">
        <Icon icon="ic:round-dashboard" style="font-size: 18px"  class="mr-1"  />
        <span>控制面板</span>


      </el-menu-item>
      <el-menu-item index="/table">
        <Icon icon="mingcute:web-line" style="font-size: 18px"  class="mr-1"  />
        <span>网站</span>
      </el-menu-item>
      <el-menu-item index="/filemanager">
        <Icon icon="clarity:directory-line" style="font-size: 18px" class="mr-1"  />
        <span>文件</span>
      </el-menu-item>
      <!-- <el-menu-item index="/about">
        <el-icon>
          <UserFilled />
        </el-icon>
        <span>About</span>
      </el-menu-item> -->
      <el-menu-item index="/console">
        <Icon icon="mdi:terminal" style="font-size: 18px" class="mr-1" :inline="true" />
        <span>终端</span>
      </el-menu-item>
      <!-- <el-menu-item index="/charts">
        <el-icon>
          <UserFilled />
        </el-icon>
        <span>Charts</span>
      </el-menu-item> -->
      <el-sub-menu index="1">
        <template #title>
          <Icon icon="mdi:server-outline" style="font-size: 18px" class="mr-1"  />
          <span>服务器</span>
        </template>
        <el-menu-item index="/system/monitor">系统监视器</el-menu-item>
        <el-menu-item index="/system/software">软件管理</el-menu-item>
        <el-menu-item index="/system/configuration">服务器配置</el-menu-item>
        <el-menu-item index="/system/security_center">安全中心</el-menu-item>
      </el-sub-menu>
      
      
      <el-menu-item @click="logout">
        <Icon icon="fe:logout" style="font-size: 18px" />
        <span>退出系统</span>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script setup>
import { useRoute,useRouter } from 'vue-router'
import { Icon } from '@iconify/vue';
const route = useRoute()
const router = useRouter()


function logout() {
  sessionStorage.removeItem('access_token')
  ElMessage({message: '退出成功',type: 'warning'})
  
  router.replace({
    path: '/login',
    query: {
      redirect: router.currentRoute.fullPath
    }
  })
  // router.push("/login")
  // serviceRequest({
  //     url:"/logout",
  //     method:"post",
  // }).then((data)=>{
  //     if(data.code === 0){
  //       router.push("/login")
  //     }else{
  //         alert("logout error")
  //     }
  //     console.log(data);
  // })
}

</script>
