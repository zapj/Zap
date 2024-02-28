<script setup>
import Header from '@/components/TopHeader.vue'
import LeftSidebar from '@/components/LeftSidebar.vue'
import FooterCopyright from '@/components/FooterCopyright.vue'
import { onMounted } from 'vue'
import apiRequest from '../httpclient/client'
import { useGlobalStore } from '../stores/global'
const globalStore = useGlobalStore()
onMounted(() => {
  apiRequest({
    url: '/v1/sync/user/settings'
  }).then((resp) => {
    globalStore.settings = resp.data
  })
})
</script>

<template>
  <!-- <el-container style="height: 100%"> -->
  <LeftSidebar></LeftSidebar>
  <el-container direction="vertical">
    <Header></Header>
    <el-scrollbar class="zap-content-scroll">
      <el-main style="margin: 0">
        <!-- <el-scrollbar class="zap-content-scroll"> -->
        <!-- <div class="content"> -->
        <RouterView></RouterView>
        <!-- </div> -->
        <!-- </el-scrollbar> -->
      </el-main>
      <FooterCopyright></FooterCopyright>
    </el-scrollbar>
  </el-container>
  <!-- </el-container> -->
</template>

<style scoped>
/* .content{
  padding-top: 84px;
  min-height: calc(100vh - 127px);
} */
</style>
