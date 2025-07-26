import request from '@/utils/request'

export function getCaptcha() {
  return request.get('/v1/auth/captcha')
}

export function login(data: {
  username: string
  password: string
  captcha: string
  captchaKey: string
}) {
  return request.post('/v1/auth/login', {
    username: data.username,
    password: data.password,
    captcha: data.captcha,
    captchaKey: data.captchaKey
  })
}

export function logout() {
  return request.delete('/v1/auth/logout')
}

export function refreshToken(refreshToken: string) {
  return request.post('/v1/auth/refresh-token', {
    refreshToken
  })
} 