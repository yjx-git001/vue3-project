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
        <circle cx="12" cy="12" r="9.25" />
        <path d="M8 12.4L10.55 14.95L16.1 9.4" />
      </svg>
      <span>完成学习后可下载学时证明</span>
    </div>

    <div v-if="loading" class="loading-wrap">
      <p>加载中...</p>
    </div>

    <div v-else-if="certificateList.length === 0" class="empty-wrap">
      <p>暂无符合条件的课程</p>
    </div>

    <main v-else class="certificate-content">
      <div
        v-for="course in certificateList"
        :key="course.courseEk"
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
        <div class="button-wrapper">
          <button
            class="download-btn"
            @click="$router.push('/certificate_detail/' + course.courseEk)"
          >
            下载学时证明
          </button>
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
const certificateList = ref<Array<{
  courseEk: number
  courseName: string
  mockCount: number
  studyTime: string
  highestScore: number
  image: string
}>>([])

const formatDuration = (seconds: number) => {
  const total = Math.max(0, Math.floor(seconds))
  const m = Math.floor(total / 60)
  const s = total % 60
  return `${m}:${s.toString().padStart(2, '0')}`
}

onMounted(async () => {
  try {
    const res: any = await courseApi.getUserCourses()
    const courses: any[] = res.data || []

    const rows = await Promise.all(
      courses.map(async (course: any) => {
        const courseEk = Number(course.courseEk)
        let mockCount = 0
        let highestScore = 0
        let totalDuration = 0

        try {
          const statsRes: any = await mockExamApi.getStats(courseEk)
          mockCount = statsRes.data?.mockCount || 0
          highestScore = statsRes.data?.highestScore || 0
        } catch {}

        try {
          const durRes: any = await studyApi.getCourseDuration(courseEk)
          totalDuration = durRes.data?.totalDuration || 0
        } catch {}

        if (totalDuration < 180) return null
        if (mockCount <= 0) return null
        if (highestScore < 60) return null

        return {
          courseEk,
          courseName: course.courseName || '',
          mockCount,
          studyTime: formatDuration(totalDuration),
          highestScore,
          image: course.courseImage ? imageBaseUrl + course.courseImage : '/images/course-detail-banner.png'
        }
      })
    )

    certificateList.value = rows.filter(Boolean) as Array<{
      courseEk: number
      courseName: string
      mockCount: number
      studyTime: string
      highestScore: number
      image: string
    }>
  } catch {
    certificateList.value = []
  } finally {
    loading.value = false
  }
})
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
  background-color: #fff1e5;
  padding: 10px 16px;
  display: flex;
  align-items: center;
  gap: 6px;
  color: #f28a1e;
  font-size: 15px;
  line-height: 1.2;
  margin-top: 8px;
  margin-bottom: 8px;
}

.notice-icon {
  width: 14px;
  height: 14px;
  stroke: currentColor;
  stroke-width: 2;
  fill: none;
  stroke-linecap: round;
  stroke-linejoin: round;
  flex-shrink: 0;
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

.certificate-content {
  padding: 0 16px 16px;
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
  flex-direction: row;
  gap: 16px;
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

.button-wrapper {
  display: flex;
  justify-content: flex-end;
  border-top: 1px solid #f0f0f0;
  margin-top: 12px;
  padding-top: 12px;
}

.download-btn {
  padding: 12px 24px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  box-shadow: none;
  outline: none;
  -webkit-appearance: none;
  appearance: none;
}
</style>
