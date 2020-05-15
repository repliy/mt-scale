import request from '@/core/request'

/**
 * get species list
 */
export function getSpecies(query) {
  return request({
    url: '/api/species/fetchall',
    method: 'get',
    params: query
  })
}
