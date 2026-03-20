<template>
  <div class="home-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <div class="logo-container">
        <img src="/logo.png" alt="Logo" class="logo" />
        <span class="app-title">港口学堂</span>
      </div>
    </header>

    <main class="main-content">
      <!-- 轮播区域 -->
      <section class="carousel-section">
        <div class="carousel-container" @touchstart="handleTouchStart" @touchmove="handleTouchMove" @touchend="handleTouchEnd">
          <div class="carousel-track" :style="{ transform: `translateX(-${currentIndex * 100}%)` }">
            <div v-for="(banner, index) in banners" :key="index" class="carousel-item">
              <img :src="imageBaseUrl + banner.image" alt="Banner" class="carousel-image" />
            </div>
          </div>
        </div>
        <div class="carousel-dots">
          <span v-for="(_, index) in banners" :key="index" :class="['dot', { active: currentIndex === index }]" @click="currentIndex = index"></span>
        </div>
      </section>

      <!-- 课程体系 -->
      <section class="course-section">
        <div class="section-header">
          <h2>课程体系</h2>
          <a href="#" class="see-all" @click.prevent="$router.push('/courses')">查看全部 ></a>
        </div>
        <CourseListContainer layout="horizontal">
          <CourseCard
            v-for="course in courseSystems"
            :key="course.id"
            :course="course"
            layout="horizontal"
            @click="handleCourseClick"
          />
        </CourseListContainer>
      </section>

      <!-- 限时秒杀 -->
      <section class="promo-section">
        <div class="section-header">
          <h2>限时秒杀</h2>
          <div class="countdown">
            <span>距离结束</span>
            <span class="timer-box">{{ countdown.hours }}</span> : 
            <span class="timer-box">{{ countdown.minutes }}</span> : 
            <span class="timer-box">{{ countdown.seconds }}</span>
          </div>
        </div>
        <div class="promo-card">
          <img :src="flashSale.image" alt="Promo" class="promo-image" />
          <div class="promo-info">
            <p class="promo-title">{{ flashSale.title }}</p>
            <div class="promo-meta">
              <span class="promo-price">¥{{ flashSale.price }}</span>
              <button
                v-if="flashSale.isPurchased"
                class="learn-button learn-button--study"
                @click="$router.push('/course_content/' + flashSale.id)"
              >立即学习</button>
              <button
                v-else
                class="learn-button"
                @click="$router.push('/course/' + flashSale.id)"
              >立即购买</button>
            </div>
          </div>
        </div>
      </section>

      <!-- 热门课程 -->
      <section class="course-section">
        <div class="section-header">
          <h2>热门课程</h2>
        </div>
        <CourseListContainer layout="grid">
          <CourseCard
            v-for="course in hotCourses"
            :key="course.id"
            :course="course"
            layout="grid"
            @click="handleCourseClick"
          />
        </CourseListContainer>
      </section>

    </main>

    <!-- 底部导航 -->
    <footer class="bottom-nav">
      <div class="nav-item active">
        <svg viewBox="0 0 24 24"><path d="M10,20V14H14V20H19V12H22L12,3L2,12H5V20H10Z" /></svg>
        <span>首页</span>
      </div>
      <div class="nav-item" @click="$router.push('/courses')">
        <svg viewBox="0 0 24 24"><path d="M12,3L1,9L12,15L21,9V17H23V9M5,13.18V17.18L12,21L19,17.18V13.18L12,17L5,13.18Z" /></svg>
        <span>课程</span>
      </div>
      <div class="nav-item" @click="$router.push('/profile')">
        <svg viewBox="0 0 24 24"><path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" /></svg>
        <span>个人中心</span>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import CourseCard from '@/components/CourseCard.vue'
import CourseListContainer from '@/components/CourseListContainer.vue'
import { useRouter } from 'vue-router'
import { bannerAdminApi, courseApi, flashSaleAdminApi, orderApi } from '@/api'

const router = useRouter()

const handleCourseClick = (course: any) => {
  if (course.isPurchased) {
    router.push(`/course_content/${course.id}`)
  } else {
    router.push(`/course/${course.id}`)
  }
}

// 轮播数据
const banners = ref<any[]>([])
const imageBaseUrl = 'http://localhost:8080'

const fetchBanners = async () => {
  try {
    const res: any = await bannerAdminApi.getList()
    banners.value = res.data || []
  } catch (error) {
    console.error('获取banner失败:', error)
  }
}

const currentIndex = ref(0)
let touchStartX = 0
let touchEndX = 0
let autoPlayTimer: number

const handleTouchStart = (e: TouchEvent) => {
  touchStartX = e.touches[0]?.clientX || 0
  clearInterval(autoPlayTimer)
}

