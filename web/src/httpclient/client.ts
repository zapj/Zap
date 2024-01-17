import axios from 'axios'

import serverConfig from './config'

import { ElLoading, ElMessage } from 'element-plus' 

const serviceRequest = axios.create({
  baseURL: serverConfig.baseURL, //请求后端数据的基本地址，自定义
  timeout: 200000, //请求超时设置，单位ms
  withCredentials: true, // 异步请求携带cookie
  headers: {
    'Content-Type': 'application/json',
    // 'Authorization': 'Bearer ',
    'X-Requested-With': 'XMLHttpRequest'
  }
})
let loading: any
function start() {
  loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    background: 'rgba(0, 0, 0, 0.7)'
  })
}
declare module 'axios' {
  interface AxiosInstance {
    (config: AxiosRequestConfig): Promise<any>
  }
}

// loading.close()
// 添加请求拦截器
serviceRequest.interceptors.request.use(
  function (config: any) {

    if (config.url != '/v1/statistics/dashboard') {
      start()
    }
    

    // if (config.url == '/api/file/upload') {
    // 	config.headers['Content-Type'] = 'multipart/form-data'
    // }
    const access_token = sessionStorage.getItem('access_token')
    if (access_token) {
      config.headers['Authorization'] = 'Bearer ' + access_token
    }
    return config
  },
  function (error) {
    return Promise.reject(error)
  }
)

// 添加响应拦截器
serviceRequest.interceptors.response.use(
  function (response) {
    if (loading){
      loading.close()
    }
    const dataAxios = response.data

    const code = dataAxios.reset
    return dataAxios
  },
  function (error) {
    if (loading){
      loading.close()
    }
    // if (error.config.url.indexOf('/api/file/list') == -1) {
    //   ElMessage.error(error.response.data.msg)
    // }
    ElMessage.error(error.response.data.msg)
    return Promise.reject(error)
  }
)

export default serviceRequest
