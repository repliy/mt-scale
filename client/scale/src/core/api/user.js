import request from '@/core/request'

/**
 * login
 */
export function login(params) {
  return request({
    url: '/api/user/login',
    method: 'post',
    data: params
  })
}