<template>
  <div class="mock-exam-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">模拟考试</span>
    </header>

    <main class="exam-content">
      <div
        v-for="course in examList"
        :key="course.id"
        class="exam-card"
      >
        <div class="card-header">
          <img :src="course.image" :alt="course.title" class="course-image" />
          <div class="course-info">
            <h3 class="course-title">{{ course.title }}</h3>
            <div class="exam-stats">
              <div class="stat-item">
                <span class="stat-label">模拟次数：</span>
                <span class="stat-value">{{ course.mockCount }}次</span>
              </div>
              <div class="stat-item">
                <span class="stat-label">学习时长：</span>
                <span class="stat-value">{{ course.studyTime }}</span>
              </div>
            </div>
            <div class="score-section">
              <span class="score-label">最高得分：</span>
              <span class="score-value">{{ course.highestScore }}分</span>
            </div>
          </div>
        </div>
        <div class="card-footer">
          <button class="mock-btn" @click="$router.push('/mock_exam_detail/' + course.id)">模拟考试</button>
          <button class="practice-btn" @click="$router.push('/exam_practice/' + course.id)">考题训练</button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const examList = ref([
  {
    id: 1,
    title: '港口特种设备检修课程',
    mockCount: 3,
    studyTime: '23.08',
    highestScore: 100,
    image: '/images/course-banner.png'
  },
  {
    id: 2,
    title: '港口特种设备检修课程',
    mockCount: 3,
    studyTime: '23.08',
    highestScore: 100,
    image: '/images/course-banner.png'
  },
  {
    id: 3,
    title: '港口特种设备检修课程',
    mockCount: 3,
    studyTime: '23.08',
    highestScore: 100,
    image: '/images/course-banner.png'
  }
])
</script>

<style scoped>
.mock-exam-page {
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

.nav-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.badge {
  background-color: #3b82f6;
  color: #fff;
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 12px;
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

.exam-content {
  padding: 15px;
}

.exam-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 12px;
}

.card-header {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
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
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.course-title {
  font-size: 15px;
  font-weight: 600;
  color: #333;
  margin: 0;
  line-height: 1.4;
}

.exam-stats {
  display: flex;
  flex-direction: row;
  gap: 16px;
}

.stat-item {
  font-size: 13px;
}

.stat-label {
  color: #666;
}

.stat-value {
  color: #333;
  font-weight: 500;
}

.score-section {
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

.card-footer {
  display: flex;
  gap: 12px;
}

.mock-btn {
  flex: 1;
  padding: 10px;
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 10px;
  font-size: 14px;
  cursor: pointer;
}

.practice-btn {
  flex: 1;
  padding: 10px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  cursor: pointer;
}
</style>
