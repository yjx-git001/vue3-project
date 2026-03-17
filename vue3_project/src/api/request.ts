// 导入 axios
import axios from 'axios'

// 创建 axios 实例
const request = axios.create({
  baseURL: '/api',  // 走 Vite 代理转发到 Go 后端 localhost:8080
  timeout: 5000  // 请求超时时间（5秒）
})

// 请求拦截器（发送请求前执行）
request.interceptors.request.use(
  (config) => {
    // 添加 token 到请求头
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器（收到响应后执行）
request.interceptors.response.use(
  (response) => {
    // 返回响应数据
    return response.data
  },
  (error) => {
    // 处理错误
    console.error('请求失败:', error.message)
    return Promise.reject(error)
  }
)

export default request
