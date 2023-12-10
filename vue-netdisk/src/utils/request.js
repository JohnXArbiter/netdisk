import { ElMessage } from 'element-plus'
import axios from 'axios'

// 1.axios二次封装
// 相当于我们创建了一个带配置的axios，然后用这个
const request = axios.create({
  // 基础路径
  // baseURL: '/api',
  // import.meta.env.BASE_URL.VITE_APP_BASE_API, //请求路径上都会写带上这个url  但是这里不知道为什么读取不到
  timeout: 5 * 1000, // 请求超时时间
})

// 2.request实例添加请求和拦截器
request.interceptors.request.use((config) => {
  // config配置对象有headers请求头
  return config
})

// 3.响应拦截器
request.interceptors.response.use(
  (response) => {
    // 这里是简化了数据
    return response.data
  },
  (error) => {
    // 一般处理http网络错误
    // 定义一个变量，存储网络错误
    const status = error.response.status
    let msg
    switch (status) {
      case 201:
        msg = '用户名或密码不对'
        break
      case 401:
        // 401 一般是token过期
        msg = 'TOKEN过期'
        break
      case 403:
        msg = '无权访问'
        break
      case 404:
        msg = '没有这个资源'
        break
      case 500:
        msg = '服务器嗝屁了，哈哈哈'
        break
      default:
        msg = '网络出现问题'
        break
    }
    // 提示错误信息
    ElMessage({
      type: 'error',
      message: msg,
    })
    return Promise.reject(error)
  },
)

export default request
