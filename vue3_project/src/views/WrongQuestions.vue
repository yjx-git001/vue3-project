<template>
  <div class="wrong-questions-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">错题记录</span>
    </header>

    <main class="questions-content">
      <div
        v-for="course in questionsList"
        :key="course.id"
        class="questions-card"
      >
        <div class="card-header">
          <img src="/images/course-detail-banner.png" alt="课程图片" class="course-image" />
          <div class="course-info">
            <h3 class="course-title">{{ course.title }}</h3>
            <div class="question-stats">
              <div class="stat-row">
                <div class="stat-item">
                  <span class="stat-label">判断题数量：</span>
                  <span class="stat-value">{{ course.judgeCount }}</span>
                </div>
                <div class="stat-item">
                  <span class="stat-label">单选题数量：</span>
                  <span class="stat-value">{{ course.singleChoiceCount }}</span>
                </div>
              </div>
              <div class="stat-row">
                <div class="stat-item">
                  <span class="stat-label">多选题数量：</span>
                  <span class="stat-value">{{ course.multipleChoiceCount }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="card-footer">
          <button class="view-btn" @click="$router.push('/wrong_question_detail/' + course.id)">查看</button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const questionsList = ref([
  {
    id: 1,
    title: '港口特种设备检修课程',
    judgeCount: 3,
    singleChoiceCount: 2,
    multipleChoiceCount: 3
  },
  {
    id: 2,
    title: '港口特种设备检修课程',
    judgeCount: 3,
    singleChoiceCount: 2,
    multipleChoiceCount: 3
  },
  {
    id: 3,
    title: '港口特种设备检修课程',
    judgeCount: 3,
    singleChoiceCount: 2,
    multipleChoiceCount: 3
  }
])
</script>

<style scoped>
.wrong-questions-page {
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

.questions-content {
  padding: 15px;
}

.questions-card {
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

.question-stats {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stat-row {
  display: flex;
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

.card-footer {
  display: flex;
  justify-content: flex-end;
}

.view-btn {
  padding: 8px 60px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  cursor: pointer;
}
</style>
