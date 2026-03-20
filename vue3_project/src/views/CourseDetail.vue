<template>
  <div class="course-detail">
    <header class="detail-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程详情</span>
      <div class="nav-placeholder"></div>
    </header>

    <main class="detail-content">
      <section class="course-info-card">
        <div class="video-section">
          <img :src="imageBaseUrl + courseDetail.image" alt="Course" class="course-banner" />
          <div class="play-overlay">
            <div class="play-btn"></div>
          </div>
        </div>

        <h1 class="course-name">{{ courseDetail.courseName }}</h1>
        <div class="price-section">
          <span class="price">¥{{ (courseDetail.price / 100).toFixed(2) }}</span>
          <span class="original-price">¥{{ (courseDetail.originalPrice / 100).toFixed(2) }}</span>
        </div>
        <div class="course-desc-wrap">
          <p class="course-desc" :class="{ 'course-desc--collapsed': !showMore }">{{ courseDetail.courseDetail }}</p>
          <button v-if="isLongDesc" class="more-btn" @click="showMore = !showMore">
            {{ showMore ? '收起 ∧' : '查看更多 ∨' }}
          </button>
        </div>
      </section>

      <!-- 单门课程显示题库和附件信息 -->
      <section v-if="courseDetail.courseCategory === 1" class="course-stats">
        <div class="stats-row">
          <div class="stat-card">
            <div class="stat-card-header stat-header--orange">
              <span class="stat-card-title">理论题库</span>
              <svg class="stat-icon" viewBox="0 0 24 24"><path d="M4 6h16v2H4zm0 5h16v2H4zm0 5h16v2H4z"/></svg>
            </div>
            <div class="stat-card-body">
              <div class="stat-row">
                <span class="stat-label">单选题</span><span class="stat-val stat-val--orange">{{ courseDetail.theoreticalQuestions?.singleQuestion ?? 0 }}</span>
                <span class="stat-label">多选题</span><span class="stat-val stat-val--orange">{{ courseDetail.theoreticalQuestions?.multipleQuestion ?? 0 }}</span>
              </div>
              <div class="stat-row">
                <span class="stat-label">判断题</span><span class="stat-val stat-val--orange">{{ courseDetail.theoreticalQuestions?.judgeQuestion ?? 0 }}</span>
              </div>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-card-header stat-header--purple">
              <span class="stat-card-title">视频课程</span>
              <svg class="stat-icon" viewBox="0 0 24 24"><path d="M17 10.5V7a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-3.5l4 4v-11l-4 4z"/></svg>
            </div>
            <div class="stat-card-body">
              <div class="stat-row">
                <span class="stat-label">数量</span><span class="stat-val stat-val--purple">{{ courseDetail.videoQuestions ?? 0 }}</span>
              </div>
            </div>
          </div>
        </div>
        <div class="stats-row">
          <div class="stat-card stat-card--half">
            <div class="stat-card-header stat-header--teal">
              <span class="stat-card-title">附　件</span>
              <svg class="stat-icon" viewBox="0 0 24 24"><path d="M16.5 6v11.5c0 2.21-1.79 4-4 4s-4-1.79-4-4V5a2.5 2.5 0 0 1 5 0v10.5c0 .83-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5V6H9v9.5a2.5 2.5 0 0 0 5 0V5c0-2.21-1.79-4-4-4S6 2.79 6 5v12.5c0 3.04 2.46 5.5 5.5 5.5s5.5-2.46 5.5-5.5V6h-1.5z"/></svg>
            </div>
            <div class="stat-card-body">
              <div class="stat-row">
                <span class="stat-label">数量</span><span class="stat-val stat-val--teal">{{ courseDetail.attachment ?? 0 }}</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 体系课程才显示课程内容列表 -->
      <section v-if="courseDetail.courseCategory === 2 && contentList.length > 0" class="course-content">
        <div class="content-header">
          <span class="content-title">课程内容</span>
        </div>
        <div class="content-list">
          <div v-for="item in contentList" :key="item.ek" class="content-item" @click="$router.push('/course/' + item.ek)">
            <span class="item-title">{{ item.courseName }}</span>
            <span class="item-price">¥{{ item.price != null ? (item.price / 100).toFixed(2) : '--' }}</span>
          </div>
        </div>
      </section>

      <!-- 体系课程才显示课程介绍图片 -->
      <section v-if="courseDetail.courseCategory === 2 && courseDetail.courseIntroduction" class="course-intro-section">
        <h2 class="section-title">课程介绍</h2>
        <img :src="imageBaseUrl + courseDetail.courseIntroduction" alt="Course Intro" class="intro-image" />
      </section>
    </main>

    <footer class="buy-footer">
      <button v-if="purchased" class="buy-btn buy-btn--study" @click="$router.push('/course_content/' + route.params.id)">立即学习</button>
      <button v-else class="buy-btn" @click="$router.push('/purchase?ek=' + route.params.id)">立即购买</button>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { courseApi, studyApi, orderApi } from '@/api'

const route = useRoute()
const courseDetail = ref<any>({})
const contentList = ref<any[]>([])
const imageBaseUrl = 'http://localhost:8080'
const showMore = ref(false)
const DESC_LIMIT = 100
const purchased = ref(false)

const isLongDesc = computed(() => (courseDetail.value.courseDetail || '').length > DESC_LIMIT)

const fetchCourseDetail = async () => {
  try {
    const ek = Number(route.params.id)
    const res: any = await courseApi.getCourseDetail(ek)
    courseDetail.value = res.data
    contentList.value = res.data.courseContent || []
  } catch (error) {
    console.error('获取课程详情失败:', error)
  }
}

