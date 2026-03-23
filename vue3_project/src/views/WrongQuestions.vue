<template>
  <div class="wrong-questions-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">错题记录</span>
      <span class="nav-placeholder"></span>
    </header>

    <!-- 加载中 -->
    <div v-if="loading" class="loading-wrap">
      <p>加载中...</p>
    </div>

    <!-- 空状态 -->
    <div v-else-if="courseList.length === 0" class="empty-wrap">
      <p>暂无错题记录</p>
    </div>

    <!-- 课程列表 -->
    <main v-else class="questions-content">
      <div
        v-for="course in courseList"
        :key="course.courseEk"
        class="questions-card"
      >
        <div class="card-header">
          <img
            :src="course.courseImage ? imageBaseUrl + course.courseImage : '/images/course-detail-banner.png'"
            alt="课程图片"
            class="course-image"
          />
          <div class="course-info">
            <h3 class="course-title">{{ course.courseName }}</h3>
            <div class="question-stats">
              <div class="stat-row">
                <div class="stat-item">
                  <span class="stat-label">单选题：</span>
                  <span class="stat-value">{{ course.singleCount }}</span>
                </div>
                <div class="stat-item">
                  <span class="stat-label">多选题：</span>
                  <span class="stat-value">{{ course.multipleCount }}</span>
                </div>
              </div>
              <div class="stat-row">
                <div class="stat-item">
                  <span class="stat-label">判断题：</span>
                  <span class="stat-value">{{ course.judgeCount }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="card-footer">
          <button class="view-btn" @click="$router.push('/wrong_question_detail/' + course.courseEk)">查看</button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { mockExamApi } from '@/api'

const imageBaseUrl = 'http://localhost:8080'
const loading = ref(true)
const courseList = ref<any[]>([])

onMounted(async () => {
  try {
    const res: any = await mockExamApi.getWrongQuestionCourses()
    courseList.value = res.data || []
  } catch {
    courseList.value = []
  } finally {
    loading.value = false
  }
})
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

.nav-placeholder {
  width: 24px;
}

.loading-wrap,
.empty-wrap {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
  color: #999;
  font-size: 15px;
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
