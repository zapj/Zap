<template>
  <el-backtop :right="100" :bottom="100" />

  <el-footer class="text-gray:70 font-size-3 border-1 bg-light-color:700">
    <el-row justify="space-between">
      <el-col :span="12">@ Copyright zap.cn</el-col>
      <el-col :span="12" class="text-right">
        <el-text> 当前版本 <el-link>更新</el-link> </el-text>
      </el-col>
    </el-row>
  </el-footer>
</template>

<style scoped>
.el-footer {
  height: 60px;
  line-height: 60px;
  /* background-color: #fff; */
}
</style>
