<template>
<el-row :gutter="20">
    <el-col :span="6" class="mb-2" v-for="(app , i) in appListRef" :key="i">
        <el-card :body-style="{ padding: '0px' }"  shadow="hover">
       
        <div style="padding: 14px">
          <span v-text="app.Title"></span>
          <div class="bottom">
            <time class="time font-size-3">{{ app.Version }}</time>
            <template v-if="app.Actions" >
              <el-button  class="button" v-for="actTitle,actionName in app.Actions" @click.prevent="installApp(app.Id,actionName)">{{ actTitle }} </el-button>
            </template>
            <template v-else>
              <el-button  class="button">安装</el-button>
            </template>
          </div>
        </div>
      </el-card>
    </el-col>

  
   
  </el-row>
</template>
<script setup>
import { onMounted } from 'vue';
import apiRequest from '../../httpclient/client';
const appListRef = ref([]) 

onMounted(()=>{
  getAppList()
})

const getAppList = ()=>{
  apiRequest({
    url:'/v1/app/appstore'
  }).then((resp)=>{
    appListRef.value = resp.data
  })
}


function installApp(Id,actionName){
  console.log(Id,actionName);
}
</script>