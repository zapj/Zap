<template>
  <el-card class="box-card filemanager">
    <template #header>
      <div class="card-header font-size-3">
        网站管理

        <el-button type="success" :icon="Plus" @click="openCreateWebSite(websiteFormDrawerRef)" />
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
          <el-button link @click.prevent="openWebSite(row.ID)">
            <template #icon>
              <Icon icon="dashicons:admin-site" />
            </template>

            {{ row.domain }}
          </el-button>
        </template>
      </el-table-column>
      <el-table-column prop="title" label="网站名称" />
      <el-table-column prop="www_root" label="站点目录" />
      <el-table-column prop="status" label="网站状态" width="120" />
      <el-table-column fixed="right" label="操作" width="100">
        <template #default="scope">
          <el-button
            :icon="Setting"
            circle
            class="mr-2 hidden-md-and-down"
            title="设置网站"
            @click.prevent="websiteSettings(scope.row.ID)"
          />
          <el-dropdown>
            <span class="el-dropdown-link">
              <Icon icon="ri:more-fill" width="24" height="24" />
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click.prevent="deleteSite(scope.$index)">删除</el-dropdown-item>
                <el-dropdown-item @click.prevent="stopSite(scope.$index)">停止</el-dropdown-item>
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

  <el-drawer
    v-model="websiteFormDrawer"
    direction="rtl"
    :size="websiteFormDrawerWidth"
    :show-close="false"
  >
    <template #header="{ close, titleId, titleClass }">
      <h4 :id="titleId" :class="titleClass">{{ drawerTitle }}</h4>
      <el-button @click="close">
        <el-icon class="el-icon--left"><CircleCloseFilled /></el-icon>
        关闭
      </el-button>
    </template>
    <template #default>
      <el-form :model="websiteForm" label-position="left" ref="websiteFormDrawerRef" :rules="rules">
        <el-form-item label="域名" :label-width="formLabelWidth" prop="domain">
          <el-input v-model="websiteForm.domain" autocomplete="off" @change="updateWwwRoot" />
        </el-form-item>
        <el-form-item label="附加域名" :label-width="formLabelWidth" prop="domain_names">
          <el-input v-model="websiteForm.domain_names" autocomplete="off" type="textarea" />
        </el-form-item>
        <el-form-item label="网站名称" :label-width="formLabelWidth" prop="title">
          <el-input v-model="websiteForm.title" autocomplete="off" />
        </el-form-item>

        <el-form-item label="站点目录" :label-width="formLabelWidth" prop="www_root">
          <el-input v-model="websiteForm.www_root" autocomplete="off">
            <template #prepend>{{ website.config.home_dir }}/</template>
          </el-input>
        </el-form-item>

        <el-form-item label="应用类型" :label-width="formLabelWidth" prop="application">
          <el-select
            v-model="websiteForm.application"
            class="m-2"
            placeholder="中间件服务器"
            style="width: 240px"
          >
            <el-option
              v-for="(item, i) in website.config.applications"
              :key="item.id"
              :label="item.title"
              :value="item.id"
            >
              <span style="float: left">{{ item.title }}</span>
              <span style="float: right; color: var(--el-text-color-secondary)">
                {{ item.expose_port }}</span
              >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="备注" :label-width="formLabelWidth" prop="description">
          <el-input v-model="websiteForm.description" autocomplete="off" type="textarea" />
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel(websiteFormDrawerRef)">取消</el-button>
        <el-button type="primary" @click="saveWebsite(websiteFormDrawerRef)">确认</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup>
