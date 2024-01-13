import axios from 'axios'
import router from '../router'
import { Message,Loading } from 'element-ui'
const service = axios.create({
  timeout: 60000,
  baseURL: process.env.BASE_URL
})
let loading = null

service.interceptors.request.use(config => {
  const token = sessionStorage.getItem('token')
  if (token) {
    config.headers['Authorization'] = 'Bearer ' + token
  }
  loading = Loading.service({
    text: '正在加载中......'
  })
  return config
}, (error) => {
  return Promise.reject(error)
})

service.interceptors.response.use(response => {
  const responseCode = response.status
  console.log(loading);
  if (loading) {
    loading.close()
  }

  
  if (responseCode === 200) {
    return Promise.resolve(response)
  } else {
    return Promise.reject(response)
  }
}, error => {
    if (loading) {
        loading.close()
    }
  
  if (!error.response) {
    if (error.message.includes('timeout')) {
      console.log('超时了')
      Message.error('请求超时，请检查网络是否连接正常')
    } else {
      console.log('断网了')
      Message.error('请求失败，请检查网络是否已连接')
    }
    return
  }


  const responseCode = error.response.status
  switch (responseCode) {
    // 401：未登录
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
      Message({
        type: 'error',
        message: '登录信息过期，请重新登录'
      })
      // 清除token
      sessionStorage.removeItem('token')
      // 跳转登录页面，并将要浏览的页面fullPath传过去，登录成功后跳转需要访问的页面
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
      Message({
        message: '网络请求不存在',
        type: 'error'
      })
      break
    // 其他错误，直接抛出错误提示
    default:
      Message({
        message: error.response.data.message,
        type: 'error'
      })
  }
  return Promise.reject(error)
})
 
export default service