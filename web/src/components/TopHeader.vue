<template>
  <el-header>
    <el-link :underline="false">
      <el-icon :size="20" @click="ExpandMenu">
        <Expand />
      </el-icon>
    </el-link>
    <div class="right-buttons mr-3">
      <!-- <el-link :underline="false" @click="dialogVisible = true">
        <Icon icon="material-symbols-light:terminal" :width="24" />
      </el-link>
      <el-link :underline="false">
        <Icon icon="tdesign:notification" :width="20" />
      </el-link> -->
      <el-dropdown>
        <el-link :underline="false">
          <span class="hidden-xs-only">{{ globalStore.settings.lname }}</span>
          <el-icon class="el-icon--right"><arrow-down /></el-icon>
        </el-link>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <el-icon><UserFilled /></el-icon>个人资料
            </el-dropdown-item>
            <el-dropdown-item><Edit />修改密码</el-dropdown-item>
            <el-dropdown-item divided @click="logout">
              <el-icon><SwitchButton /></el-icon>退出系统
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </el-header>
  <el-dialog
    modal-class="zeditor"
    v-model="dialogVisible"
    title="Tips"
    draggable
    :modal="false"
    :append-to-body="true"
    :close-on-click-modal="false"
    class="te-window"
  >
    <span>It's a draggable Dialog</span>
    <EditVue></EditVue>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import {
  ArrowDown,
  Expand,
  Notification,
  Edit,
  SwitchButton,
  UserFilled
} from '@element-plus/icons-vue'
import { ref, onMounted, defineAsyncComponent } from 'vue'
import { Icon } from '@iconify/vue'
import router from '../router'
import { useGlobalStore } from '../stores/global'
// const leftMenuStore = useLeftMenuStore()
const globalStore = useGlobalStore()

const EditVue = defineAsyncComponent(() => import('../components/editor/CodeEditor.vue'))
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
<style scoped>
.fixed-header {
  width: calc(100% - 200px);
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
}
</style>
