import axios from 'axios'
import { requestCode } from '@/core/config.js'
import log from '@/utils/log'
import store from '@/store'
// import router from '@/router'

// 超时时间
const requestTimeout = 10000

// 错误码处理器
const errorCoderHandler = (response, config) => {
  log.d('response:', response)
  const code = response.code
  switch (code) {
    case requestCode.TOKEN_AUTH_FAILED:
      store.dispatch('FedLogOut').then(() => {
        // router.push('/login')
        location.href = '/login'
      })
  }
}

// 创建axios实例
const service = axios.create({
  // baseURL: process.env.BASE_API, // api 的 base_url // 使用当前域
  timeout: requestTimeout // 请求超时时间
})

// request
service.interceptors.request.use(
  config => {
    if (store.getters.getAccessToken) {
      config.headers.Authorization = `Bearer ${store.getters.getAccessToken}`
    }
    return config
  },
  error => {
    // Do something with request error
    log.d(error) // for debug
    Promise.reject(error)
  }
)

// response
service.interceptors.response.use(
  response => {
    const contentType = response.headers['content-type']
    if (!contentType.startsWith('application/json')) {
      return response
    }
    /**
     * code为非99990200是抛错 可结合自己业务进行修改
     */
    const res = response.data
    if (res.code !== requestCode.SUCCESS) {
      const handleResult = errorCoderHandler(res, response.config)
      if (handleResult) {
        return handleResult
      }
      return Promise.reject(res)
    } else {
      return res
    }
  },
  error => {
    const response = error.response
    let message = error.message
    let code
    if (response) {
      // 业务数据
      if (response.data) {
        const temp = response.data.message
        if (temp && temp.length > 0) {
          message = temp
        }
        code = response.data.code
      } else {
        code = response.status
      }
    } else {
      code = requestCode.UNKNOW
      message = '网络异常，请稍后重试'
      log.d(`message=${message}`)
    }

    const wrapper = {
      message,
      code,
      response
    }

    const handleResult = errorCoderHandler(wrapper, error.config)
    if (handleResult) {
      return handleResult
    }

    return Promise.reject(wrapper)
  }
)

export default service
