import { enums } from '@/core/config'
import IngotCookie from '@/store/cookie'
import router from '@/router'

const user = {
  state: {
    username: '',
    authToken: {
      access_token: '',
      expires_in: 2 / 24,
      timestamp: ''
    }
  },
  getters: {
    getAccessToken: state => {
      if (!state.authToken || !state.authToken.access_token || state.authToken.access_token.length === 0) {
        state.authToken = IngotCookie.get(enums.USER.AUTH_TOKEN) ? JSON.parse(IngotCookie.get(enums.USER.AUTH_TOKEN)) : {}
      }
      return state.authToken.access_token
    },
    getUsername: state => {
      if (!state.username || state.username.length === 0) {
        state.username = IngotCookie.get(enums.USER.LOGIN_NAME)
      }
      return state.username
    }
  },
  mutations: {
    updateAuthToken(state, authToken) {
      const timestamp = new Date().getTime()
      // authToken 中返回的 expires 时间单位为秒，IngotCookie接收单位为天，如果小时可以用小数，比如2小时为2/24天
      if (authToken.access_token) {
        // access token 默认保留7天，该时间不是token有效时间，是保留在浏览器cookies的时间
        const accessTokenExpires = 7
        // 如果存在 refresh token，则保存时间和 refresh token 失效时间相同，为了access token 失效后，刷新token
        const accessToken = {
          access_token: authToken.access_token,
          expires_in: accessTokenExpires,
          timestamp: timestamp
        }
        state.authToken = accessToken
        IngotCookie.set({
          key: enums.USER.AUTH_TOKEN,
          value: accessToken,
          expires: accessToken.expires_in
        })
      }
    },
    SET_USERNAME(state, name) {
      const usernameExpires = 7
      state.username = name
      IngotCookie.set({
        key: enums.USER.LOGIN_NAME,
        value: name,
        expires: usernameExpires
      })
    },
    logout(state) {
      state.username = ''
      state.authToken = {}
      IngotCookie.remove({
        key: enums.USER.AUTH_TOKEN
      })
      IngotCookie.remove({
        key: enums.USER.LOGIN_NAME
      })
      router.push('/login')
    }
  },
  actions: {
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('logout')
        resolve()
      })
    }
  }
}

export default user
