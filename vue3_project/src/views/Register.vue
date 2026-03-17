<template>
  <div class="register-container">
    <div class="register-card">
      <div class="logo-section">
        <img src="/logo.png" alt="Logo" class="logo" />
        <h1 class="app-name">港口学堂</h1>
        <p class="subtitle">欢迎注册</p>
      </div>

      <!-- 提示信息 -->
      <div class="message-container">
        <div v-if="message" class="message" :class="messageType">
          {{ message }}
        </div>
      </div>

      <div class="input-section">
        <div class="input-item">
          <span class="icon">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
            </svg>
          </span>
          <input
            type="text"
            placeholder="请输入用户名"
            v-model="username"
          />
        </div>

        <div class="input-item">
          <span class="icon">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M17 1H7c-1.1 0-2 .9-2 2v18c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2zm-5 20c-.55 0-1-.45-1-1s.45-1 1-1 1 .45 1 1-.45 1-1 1zm5-4H7V4h10v13z"/>
            </svg>
          </span>
          <input
            type="tel"
            placeholder="请输入手机号"
            v-model="phoneNumber"
            maxlength="11"
          />
        </div>

        <div class="input-item">
          <span class="icon">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M18 8h-1V6c0-2.76-2.24-5-5-5S7 3.24 7 6v2H6c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V10c0-1.1-.9-2-2-2zm-6 9c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm3.1-9H8.9V6c0-1.71 1.39-3.1 3.1-3.1 1.71 0 3.1 1.39 3.1 3.1v2z"/>
            </svg>
          </span>
          <input
            :type="showPassword ? 'text' : 'password'"
            placeholder="请设置密码（6-20位）"
            v-model="password"
          />
          <button class="toggle-password" @click="showPassword = !showPassword">
            <svg v-if="!showPassword" viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M12 4.5C7 4.5 2.73 7.61 1 12c1.73 4.39 6 7.5 11 7.5s9.27-3.11 11-7.5c-1.73-4.39-6-7.5-11-7.5zM12 17c-2.76 0-5-2.24-5-5s2.24-5 5-5 5 2.24 5 5-2.24 5-5 5zm0-8c-1.66 0-3 1.34-3 3s1.34 3 3 3 3-1.34 3-3-1.34-3-3-3z"/>
            </svg>
            <svg v-else viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M12 7c2.76 0 5 2.24 5 5 0 .65-.13 1.26-.36 1.83l2.92 2.92c1.51-1.26 2.7-2.89 3.43-4.75-1.73-4.39-6-7.5-11-7.5-1.4 0-2.74.25-3.98.7l2.16 2.16C10.74 7.13 11.35 7 12 7zM2 4.27l2.28 2.28.46.46C3.08 8.3 1.78 10.02 1 12c1.73 4.39 6 7.5 11 7.5 1.55 0 3.03-.3 4.38-.84l.42.42L19.73 22 21 20.73 3.27 3 2 4.27zM7.53 9.8l1.55 1.55c-.05.21-.08.43-.08.65 0 1.66 1.34 3 3 3 .22 0 .44-.03.65-.08l1.55 1.55c-.67.33-1.41.53-2.2.53-2.76 0-5-2.24-5-5 0-.79.2-1.53.53-2.2zm4.31-.78l3.15 3.15.02-.16c0-1.66-1.34-3-3-3l-.17.01z"/>
            </svg>
          </button>
        </div>

        <div class="input-item">
          <span class="icon">
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M18 8h-1V6c0-2.76-2.24-5-5-5S7 3.24 7 6v2H6c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V10c0-1.1-.9-2-2-2zm-6 9c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm3.1-9H8.9V6c0-1.71 1.39-3.1 3.1-3.1 1.71 0 3.1 1.39 3.1 3.1v2z"/>
            </svg>
          </span>
          <input
            :type="showConfirmPassword ? 'text' : 'password'"
            placeholder="请再次输入密码"
            v-model="confirmPassword"
          />
          <button class="toggle-password" @click="showConfirmPassword = !showConfirmPassword">
            <svg v-if="!showConfirmPassword" viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M12 4.5C7 4.5 2.73 7.61 1 12c1.73 4.39 6 7.5 11 7.5s9.27-3.11 11-7.5c-1.73-4.39-6-7.5-11-7.5zM12 17c-2.76 0-5-2.24-5-5s2.24-5 5-5 5 2.24 5 5-2.24 5-5 5zm0-8c-1.66 0-3 1.34-3 3s1.34 3 3 3 3-1.34 3-3-1.34-3-3-3z"/>
            </svg>
            <svg v-else viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M12 7c2.76 0 5 2.24 5 5 0 .65-.13 1.26-.36 1.83l2.92 2.92c1.51-1.26 2.7-2.89 3.43-4.75-1.73-4.39-6-7.5-11-7.5-1.4 0-2.74.25-3.98.7l2.16 2.16C10.74 7.13 11.35 7 12 7zM2 4.27l2.28 2.28.46.46C3.08 8.3 1.78 10.02 1 12c1.73 4.39 6 7.5 11 7.5 1.55 0 3.03-.3 4.38-.84l.42.42L19.73 22 21 20.73 3.27 3 2 4.27zM7.53 9.8l1.55 1.55c-.05.21-.08.43-.08.65 0 1.66 1.34 3 3 3 .22 0 .44-.03.65-.08l1.55 1.55c-.67.33-1.41.53-2.2.53-2.76 0-5-2.24-5-5 0-.79.2-1.53.53-2.2zm4.31-.78l3.15 3.15.02-.16c0-1.66-1.34-3-3-3l-.17.01z"/>
            </svg>
          </button>
        </div>
      </div>

      <button
        class="register-button"
        v-bind:class="{ disabled: !agreeToTerms }"
        @click="handleRegister"
      >
        立即注册
      </button>

      <div class="terms-section">
        <input type="checkbox" id="terms" v-model="agreeToTerms" />
        <label for="terms">
          我已阅读并同意
          <a href="#" class="link">《服务协议》</a>
          和
          <a href="#" class="link">《隐私协议》</a>
        </label>
      </div>

      <div class="login-link">
        已有账号？
        <router-link to="/login" class="link">立即登录</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { userApi } from '@/api'

