<template>
  <div class="zap-overlay "></div>
    <el-aside width="200px" >
      <div class="font-bold color-green h-60px lh-60px text-center ">
        <span>ZAP</span>
      </div>

    <el-menu
        :default-active="route.path"
        background-color="#363636"
        text-color="#fff"
        router
    >
      <el-menu-item index="/">
        <el-icon>
          <HomeFilled/>
        </el-icon>
        <span>控制面板</span>
      </el-menu-item>
      <el-menu-item index="/table">
        <el-icon>
          <Grid/>
        </el-icon>
        <span>Table</span>
      </el-menu-item>
      <el-menu-item index="/about">
        <el-icon>
          <UserFilled/>
        </el-icon>
        <span>About</span>
      </el-menu-item>
      <el-menu-item index="/console">
        <el-icon>
          <UserFilled/>
        </el-icon>
        <span>Console</span>
      </el-menu-item>
      <el-sub-menu index="1">
        <template #title>
          <el-icon>
            <location/>
          </el-icon>
          <span>多级菜单</span>
        </template>
        <el-menu-item-group title="Group One">
          <el-menu-item index="1-1">item one</el-menu-item>
          <el-menu-item index="1-2">item two</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="Group Two">
          <el-menu-item index="1-3">item three</el-menu-item>
        </el-menu-item-group>
        <el-sub-menu index="1-4">
          <template #title>item four</template>
          <el-menu-item index="1-4-1">item one</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
      <el-menu-item index="2">
        <el-icon>
          <setting/>
        </el-icon>
        <span>Navigator Two</span>
      </el-menu-item>
      <el-menu-item index="3" disabled>
        <el-icon>
          <document/>
        </el-icon>
        <span>Navigator Three</span>
      </el-menu-item>
      <el-menu-item index="4">
        <el-icon>
          <setting/>
        </el-icon>
        <span>Navigator Four</span>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script setup>
import {useRoute} from "vue-router";
import {Document, Grid, HomeFilled, Location, Setting, UserFilled} from "@element-plus/icons-vue";
const route = useRoute();



</script>