import { onMounted, onUnmounted, provide, ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import apiRequest from '../../httpclient/client'
import { Setting, Plus, CircleCloseFilled } from '@element-plus/icons-vue'
import { Icon } from '@iconify/vue'

let wOp = false
//table
const tableData = ref([])
const tableHeight = ref(400)

// drawer
const websiteFormDrawerRef = ref()
const websiteFormDrawer = ref(false)
const websiteFormDrawerWidth = ref('50%')
//form
const formLabelWidth = '140px'
const drawerTitle = ref('添加网站')
const websiteForm = reactive({
  domain: '',
  title: '',
  description: '',
  domain_names: '',
  www_root: '',
  run_directory: '',
  application: 0,
  website_id: 0
})

const website = reactive({
  config: {}
})

const pageState = reactive({
  total: 0,
  pagesize: 50,
  currentpage: 1,
  tbLoading: false
})

const rules = {
  domain: [
    {
      required: true,
      trigger: 'blur',
      validator: (rule, value, callback) => {
        if (value == '' && value === undefined) {
          callback(new Error('域名不能为空'))
        }
        apiRequest({
          url: 'v1/website/check_domain',
          dataType: 'form',
          data: { domain: value,website_id:websiteForm.website_id },
          method: 'POST',
          loading:false
        })
          .then((res) => {
            if (res.code === 0) {
              callback()
            } else {
              callback(new Error(res.msg))
            }
          })
          .catch((error) => {
            callback(new Error('域名验证失败'))
          })
      }
    },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  title: [{ required: true, message: '请输入网站名称', trigger: 'blur' }],
  www_root: [{ required: true, message: '请输入站点目录', trigger: 'blur' }],
  application: [{ required: true, message: '请选择应用类型', trigger: 'change' }]
}


onMounted(() => {
  getWebSites({
    page: 1,
    pagesize: pageState.pagesize
  })

  window.addEventListener('resize', resize)
})
onUnmounted(() => {
  window.removeEventListener('resize', resize)
})

const updateWwwRoot = () => {
  if (websiteForm.www_root === '') {
    websiteForm.www_root = websiteForm.domain.replaceAll('*.', '')
  }
  if (websiteForm.title === '') {
    websiteForm.title = websiteForm.domain.replaceAll('www.', '')
  }
}
const syncWebsiteConfig = () => {
  return apiRequest({
    url: '/v1/website/config',
    method: 'get'
  }).then((res) => {
    if (res.code === 0) {
      website.config = res.data
    } else {
      ElMessage.error(res.msg)
      websiteFormDrawer.value = false
      return
    }
  })
}

const stopSite = () => {
  apiRequest({
    url: '/v1/website/stop',
    method: 'post'
  }).then((res) => {
    if (res.code === 0) {
      ElMessage.success('停止成功')
    } else {
      ElMessage.error(res.msg)
    }
  })
}

const websiteSettings = (website_id) => {
  wOp = true
  websiteFormDrawer.value = true
  drawerTitle.value = '网站设置'
  websiteForm.website_id = website_id
  apiRequest({
    url: '/v1/website/settings',
    method: 'get',
    params: { id: website_id }
  }).then((res) => {
    if (res.code === 0) {
      website.config = res.data.config
      websiteForm.domain = res.data.website.domain
      websiteForm.domain_names = res.data.website.domain_names
      websiteForm.www_root = res.data.website.www_root
      websiteForm.run_directory = res.data.website.run_directory
      websiteForm.application = res.data.website.application_id
      websiteForm.title = res.data.website.title
      websiteForm.description = res.data.website.description
      websiteForm.website_id = res.data.website.ID
      if (websiteForm.www_root === '') {
        websiteForm.www_root = websiteForm.domain.replaceAll('*.', '')
      } else {
        websiteForm.www_root = res.data.website.www_root.replace(website.config.home_dir + '/', '')
      }
      if (websiteForm.title === '') {
        websiteForm.title = websiteForm.domain.replaceAll('www.', '')
      }
    } else {
      ElMessage.error(res.msg)
      websiteFormDrawer.value = false
      return
    }
  })
}

//reset fields
const cancel = (formRef) => {
  if (!formRef) return
  formRef.resetFields()
  websiteFormDrawer.value = false
}

// open drawer
const openCreateWebSite = (formRef) => {
  syncWebsiteConfig().finally(() => {
    websiteFormDrawer.value = true
  })
  wOp = false //创建网站
  // websiteFormDrawer.value = true
  websiteForm.website_id = 0
  if (!formRef) return
  formRef.resetFields()
 
}

// save website
const saveWebsite = (formRef) => {
  if (!formRef) return

  formRef.validate((valid) => {
    if (!valid) {
      return false
    }
  })

  let wform = Object.assign({}, websiteForm)
  if (website.config.applications.length === 0) {
    ElMessage.error('请先添加应用')
    return
  }

  let app = null
  website.config.applications.forEach((item, index) => {
    if (item.id === websiteForm.application) {
      app = item
    }
  })
  if (!app) {
    ElMessage.error('应用不存在')
    return
  }

  wform.expose_port = app.expose_port
  wform.expose_proto = app.expose_proto
  wform.application = app.id

  // console.log(wform);
  // return

  apiRequest({
    url: wOp ? '/v1/website/update' : '/v1/website/create',
    method: 'post',
    data: wform
  })
    .then((res) => {
      if (res.code === 0) {
        ElMessage.success('添加成功')
        websiteFormDrawer.value = false
        getWebSites({
          page: 1,
          pagesize: pageState.pagesize
        })
      }
    })
    .finally(() => {
      formRef.resetFields()
    })
}

const openWebSite = (website_id) => {
  websiteSettings(website_id)
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
  if (document.body.clientWidth < 1200 && document.body.clientWidth > 768) {
    websiteFormDrawerWidth.value = '60%'
  } else if (document.body.clientWidth < 768) {
    websiteFormDrawerWidth.value = '100%'
  } else {
    websiteFormDrawerWidth.value = '50%'
  }

  //resize table
  let tbHeight = document.body.clientHeight - 215 - 140
  if (tbHeight < 300) {
    tableHeight.value = 300
  } else {
    tableHeight.value = tbHeight
  }
}



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
    dataType: 'form'
  })
    .then((res) => {
      if (res.code === 0) {
        ElMessage.success('网站删除成功')
        getWebSites({
          page: 1,
          pagesize: pageState.pagesize
        })
      }
    })
    .catch((error) => {
      ElMessage.error('网站删除失败')
    })
}
</script>
<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
