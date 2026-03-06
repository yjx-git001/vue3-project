<template>
  <div class="profile-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <div class="logo-container">
        <img src="/logo.png" alt="Logo" class="logo" />
        <span class="app-title">港口学堂</span>
      </div>
    </header>

    <main class="profile-content">
      <!-- 用户信息卡片 -->
      <section class="user-card">
        <div class="user-info">
          <div class="avatar">
            <svg viewBox="0 0 24 24"><path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" fill="white" /></svg>
          </div>
          <div class="user-details">
            <h2 class="username">爱学习的老铁</h2>
            <p class="join-date">入学日期：2025.12.30</p>
          </div>
          <div class="user-id">ID: 26372</div>
        </div>
        <div class="user-stats">
          <div class="stat-item">
            <div class="stat-value">0</div>
            <div class="stat-label">今日学时</div>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <div class="stat-value">68 小时</div>
            <div class="stat-label">累计学时</div>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <div class="stat-value">4</div>
            <div class="stat-label">获取证书</div>
          </div>
        </div>
      </section>

      <!-- 我的课程 -->
      <section class="my-courses">
        <h3 class="section-title">我的课程</h3>
        <div class="course-actions">
          <div class="action-item" @click="$router.push('/course-study')">
            <div class="action-icon books">
              <svg viewBox="0 0 24 24"><path d="M19,2L14,6.5V17.5L19,13V2M6.5,5C4.55,5 2.45,5.4 1,6.5V21.16C1,21.41 1.25,21.66 1.5,21.66C1.6,21.66 1.65,21.59 1.75,21.59C3.1,20.94 5.05,20.5 6.5,20.5C8.45,20.5 10.55,20.9 12,22C13.35,21.15 15.8,20.5 17.5,20.5C19.15,20.5 20.85,20.81 22.25,21.56C22.35,21.61 22.4,21.59 22.5,21.59C22.75,21.59 23,21.34 23,21.09V6.5C22.4,6.05 21.75,5.75 21,5.5V7.5L21,13V19C19.9,18.65 18.7,18.5 17.5,18.5C15.8,18.5 13.35,19.15 12,20V6.5C10.55,5.4 8.45,5 6.5,5Z" /></svg>
            </div>
            <span class="action-label">课程学习</span>
          </div>
          <div class="action-item" @click="$router.push('/mock-exam')">
            <div class="action-icon exam">
              <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20M10,19L12,15H9V10L7,14H10V19Z" /></svg>
            </div>
            <span class="action-label">模拟考试</span>
          </div>
          <div class="action-item" @click="$router.push('/wrong-questions')">
            <div class="action-icon notes">
              <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20M15,18V16H6V18H15M18,14V12H6V14H18Z" /></svg>
            </div>
            <span class="action-label">错题记录</span>
          </div>
        </div>
      </section>

      <!-- 功能菜单 -->
      <section class="menu-section">
        <div class="menu-grid">
          <div class="menu-item" @click="$router.push('/course-orders')">
            <div class="menu-icon blue">
              <svg viewBox="0 0 24 24"><path d="M19,3H14.82C14.4,1.84 13.3,1 12,1C10.7,1 9.6,1.84 9.18,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M12,3A1,1 0 0,1 13,4A1,1 0 0,1 12,5A1,1 0 0,1 11,4A1,1 0 0,1 12,3Z" /></svg>
            </div>
            <span class="menu-label">课程订单</span>
          </div>
          <div class="menu-item" @click="$router.push('/learning-certificate')">
            <div class="menu-icon purple">
              <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" /></svg>
            </div>
            <span class="menu-label">学习证书</span>
          </div>
        </div>
        <div class="menu-grid">
          <div class="menu-item" @click="$router.push('/coupon-list')">
            <div class="menu-icon purple">
              <svg viewBox="0 0 24 24"><path d="M21.47,4.35L20.13,3.79V12.82L22.56,6.96C22.97,5.94 22.5,4.77 21.47,4.35M1.97,8.05L6.93,20C7.24,20.77 7.97,21.24 8.74,21.26C9,21.26 9.27,21.21 9.53,21.1L16.9,18.05C17.65,17.74 18.11,17 18.13,16.26C18.14,16 18.09,15.71 18,15.45L13,3.5C12.71,2.73 11.97,2.26 11.19,2.25C10.93,2.25 10.67,2.31 10.42,2.4L3.06,5.45C2.04,5.87 1.55,7.04 1.97,8.05M18.12,4.25A2,2 0 0,0 16.12,2.25H14.67L18.12,10.59" /></svg>
            </div>
            <span class="menu-label">优惠卡券</span>
          </div>
          <div class="menu-item" @click="$router.push('/account-settings')">
            <div class="menu-icon blue">
              <svg viewBox="0 0 24 24"><path d="M12,15.5A3.5,3.5 0 0,1 8.5,12A3.5,3.5 0 0,1 12,8.5A3.5,3.5 0 0,1 15.5,12A3.5,3.5 0 0,1 12,15.5M19.43,12.97C19.47,12.65 19.5,12.33 19.5,12C19.5,11.67 19.47,11.34 19.43,11L21.54,9.37C21.73,9.22 21.78,8.95 21.66,8.73L19.66,5.27C19.54,5.05 19.27,4.96 19.05,5.05L16.56,6.05C16.04,5.66 15.5,5.32 14.87,5.07L14.5,2.42C14.46,2.18 14.25,2 14,2H10C9.75,2 9.54,2.18 9.5,2.42L9.13,5.07C8.5,5.32 7.96,5.66 7.44,6.05L4.95,5.05C4.73,4.96 4.46,5.05 4.34,5.27L2.34,8.73C2.21,8.95 2.27,9.22 2.46,9.37L4.57,11C4.53,11.34 4.5,11.67 4.5,12C4.5,12.33 4.53,12.65 4.57,12.97L2.46,14.63C2.27,14.78 2.21,15.05 2.34,15.27L4.34,18.73C4.46,18.95 4.73,19.03 4.95,18.95L7.44,17.94C7.96,18.34 8.5,18.68 9.13,18.93L9.5,21.58C9.54,21.82 9.75,22 10,22H14C14.25,22 14.46,21.82 14.5,21.58L14.87,18.93C15.5,18.67 16.04,18.34 16.56,17.94L19.05,18.95C19.27,19.03 19.54,18.95 19.66,18.73L21.66,15.27C21.78,15.05 21.73,14.78 21.54,14.63L19.43,12.97Z" /></svg>
            </div>
            <span class="menu-label">账号设置</span>
          </div>
        </div>
      </section>

      <!-- 客服区域 -->
      <section class="service-banner">
        <div class="service-content">
          <div class="robot-icon">
            <svg viewBox="0 0 24 24"><path d="M12,2A2,2 0 0,1 14,4C14,4.74 13.6,5.39 13,5.73V7H14A7,7 0 0,1 21,14H22A1,1 0 0,1 23,15V18A1,1 0 0,1 22,19H21V20A2,2 0 0,1 19,22H5A2,2 0 0,1 3,20V19H2A1,1 0 0,1 1,18V15A1,1 0 0,1 2,14H3A7,7 0 0,1 10,7H11V5.73C10.4,5.39 10,4.74 10,4A2,2 0 0,1 12,2M7.5,13A2.5,2.5 0 0,0 5,15.5A2.5,2.5 0 0,0 7.5,18A2.5,2.5 0 0,0 10,15.5A2.5,2.5 0 0,0 7.5,13M16.5,13A2.5,2.5 0 0,0 14,15.5A2.5,2.5 0 0,0 16.5,18A2.5,2.5 0 0,0 19,15.5A2.5,2.5 0 0,0 16.5,13Z" fill="#00838f" /></svg>
          </div>
          <span class="service-text">有任何问题联系我～</span>
        </div>
        <button class="contact-btn" @click="$router.push('/chat-support')">
          <svg viewBox="0 0 24 24"><path d="M20,2H4A2,2 0 0,0 2,4V22L6,18H20A2,2 0 0,0 22,16V4A2,2 0 0,0 20,2M6,9H18V11H6M14,14H6V12H14M18,8H6V6H18" /></svg>
        </button>
      </section>

      <!-- 推荐课程 -->
      <section class="recommend-section">
        <h3 class="section-title">推荐课程</h3>
        <div class="course-grid">
          <div
            v-for="course in recommendCourses"
            :key="course.id"
            class="course-card"
            @click="$router.push('/course/' + course.id)"
          >
            <div class="course-image-wrapper">
              <img :src="course.image" :alt="course.title" class="course-image" />
              <span class="course-tag">{{ course.tag }}</span>
            </div>
            <div class="course-info">
              <h4 class="course-title">{{ course.title }}</h4>
            </div>
          </div>
        </div>
      </section>
    </main>

    <!-- 底部导航 -->
    <footer class="bottom-nav">
      <div class="nav-item" @click="$router.push('/')">
        <svg viewBox="0 0 24 24"><path d="M10,20V14H14V20H19V12H22L12,3L2,12H5V20H10Z" /></svg>
        <span>首页</span>
      </div>
      <div class="nav-item" @click="$router.push('/courses')">
        <svg viewBox="0 0 24 24"><path d="M12,3L1,9L12,15L21,9V17H23V9M5,13.18V17.18L12,21L19,17.18V13.18L12,17L5,13.18Z" /></svg>
        <span>课程</span>
      </div>
      <div class="nav-item active">
        <svg viewBox="0 0 24 24"><path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" /></svg>
        <span>个人中心</span>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const recommendCourses = ref([
  {
    id: 1,
    title: '港口特种设备检修课程',
    tag: '体系课程',
    image: '/images/course-banner.png'
  },
  {
    id: 2,
    title: '港口特种设备检修课程',
    tag: '热门课程',
    image: '/images/course-banner.png'
  }
])
</script>

