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
      <!-- 视频播放区域 -->
      <section class="video-section">
        <img src="/images/course-banner.png" alt="Video Thumbnail" class="video-thumb" />
        <div class="video-overlay">
          <div class="play-button"></div>
        </div>
        <span class="video-tag">大师课</span>
        <span class="video-title">智慧港口运营与自动化管理</span>
      </section>

      <!-- 课程体系 -->
      <section class="course-section">
        <div class="section-header">
          <h2>课程体系</h2>
          <a href="#" class="see-all" @click.prevent="$router.push('/courses')">查看全部 ></a>
        </div>
        <div class="course-list">
          <div v-for="course in courseSystems" :key="course.id" class="course-card">
            <div class="course-image-wrapper">
              <img :src="course.image" :alt="course.title" class="course-image" />
              <span class="course-tag">{{ course.tag }}</span>
            </div>
            <div class="course-info">
              <p class="course-title">{{ course.title }}</p>
              <div class="course-footer">
                <div class="course-meta">
                  <span class="original-price">¥{{ course.originalPrice }}</span>
                  <span class="course-price">¥{{ course.price }}</span>
                </div>
                <button v-if="course.isPurchased" class="learn-button" @click="$router.push('/course_content/' + course.id)">立即学习</button>
                <button v-else class="buy-button" @click="$router.push('/course/' + course.id)">立即购买</button>
              </div>
            </div>
          </div>
        </div>
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
              <button class="learn-button" @click="$router.push('/course_content/' + flashSale.id)">立即学习</button>
            </div>
          </div>
        </div>
      </section>

      <!-- 热门课程 -->
      <section class="course-section">
        <div class="section-header">
          <h2>热门课程</h2>
        </div>
        <div class="course-grid">
          <div v-for="course in hotCourses" :key="course.id" class="course-card-grid">
            <img :src="course.image" :alt="course.title" class="course-image" />
            <p class="course-title">{{ course.title }}</p>
            <div class="course-meta-grid">
              <span class="course-price">¥{{ course.price }}</span>
              <span class="original-price">¥{{ course.originalPrice }}</span>
            </div>
            <button class="buy-button" @click="$router.push('/course/' + course.id)">立即购买</button>
          </div>
        </div>
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

// 模拟数据
const courseSystems = ref([
  {
    id: 1,
    title: '港口特种设备检修课程',
    tag: '12课时',
    price: 499,
    originalPrice: 500,
    image: '/images/course-banner.png',
    isPurchased: false
  },
  {
    id: 2,
    title: '港口特种设备检修课程',
    tag: '12课时',
    price: 499,
    originalPrice: 500,
    image: '/images/course-banner.png',
    isPurchased: true
  }
])

const flashSale = ref({
  id: 1,
  title: '港口安全知识-倒运司机',
  price: 0,
  image: '/images/promo-thumb.png'
})

const hotCourses = ref([
  { id: 1, title: '港口特种设备检修课程', price: 499, originalPrice: 500, image: '/images/course-banner.png' },
  { id: 2, title: '港口特种设备检修课程', price: 499, originalPrice: 500, image: '/images/course-banner.png' },
  { id: 3, title: '港口特种设备检修课程', price: 499, originalPrice: 500, image: '/images/course-banner.png' },
  { id: 4, title: '港口特种设备检修课程', price: 499, originalPrice: 500, image: '/images/course-banner.png' }
])


// 倒计时逻辑
const countdown = ref({ hours: '02', minutes: '14', seconds: '55' })
let timer: number

function startCountdown() {
  let totalSeconds = 2 * 3600 + 14 * 60 + 55
  timer = window.setInterval(() => {
    if (totalSeconds > 0) {
      totalSeconds--
      const hours = Math.floor(totalSeconds / 3600)
      const minutes = Math.floor((totalSeconds % 3600) / 60)
      const seconds = totalSeconds % 60
      countdown.value = {
        hours: String(hours).padStart(2, '0'),
        minutes: String(minutes).padStart(2, '0'),
        seconds: String(seconds).padStart(2, '0')
      }
    } else {
      clearInterval(timer)
    }
  }, 1000)
}

onMounted(() => {
  startCountdown()
})

onUnmounted(() => {
  clearInterval(timer)
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

.video-section {
  position: relative;
  border-radius: 10px;
  overflow: hidden;
  margin-bottom: 20px;
}

.video-thumb {
  width: 100%;
  display: block;
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
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 10px;
  cursor: pointer;
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
