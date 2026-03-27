<template>
  <div class="profile-page">
    <main class="profile-content">
      <section class="hero-section">
        <header class="top-nav">
          <div class="logo-container">
            <img src="/logo.png" alt="Logo" class="logo" />
            <span class="app-title">港口学堂</span>
          </div>
        </header>

        <section class="user-card">
          <div class="user-info">
            <div class="avatar">
              <img v-if="userInfo.avatar" :src="imageBaseUrl + userInfo.avatar" class="avatar-img" />
              <svg v-else viewBox="0 0 24 24"><path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" fill="white" /></svg>
            </div>
            <div class="user-details">
              <h2 class="username">{{ userInfo.nickname || userInfo.phone }}</h2>
              <p class="join-date">入学日期：{{ userInfo.createdAt }}</p>
            </div>
            <div class="user-id">ID: {{ userInfo.id }}</div>
          </div>
          <div class="user-stats">
            <div class="stat-item">
              <div class="stat-value">{{ formatHours(todayDuration) }}</div>
              <div class="stat-label">今日学时</div>
            </div>
            <div class="stat-item">
              <div class="stat-value">{{ formatHours(totalDuration) }}</div>
              <div class="stat-label">累计学时</div>
            </div>
            <div class="stat-item">
              <div class="stat-value">{{ certificateCount }}</div>
              <div class="stat-label">获取证书</div>
            </div>
          </div>
        </section>
      </section>

      <!-- 我的课程 -->
      <section class="my-courses">
        <h3 class="section-title">我的课程</h3>
        <div class="course-actions">
          <div class="action-item" @click="$router.push('/course_study')">
            <div class="action-icon">
              <img src="/images/course-study.png" alt="课程学习" />
            </div>
            <span class="action-label">课程学习</span>
          </div>
          <div class="action-item" @click="$router.push('/mock_exam')">
            <div class="action-icon">
              <img src="/images/mock-exam.png" alt="模拟考试" />
            </div>
            <span class="action-label">模拟考试</span>
          </div>
          <div class="action-item" @click="$router.push('/wrong_questions')">
            <div class="action-icon">
              <img src="/images/error-record.png" alt="错题记录" />
            </div>
            <span class="action-label">错题记录</span>
          </div>
        </div>
      </section>

      <!-- 功能菜单 -->
      <section class="menu-section">
        <div class="menu-grid">
          <div class="menu-item" @click="$router.push('/course_orders')">
            <div class="menu-icon">
              <img src="/images/course-order.png" alt="课程订单" />
            </div>
            <span class="menu-label">课程订单</span>
          </div>
          <div class="menu-item" @click="$router.push('/learning_certificate')">
            <div class="menu-icon">
              <img src="/images/study-cert.png" alt="学习证书" />
            </div>
            <span class="menu-label">学习证书</span>
          </div>
        </div>
        <div class="menu-grid">
          <div class="menu-item" @click="$router.push('/coupon_list')">
            <div class="menu-icon">
              <img src="/images/coupon.png" alt="优惠卡券" />
            </div>
            <span class="menu-label">优惠卡券</span>
          </div>
          <div class="menu-item" @click="$router.push('/account_settings')">
            <div class="menu-icon">
              <img src="/images/account-settings.png" alt="账号设置" />
            </div>
            <span class="menu-label">账号设置</span>
          </div>
        </div>
      </section>

      <!-- 客服区域 -->
      <section class="service-banner" @click="$router.push('/chat_support')">
        <img src="/images/feedback-banner.png" alt="客服横幅" class="banner-image" />
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
              <span :class="['course-tag', course.tag === '单门课程' ? 'course-tag--single' : 'course-tag--system']">{{ course.tag }}</span>
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
import { ref, onMounted } from 'vue'
import { courseApi, mockExamApi, studyApi } from '@/api'

const userInfo = ref({ id: '', nickname: '', phone: '', avatar: '', createdAt: '' })
const imageBaseUrl = 'http://localhost:8080'
const todayDuration = ref(0)
const totalDuration = ref(0)
const certificateCount = ref(0)

function formatHours(seconds: number) {
  const h = seconds / 3600
  return `${parseFloat(h.toFixed(1))} 小时`
}

onMounted(async () => {
  const stored = localStorage.getItem('userInfo')
  if (stored) userInfo.value = JSON.parse(stored)
  const token = localStorage.getItem('token')
  if (token) {
    try {
      const [studyRes, coursesRes]: any[] = await Promise.all([
        studyApi.getStats(),
        courseApi.getUserCourses()
      ])
      todayDuration.value = studyRes.data?.todayDuration || 0
      totalDuration.value = studyRes.data?.totalDuration || 0

      const courses: any[] = coursesRes.data || []
      const rows = await Promise.all(
        courses.map(async (course: any) => {
          const ek = Number(course.courseEk)
          if (!Number.isFinite(ek)) return { qualified: false, duration: 0, course }
          try {
            const [statsRes, durRes]: any[] = await Promise.all([
              mockExamApi.getStats(ek),
              studyApi.getCourseDuration(ek)
            ])
            const highestScore = statsRes.data?.highestScore || 0
            const mockCount = statsRes.data?.mockCount || 0
            const totalDurationSec = durRes.data?.totalDuration || 0
            return {
              qualified: totalDurationSec >= 180 && mockCount > 0 && highestScore >= 60,
              duration: totalDurationSec,
              course
            }
          } catch {
            return { qualified: false, duration: 0, course }
          }
        })
      )

      certificateCount.value = rows.filter(r => r.qualified).length

      const topCourses = rows
        .sort((a, b) => b.duration - a.duration)
        .slice(0, 2)
        .map(({ course }) => ({
          id: Number(course.courseEk),
          title: course.courseName || '',
          tag: Number(course.courseCategory) === 2 ? '体系课程' : '单门课程',
          image: course.courseImage ? imageBaseUrl + course.courseImage : '/images/course-banner.png'
        }))
      recommendCourses.value = topCourses
    } catch (e) {}
  }
})

