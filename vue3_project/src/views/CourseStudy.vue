<template>
  <div class="course-study-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程学习</span>
      <span class="nav-placeholder"></span>
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
              <span class="hours">{{ course.hours }}</span>
            </div>
            <div v-if="course.progress > 0" class="progress-section">
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
import { onMounted, ref } from 'vue'
import { courseApi, courseContentApi, questionApi, studyApi } from '@/api'

const imageBaseUrl = 'http://localhost:8080'

const courseList = ref<Array<{
  id: number
  title: string
  tag: string
  hours: string
  progress: number
  studyTime: string
  image: string
  canDownload: boolean
}>>([])

const parseVideoDurationToSeconds = (value: string) => {
  if (!value) return 0
  const text = String(value).trim()
  if (!text) return 0

  if (/^\d+$/.test(text)) return Number(text)

  if (text.includes(':')) {
    const parts = text.split(':').map(v => Number(v))
    if (parts.some(v => Number.isNaN(v))) return 0
    if (parts.length === 2) return parts[0] * 60 + parts[1]
    if (parts.length === 3) return parts[0] * 3600 + parts[1] * 60 + parts[2]
  }

  return 0
}

const formatSeconds = (seconds: number) => {
  const total = Math.max(0, Math.floor(seconds))
  const h = Math.floor(total / 3600)
  const m = Math.floor((total % 3600) / 60)
  const s = total % 60

  if (h > 0) {
    return `${String(h).padStart(2, '0')}:${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`
  }

  return `${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`
}

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) {
    courseList.value = []
    return
  }

  try {
    const coursesRes: any = await courseApi.getUserCourses()
    const courses: any[] = coursesRes.data || []

    const rows = await Promise.all(
      courses.map(async (course: any) => {
        const ek = Number(course.courseEk)
        if (!Number.isFinite(ek)) return null

        try {
          const [questionRes, videoRes, durationRes]: any[] = await Promise.all([
            questionApi.getList(ek),
            courseContentApi.getVideos(ek),
            studyApi.getCourseDuration(ek)
          ])

          const questions: any[] = questionRes.data || []
          if (questions.length === 0) return null

          const videos: Array<{ duration: string }> = videoRes.data || []
          const totalSeconds = videos.reduce((sum, v) => sum + parseVideoDurationToSeconds(v.duration), 0)
          const studySeconds = durationRes.data?.totalDuration || 0
          const progress = totalSeconds > 0 ? Math.min(100, Math.round((studySeconds / totalSeconds) * 100)) : 0

          const canDownload = totalSeconds > 0 && studySeconds >= totalSeconds

          return {
            id: ek,
            title: course.courseName || '',
            tag: Number(course.courseCategory) === 2 ? '体系课程' : '单门课程',
            hours: course.courseTime ? `${course.courseTime}课时` : '0课时',
            progress,
            studyTime: studySeconds > 0 ? formatSeconds(studySeconds) : '',
            image: course.courseImage ? imageBaseUrl + course.courseImage : '/images/course-banner.png',
            canDownload
          }
        } catch {
          return null
        }
      })
    )

    courseList.value = rows.filter(Boolean) as Array<{
      id: number
      title: string
      tag: string
      hours: string
      progress: number
      studyTime: string
      image: string
      canDownload: boolean
    }>
  } catch {
    courseList.value = []
  }
})
</script>

<style scoped>
.course-study-page {
  background-color: #f5f5f5;
  position: fixed;
  inset: 0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  flex-shrink: 0;
}

.back-btn {
  background: none;
  border: none;
  padding: 0;
  width: 24px;
  height: 24px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
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
  height: 24px;
  flex-shrink: 0;
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
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
  overscroll-behavior-y: contain;
  padding: 10px 12px 16px;
}

.study-card {
  background-color: #fff;
  border-radius: 14px;
  padding: 14px;
  margin-bottom: 12px;
  box-shadow: 0 1px 2px rgba(17, 24, 39, 0.04);
}

.card-header {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
}

.course-image {
  width: 96px;
  height: 72px;
  border-radius: 8px;
  object-fit: cover;
  flex-shrink: 0;
}

.course-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 0;
}

.course-title {
  font-size: 17px;
  font-weight: 600;
  color: #333;
  margin: 0;
  line-height: 1.3;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.course-tags {
  display: flex;
  gap: 8px;
  align-items: center;
}

.tag {
  background-color: #eaf2ff;
  color: #4f72ff;
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 12px;
  line-height: 1;
}

.hours {
  background-color: #eaf2ff;
  color: #4f72ff;
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 12px;
  line-height: 1;
}

.progress-section {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 2px;
}

.progress-bar {
  flex: 1;
  height: 8px;
  background-color: #e8ebf2;
  border-radius: 999px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: #3b82f6;
  border-radius: 999px;
  transition: width 0.3s;
}

.progress-text {
  font-size: 14px;
  color: #3b82f6;
  font-weight: 600;
  min-width: 44px;
  text-align: right;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: nowrap;
  gap: 10px;
  border-top: 1px solid #f1f3f7;
  padding-top: 12px;
}

.study-time {
  font-size: 13px;
  color: #6b7280;
  line-height: 1.4;
  white-space: nowrap;
  flex-shrink: 0;
}

.action-buttons {
  display: flex;
  gap: 10px;
  margin-left: auto;
  flex-shrink: 0;
}

.download-btn {
  padding: 6px 16px;
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 10px;
  font-size: 13px;
  cursor: pointer;
  white-space: nowrap;
}

.study-btn {
  padding: 6px 20px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 13px;
  cursor: pointer;
  white-space: nowrap;
}
</style>
