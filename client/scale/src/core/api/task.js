import request from '@/core/request'

/**
 * get latest task
 */
export function getLatestTask() {
  return request({
    url: '/api/task/latest',
    method: 'get',
    params: {}
  })
}