const recommendCourses = ref<Array<{
  id: number
  title: string
  tag: string
  image: string
}>>([])
</script>

<style scoped>
.profile-page {
  background-color: #eef0f4;
  min-height: 100vh;
}

.profile-content {
  padding: 0 0 76px;
}

.hero-section {
  background: linear-gradient(180deg, #c9e2f9 0%, #e8eef9 72%, #f0f2f7 100%);
  margin-bottom: 6px;
  padding: 10px 0 4px;
}

.top-nav {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding: 8px 18px 0;
  margin-bottom: 16px;
}

.logo-container {
  display: flex;
  align-items: center;
}

.logo {
  width: 44px;
  height: 44px;
  margin-right: 10px;
}

.app-title {
  font-size: 24px;
  font-weight: 700;
  line-height: 1.05;
  color: #2f3540;
}

.user-card {
  background-color: transparent;
  padding: 0 0 4px;
  margin-bottom: 0;
}

.user-info {
  display: grid;
  grid-template-columns: 56px minmax(0, 1fr) auto;
  align-items: center;
  column-gap: 12px;
  margin-bottom: 14px;
  padding: 0 0 0 20px;
}

.avatar {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  border: 1px solid #e0e0e0;
  background: linear-gradient(135deg, #6b80f0 0%, #8a6fd3 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.avatar svg {
  width: 32px;
  height: 32px;
  fill: white;
}

.user-details {
  min-width: 0;
  padding-left: 4px;
}

.username {
  font-size: 19px;
  font-weight: 600;
  margin: 0 0 5px 0;
  color: #242a34;
  line-height: 1.15;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.join-date {
  font-size: 14px;
  color: #2f3440;
  margin: 0;
  line-height: 1.2;
}

.user-id {
  background-color: #7395fb;
  color: #eef2ff;
  min-width: 108px;
  padding: 0 16px;
  border-radius: 8px 0 0 8px;
  border: 1px solid rgba(133, 161, 252, 0.95);
  font-size: 14px;
  font-weight: 500;
  line-height: 30px;
  height: 30px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  justify-self: end;
  box-sizing: border-box;
}

.user-stats {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  align-items: start;
  justify-items: center;
  padding: 0 6px 2px;
}

.stat-item {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
}

.stat-value {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 4px;
  line-height: 1;
  color: #202833;
}

.stat-label {
  font-size: 13px;
  color: #5f6470;
  margin-top: 0;
  line-height: 1.2;
  white-space: nowrap;
}

.section-title {
  font-size: 18px;
  font-weight: 700;
  margin: 0 0 14px;
  color: #1f242d;
}

.my-courses {
  margin: 0 16px 10px;
  background-color: #f7f7fa;
  border-radius: 20px;
  padding: 16px 16px 14px;
}

.course-actions {
  display: flex;
  justify-content: space-between;
  gap: 10px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  cursor: pointer;
}

.action-icon {
  width: 62px;
  height: 62px;
  margin-bottom: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.action-icon img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.action-label {
  font-size: 14px;
  color: #242a34;
}

.menu-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin: 0 16px 10px;
}

.menu-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}

.menu-item {
  background-color: #f7f7fa;
  border-radius: 14px;
  padding: 0 12px;
  min-height: 64px;
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  transition: transform 0.2s;
}

.menu-item:active {
  transform: scale(0.98);
}

.menu-icon {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu-icon img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.menu-label {
  font-size: 14px;
  color: #2d323c;
  font-weight: 500;
  line-height: 1;
}

.service-banner {
  margin: 0 16px 14px;
  border-radius: 28px;
  overflow: hidden;
  cursor: pointer;
}

.banner-image {
  width: 100%;
  display: block;
}

.recommend-section {
  margin: 0 16px 12px;
}

.course-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.course-card {
  background-color: transparent;
  cursor: pointer;
}

.course-image-wrapper {
  position: relative;
  border-radius: 14px;
  overflow: hidden;
}

.course-image {
  width: 100%;
  height: 132px;
  object-fit: cover;
}

.course-tag {
  position: absolute;
  top: 7px;
  left: 8px;
  padding: 4px 11px;
  border-radius: 12px;
  font-size: 13px;
  font-weight: 500;
}

.course-tag--system {
  background-color: rgba(255, 255, 255, 0.86);
  color: #3f7cea;
}

.course-tag--single {
  background-color: rgba(255, 255, 255, 0.86);
  color: #e5782f;
}

.course-info {
  padding: 8px 3px 0;
}

.course-title {
  font-size: 14px;
  font-weight: 500;
  color: #1e2530;
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 100%;
  max-width: 100%;
  display: flex;
  justify-content: space-around;
  background-color: #fff;
  padding: 8px 0 10px;
  border-top: 1px solid #ececf1;
  z-index: 100;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 13px;
  color: #8d929d;
  cursor: pointer;
  transition: color 0.3s;
}

.nav-item.active {
  color: #4c7ff6;
}

.nav-item svg {
  width: 24px;
  height: 24px;
  margin-bottom: 3px;
  fill: currentColor;
}

</style>