const router = useRouter()

const username = ref('')
const phoneNumber = ref('')
const password = ref('')
const confirmPassword = ref('')
const agreeToTerms = ref(false)
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const message = ref('')
const messageType = ref<'error' | 'success'>('error')

let messageTimer: number | null = null

function showMessage(msg: string, type: 'error' | 'success' = 'error') {
  message.value = msg
  messageType.value = type

  if (messageTimer) {
    clearTimeout(messageTimer)
  }

  messageTimer = window.setTimeout(() => {
    message.value = ''
    messageTimer = null
  }, 3000)
}

function validateForm() {
  if (!username.value.trim()) {
    showMessage('请输入用户名', 'error')
    return false
  }

  if (!phoneNumber.value) {
    showMessage('请输入手机号', 'error')
    return false
  }

  const phoneRegex = /^1[3-9]\d{9}$/
  if (!phoneRegex.test(phoneNumber.value)) {
    showMessage('请输入正确的手机号', 'error')
    return false
  }

  if (!password.value) {
    showMessage('请设置密码', 'error')
    return false
  }

  if (password.value.length < 6 || password.value.length > 20) {
    showMessage('密码长度应为6-20位', 'error')
    return false
  }

  if (password.value !== confirmPassword.value) {
    showMessage('两次输入的密码不一致', 'error')
    return false
  }

  if (!agreeToTerms.value) {
    showMessage('请先阅读并同意《服务协议》和《隐私协议》', 'error')
    return false
  }

  return true
}

