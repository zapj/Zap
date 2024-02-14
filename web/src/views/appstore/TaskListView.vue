<template>
<el-row :gutter="20">
    <el-col :span="24" class="mb-2">
        <el-card :body-style="{ padding: '0px' }"  shadow="hover">
          <el-table :data="tbDataRef" stripe style="width: 100%">

            <el-table-column prop="id" label="ID" width="180" />
            <el-table-column prop="title" label="任务名称"  />
            <el-table-column prop="start_time" label="开始时间" width="180" :formatter="formatDate" />
            <el-table-column prop="end_time" label="完成时间" width="180" :formatter="formatDate" />
            <el-table-column prop="Status" label="状态" width="120" />
            <el-table-column prop="error" label="Error"  />
            <el-table-column label="操作" width="120">
              <template #default="scope">
                <el-button link type="primary" size="small" @click.prevent="removeTask(scope.$index,scope.row.id)" > 删除 </el-button>
                <el-button link type="primary" size="small" @click.prevent="cancelTask(scope.$index,scope.row.id)" > 取消 </el-button>
                <el-button link type="primary" size="small" @click.prevent="tailLog(scope.$index,scope.row.id)" > View Log </el-button>
              </template>
            </el-table-column>
          </el-table>
          <template #footer>
            <el-button type="primary" @click="genTask">安装</el-button>
            <el-button type="primary" @click="getAppList">刷新</el-button>
          </template>
       </el-card>
    </el-col>

  
   
  </el-row>

  <el-dialog
    v-model="dialogVisible"
    width="80%"
    title="LogViewer"
    >
    <Suspense>
      <template #default>
        <LogViewer :logfile="logfile"></LogViewer>
      </template>
      <template #fallback>
        <ElSkeleton :rows="4"></ElSkeleton>
      </template>
    </Suspense>
    
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">关闭</el-button>
      </div>
    </template>
  </el-dialog>



</template>
<script setup>
import { defineAsyncComponent, onMounted } from 'vue';
import apiRequest from '../../httpclient/client';
import { ElMessage, ElSkeleton } from 'element-plus';
import { formatDate } from '../../commons/commons.js'
const tbDataRef = ref([]) 
const dialogVisible = ref(false)
const logfile = ref("")
const LogViewer = defineAsyncComponent(()=>import('../../components/logviewer/LogViewer.vue'))
onMounted(()=>{
  getAppList()
})

const handleClose = () => {
  return true
 
}
const getAppList = ()=>{
  apiRequest({
    url:'/v1/task/appinstall/tasklist'
  }).then((resp)=>{
    tbDataRef.value = resp.data
  })
}

const removeTask = (index,id) => {
  apiRequest({
    url:'/v1/task/appinstall/removetask',
    data:{id:id},
    dataType:"form",
    method  :"post"
  }).then((resp)=>{
    tbDataRef.value.splice(index, 1)
    // tbDataRef.value = resp.data
    ElMessage({message:"删除成功",type:"success"})
  })
}

const cancelTask = (index,id) => {
  apiRequest({
    url:'/v1/task/appinstall/canceltask',
    data:{id:id},
    dataType:"form",
    method  :"post"
  }).then((resp)=>{
    getAppList()
    ElMessage({message:"取消成功",type:"success"})
  })
}

const tailLog = (index,id) => {
  logfile.value = "install_"+id+".log"
  dialogVisible.value = true
  
  // apiRequest({
  //   url:'/v1/task/appinstall/taillog',
  //   data:{id:id},
  //   dataType:"form",
  //   method  :"post"
  // }).then((resp)=>{
  //   ElMessage({message:"查看成功",type:"success"})
  // })
}

const genTask = ()=>{
  apiRequest({
    url:'/v1/task/appinstall/gentask',
    dataType:"form",
    method  :"post"
  }).then((resp)=>{
    ElMessage({message:"生成成功",type:"success"})
  })
 
}

function installApp(Id,actionName){
  console.log(Id,actionName);
}
</script>