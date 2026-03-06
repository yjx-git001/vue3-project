<template>
  <div class="learning-certificate-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">学习证书</span>
    </header>

    <!-- 提示信息 -->
    <div class="notice-banner">
      <svg viewBox="0 0 24 24" class="notice-icon">
        <path d="M13,13H11V7H13M13,17H11V15H13M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" />
      </svg>
      <span>完成学习后可下载学时证明</span>
    </div>

    <main class="certificate-content">
      <div
        v-for="course in certificateList"
        :key="course.id"
        class="certificate-card"
      >
        <div class="card-content">
          <img :src="course.image" :alt="course.courseName" class="course-image" />
          <div class="course-info">
            <h3 class="course-name">{{ course.courseName }}</h3>
            <div class="course-stats">
              <span class="stat-item">模拟次数：{{ course.mockCount }}次</span>
              <span class="stat-item">学习时长：{{ course.studyTime }}</span>
            </div>
            <div class="score-info">
              <span class="score-label">最高得分：</span>
              <span class="score-value">{{ course.highestScore }}分</span>
            </div>
          </div>
        </div>
        <button
          class="download-btn"
          @click="$router.push('/certificate-detail/' + course.id)"
        >
          下载学时证明
        </button>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const certificateList = ref([
  {
    id: 1,
    courseName: '港口特种设备检修课程',
    mockCount: 3,
    studyTime: '23:08',
    highestScore: 100,
    image: '/images/course-detail-banner.png'
  }
])
</script>

<style scoped>
.learning-certificate-page {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.back-btn {
  background: none;
  border: none;
  padding: 0;
  cursor: pointer;
}

.back-btn svg {
  width: 24px;
  height: 24px;
  fill: #333;
}

.nav-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  flex: 1;
  text-align: center;
}

.notice-banner {
  background-color: #fff3e0;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  color: #ff9800;
  font-size: 14px;
}

.notice-icon {
  width: 18px;
  height: 18px;
  fill: currentColor;
  flex-shrink: 0;
}

.certificate-content {
  padding: 16px;
}

.certificate-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 12px;
}

.card-content {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
}

.course-image {
  width: 100px;
  height: 75px;
  border-radius: 8px;
  object-fit: cover;
  flex-shrink: 0;
}

.course-info {
  flex: 1;
}

.course-name {
  font-size: 15px;
  font-weight: 600;
  color: #333;
  margin: 0 0 8px 0;
  line-height: 1.4;
}

.course-stats {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-bottom: 8px;
}

.stat-item {
  font-size: 13px;
  color: #666;
}

.score-info {
  font-size: 13px;
}

.score-label {
  color: #666;
}

.score-value {
  color: #ff4d4f;
  font-weight: 600;
  font-size: 14px;
}

.download-btn {
  width: 100%;
  padding: 12px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
}
</style>
