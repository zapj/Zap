<template>
  <el-header>
    <el-link :underline="false">
      <el-icon :size="20" @click="ExpandMenu">
        <Expand />
      </el-icon>
    </el-link>
    <div class="right-buttons">
      <el-link :underline="false">Without Underline</el-link>
      <el-link :underline="false">
        <el-icon :size="20"><notification/></el-icon>
      </el-link>
      <el-dropdown>
        <el-link :underline="false">
          <el-avatar src="assets/img/avatar.png" fit="cover" />
          Account
          <el-icon class="el-icon--right"><arrow-down /></el-icon>
        </el-link>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>Account Settings</el-dropdown-item>
            <el-dropdown-item><el-icon><UserFilled /></el-icon>个人资料</el-dropdown-item>
            <el-dropdown-item>修改密码</el-dropdown-item>
            <el-dropdown-item disabled>Action 4</el-dropdown-item>
            <el-dropdown-item divided><el-icon><SwitchButton /></el-icon>退出系统</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>

    </div>
  </el-header>

</template>


<script setup>
import {ArrowDown, Expand, Notification, SwitchButton, UserFilled} from "@element-plus/icons-vue";
import {onMounted} from "vue";
// const leftMenuStore = useLeftMenuStore()

onMounted(function(){
  const overlay = document.querySelector('.zap-overlay');
  overlay.addEventListener('click',function(){
    ExpandMenu();
  })
  window.addEventListener('resize',()=>{
    if(document.body.clientWidth < 768){
      overlay.style.visibility = 'hidden';
      document.querySelector('aside').style.marginLeft = '-200px';
    }else{
      document.querySelector('aside').style.marginLeft = '0';
    }
  })
})

function ExpandMenu(){

  const overlay = document.querySelector('.zap-overlay');
  const aside = document.querySelector('aside');

  if(aside.style.marginLeft === '0px' || aside.style.marginLeft === ""){
    aside.style.marginLeft = '-200px';
    overlay.style.visibility = 'hidden';
  }else{
    overlay.style.visibility = 'visible';
    aside.style.marginLeft = '0px';
  }

}

</script>