const handleTouchMove = (e: TouchEvent) => {
  touchEndX = e.touches[0]?.clientX || 0
}

const handleTouchEnd = () => {
  if (touchStartX - touchEndX > 50 && currentIndex.value < banners.value.length - 1) {
    currentIndex.value++
  }
  if (touchEndX - touchStartX > 50 && currentIndex.value > 0) {
    currentIndex.value--
  }
  startAutoPlay()
}

const startAutoPlay = () => {
  autoPlayTimer = window.setInterval(() => {
    currentIndex.value = (currentIndex.value + 1) % banners.value.length
  }, 3000)
}

// 课程数据
const courseSystems = ref<any[]>([])
const hotCourses = ref<any[]>([])

const mapCourse = (item: any) => ({
  id: item.ek,
  title: item.title,
  tag: item.courseTime ? `${item.courseTime}课时` : '',
  price: (item.price / 100).toFixed(2),
  originalPrice: (item.originalPrice / 100).toFixed(2),
  image: imageBaseUrl + item.image,
  isPurchased: item.purchased || false
})

const fetchCourses = async () => {
  try {
    const [systemRes, hotRes]: any[] = await Promise.all([
      courseApi.getCourseList({ pageIndex: 1, pageSize: 10 }),
      courseApi.getHotCourses()
    ])
    courseSystems.value = (systemRes.data?.list || []).map(mapCourse)
    hotCourses.value = (hotRes.data || []).map(mapCourse)

    // 检查购买状态
    if (localStorage.getItem('token')) {
      try {
        const systemChecks = await Promise.all(
          courseSystems.value.map(c => orderApi.hasPurchased(c.id).catch(() => ({ data: false })))
        )
        courseSystems.value = courseSystems.value.map((c, i) => ({ ...c, isPurchased: systemChecks[i]?.data || false }))
      } catch (e) {}
    }
  } catch (error) {
    console.error('获取课程失败:', error)
  }
}

const flashSale = ref<any>({ id: 0, title: '', price: 0, image: '', isPurchased: false })
const countdown = ref({ hours: '00', minutes: '00', seconds: '00' })
let timer: number

const fetchFlashSale = async () => {
  try {
    const res: any = await flashSaleAdminApi.getActive()
    if (res.data) {
      const ek = res.data.course?.ek || 0
      let isPurchased = false
      if (ek && localStorage.getItem('token')) {
        try {
          const purchaseRes: any = await orderApi.hasPurchased(ek)
          isPurchased = purchaseRes.data || false
        } catch {}
      }
      flashSale.value = {
        id: ek,
        title: res.data.course?.courseName || '',
        price: (res.data.price / 100).toFixed(2),
        image: res.data.course?.image ? imageBaseUrl + res.data.course.image : '',
        isPurchased
      }
      if (res.data.endTime) {
        startCountdown(new Date(res.data.endTime).getTime())
      }
    }
  } catch {}
}

function startCountdown(endTimestamp: number) {
  clearInterval(timer)
  timer = window.setInterval(() => {
    const diff = Math.max(0, Math.floor((endTimestamp - Date.now()) / 1000))
    const hours = Math.floor(diff / 3600)
    const minutes = Math.floor((diff % 3600) / 60)
    const seconds = diff % 60
    countdown.value = {
      hours: String(hours).padStart(2, '0'),
      minutes: String(minutes).padStart(2, '0'),
      seconds: String(seconds).padStart(2, '0')
    }
    if (diff === 0) clearInterval(timer)
  }, 1000)
}

onMounted(() => {
  startAutoPlay()
  fetchBanners()
  fetchCourses()
  fetchFlashSale()
})

onUnmounted(() => {
  clearInterval(timer)
  clearInterval(autoPlayTimer)
})
</script>

<style scoped>
.home-page {
  background-color: #f5f5f5;
  padding-bottom: 70px;
  overflow-x: hidden;
  max-width: 100vw;
  margin: 0 auto;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  background-color: #fff;
}

.logo-container {
  display: flex;
  align-items: center;
}

.logo {
  width: 40px;
  height: 40px;
  margin-right: 8px;
}

.app-title {
  font-size: 28px;
  font-weight: bold;
}

.actions .icon {
  width: 24px;
  height: 24px;
  margin-left: 15px;
}

.main-content {
  padding: 15px;
  max-width: 1200px;
  margin: 0 auto;
}

.carousel-section {
  margin-bottom: 20px;
}

.carousel-container {
  position: relative;
  overflow: hidden;
  border-radius: 10px;
}

.carousel-track {
  display: flex;
  transition: transform 0.3s ease;
}