const checkPurchased = async () => {
  const token = localStorage.getItem('token')
  if (!token) return
  try {
    const ek = Number(route.params.id)
    const res: any = await orderApi.hasPurchased(ek)
    purchased.value = res.data || false
  } catch (e) {}
}

// 学时计时
let startTime = 0

onMounted(() => {
  fetchCourseDetail()
  checkPurchased()
  startTime = Date.now()
})

watch(() => route.params.id, () => {
  fetchCourseDetail()
  checkPurchased()
  startTime = Date.now()
})

onUnmounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) return
  const duration = Math.floor((Date.now() - startTime) / 1000)
  if (duration < 1) return
  try {
    await studyApi.addRecord(duration)
  } catch (e) {
    console.error('上报学时失败:', e)
  }
})
</script>

<style scoped>
.course-detail {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 70px;
  overflow-x: hidden;
  max-width: 100vw;
  margin: 0 auto;
}

.detail-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.nav-placeholder {
  width: 24px;
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
  width: 20px;
  height: 20px;
  fill: #666;
}

.detail-content {
  padding-bottom: 80px;
  max-width: 1200px;
  margin: 0 auto;
}

.course-info-card {
  background-color: #fff;
  margin: 10px 12px;
  border-radius: 12px;
  padding: 12px;
  margin-bottom: 10px;
}

.video-section {
  position: relative;
  border-radius: 10px;
  overflow: hidden;
  margin-bottom: 12px;
}

.course-banner {
  width: 100%;
  max-height: 200px;
  object-fit: cover;
  display: block;
  border-radius: 10px;
}

.play-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(0,0,0,0.2);
}

.play-btn {
  width: 60px;
  height: 60px;
  background-color: rgba(255,255,255,0.9);
  border-radius: 50%;
  position: relative;
}

.play-btn::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-35%, -50%);
  width: 0;
  height: 0;
  border-top: 12px solid transparent;
  border-bottom: 12px solid transparent;
  border-left: 18px solid #3b82f6;
}

.course-info {
  background-color: #fff;
  padding: 16px;
  margin-bottom: 10px;
}

.course-name {
  font-size: 20px;
  font-weight: 600;
  color: #333;
  margin-bottom: 12px;
}

.price-section {
  display: flex;
  align-items: baseline;
  gap: 10px;
  margin-bottom: 12px;
}

.price {
  font-size: 24px;
  color:  #ff6b35;
  font-weight: 600;
}

.price-section .original-price {
  font-size: 14px;
  color: #999;
  text-decoration: line-through;
}

.course-desc {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin-bottom: 4px;
}

.course-desc--collapsed {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.more-btn {
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 14px;
  cursor: pointer;
  padding: 0;
  margin-bottom: 8px;
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.course-stats {
  background: #fff;
  padding: 14px 12px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 10px;
}

.stats-row {
  display: flex;
  gap: 10px;
}

.stat-card {
  flex: 1;
  border-radius: 10px;
  overflow: hidden;
  background: #f0f4ff;
}

.stat-card--half {
  flex: none;
  width: calc(50% - 5px);
  border-radius: 10px;
  overflow: hidden;
  background: #f0faf6;
}

.stat-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 7px 12px;
}

.stat-header--orange {
  background: linear-gradient(90deg, #ff9a3c, #ffb86c);
}

.stat-header--purple {
  background: linear-gradient(90deg, #a78bfa, #c4b5fd);
}

.stat-header--teal {
  background: linear-gradient(90deg, #2dd4bf, #5eead4);
}

.stat-card-title {
  font-size: 14px;
  font-weight: 600;
  color: #fff;
}

.stat-icon {
  width: 20px;
  height: 20px;
  fill: rgba(255,255,255,0.85);
}

.stat-card-body {
  padding: 10px 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.stat-row {
  display: flex;
  align-items: center;
  gap: 6px;
  flex-wrap: wrap;
}

.stat-label {
  font-size: 13px;
  color: #666;
}

.stat-val {
  font-size: 16px;
  font-weight: 700;
  color: #333;
  margin-right: 8px;
}

.stat-val--orange { color: #ff9a3c; }
.stat-val--purple { color: #a78bfa; }
.stat-val--teal   { color: #2dd4bf; }

.course-content {
  background-color: #fff;
  padding: 16px;
  margin-bottom: 10px;
}

.content-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.content-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.content-list {
  display: flex;
  flex-direction: column;
  gap: 0;
}

.content-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  padding: 12px 8px;
  transition: background-color 0.2s;
}

.content-item:hover {
  background-color: #f0f0f0;
}

.item-title {
  font-size: 14px;
  color: #3b82f6;
  text-decoration: underline;
  text-decoration-color: #3b82f6;
  text-underline-offset: 2px;
}

.item-price {
  font-size: 14px;
  color: #ff6b35;
  font-weight: 600;
}

.course-intro-section,
.course-outline-section {
  width: 100%;
  background-color: #fff;
  padding: 16px;
  margin-bottom: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
  align-self: flex-start;
}

.intro-image,
.outline-image {
  width: 100%;
  display: block;
  margin: 0 auto;
}

.buy-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 -2px 8px rgba(0,0,0,0.1);
}

.price-info {
  display: flex;
  align-items: baseline;
  gap: 8px;
}

.current-price {
  font-size: 24px;
  color: #ff4d4f;
  font-weight: 600;
}

.original-price {
  font-size: 14px;
  color: #999;
  text-decoration: line-through;
}

.buy-btn {
  padding: 12px 40px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 25px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  width: 100%;
  max-width: 400px;
}

.buy-btn--study {
  background: #fff;
  color: #3b82f6;
  border: 1.5px solid #3b82f6;
}
</style>
