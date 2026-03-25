<template>
  <div class="account-settings-page">
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">账号设置</span>
      <span class="nav-placeholder"></span>
    </header>

    <main class="settings-content">
      <!-- 更换头像 -->
      <div class="setting-item" @click="triggerAvatarUpload">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon">
            <path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" fill="#3b82f6" />
          </svg>
          <span class="setting-label">更换头像</span>
        </div>
        <div class="setting-right">
          <img :src="avatarUrl || '/logo.png'" alt="头像" class="avatar-preview" />
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>
      <input ref="avatarInput" type="file" accept="image/*" style="display:none" @change="handleAvatarChange" />

      <!-- 昵称 -->
      <div class="setting-item" @click="openEdit('nickname', '昵称', profile.nickname)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M20.71,4.63L19.37,3.29C19,2.9 18.35,2.9 17.96,3.29L9,12.25L11.75,15L20.71,6.04C21.1,5.65 21.1,5 20.71,4.63M7,14A3,3 0 0,0 4,17C4,18.31 2.84,19 2,19C2.92,20.22 4.5,21 6,21A4,4 0 0,0 10,17A3,3 0 0,0 7,14Z" fill="#3b82f6" /></svg>
          <span class="setting-label">昵称</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.nickname }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 手机号 -->
      <div class="setting-item" @click="openEdit('phone', '手机号', profile.phone)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M6.62,10.79C8.06,13.62 10.38,15.94 13.21,17.38L15.41,15.18C15.69,14.9 16.08,14.82 16.43,14.93C17.55,15.3 18.75,15.5 20,15.5A1,1 0 0,1 21,16.5V20A1,1 0 0,1 20,21A17,17 0 0,1 3,4A1,1 0 0,1 4,3H7.5A1,1 0 0,1 8.5,4C8.5,5.25 8.7,6.45 9.07,7.57C9.18,7.92 9.1,8.31 8.82,8.59L6.62,10.79Z" fill="#3b82f6" /></svg>
          <span class="setting-label">手机号</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.phone }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 真实姓名 -->
      <div class="setting-item" @click="openEdit('name', '真实姓名', profile.name)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" fill="#3b82f6" /></svg>
          <span class="setting-label">真实姓名</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.name }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 所在城市 -->
      <div class="setting-item" @click="openEdit('city', '所在城市', profile.city)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5M12,2A7,7 0 0,0 5,9C5,14.25 12,22 12,22C12,22 19,14.25 19,9A7,7 0 0,0 12,2Z" fill="#3b82f6" /></svg>
          <span class="setting-label">所在城市</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.city }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 所属机构 -->
      <div class="setting-item" @click="openEdit('organization', '所属机构', profile.organization)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M12,3L1,9L12,15L21,10.09V17H23V9M5,13.18V17.18L12,21L19,17.18V13.18L12,17L5,13.18Z" fill="#3b82f6" /></svg>
          <span class="setting-label">所属机构</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.organization }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 工作单位 -->
      <div class="setting-item" @click="openEdit('company', '工作单位', profile.company)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" fill="#3b82f6" /></svg>
          <span class="setting-label">工作单位</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.company }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 语言设置 -->
      <div class="setting-item" @click="openEdit('language', '语言设置', profile.language)">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M12.87,15.07L10.33,12.56L10.36,12.53C12.1,10.59 13.34,8.36 14.07,6H17V4H10V2H8V4H1V6H12.17C11.5,7.92 10.44,9.75 9,11.35C8.07,10.32 7.3,9.19 6.69,8H4.69C5.42,9.63 6.42,11.17 7.67,12.56L2.58,17.58L4,19L9,14L12.11,17.11L12.87,15.07M18.5,10H16.5L12,22H14L15.12,19H19.87L21,22H23L18.5,10M15.88,17L17.5,12.67L19.12,17H15.88Z" fill="#3b82f6" /></svg>
          <span class="setting-label">语言设置</span>
        </div>
        <div class="setting-right">
          <span class="setting-value">{{ profile.language || '中文' }}</span>
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 隐私政策 -->
      <div class="setting-item">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M12,1L3,5V11C3,16.55 6.84,21.74 12,23C17.16,21.74 21,16.55 21,11V5L12,1Z" fill="#3b82f6" /></svg>
          <span class="setting-label">隐私政策</span>
        </div>
        <div class="setting-right">
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <!-- 关于我们 -->
      <div class="setting-item">
        <div class="setting-left">
          <svg viewBox="0 0 24 24" class="setting-icon"><path d="M13,9H11V7H13M13,17H11V11H13M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" fill="#3b82f6" /></svg>
          <span class="setting-label">关于我们</span>
        </div>
        <div class="setting-right">
          <svg viewBox="0 0 24 24" class="arrow-icon"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" fill="#999" /></svg>
        </div>
      </div>

      <button class="logout-btn" @click="handleLogout">退出登录</button>
    </main>

    <!-- 编辑弹窗 -->
    <div v-if="editModal.show" class="modal-mask" @click.self="editModal.show = false">
      <div class="modal">
        <div class="modal-title">修改{{ editModal.label }}</div>
        <input class="modal-input" v-model="editModal.value" :placeholder="'请输入' + editModal.label" />
        <div class="modal-actions">
          <button class="modal-cancel" @click="editModal.show = false">取消</button>
          <button class="modal-confirm" @click="confirmEdit">确认</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { userApi, uploadApi } from '@/api'