.carousel-item {
  min-width: 100%;
  position: relative;
}

.carousel-image {
  width: 100%;
  display: block;
}

.carousel-dots {
  display: flex;
  justify-content: center;
  gap: 8px;
  margin-top: 10px;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #ddd;
  cursor: pointer;
  transition: background-color 0.3s;
}

.dot.active {
  background-color: #3b82f6;
}

.video-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.2);
  display: flex;
  justify-content: center;
  align-items: center;
}

.play-button {
  width: 50px;
  height: 50px;
  background-color: rgba(0,0,0,0.5);
  border-radius: 50%;
  position: relative;
}
.play-button::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-40%, -50%);
  width: 0;
  height: 0;
  border-top: 10px solid transparent;
  border-bottom: 10px solid transparent;
  border-left: 15px solid #fff;
}

.video-title {
  position: absolute;
  bottom: 15px;
  left: 15px;
  right: 15px;
  font-size: 16px;
  font-weight: bold;
  color: #fff;
}

.video-tag {
  position: absolute;
  bottom: 50px;
  left: 15px;
  background-color: #3b82f6;
  color: #fff;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
}

.course-section {
  margin-bottom: 20px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5px;
}

.section-header h2 {
  font-size: 18px;
}

.promo-section .section-header h2 {
  color: #fff;
}

.see-all {
  font-size: 14px;
  color: #666;
  text-decoration: none;
}

.course-list {
  display: flex;
  gap: 15px;
  overflow-x: auto;
}

.course-card {
  flex: 0 0 200px;
  background-color: #fff;
  border-radius: 10px;
  overflow: hidden;
}

.course-image-wrapper {
  position: relative;
}

.course-image {
  width: 100%;
  height: 100px;
  object-fit: cover;
}

.course-info {
  padding: 8px;
}

.course-tag {
  position: absolute;
  top: 8px;
  left: 8px;
  font-size: 12px;
  color: #fff;
  background-color: rgba(0,0,0,0.6);
  padding: 4px 8px;
  border-radius: 4px;
}

.course-title {
  font-size: 14px;
  font-weight: bold;
  margin: 8px 0;
}

.course-footer {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
}

.course-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.course-price {
  font-size: 20px;
  color: #ff6b35;
  font-weight: bold;
}

.original-price {
  font-size: 12px;
  color: #999;
  text-decoration: line-through;
  margin-bottom: 4px;
}

.buy-button {
  padding: 8px 20px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  white-space: nowrap;
}

.promo-section {
  background-color: #cce5ff;
  padding: 5px 20px 15px 20px;
  border-radius: 10px;
  margin-bottom: 20px;
}

.countdown {
  display: flex;
  align-items: center;
  font-size: 14px;
  color: #fff;
}

.timer-box {
  background-color: #333;
  color: #fff;
  padding: 2px 5px;
  border-radius: 3px;
  margin: 0 5px;
}

.promo-card {
  display: flex;
  align-items: center;
  margin-top: 8px;
  background-color: #fff;
  padding: 15px;
  border-radius: 10px;
}

.promo-image {
  width: 100px;
  height: 100px;
  border-radius: 10px;
  margin-right: 15px;
}

.promo-title {
  font-size: 16px;
  font-weight: bold;
}

.promo-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
}

.promo-meta .learn-button {
  width: auto;
  padding: 8px 24px;
  margin-left: 15px;
}

.promo-price {
  font-size: 20px;
  color: #ff6b35;
  font-weight: bold;
}

.learn-button {
  width: 100%;
  padding: 8px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  cursor: pointer;
}

.learn-button--study {
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
}

.course-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.course-card-grid {
  background-color: #fff;
  border-radius: 10px;
  overflow: hidden;
  padding: 8px;
}

.course-card-grid .course-image {
  width: 100%;
  height: 120px;
  object-fit: cover;
  border-radius: 10px;
  margin-bottom: 8px;
}

.course-card-grid .course-title {
  font-size: 14px;
  font-weight: bold;
  margin: 0 0 8px 0;
}

.course-meta-grid {
  display: flex;
  justify-content: flex-start;
  align-items: baseline;
  gap: 10px;
  margin-bottom: 10px;
}

.course-card-grid .buy-button {
  width: 85%;
  margin: 0 auto;
  display: block;
  padding: 8px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 14px;
}


.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-around;
  background-color: #fff;
  padding: 10px 0;
  border-top: 1px solid #eee;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 12px;
  color: #999;
  cursor: pointer;
  transition: color 0.3s;
}

.nav-item.active {
  color: #3b82f6;
}

.nav-item svg {
  width: 24px;
  height: 24px;
  margin-bottom: 4px;
  fill: currentColor;
}
</style>
