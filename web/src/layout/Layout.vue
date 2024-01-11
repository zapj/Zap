<script setup>

import Header from "@/components/TopHeader.vue";
import LeftSidebar from "@/components/LeftSidebar.vue";
import FooterCopyright from "@/components/FooterCopyright.vue";
</script>

<template>



    <el-container>

      <LeftSidebar ></LeftSidebar>
      <el-container direction="vertical">
        <Header></Header>
        <el-scrollbar class="zap-content-scroll">
          <el-main>
            <RouterView></RouterView>
          </el-main>
          <FooterCopyright></FooterCopyright>
        </el-scrollbar>

      </el-container>

    </el-container>
</template>

<style scoped>

</style>
