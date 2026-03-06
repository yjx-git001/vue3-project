<template>
  <div class="course-study-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程学习</span>
    </header>

    <main class="study-content">
      <div
        v-for="course in courseList"
        :key="course.id"
        class="study-card"
      >
        <div class="card-header">
          <img :src="course.image" :alt="course.title" class="course-image" />
          <div class="course-info">
            <h3 class="course-title">{{ course.title }}</h3>
            <div class="course-tags">
              <span class="tag">{{ course.tag }}</span>
              <span class="hours">{{ course.hours }}课时</span>
            </div>
            <div class="progress-section">
              <div class="progress-bar">
                <div class="progress-fill" :style="{ width: course.progress + '%' }"></div>
              </div>
              <span class="progress-text">{{ course.progress }}%</span>
            </div>
          </div>
        </div>
        <div class="card-footer">
          <span class="study-time" v-if="course.studyTime">已学习时长：{{ course.studyTime }}</span>
          <span class="study-time" v-else>暂未开始学习</span>
          <div class="action-buttons">
            <button v-if="course.canDownload" class="download-btn" @click="$router.push('/certificate_detail/' + course.id)">下载学时证明</button>
            <button class="study-btn" @click="$router.push('/course_content/' + course.id)">去学习</button>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const courseList = ref([
  {
    id: 1,
    title: '港口特种设备检修课程',
    tag: '体系课程',
    hours: 48,
    progress: 40,
    studyTime: '34:09',
    image: '/images/course-banner.png',
    canDownload: false
  },
  {
    id: 2,
    title: '港口特种设备检修课程',
    tag: '单门课程',
    hours: 48,
    progress: 100,
    studyTime: '34:09',
    image: '/images/course-banner.png',
    canDownload: true
  },
  {
    id: 3,
    title: '港口特种设备检修课程',
    tag: '体系课程',
    hours: 48,
    progress: 0,
    studyTime: '',
    image: '/images/course-banner.png',
    canDownload: false
  }
])
</script>

<style scoped>
.course-study-page {
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

.study-content {
  padding: 15px;
}

.study-card {
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
  gap: 8px;
}

.course-title {
  font-size: 15px;
  font-weight: 600;
  color: #333;
  margin: 0;
  line-height: 1.4;
}

.course-tags {
  display: flex;
  gap: 8px;
  align-items: center;
}

.tag {
  background-color: #e3f2fd;
  color: #2196f3;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.hours {
  font-size: 12px;
  color: #999;
}

.progress-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.progress-bar {
  flex: 1;
  height: 6px;
  background-color: #e0e0e0;
  border-radius: 3px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: #3b82f6;
  transition: width 0.3s;
}

.progress-text {
  font-size: 12px;
  color: #3b82f6;
  font-weight: 500;
  min-width: 35px;
  text-align: right;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.study-time {
  font-size: 13px;
  color: #666;
}

.action-buttons {
  display: flex;
  gap: 8px;
}

.download-btn {
  padding: 6px 16px;
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 20px;
  font-size: 13px;
  cursor: pointer;
  white-space: nowrap;
}

.study-btn {
  padding: 6px 20px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 20px;
  font-size: 13px;
  cursor: pointer;
  white-space: nowrap;
}
</style>
