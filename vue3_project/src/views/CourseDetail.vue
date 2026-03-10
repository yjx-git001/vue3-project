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
      <section class="video-section">
        <img src="/images/course-detail-banner.png" alt="Course" class="course-banner" />
        <div class="play-overlay">
          <div class="play-btn"></div>
        </div>
      </section>

      <section class="course-info">
        <h1 class="course-name">门机操作初级课程</h1>
        <div class="price-section">
          <span class="price">¥129</span>
          <span class="original-price">¥350</span>
        </div>
        <p class="course-desc">课程详情：课程学习时长，基于课程的实操教学服务。本文档为初级课程，为保证学习效果，建议按照顺序进行学习。<span v-if="showMore">本课程涵盖门机操作的核心知识点，包括设备结构认识、安全操作规范、日常维护保养等内容，适合零基础学员系统学习。完成本课程后可参加初级资格认证考试。</span></p>
        <button class="more-btn" @click="showMore = !showMore">{{ showMore ? '收起 ∧' : '查看更多 >' }}</button>
      </section>

      <section class="course-content">
        <div class="content-header">
          <span class="content-title">课程内容</span>
          <span class="save-text">立省221↓</span>
        </div>
        <div class="content-list">
          <div v-for="item in contentList" :key="item.id" class="content-item" @click="$router.push('/single_course/' + item.id)">
            <span class="item-title">{{ item.title }}</span>
            <span class="item-price">¥50</span>
          </div>
        </div>
      </section>

      <section class="course-intro-section">
        <h2 class="section-title">课程介绍</h2>
        <img src="/images/course-intro.png" alt="Course Intro" class="intro-image" />
      </section>

      <section class="course-outline-section">
        <img src="/images/course-outline.png" alt="Course Outline" class="outline-image" />
      </section>
    </main>

    <footer class="buy-footer">
      <button class="buy-btn" @click="$router.push('/purchase')">立即购买</button>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const showMore = ref(false)

const contentList = ref([
  { id: 1, title: '港机设备-起重机械' },
  { id: 2, title: '港机设备-年度装卸起重机' },
  { id: 3, title: '港机设备-轮胎式集装箱门机' },
  { id: 4, title: '港机设备-门座式起重机' },
  { id: 5, title: '港机设备-门机司机' },
  { id: 6, title: '港机设备-门机司机' },
  { id: 7, title: '港机设备-门机司机' }
])
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

.video-section {
  position: relative;
}

.course-banner {
  width: 100%;
  display: block;
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
  margin-bottom: 12px;
}

.more-btn {
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 14px;
  cursor: pointer;
  padding: 0;
}

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

.save-text {
  font-size: 14px;
  color: #ff4d4f;
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
  color: #ff4d4f;
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
</style>
