<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">订单管理系统登录</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="账号" prop="username">
          <el-input v-model="form.username" placeholder="请输入账号" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码" show-password />
        </el-form-item>
        <el-form-item label="验证码" prop="captcha">
          <el-row :gutter="10">
            <el-col :span="12">
              <el-input v-model="form.captcha" placeholder="请输入验证码" maxlength="6" />
            </el-col>
            <el-col :span="12">
              <img
                :src="captchaImg"
                @click="fetchCaptcha"
                class="captcha-img"
                title="点击刷新验证码"
                v-if="captchaImg"
              />
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" @click="onLogin">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter, useRoute } from 'vue-router'
import { login, getCaptcha } from '@/api/auth'

const router = useRouter()
const route = useRoute()
const formRef = ref()
const form = reactive({
  username: '',
  password: '',
  captcha: '',
  captchaKey: ''
})
const rules = {
  username: [{ required: true, message: '请输入账号', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }]
}
const captchaImg = ref('')
const loading = ref(false)

const fetchCaptcha = async () => {
  const res = await getCaptcha()
  if (res?.data?.captchaBase64 && res?.data?.captchaKey) {
    captchaImg.value = res.data.captchaBase64
    form.captchaKey = res.data.captchaKey
    form.captcha = ''
  }
}

const onLogin = async () => {
  await formRef.value?.validate()
  loading.value = true
  try {
    const res = await login({
      username: form.username,
      password: form.password,
      captcha: form.captcha,
      captchaKey: form.captchaKey
    })
    if (res.code === '0') {
      // 保存token和用户信息
      localStorage.setItem('token', res.data.accessToken)
      localStorage.setItem('userInfo', JSON.stringify({
        username: form.username,
        nickname: res.data.nickname || form.username
      }))
      
      ElMessage.success('登录成功')
      
      // 自动跳转：如果有redirect参数则跳转到指定页面，否则跳转到首页
      const redirect = route.query.redirect as string
      router.push(redirect || '/')
    } else {
      ElMessage.error(res.msg || '登录失败')
      fetchCaptcha()
    }
  } catch (e) {
    ElMessage.error('登录异常')
    fetchCaptcha()
  } finally {
    loading.value = false
  }
}

onMounted(fetchCaptcha)
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f6fa;
}
.login-card {
  width: 380px;
  padding: 32px 24px 16px 24px;
}
.login-title {
  text-align: center;
  margin-bottom: 24px;
  font-weight: bold;
  font-size: 22px;
  color: #333;
}
.captcha-img {
  height: 38px;
  cursor: pointer;
  border-radius: 4px;
  border: 1px solid #eee;
  background: #fff;
}
</style> 