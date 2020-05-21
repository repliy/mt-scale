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

/**
 * update task status
 */
export function updateTaskStatus(params) {
  return request({
    url: '/api/task/status/upd',
    method: 'post',
    data: params
  })
}