async function handleRegister() {
  if (!validateForm()) {
    return
  }

  try {
    const response: any = await userApi.register(
      username.value,
      phoneNumber.value,
      password.value
    )

    if (response.code === 200) {
      showMessage('注册成功！', 'success')
      setTimeout(() => {
        router.push('/login')
      }, 1000)
    } else {
      showMessage(response.msg || '注册失败', 'error')
    }
  } catch (error: any) {
    console.error('注册失败:', error)
    showMessage(error.response?.data?.msg || '注册失败，请稍后重试', 'error')
  }
}
</script>

<style scoped>
.register-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(180deg, #eef4ff 0%, #ffffff 100%);
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  padding: 20px 0;
}

.register-card {
  width: 100%;
  max-width: 400px;
  padding: 40px 24px;
  text-align: center;
  position: relative;
}

.logo-section {
  margin-bottom: 40px;
}

.logo {
  width: 100px;
  height: 100px;
  border-radius: 20px;
  margin-bottom: 16px;
  box-shadow: 0 4px 12px rgba(66, 133, 244, 0.2);
}

.app-name {
  color: #3b82f6;
  font-size: 28px;
  font-weight: 900;
  font-style: italic;
  letter-spacing: 1px;
  margin-bottom: 8px;
}

.subtitle {
  color: #64748b;
  font-size: 16px;
  margin: 0;
}

.input-section {
  margin-bottom: 20px;
}

.input-item {
  display: flex;
  align-items: center;
  background: white;
  border-radius: 16px;
  margin-bottom: 20px;
  padding: 16px 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
  border: 1px solid #f8fafc;
}

.icon {
  color: #94a3b8;
  margin-right: 12px;
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

input {
  border: none;
  outline: none;
  flex: 1;
  font-size: 16px;
  color: #333;
}

input::placeholder {
  color: #cbd5e1;
}

.get-code {
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 16px;
  cursor: pointer;
  padding-left: 12px;
  border-left: 1px solid #f0f0f0;
  white-space: nowrap;
  flex-shrink: 0;
}

.get-code:disabled {
  color: #94a3b8;
  cursor: not-allowed;
}

.toggle-password {
  background: none;
  border: none;
  color: #94a3b8;
  cursor: pointer;
  padding: 0;
  margin-left: 8px;
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.toggle-password:hover {
  color: #3b82f6;
}

.register-button {
  width: 100%;
  padding: 16px;
  background: #3b82f6;
  color: white;
  border: none;
  border-radius: 30px;
  font-size: 20px;
  font-weight: bold;
  cursor: pointer;
  box-shadow: 0 8px 16px rgba(59, 130, 246, 0.2);
  margin-top: 20px;
  margin-bottom: 24px;
  transition: all 0.3s;
}

.register-button:hover {
  background: #2563eb;
  transform: translateY(-2px);
  box-shadow: 0 12px 20px rgba(59, 130, 246, 0.3);
}

.register-button.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.register-button.disabled:hover {
  transform: none;
  background: #3b82f6;
}

.terms-section {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: #64748b;
  margin-bottom: 20px;
}

.terms-section input[type="checkbox"] {
  width: 16px;
  height: 16px;
  margin-right: 8px;
  cursor: pointer;
  accent-color: #3b82f6;
  flex-shrink: 0;
}

.link {
  color: #3b82f6;
  text-decoration: none;
}

.link:hover {
  text-decoration: underline;
}

.login-link {
  font-size: 14px;
  color: #64748b;
}

.message-container {
  min-height: 44px;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.message {
  font-size: 16px;
  text-align: center;
  opacity: 0;
  animation: fadeIn 0.3s ease forwards;
}

.message.error {
  color: #ef4444;
}

.message.success {
  color: #10b981;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@media (max-width: 480px) {
  .register-card {
    padding: 20px;
  }

  .logo {
    width: 80px;
    height: 80px;
  }

  .app-name {
    font-size: 24px;
  }
}
</style>
