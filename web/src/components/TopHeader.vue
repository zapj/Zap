<template>
  <el-header>
    <el-link :underline="false">
      <el-icon :size="20" @click="ExpandMenu">
        <Expand />
      </el-icon>
    </el-link>
    <div class="right-buttons">
      <el-link :underline="false" @click="dialogVisible = true">Console</el-link>
      <el-link :underline="false" class="hidden-xs-only">Without Underline</el-link>
      <el-link :underline="false">
        <el-icon :size="20"><notification /></el-icon>
      </el-link>
      <el-dropdown>
        <el-link :underline="false">
          <!-- <el-avatar src="assets/img/avatar.png" fit="cover" /> -->
          <el-avatar :icon="UserFilled" />
          <span class="hidden-xs-only">Account</span>
          <el-icon class="el-icon--right"><arrow-down /></el-icon>
        </el-link>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>Account Settings</el-dropdown-item>
            <el-dropdown-item
              ><el-icon><UserFilled /></el-icon>个人资料</el-dropdown-item
            >
            <el-dropdown-item>修改密码</el-dropdown-item>
            <el-dropdown-item disabled>Action 4</el-dropdown-item>
            <el-dropdown-item divided @click="logout"><el-icon><SwitchButton /></el-icon>退出系统</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </el-header>

  <el-dialog
    v-model="dialogVisible"
    title="Tips"
    draggable
    :modal="false"
    append-to-body
    :close-on-click-modal="false"
    class="te-window"
  >
    <span>It's a draggable Dialog</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ArrowDown, Expand, Notification, SwitchButton, UserFilled } from '@element-plus/icons-vue'
import { ref, onMounted } from 'vue'
import router from '../router'
// const leftMenuStore = useLeftMenuStore()
const dialogVisible = ref(false)
onMounted(function () {
  const overlay = document.querySelector('.zap-overlay')
  overlay.addEventListener('click', function () {
    ExpandMenu()
  })
  window.addEventListener('resize', () => {
    if (document.body.clientWidth < 768) {
      overlay.style.visibility = 'hidden'
      document.querySelector('aside').style.marginLeft = '-200px'
    } else {
      document.querySelector('aside').style.marginLeft = '0'
    }
  })
})

function ExpandMenu() {
  const overlay = document.querySelector('.zap-overlay')
  const aside = document.querySelector('aside')

  if (aside.style.marginLeft === '0px' || aside.style.marginLeft === '') {
    aside.style.marginLeft = '-200px'
    overlay.style.visibility = 'hidden'
  } else {
    overlay.style.visibility = 'visible'
    aside.style.marginLeft = '0px'
  }
}

function logout() {
sessionStorage.removeItem('access_token')
router.replace({
  path: '/login',
  query: {
    redirect: router.currentRoute.fullPath
  }
})
// router.push("/login")
// apiRequest({
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
