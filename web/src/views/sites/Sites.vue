<template>
  
  <el-card class="box-card filemanager" >
    <template #header>
      <div class="card-header font-size-3">
        网站管理

        <el-button
          type="primary"
          :icon="Plus"
          @click="addSite"
        />
      </div>
    </template>

    <el-table
      :data="tableData"
      min-height="200"
      style="width: 100%"
      :height="tableHeight"
      v-loading="pageState.tbLoading"
      fit 
    >
      <el-table-column type="selection" width="30" />
      <el-table-column prop="domain" label="域名" show-overflow-tooltip>
        <template #default="{ row, $index }">
          <el-button link @click.prevent="openWebSite($index)">
            <template #icon>
              <span v-if="row.is_file">
                <Icon icon="ic:round-insert-drive-file" width="28" height="28"
              /></span>
              <span v-else>
                <Icon icon="ic:baseline-folder" width="28" height="28" :inline="true"
              /></span>
            </template>

            {{ row.domain }}
          </el-button>
          
        </template>
      </el-table-column>
      <el-table-column prop="title" label="网站名称"  />
      <el-table-column prop="description" label="描述"  />
      <el-table-column prop="home" label="站点目录" width="180" />
      <el-table-column prop="status" label="网站状态" width="120" />
      <el-table-column fixed="right" label="操作" width="100">
        <template #default="scope">
          <el-button
            :icon="Edit"
            circle
            class="mr-2 hidden-md-and-down"
            title="设置网站"
            @click.prevent="openFile(scope.$index)"
          />
          <el-dropdown>
            <span class="el-dropdown-link">
              <Icon icon="ri:more-fill" width="24" height="24" />
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click.prevent="deleteRow(scope.$index)">删除</el-dropdown-item>
                <el-dropdown-item>停止</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty description="暂无数据">
          <el-button type="primary" @click="refreshData">刷新</el-button>
        </el-empty>
      </template>
    </el-table>

    <template #footer>
      <el-pagination
        background
        layout="total,prev, pager, next"
        :total="pageState.total"
        :page-size="pageState.pagesize"
        v-model:current-page="pageState.currentpage"
        :default-current-page="1"
        @change="handlePaginration"
      />
    </template>
  </el-card>
  <el-dialog v-model="dialogFormVisible" title="新建网站" width="500">
    <el-form :model="websiteForm">
      <el-form-item label="域名" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.domain" autocomplete="off" />
      </el-form-item>
      <el-form-item label="附加域名" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.other_domain" autocomplete="off" type="textarea" />
      </el-form-item>
      <el-form-item label="网站标题" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.title" autocomplete="off" />
      </el-form-item>
      
      <el-form-item label="站点目录" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.home" autocomplete="off" />
      </el-form-item>
      <el-form-item label="备注" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.description" autocomplete="off"  type="textarea" />
      </el-form-item>
    
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="createWebSite">
          添加
        </el-button>
      </div>
    </template>
  </el-dialog>


</template>

<script setup>
import { onMounted, onUnmounted, provide, ref ,reactive } from 'vue'
import { ElMessage } from 'element-plus'
import apiRequest from '../../httpclient/client'
import { Edit, Plus } from '@element-plus/icons-vue'
import { Icon } from '@iconify/vue'

const tableData = ref([])
const tableHeight = ref(300)
const dialogFormVisible = ref(false)
const formLabelWidth = '140px'
const websiteForm = reactive({
  domain: '',
  title: '',
  description: '',
  other_domain: '',
  home: '',
})


const pageState = reactive({
  total: 0,
  pagesize: 50,
  currentpage: 1,
  tbLoading: false
})

const addSite = () => {
  dialogFormVisible.value = true
}
const createWebSite = () => {
  apiRequest({
    url: '/v1/website/create',
    method: 'post',
    data: websiteForm
  }).then((res) => {
    if (res.code === 0) {
      ElMessage.success('添加成功')
      dialogFormVisible.value = false
      getWebSites({
        page: 1,
        pagesize: pageState.pagesize
      })
    }
  })
}

const openWebSite = (index) => {
  ElMessage.success('打开网站')
}

const refreshData = () => {
  getWebSites({
    page: pageState.currentpage,
    pagesize: pageState.pagesize
  })
}

const handlePaginration = (currentPage, pageSize) => {
  pageState.currentpage = currentPage
  pageState.pagesize = pageSize
  getWebSites({
    page: currentPage,
    pagesize: pageSize
  })
}

const resizeTable = () => {
  let tbHeight = document.body.clientHeight - 215 - 140
  if (tbHeight < 300) {
    tableHeight.value = 300
  } else {
    tableHeight.value = tbHeight
  }
}

onMounted(() => {
  
  getWebSites({
    page: 1,
    pagesize: pageState.pagesize
  })

  window.addEventListener('resize',resizeTable)
})
onUnmounted(()=>{
  window.removeEventListener('resize',resizeTable)
})

function getWebSites(options) {
  const defaultOptions = {
    page: 1,
    pagesize: 50
  }
  options = Object.assign(defaultOptions, options)
  pageState.tbLoading = true
  apiRequest({
    url: '/v1/website/list',
    method: 'get',
    params: {
      path: options.path,
      page: options.page,
      limit: options.pagesize
    },
    dataType: 'form',
    loading: false
  })
    .then((resp) => {
      tableData.value = []
      tableData.value.push(...resp.data)
      pageState.total = resp.total
      pageState.currentpage = resp.page
      resizeTable()
    })
    .catch((error) => {
      console.log(error)
    })
    .finally(() => (pageState.tbLoading = false))
}


</script>
