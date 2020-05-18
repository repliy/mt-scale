import request from '@/core/request'

/**
 * create box list
 */
export function createBoxList(params) {
  return request({
    url: '/api/box/crtboxes',
    method: 'post',
    data: params
  })
}
