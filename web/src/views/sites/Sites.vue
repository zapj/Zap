<template>
  
  <el-card class="box-card filemanager" >
    <template #header>
      <div class="card-header font-size-3">
        网站管理

        <el-button
          type="primary"
          :icon="Plus"
          @click="websiteFormDrawer = true"
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
              <Icon icon="dashicons:admin-site" />
            </template>

            {{ row.domain }}
          </el-button>
          
        </template>
      </el-table-column>
      <el-table-column prop="title" label="网站名称"  />
      <el-table-column prop="description" label="描述"  />
      <el-table-column prop="home" label="站点目录"  />
      <el-table-column prop="status" label="网站状态" width="120" />
      <el-table-column fixed="right" label="操作" width="100">
        <template #default="scope">
          <el-button
            :icon="Setting"
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
                <el-dropdown-item @click.prevent="deleteSite(scope.$index)">删除</el-dropdown-item>
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
 
  <el-drawer v-model="websiteFormDrawer" direction="rtl" :size="websiteFormDrawerWidth" :show-close="false" @open="syncWebsiteConfig">
    <template #header="{ close, titleId, titleClass }">
      <h4 :id="titleId" :class="titleClass">添加网站</h4>
      <el-button @click="close">
        <el-icon class="el-icon--left"><CircleCloseFilled /></el-icon>
        关闭
      </el-button>
    </template>
    <template #default>
      <el-form :model="websiteForm" label-position="left">
      <el-form-item label="域名" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.domain" autocomplete="off" @change="updateWwwRoot"  />
      </el-form-item>
      <el-form-item label="附加域名" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.domain_names" autocomplete="off" type="textarea" />
      </el-form-item>
      <el-form-item label="网站名称" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.title" autocomplete="off" />
      </el-form-item>
      
      <el-form-item label="站点目录" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.www_root" autocomplete="off"  >
          <template #prepend>{{ website.config.home_dir }}/</template>
        </el-input>
      </el-form-item>

      
      <el-form-item label="中间件" :label-width="formLabelWidth">
        <el-select
            v-model="websiteForm.application"
            class="m-2"
            placeholder="Select"
            style="width: 240px"
          >

          
            <el-option
              v-for="item,i in website.config.applications"
              :key="i"
              :label="item.title"
              :value="i"
            >
            <span style="float: left">{{ item.title }}</span>
              <span style="float: right;color: var(--el-text-color-secondary);"> {{ item.expose_port }}</span>
          </el-option>
        </el-select>
      </el-form-item>


      <el-form-item label="备注" :label-width="formLabelWidth">
        <el-input v-model="websiteForm.description" autocomplete="off"  type="textarea" />
      </el-form-item>
    
    </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancelClick">取消</el-button>
        <el-button type="primary" @click="createWebSite">确认</el-button>
      </div>
    </template>
  </el-drawer>

</template>

<script setup>
import { onMounted, onUnmounted, provide, ref ,reactive } from 'vue'
import { ElMessage } from 'element-plus'
import apiRequest from '../../httpclient/client'
import { Setting, Plus,CircleCloseFilled  } from '@element-plus/icons-vue'
import { Icon } from '@iconify/vue'

const websiteFormDrawer = ref(false)
const websiteFormDrawerWidth = ref("50%")
const tableData = ref([])
const tableHeight = ref(400)
 
const formLabelWidth = '140px'
const websiteForm = reactive({
  domain: '',
  title: '',
  description: '',
  domain_names: '',
  www_root: '',
  run_directory: '',
  application:''
})

const website = reactive({
  config : {}
})


const pageState = reactive({
  total: 0,
  pagesize: 50,
  currentpage: 1,
  tbLoading: false
})
const updateWwwRoot = () => {
  websiteForm.www_root = websiteForm.domain.replaceAll('*.', '')
}
const syncWebsiteConfig = () => {
    apiRequest({
      url: '/v1/website/config',
      method: 'get'
    }).then((res) => {
      if (res.code === 0) {
        website.config = res.data
      }else{
        ElMessage.error(res.msg)
        websiteFormDrawer.value = false
        return
      }

      console.log( website.config );
  })
}

const createWebSite = () => {
  console.log(websiteForm);
  console.log(website.config.applications[websiteForm.application]);
  return
  apiRequest({
    url: '/v1/website/create',
    method: 'post',
    data: websiteForm
  }).then((res) => {
    if (res.code === 0) {
      ElMessage.success('添加成功')
      websiteFormDrawer.value = false
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

const resize = () => {
  //resize drawer width
  if(document.body.clientWidth < 1200 && document.body.clientWidth > 768){
    websiteFormDrawerWidth.value = "60%"
  }else if(document.body.clientWidth < 768){
    websiteFormDrawerWidth.value = "100%"
  }else{
    websiteFormDrawerWidth.value = "50%"
  }

  //resize table
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

  window.addEventListener('resize',resize)
})
onUnmounted(()=>{
  window.removeEventListener('resize',resize)
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
      resize()
    })
    .catch((error) => {
      console.log(error)
    })
    .finally(() => (pageState.tbLoading = false))
}


const deleteSite = (index) => {
  console.log(tableData.value[index])
  apiRequest({
    url: '/v1/website/delete',
    method: 'post',
    data: {
      id: tableData.value[index].ID
    },
    dataType: 'form',
  }).then((res) => {
    if (res.code === 0) {
      ElMessage.success('网站删除成功')
      getWebSites({
        page: 1,
        pagesize: pageState.pagesize
      })
    }
  }).catch((error) => {
    ElMessage.error('网站删除失败')
  })
}


</script>
