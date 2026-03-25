<template>
  <div class="learning-certificate-page">
    <header class="header-shell">
      <div class="top-nav">
        <button class="back-btn" @click="$router.back()" aria-label="&#x8FD4;&#x56DE;">
          <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
        </button>
        <span class="nav-title">&#x5B66;&#x4E60;&#x8BC1;&#x4E66;</span>
        <span class="nav-placeholder"></span>
      </div>
    </header>

    <div class="notice-banner">
      <svg viewBox="0 0 24 24" class="notice-icon">
        <circle cx="12" cy="12" r="9.25" />
        <path d="M8 12.4L10.55 14.95L16.1 9.4" />
      </svg>
      <span>&#x5B8C;&#x6210;&#x5B66;&#x4E60;&#x540E;&#x53EF;&#x4E0B;&#x8F7D;&#x5B66;&#x65F6;&#x8BC1;&#x660E;</span>
    </div>

    <div v-if="loading" class="loading-wrap">
      <p>&#x52A0;&#x8F7D;&#x4E2D;...</p>
    </div>

    <div v-else-if="certificateList.length === 0" class="empty-wrap">
      <p>&#x6682;&#x65E0;&#x7B26;&#x5408;&#x6761;&#x4EF6;&#x7684;&#x8BFE;&#x7A0B;</p>
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
              <span class="stat-item">&#x6A21;&#x62DF;&#x6B21;&#x6570;&#xFF1A;{{ course.mockCount }}&#x6B21;</span>
              <span class="stat-item">&#x5B66;&#x4E60;&#x65F6;&#x957F;&#xFF1A;{{ course.studyTime }}</span>
            </div>
            <div class="score-info">
              <span class="score-label">&#x6700;&#x9AD8;&#x5F97;&#x5206;&#xFF1A;</span>
              <span class="score-value">{{ course.highestScore }}&#x5206;</span>
            </div>
          </div>
        </div>

        <div class="button-wrapper">
          <button
            class="download-btn"
            @click="$router.push('/certificate_detail/' + course.courseEk)"
          >
            &#x4E0B;&#x8F7D;&#x5B66;&#x65F6;&#x8BC1;&#x660E;
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
  position: fixed;
  inset: 0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  background: #eef1f6;
  width: 100%;
  max-width: 390px;
  margin: 0 auto;
  box-sizing: border-box;
}

.header-shell {
  background: #fff;
  border-bottom: 1px solid #edf1f6;
  flex-shrink: 0;
}

.top-nav {
  height: 56px;
  padding: 0 14px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
}

.back-btn {
  background: none;
  border: none;
  width: 24px;
  height: 24px;
  padding: 0;
  cursor: pointer;
}

.back-btn svg {
  width: 24px;
  height: 24px;
  fill: #1f2937;
}

.nav-title {
  flex: 1;
  text-align: center;
  color: #1f2937;
  font-size: 17px;
  font-weight: 700;
  line-height: 1;
}

.nav-placeholder {
  width: 24px;
  height: 24px;
}

.notice-banner {
  margin-top: 8px;
  background: #fdf2e6;
  color: #f29543;
  font-size: 14px;
  font-weight: 500;
  padding: 10px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}

.notice-icon {
  width: 15px;
  height: 15px;
  stroke: currentColor;
  stroke-width: 2;
  fill: none;
  stroke-linecap: round;
  stroke-linejoin: round;
  flex-shrink: 0;
}

.loading-wrap,
.empty-wrap {
  flex: 1;
  min-height: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9aa3af;
  font-size: 14px;
}

.certificate-content {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
  overscroll-behavior-y: contain;
  padding: 10px;
}

.certificate-card {
  background: #fff;
  border-radius: 16px;
  padding: 14px;
}

.card-content {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  margin-bottom: 10px;
}

.course-image {
  width: 92px;
  height: 68px;
  border-radius: 10px;
  object-fit: cover;
  flex-shrink: 0;
}

.course-info {
  flex: 1;
  min-width: 0;
}

.course-name {
  margin: 0 0 10px;
  font-size: 16px;
  line-height: 1.1;
  color: #1f2937;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.course-stats {
  display: flex;
  gap: 8px 14px;
  flex-wrap: wrap;
  margin-bottom: 6px;
}

.stat-item {
  font-size: 14px;
  color: #4b5563;
  line-height: 1.25;
}

.score-info {
  font-size: 14px;
}

.score-label {
  color: #4b5563;
}

.score-value {
  color: #f97316;
  font-weight: 700;
}

.button-wrapper {
  border-top: 1px solid #edf1f6;
  padding-top: 10px;
  display: flex;
  justify-content: flex-end;
}

.download-btn {
  width: 192px;
  height: 44px;
  border: none;
  border-radius: 12px;
  background: #4f72ff;
  color: #fff;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
}
</style>
