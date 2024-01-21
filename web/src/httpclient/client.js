import axios from 'axios'

import serverConfig from './config'
import router from '../router'
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
let loading = null
function start() {
  loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    background: 'rgba(0, 0, 0, 0.7)'
  })
}

// loading.close()
// 添加请求拦截器
serviceRequest.interceptors.request.use(
  (config) => {
    if (config.url != '/v1/statistics/dashboard') {
      start()
    }
    if(config.dataType == 'form'){
      config.headers['Content-Type'] = 'application/x-www-form-urlencoded'
    } else if(config.dataType == 'upload'){
      config.headers['Content-Type'] = 'multipart/form-data'
    }
   
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
    if (loading) {
      loading.close()
    }
    const dataAxios = response.data

    if (dataAxios.code) {
      ElMessage({
        type: 'error',
        message: dataAxios.msg
      })
      // Promise.reject(dataAxios)
    }

    return dataAxios
  },
  function (error) {
    if (loading) {
      loading.close()
    }

    const responseCode = error.response.status
    switch (responseCode) {
      case 400:
        ElMessage({
          type: 'error',
          message: '无效的Token，请重新登录'
        })
        router.replace({
          path: '/login',
          query: {
            redirect: router.currentRoute.fullPath
          }
        })
        break
      case 401:
        // 跳转登录页
        router.replace({
          path: '/login',
          query: {
            redirect: router.currentRoute.fullPath
          }
        })
        break
      // 403: token过期
      case 403:
        // 弹出错误信息
        ElMessage({
          type: 'error',
          message: '登录信息过期，请重新登录'
        })
        // 清除token
        sessionStorage.removeItem('access_token')
        setTimeout(() => {
          router.replace({
            path: '/login',
            query: {
              redirect: router.currentRoute.fullPath
            }
          })
        }, 1000)
        break
      // 404请求不存在
      case 404:
        ElMessage({
          message: '网络请求不存在',
          type: 'error'
        })
        break
      // 其他错误，直接抛出错误提示
      default:
        ElMessage({
          message: error.response.data.message,
          type: 'error'
        })
    }

    return Promise.reject(error)
  }
)

export default serviceRequest
