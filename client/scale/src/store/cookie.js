import Cookies from 'js-cookie'
import { config } from '@/core/config'

class IngotCookie {
  constructor() {
    this.keyPrefix = config.key_prefix
    this.domain = config.domain
    this.expireTime = 7
  }

  set(cookie) {
    const { key, value, expires, path, success } = cookie
    IngotCookie.checkKey(key)
    const finalKey = this.keyPrefix + key
    Cookies.set(finalKey, value, {
      expires: expires || this.expireTime,
      path: path || '/',
      domain: this.domain
    })
    success && success()
  }

  get(key) {
    IngotCookie.checkKey(key)
    return Cookies.get(this.keyPrefix + key)
  }

  remove(cookie) {
    const { key, path, success } = cookie
    IngotCookie.checkKey(key)
    Cookies.remove(this.keyPrefix + key, {
      path: path || '/',
      domain: this.domain
    })
    success && success()
  }

  getAll() {
    return Cookies.get()
  }

  static checkKey(key) {
    if (!key) {
      throw new Error('给定的参数key无效')
    }
    if (typeof key === 'object') {
      throw new Error('key不能是一个对象。')
    }
  }
}

export default new IngotCookie()
