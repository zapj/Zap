<template>
   <el-card class="box-card">

    <template #header>
      <div class="card-header">
        <el-row :gutter="20">
            <el-col :sm="20" :md="14" :lg="14" >
                <el-input  placeholder="文件路径">
                    <template #prepend>/home/admin</template>
                    <template #append><el-button><Icon icon="ic:outline-search" :width="20"  /></el-button></template>
                </el-input>
            </el-col>
          
           
        </el-row>
      </div>
    </template>

        
    <el-table :data="tableData" v-loading="loading" style="width: 100%" height="500">
    <el-table-column prop="name" label="文件名"   >
        <template #default="{ row, column, $index }">
            <el-button :type="success" link @click.prevent="openFile($index)">
                <template #icon>
                    <span v-if="row.is_file" >
                    <Icon icon="ic:round-insert-drive-file" :size="24"/></span>
                    <span v-else >
                    <Icon icon="ic:baseline-folder" :size="24"  :inline="treu" /></span>
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
        <el-button
          link
          type="primary"
          size="small"
          @click.prevent="deleteRow(scope.$index)"
        >
          Remove
        </el-button>
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
const now = new Date()

const tableData = ref([])

const deleteRow = (index) => {
  tableData.value.splice(index, 1)
}

onMounted(()=>{
    serviceRequest({
        url:"/v1/filemanager/list"
    }).then((data)=>{
        data.data.forEach(item => {
            tableData.value.push(item)
        });
    }).catch(()=>{

    })
})

const openFile = (index)=>{
    
    console.log(tableData.value[index]);
}
</script>