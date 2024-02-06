<template>
<el-row :gutter="20">
    <el-col :span="6" class="mb-2" v-for="(app , i) in appListRef" :key="i">
        <el-card :body-style="{ padding: '0px' }"  shadow="hover">
       
        <div style="padding: 14px">
          <span v-text="app.title"></span><el-tag class="tag ml-2" type="success" effect="dark" size="small" v-text="app.category"></el-tag>
          <div class="bottom mt-2">
            
            <time class="time font-size-3 me-2" >{{ app.version[0] }}</time>
            <template v-if="app.actions" >
              <el-button  class="button" v-for="actTitle,actionName in app.actions" @click.prevent="installApp(app.id,actionName)">{{ actTitle }} </el-button>
            </template>
            <template v-else>
              <el-button  class="button" @click.prevent="installApp(app.id,)">安装</el-button>
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
  apiRequest({
    url:'/v1/app/appstore/install',
    method:'post',
    data:{
      id:Id,
      action:actionName
    },
    dataType:'form'
  }).then((resp)=>{
    console.log(resp.data);
  })
}
</script>