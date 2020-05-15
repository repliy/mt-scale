const system = {
  // 系统编码
  code: process.env.VUE_APP_SYSTEM_CODE
}

// 字典
const dic = {
  token: 'Ingot-Cloud-Auth-Token'
}

const requestCode = {
  SUCCESS: 200000,
  // 未知异常
  UNKNOW: -1,
  // 10010004:Token 过期了;
  TOKEN_EXPIRE: 10010004,
  // 10010007:Token 已签退
  TOKEN_SIGN_BACK: 10010007
}

const config = {
  key_prefix: 'Ingot-Cloud-',
  domain: process.env.VUE_APP_COOKIE_DOMAIN
}

const enums = {
  USER: {
    LOGIN_NAME: 'Login-Name',
    REMEMBER_ME: 'Remember-Me',
    AUTH_TOKEN: 'Auth-Token',
    REFRESH_TOKEN: 'Refresh-Token',
    REDIRECT_URI: 'Redirect-Uri',
    BASIC_TOKEN: 'Basic-Token',
    NEED_CAPTCHA_CODE: 'Need-Captcha-Code'
  }
}

const token = {
  // basic token
  basic: process.env.VUE_APP_WEB_BASIC_TOKEN
}

const app = {
  title: process.env.VUE_APP_TITLE
}

export {
  token,
  config,
  enums,
  system,
  dic,
  requestCode,
  app
}
