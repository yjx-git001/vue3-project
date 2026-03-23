<template>
  <div class="mock-exam-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">模拟考试</span>
      <span class="nav-placeholder"></span>
    </header>

    <!-- 加载中 -->
    <div v-if="loading" class="loading-wrap">
      <p>加载中...</p>
    </div>

    <!-- 无已购课程 -->
    <div v-else-if="courseList.length === 0" class="empty-wrap">
      <p class="empty-text">暂无已购课程</p>
    </div>

    <!-- 课程列表 -->
    <main v-else class="exam-content">
      <div
        v-for="course in courseList"
        :key="course.courseEk"
        class="exam-card"
      >
        <div class="card-header">
          <img
            :src="course.courseImage ? imageBaseUrl + course.courseImage : '/images/course-detail-banner.png'"
            :alt="course.courseName"
            class="course-image"
          />
          <div class="course-info">
            <h3 class="course-title">{{ course.courseName }}</h3>
            <div class="exam-stats">
              <div class="stat-row">
                <span class="stat-item">模拟次数：{{ statsMap[course.courseEk]?.mockCount ?? 0 }}次</span>
                <span class="stat-item">学习时长：{{ formatDuration(durationMap[course.courseEk] || 0) }}</span>
              </div>
              <div class="stat-row">
                <span class="stat-item">最高得分：<span class="score-val">{{ statsMap[course.courseEk]?.highestScore ?? 0 }}分</span></span>
              </div>
            </div>
          </div>
        </div>
        <div class="card-footer">
          <button
            class="mock-btn"
            @click="$router.push('/mock_exam_detail/' + course.courseEk)"
          >模拟考试</button>
          <button
            class="practice-btn"
            @click="$router.push('/exam_practice/' + course.courseEk + '?type=1')"
          >考题训练</button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { courseApi, mockExamApi, studyApi } from '@/api'

const imageBaseUrl = 'http://localhost:8080'
const loading = ref(true)
const courseList = ref<any[]>([])
const statsMap = ref<Record<number, { mockCount: number; highestScore: number }>>({})
const durationMap = ref<Record<number, number>>({}) // courseEk -> 秒

const formatDuration = (seconds: number) => {
  const h = Math.floor(seconds / 3600)
  const m = Math.floor((seconds % 3600) / 60)
  if (h > 0) return `${h}h${m}m`
  return `${m}分钟`
}

onMounted(async () => {
  try {
    const res: any = await courseApi.getUserCourses()
    courseList.value = res.data || []
    // 并发获取每个课程的统计和学习时长
    await Promise.all(
      courseList.value.map(async (course: any) => {
        const ek = course.courseEk
        try {
          const statsRes: any = await mockExamApi.getStats(ek)
          statsMap.value[ek] = statsRes.data || { mockCount: 0, highestScore: 0 }
        } catch {
          statsMap.value[ek] = { mockCount: 0, highestScore: 0 }
        }
        try {
          const durRes: any = await studyApi.getCourseDuration(ek)
          durationMap.value[ek] = durRes.data?.totalDuration || 0
        } catch {
          durationMap.value[ek] = 0
        }
      })
    )
  } catch {
    courseList.value = []
  } finally {
    loading.value = false
  }
})
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

.nav-placeholder {
  width: 24px;
}

.loading-wrap,
.empty-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
  gap: 12px;
  color: #999;
  font-size: 15px;
}

.empty-text {
  color: #999;
  font-size: 15px;
  margin: 0;
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
  margin-bottom: 14px;
}

.course-image {
  width: 100px;
  height: 70px;
  border-radius: 8px;
  object-fit: cover;
  flex-shrink: 0;
}

.course-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
  justify-content: center;
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
  flex-direction: column;
  gap: 4px;
  margin-top: 6px;
}

.stat-row {
  display: flex;
  gap: 16px;
}

.stat-item {
  font-size: 13px;
  color: #666;
}

.score-val {
  color: #ff4d4f;
  font-weight: 600;
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
