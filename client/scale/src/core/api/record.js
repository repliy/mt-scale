/* jshint esversion: 6 */
import request from '@/core/request'

/**
 * add a weight record
 */
export function AddWeightRecord(params) {
  return request({
    url: '/api/record/crt',
    method: 'post',
    data: params
  })
}

/**
 * update a weight record
 */
export function UpdWeightRecord(params) {
  return request({
    url: '/api/record/upd',
    method: 'post',
    data: params
  })
}

/**
 * get weight record list
 */
export function FetchWeightRecord(query) {
  return request({
    url: '/api/record/fetch',
    method: 'get',
    params: query
  })
}

/**
 * delete weight record
 */
export function DeleteWeightRecord(params) {
  return request({
    url: '/api/record/del',
    method: 'post',
    data: params
  })
}

/**
 * statis weight
 */
export function StatWeight(query) {
  return request({
    url: '/api/stat/weight',
    method: 'get',
    params: query
  })
}
