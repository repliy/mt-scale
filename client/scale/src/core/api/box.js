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

/**
 * get latest boxes
 */
export function getLatestBoxes(query) {
  return request({
    url: '/api/box/latestbox',
    method: 'get',
    params: query
  })
}
