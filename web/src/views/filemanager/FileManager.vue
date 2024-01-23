<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <el-row :gutter="20">
          <el-col :sm="24" :md="12" :lg="12">
            <el-input v-model="inputFileDynPath">
              <template #append>
                <el-button><Icon icon="ic:outline-search" size="24" /></el-button>
              </template>
            </el-input>

            <ul style="position: absolute; top: 8px" class="fm-breadcrumb">
              <li>
                <el-link :underline="false" @click="changePath(0)">
                  <Icon icon="ic:round-home" width="20" height="20" />
                </el-link>
              </li>
              <li v-for="(item, index) in pageState.breadcrumbPaths">
                <el-link :underline="false" @click="changePath(index)">{{ item }}</el-link>
              </li>
            </ul>
          </el-col>
          <el-col :sm="24" :md="12" :lg="12">
            <el-button title="上一级"
              ><el-icon><ArrowLeftBold /></el-icon
            ></el-button>
          </el-col>
        </el-row>
      </div>
    </template>

    <el-table
      :data="tableData"
      min-height="200"
      style="width: 100%"
      :height="tableHeight"
      v-loading="pageState.tbLoading"
    >
      <el-table-column type="selection" width="50" />
      <el-table-column prop="name" label="文件名" min-width="250" show-overflow-tooltip>
        <template #default="{ row, $index }">
          <el-button link @click.prevent="openFile($index)">
            <template #icon>
              <span v-if="row.is_file">
                <Icon icon="ic:round-insert-drive-file" width="28" height="28"
              /></span>
              <span v-else>
                <Icon icon="ic:baseline-folder" width="28" height="28" :inline="true"
              /></span>
            </template>

            {{ row.name }}
          </el-button>
          <template v-if="row.is_symlink === true">
            <el-text class="mx-1 color-zinc-200" type="info" size="small">
              > {{ row.symlink_to }}</el-text
            ></template
          >
        </template>
      </el-table-column>
      <el-table-column prop="perm" label="权限" width="70" />
      <el-table-column prop="uid" label="UID" width="60" />
      <el-table-column prop="gid" label="GID" width="60" />
      <el-table-column prop="mod_time" label="修改时间" width="180" />
      <el-table-column prop="filesize" label="文件大小" width="120" />
      <el-table-column fixed="right" label="操作" width="120">
        <template #default="scope">
          <el-button
            :icon="Edit"
            circle
            class="mr-2"
            title="编辑"
            @click="editFile(scope.$index)"
          />
          <el-dropdown>
            <span class="el-dropdown-link">
              <Icon icon="ri:more-fill" width="24" height="24" />
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click.prevent="deleteRow(scope.$index)">删除</el-dropdown-item>
                <el-dropdown-item @click.prevent="coypFilePath(scope.$index)"
                  >复制文件路径</el-dropdown-item
                >
                <el-dropdown-item>重命名</el-dropdown-item>
                <el-dropdown-item>压缩文件</el-dropdown-item>
                <el-dropdown-item>权限</el-dropdown-item>
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
</template>
<script setup>
import { Icon } from '@iconify/vue'
import { onMounted, ref } from 'vue'
import apiRequest from '../../httpclient/client'
import { useGlobalStore } from '../../stores/global'
import { ElMessage } from 'element-plus'
import { Edit, ArrowLeftBold } from '@element-plus/icons-vue'
import { copyText } from 'vue3-clipboard'
const globalStore = useGlobalStore()
const tableData = ref([])
const tableHeight = ref(300)

const inputFileDynPath = ref('')

const handlePaginration = (currentPage, pageSize) => {
  pageState.currentpage = currentPage
  pageState.pagesize = pageSize
  getFiles({
    path: globalStore.lastFilePath,
    page: currentPage,
    pagesize: pageSize
  })
}
const pageState = reactive({
  inputFileDynPath: '',
  breadcrumbPaths: [],
  total: 0,
  pagesize: 50,
  currentpage: 1,
  tbLoading: false
})
const resizeTable = () => {
  let tbHeight = document.body.clientHeight - 215 - 140
  if (tbHeight < 300) {
    tableHeight.value = 300
  } else {
    tableHeight.value = tbHeight
  }
}
onMounted(() => {
  if (globalStore.lastFilePath === '') {
    globalStore.lastFilePath = '/'
  }
  getFiles({
    path: globalStore.lastFilePath,
    page: 1,
    pagesize: pageState.pagesize
  })

  window.onresize = resizeTable
})

const openFile = (index) => {
  const row = tableData.value[index]
  if (!row) {
    return
  }
  if (row.is_dir) {
    getFiles({ path: globalStore.lastFilePath + '/' + row.name })
  } else {
    ElMessage({
      message: 'open file'
    })
  }
}

const changePath = (index) => {
  const pageSegs = [].slice.call(pageState.breadcrumbPaths)
  let fullpath = pageSegs.splice(0, index + 1).join('')
  if (fullpath === '') {
    fullpath = '/'
  }
  getFiles({
    path: fullpath
  })
}
//刷新
const refreshData = () => {
  updateBreadCrumb(globalStore.lastFilePath)
}
//更新Path
function updateBreadCrumb(path) {
  path = path || ''
  const pathArr = path.split('/')
  pageState.breadcrumbPaths.splice(0, pageState.breadcrumbPaths.length)
  if (pathArr.length > 0) {
    for (let index = 0; index < pathArr.length; index++) {
      const element = pathArr[index]
      if (element === '') {
        continue
      }

      pageState.breadcrumbPaths.push('/')
      pageState.breadcrumbPaths.push(element)
    }
  }
}

const deleteRow = (index) => {
  console.log(tableData.value[index])
}
const editFile = (index) => {
  console.log(tableData.value[index])
}
const coypFilePath = (index) => {
  copyText(
    globalStore.lastFilePath + '/' + tableData.value[index].name,
    undefined,
    (error, event) => {
      if (error) {
        ElMessage({ message: '无法复制', type: 'warning' })
      } else {
        ElMessage({ message: '复制成功', type: 'success' })
      }
    }
  )
}

function getFiles(options) {
  console.log('request getFIle', options)
  const defaultOptions = {
    path: globalStore.lastFilePath,
    page: 1,
    pagesize: 50
  }
  options = Object.assign(defaultOptions, options)
  pageState.tbLoading = true
  apiRequest({
    url: '/v1/filemanager/list',
    method: 'post',
    data: {
      path: options.path,
      page: options.page,
      limit: options.pagesize
    },
    dataType: 'form',
    loading: false
  })
    .then((data) => {
      tableData.value = []
      tableData.value.push(...data.data)

      console.log('push data')
      globalStore.lastFilePath = data.path
      pageState.total = data.total
      pageState.currentpage = data.page
      updateBreadCrumb(data.path)
      resizeTable()
    })
    .catch((error) => {
      console.log(error)
    })
    .finally(() => (pageState.tbLoading = false))
}
</script>
<style scoped>
.fm-breadcrumb {
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
