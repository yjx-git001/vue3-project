<template>
  <div class="course-list-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <div class="logo-container">
        <img src="/logo.png" alt="Logo" class="logo" />
        <span class="app-title">港口学堂</span>
      </div>
    </header>

    <!-- 课程分类标签 -->
    <div class="course-tabs">
      <div
        v-for="tab in tabs"
        :key="tab.id"
        :class="['tab-item', { active: activeTab === tab.id }]"
        @click="activeTab = tab.id"
      >
        {{ tab.name }}
      </div>
    </div>

    <!-- 筛选按钮 -->
    <div class="filter-section">
      <button
        v-for="filter in filters"
        :key="filter.id"
        :class="['filter-btn', { active: activeFilter === filter.id }]"
        @click="activeFilter = filter.id"
      >
        {{ filter.name }}
      </button>
    </div>

    <!-- 课程列表 -->
    <main class="course-content">
      <div
        v-for="course in courseList"
        :key="course.id"
        class="course-card"
      >
        <div class="course-image-wrapper">
          <img :src="course.image" :alt="course.title" class="course-image" />
          <span v-if="course.tag" :class="['course-tag', course.tag === '热门课程' ? 'course-tag--hot' : 'course-tag--system']">{{ course.tag }}</span>
        </div>
        <div class="course-info">
          <h3 class="course-title">{{ course.title }}</h3>
          <div v-if="!course.purchased" class="course-meta">
            <span class="course-price">¥{{ course.price }}</span>
            <span class="original-price">¥{{ course.originalPrice }}</span>
          </div>
          <div v-else class="purchased-tag">已购买</div>
          <button
            v-if="!course.purchased"
            class="buy-btn"
            @click="$router.push('/course/' + course.id)"
          >
            立即购买
          </button>
          <button
            v-else
            class="learn-btn"
            @click="$router.push('/course_content/' + course.id)"
          >
            立即学习
          </button>
        </div>
      </div>
    </main>

    <!-- 底部导航 -->
    <footer class="bottom-nav">
      <div class="nav-item" @click="$router.push('/')">
        <svg viewBox="0 0 24 24"><path d="M10,20V14H14V20H19V12H22L12,3L2,12H5V20H10Z" /></svg>
        <span>首页</span>
      </div>
      <div class="nav-item active">
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
import { ref } from 'vue'

const activeTab = ref('all')
const activeFilter = ref('all')

const tabs = ref([
  { id: 'all', name: '全部课程' },
  { id: 'single', name: '单门课程' },
  { id: 'system', name: '体系课程' }
])

const filters = ref([
  { id: 'all', name: '全部' },
  { id: 'equipment', name: '港机设备' },
  { id: 'safety', name: '安全风险' },
  { id: 'production', name: '生产装卸' },
  { id: 'maintenance', name: '维护保养' },
  { id: 'other', name: '其他' }
])

const courseList = ref([
  {
    id: 1,
    title: '港口特种设备检修课程课程课程',
    description: '12课时 · 专业讲师授课',
    price: 499,
    originalPrice: 500,
    image: '/images/course-banner.png',
    tag: '体系课程',
    purchased: false
  },
  {
    id: 2,
    title: '港口特种设备检修课程',
    description: '15课时 · 实战经验分享',
    price: 499,
    originalPrice: 500,
    image: '/images/course-banner.png',
    tag: '热门课程',
    purchased: true
  },
  {
    id: 3,
    title: '港口特种设备检修课程课程课程',
    description: '10课时 · 操作技能培训',
    price: 499,
    originalPrice: 500,
    image: '/images/course-banner.png',
    tag: '体系课程',
    purchased: false
  },
  {
    id: 4,
    title: '港口特种设备检修课程',
    description: '18课时 · 维护技能提升',
    price: 499,
    originalPrice: 500,
    image: '/images/course-banner.png',
    tag: '热门课程',
    purchased: false
  }
])
</script>

<style scoped>
.course-list-page {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 70px;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
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


.course-tabs {
  display: flex;
  background-color: #fff;
  padding: 12px 0;
  border-bottom: 1px solid #eee;
}

.tab-item {
  flex: 1;
  text-align: center;
  font-size: 15px;
  color: #666;
  cursor: pointer;
  padding-bottom: 8px;
  position: relative;
  transition: color 0.3s;
}

.tab-item.active {
  color: #3b82f6;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background-color: #3b82f6;
}

.filter-section {
  display: flex;
  gap: 12px;
  padding: 12px 16px;
  background-color: #fff;
  overflow-x: auto;
  margin-bottom: 10px;
}

.filter-btn {
  padding: 6px 16px;
  border: 1px solid #ddd;
  background-color: #fff;
  border-radius: 20px;
  font-size: 14px;
  color: #666;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.3s;
}

.filter-btn.active {
  background-color: #3b82f6;
  color: #fff;
  border-color: #3b82f6;
}

.course-content {
  padding: 0 16px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.course-card {
  background-color: #fff;
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.course-image-wrapper {
  position: relative;
  width: 100%;
}

.course-image {
  width: 100%;
  height: 120px;
  object-fit: cover;
}

.course-tag {
  position: absolute;
  top: 8px;
  left: 8px;
  color: #fff;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.course-tag--system {
  background-color: rgba(255,255,255,0.8);
  color: #3b82f6;
  border-radius: 10px;
}

.course-tag--hot {
  background-color: rgba(255,255,255,0.8);
  color: #ff6b35;
  border-radius: 10px;
}

.course-info {
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.course-title {
  font-size: 14px;
  font-weight: 600;
  color: #333;
  margin: 0;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1.4;
  min-height: 40px;
}

.course-meta {
  display: flex;
  align-items: baseline;
  gap: 6px;
}

.purchased-tag {
  font-size: 14px;
  color: #3b82f6;
  font-weight: 500;
}

.course-price {
  font-size: 20px;
  color: #ff6b35;
  font-weight: 600;
}

.original-price {
  font-size: 12px;
  color: #999;
  text-decoration: line-through;
}

.buy-btn,
.learn-btn {
  width: 100%;
  padding: 8px;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
}

.buy-btn {
  background-color: #3b82f6;
  color: #fff;
}

.buy-btn:active {
  background-color: #2563eb;
}

.learn-btn {
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
}

.learn-btn:active {
  background-color: #eff6ff;
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
  z-index: 100;
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
