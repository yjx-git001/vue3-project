<template>
  <div class="login-container">
    <div class="login-card">
      <div class="logo-section">
        <img src="/logo.png" alt="Logo" class="logo" />
        <h1 class="app-name">港口学堂</h1>
      </div>

      <div class="input-section">
        <div class="input-item">
          <span class="icon">
            <!-- Mobile Icon -->
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M17 1H7c-1.1 0-2 .9-2 2v18c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2zm-5 20c-.55 0-1-.45-1-1s.45-1 1-1 1 .45 1 1-.45 1-1 1zm5-4H7V4h10v13z"/>
            </svg>
          </span>
          <input 
            type="tel" 
            placeholder="请输入手机号" 
            v-model="phoneNumber"
          />
        </div>

        <div class="input-item">
          <span class="icon">
            <!-- Shield/Code Icon -->
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
              <path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm-2 16l-4-4 1.41-1.41L10 14.17l6.59-6.59L18 9l-8 8z"/>
            </svg>
          </span>
          <input 
            type="text" 
            placeholder="请输入验证码" 
            v-model="verificationCode"
          />
          <button class="get-code" @click="getVerificationCode">获取</button>
        </div>
      </div>

      <button 
        class="login-button" 
        :class="{ disabled: !agreeToTerms }" 
        @click="handleLogin"
      >
        立即登录
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const phoneNumber = ref('')
const verificationCode = ref('')
const agreeToTerms = ref(false)

function handleLogin() {
  if (!agreeToTerms.value) {
    alert('请先阅读并同意《服务协议》和《隐私协议》')
    return
  }
  
  if (phoneNumber.value && verificationCode.value) {
    alert('登录成功！')
    console.log('Login successful with:', phoneNumber.value)
  } else {
    alert('请输入手机号和验证码')
  }
}

function getVerificationCode() {
  if (!phoneNumber.value) {
    alert('请输入手机号')
    return
  }
  alert('验证码已发送')
  console.log('Getting verification code for:', phoneNumber.value)
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(180deg, #eef4ff 0%, #ffffff 100%);
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 40px 24px;
  text-align: center;
}

.logo-section {
  margin-bottom: 60px;
}

.input-section {
  margin-bottom: 20px;
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
}

.input-item {
  display: flex;
  align-items: center;
  background: white;
  border-radius: 16px;
  margin-bottom: 24px;
  padding: 16px 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
  border: 1px solid #f8fafc;
}

.icon {
  color: #94a3b8;
  margin-right: 12px;
  display: flex;
  align-items: center;
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
}

.login-button {
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
  margin-top: 40px;
  margin-bottom: 24px;
  transition: all 0.3s;
}

.login-button.disabled {
  opacity: 0.7;
}

.terms-section {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: #64748b;
}

.terms-section input[type="checkbox"] {
  width: 16px;
  height: 16px;
  margin-right: 8px;
  cursor: pointer;
  accent-color: #3b82f6;
}

.link {
  color: #3b82f6;
  text-decoration: none;
}

@media (max-width: 480px) {
  .login-card {
    padding: 20px;
  }
}
</style>