<style scoped>
.profile-page {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 70px;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  background-color: #fff;
}

.logo-container {
  display: flex;
  align-items: center;
}

.logo {
  width: 30px;
  height: 30px;
  margin-right: 8px;
}

.app-title {
  font-size: 18px;
  font-weight: bold;
}

.nav-actions {
  display: flex;
  gap: 12px;
}

.icon-btn {
  background: none;
  border: none;
  padding: 0;
  cursor: pointer;
}

.icon-btn svg {
  width: 22px;
  height: 22px;
  fill: #666;
}

.profile-content {
  padding: 15px;
}

.user-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 15px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.user-info {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  position: relative;
}

.avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  margin-right: 15px;
  border: 2px solid #e0e0e0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar svg {
  width: 36px;
  height: 36px;
  fill: white;
}

.user-details {
  flex: 1;
}

.username {
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 6px 0;
  color: #333;
}

.join-date {
  font-size: 13px;
  color: #666;
  margin: 0;
}

.user-id {
  background-color: #e3f2fd;
  color: #2196f3;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 13px;
  font-weight: 500;
}

.user-stats {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.stat-item {
  text-align: center;
  flex: 1;
}

.stat-value {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 4px;
  color: #333;
}

.stat-label {
  font-size: 12px;
  color: #999;
}

.stat-divider {
  width: 1px;
  height: 30px;
  background-color: #e0e0e0;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  margin: 0 0 15px 0;
  color: #333;
}

.my-courses {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 15px;
}

.course-actions {
  display: flex;
  justify-content: space-around;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
}

.action-icon {
  width: 50px;
  height: 50px;
  margin-bottom: 8px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.action-icon.books {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.action-icon.exam {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.action-icon.notes {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.action-icon svg {
  width: 28px;
  height: 28px;
  fill: white;
}

.action-label {
  font-size: 13px;
  color: #666;
}

.menu-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 15px;
}

.menu-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.menu-item {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  transition: transform 0.2s;
}

.menu-item:active {
  transform: scale(0.98);
}

.menu-icon {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu-icon.blue {
  background-color: #e3f2fd;
}

.menu-icon.purple {
  background-color: #f3e5f5;
}

.menu-icon svg {
  width: 22px;
  height: 22px;
}

.menu-icon.blue svg {
  fill: #2196f3;
}

.menu-icon.purple svg {
  fill: #9c27b0;
}

.menu-label {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.service-banner {
  background: linear-gradient(135deg, #e0f7fa 0%, #b2ebf2 100%);
  border-radius: 12px;
  padding: 16px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.service-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.robot-icon {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.robot-icon svg {
  width: 100%;
  height: 100%;
}

.service-text {
  font-size: 14px;
  color: #00838f;
  font-weight: 500;
}

.contact-btn {
  background-color: #fff;
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.contact-btn svg {
  width: 20px;
  height: 20px;
  fill: #00838f;
}

.recommend-section {
  margin-bottom: 20px;
}

.course-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.course-card {
  background-color: #fff;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
}

.course-image-wrapper {
  position: relative;
}

.course-image {
  width: 100%;
  height: 100px;
  object-fit: cover;
}

.course-tag {
  position: absolute;
  top: 8px;
  left: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
}

.course-info {
  padding: 10px;
}

.course-title {
  font-size: 13px;
  font-weight: 500;
  color: #333;
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-around;
  background-color: #fff;
  padding: 10px 0;
  border-top: 1px solid #eee;
  z-index: 100;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 12px;
  color: #999;
  cursor: pointer;
  transition: color 0.3s;
}

.nav-item.active {
  color: #3b82f6;
}

.nav-item svg {
  width: 24px;
  height: 24px;
  margin-bottom: 4px;
  fill: currentColor;
}
</style>
