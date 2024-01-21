<template>
   <el-card class="box-card">

    <template #header>
      <div class="card-header">
        <el-row :gutter="20">
            <el-col :sm="20" :md="14" :lg="14" >
                <el-input v-model="inputFileDynPath" >
                    <template #append><el-button><Icon icon="ic:outline-search" :width="20"  /></el-button></template>
                </el-input>
                <ul style="position: absolute;top: 8px;" class="fm-breadcrumb">
                  <li><el-link :underline="false">
                  <Icon icon="ic:outline-home" :size="24" />
                  </el-link>
                  </li>
                  <li v-for="item in breadcrumbPaths" ><el-link :underline="false">{{ item }}</el-link></li>
                  
                </ul>
            </el-col>
          
           
        </el-row>
      </div>
    </template>

        
    <el-table :data="tableData" v-loading="loading" style="width: 100%" height="500">
        <el-table-column prop="name" label="文件名"   >
            <template #default="{ row, $index }">
                <el-button type="success" link @click.prevent="openFile($index)">
                    <template #icon>
                        <span v-if="row.is_file" >
                        <Icon icon="ic:round-insert-drive-file" :size="24"/></span>
                        <span v-else >
                        <Icon icon="ic:baseline-folder" :size="24"  :inline="true" /></span>
                    </template>
                    
                    {{ row.name }}
                    </el-button>
                
            
          </template>
        </el-table-column>
        <el-table-column prop="perm" label="权限"  />
        <el-table-column prop="uid" label="UID"  />
        <el-table-column prop="gid" label="GID"  />
        <el-table-column prop="mod_time" label="修改时间"  />
        <el-table-column prop="filesize" label="文件大小"  />
        <el-table-column fixed="right" label="操作" width="120">
          <template #default="scope">
            <el-button link type="primary" size="small" @click.prevent="deleteRow(scope.$index)" > Remove </el-button>
          </template>
        </el-table-column>
      </el-table>


      <template #footer></template>
</el-card>
</template>
<script setup>
import { Icon } from '@iconify/vue';
import { onMounted, ref } from 'vue'
import serviceRequest from '../../httpclient/client';
import {useGlobalStore} from '../../stores/global'
const now = new Date()
const globalStore = useGlobalStore()
const tableData = ref([])
const inputFileDynPath = ref('')
const deleteRow = (index) => {
  tableData.value.splice(index, 1)
}
const breadcrumbPaths = ref([])
onMounted(()=>{
  if (globalStore.lastFilePath === ""){
    globalStore.lastFilePath = "/"
  }
  getFiles(globalStore.lastFilePath)
  breadcrumbPaths.value.push('/','home') 
})

const openFile = (index)=>{
    const row = tableData.value[index]
    if (!row) {
      return
    }
    getFiles(row.name)

    
}

function getFiles(path){
  serviceRequest({
        url:"/v1/filemanager/list",
        method:'post',
        data:{
          path:  (globalStore.lastFilePath === "/") ? globalStore.lastFilePath + path : globalStore.lastFilePath + "/" + path
        },
        dataType:'form'
    }).then((data)=>{
        tableData.value = []
        data.data.forEach(item => {
            tableData.value.push(item)
        });
    }).catch(()=>{

    })
}
</script>
<style scoped>
.fm-breadcrumb{
  position: absolute;
    top: 8px;
    margin: -5px 0 0 0;
    padding-left: 10px;
}
.fm-breadcrumb li {
  display: inline-block;
    padding-left: 7px;
}
</style>