const router = useRouter()
const imageBaseUrl = 'http://localhost:8080'
const avatarInput = ref<HTMLInputElement | null>(null)
const avatarUrl = ref('')

const profile = reactive({
  nickname: '', name: '', phone: '', city: '',
  organization: '', company: '', language: '', avatar: ''
})

const editModal = reactive({ show: false, field: '', label: '', value: '' })

onMounted(() => {
  const stored = localStorage.getItem('userInfo')
  if (stored) {
    const info = JSON.parse(stored)
    Object.assign(profile, info)
    if (info.avatar) avatarUrl.value = imageBaseUrl + info.avatar
  }
})

function openEdit(field: string, label: string, current: string) {
  editModal.field = field
  editModal.label = label
  editModal.value = current || ''
  editModal.show = true
}

async function confirmEdit() {
  const updated = { ...profile, [editModal.field]: editModal.value }
  await userApi.updateProfile(updated)
  ;(profile as any)[editModal.field] = editModal.value
  const stored = localStorage.getItem('userInfo')
  if (stored) {
    const info = JSON.parse(stored)
    info[editModal.field] = editModal.value
    localStorage.setItem('userInfo', JSON.stringify(info))
  }
  editModal.show = false
}

function triggerAvatarUpload() {
  avatarInput.value?.click()
}

async function handleAvatarChange(e: Event) {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  const res: any = await uploadApi.uploadFile(file)
  const url = res.data?.url
  if (!url) return
  await userApi.updateProfile({ ...profile, avatar: url })
  avatarUrl.value = imageBaseUrl + url
  profile.avatar = url
  const stored = localStorage.getItem('userInfo')
  if (stored) {
    const info = JSON.parse(stored)
    info.avatar = url
    localStorage.setItem('userInfo', JSON.stringify(info))
  }
}

function handleLogout() {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  router.push('/login')
}
</script>

<style scoped>
.account-settings-page { background-color: #f5f5f5; position: fixed; inset: 0; overflow: hidden; display: flex; flex-direction: column; box-sizing: border-box; }
.top-nav { display: flex; align-items: center; justify-content: space-between; padding: 12px 16px; background: #fff; box-shadow: 0 2px 4px rgba(0,0,0,0.05); flex-shrink: 0; }
.back-btn { background: none; border: none; padding: 0; width: 24px; height: 24px; flex-shrink: 0; display: flex; align-items: center; justify-content: center; cursor: pointer; }
.back-btn svg { width: 24px; height: 24px; fill: #333; }
.nav-title { font-size: 18px; font-weight: 600; color: #333; flex: 1; text-align: center; }
.nav-placeholder { width: 24px; height: 24px; flex-shrink: 0; }
.settings-content { flex: 1; min-height: 0; overflow-y: auto; -webkit-overflow-scrolling: touch; overscroll-behavior-y: contain; padding: 16px; }
.setting-item { background: #fff; padding: 16px; display: flex; justify-content: space-between; align-items: center; margin-bottom: 1px; cursor: pointer; }
.setting-item:first-child { border-radius: 12px 12px 0 0; }
.setting-item:last-of-type { border-radius: 0 0 12px 12px; margin-bottom: 20px; }
.setting-left { display: flex; align-items: center; gap: 12px; }
.setting-icon { width: 20px; height: 20px; }
.setting-label { font-size: 15px; color: #333; }
.setting-right { display: flex; align-items: center; gap: 8px; }
.setting-value { font-size: 14px; color: #666; }
.avatar-preview { width: 40px; height: 40px; border-radius: 50%; object-fit: cover; }
.arrow-icon { width: 20px; height: 20px; }
.logout-btn { width: 100%; padding: 14px; background: #fff; color: #3b82f6; border: 1px solid #3b82f6; border-radius: 12px; font-size: 16px; font-weight: 500; cursor: pointer; }

.modal-mask { position: fixed; inset: 0; background: rgba(0,0,0,0.4); display: flex; align-items: center; justify-content: center; z-index: 999; }
.modal { background: #fff; border-radius: 16px; padding: 24px; width: 80%; max-width: 320px; }
.modal-title { font-size: 16px; font-weight: 600; margin-bottom: 16px; color: #333; }
.modal-input { width: 100%; padding: 10px 12px; border: 1px solid #e0e0e0; border-radius: 8px; font-size: 15px; outline: none; box-sizing: border-box; }
.modal-actions { display: flex; gap: 12px; margin-top: 20px; }
.modal-cancel { flex: 1; padding: 10px; border: 1px solid #e0e0e0; border-radius: 8px; background: #fff; color: #666; font-size: 15px; cursor: pointer; }
.modal-confirm { flex: 1; padding: 10px; border: none; border-radius: 8px; background: #3b82f6; color: #fff; font-size: 15px; cursor: pointer; }
</style>
