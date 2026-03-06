<template>
  <div class="course-content-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程详情</span>
      <div class="nav-actions">
        <button class="icon-btn">
          <svg viewBox="0 0 24 24"><path d="M12,16A2,2 0 0,1 14,18A2,2 0 0,1 12,20A2,2 0 0,1 10,18A2,2 0 0,1 12,16M12,10A2,2 0 0,1 14,12A2,2 0 0,1 12,14A2,2 0 0,1 10,12A2,2 0 0,1 12,10M12,4A2,2 0 0,1 14,6A2,2 0 0,1 12,8A2,2 0 0,1 10,6A2,2 0 0,1 12,4Z" /></svg>
        </button>
        <button class="icon-btn">
          <svg viewBox="0 0 24 24"><path d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z" /></svg>
        </button>
      </div>
    </header>

    <!-- 课程横幅 -->
    <div class="course-banner">
      <img src="/images/course-detail-banner.png" alt="课程横幅" class="banner-image" />
      <div class="play-overlay">
        <div class="play-button">
          <svg viewBox="0 0 24 24"><path d="M8,5.14V19.14L19,12.14L8,5.14Z" fill="white" /></svg>
        </div>
      </div>
    </div>

    <!-- 课程信息 -->
    <div class="course-info-section">
      <h2 class="course-title">港口特种设备检修课程</h2>
      <span class="course-tag">体系课程</span>
    </div>

    <!-- 标签页 -->
    <div class="tabs">
      <div
        v-for="tab in tabs"
        :key="tab.value"
        :class="['tab-item', { active: activeTab === tab.value }]"
        @click="activeTab = tab.value"
      >
        {{ tab.label }}
      </div>
    </div>

    <main class="content-area">
      <!-- 目录 -->
      <div v-if="activeTab === 'catalog'" class="catalog-content">
        <div class="chapter-header">
          <span class="chapter-title">港机设备-起重机械</span>
        </div>
        <div
          v-for="lesson in lessons"
          :key="lesson.id"
          class="lesson-item"
        >
          <div class="lesson-number">{{ String(lesson.id).padStart(2, '0') }}</div>
          <div class="lesson-info">
            <div class="lesson-title">{{ lesson.title }}</div>
            <div class="lesson-meta">
              <span class="lesson-type">{{ lesson.type }}</span>
              <span class="lesson-duration">{{ lesson.duration }}</span>
            </div>
          </div>
          <button class="play-btn" :class="{ playing: lesson.isPlaying }">
            <svg v-if="lesson.isPlaying" viewBox="0 0 24 24"><path d="M14,19H18V5H14M6,19H10V5H6V19Z" /></svg>
            <svg v-else viewBox="0 0 24 24"><path d="M8,5.14V19.14L19,12.14L8,5.14Z" /></svg>
          </button>
        </div>
      </div>

      <!-- 考题 -->
      <div v-if="activeTab === 'exam'" class="exam-content">
        <div class="chapter-header">
          <span class="chapter-title">港机设备-起重机械</span>
          <button class="switch-course-btn">
            切换课程
            <svg viewBox="0 0 24 24" class="dropdown-icon"><path d="M7,10L12,15L17,10H7Z" /></svg>
          </button>
        </div>
        <div class="exam-grid">
          <div class="exam-card" @click="showExamModal = true">
            <div class="exam-icon blue">
              <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" fill="currentColor" /></svg>
            </div>
            <span class="exam-label">考题训练</span>
          </div>
          <div class="exam-card">
            <div class="exam-icon red">
              <svg viewBox="0 0 24 24"><path d="M11,15H13V17H11V15M11,7H13V13H11V7M12,2C6.47,2 2,6.5 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20Z" fill="currentColor" /></svg>
            </div>
            <span class="exam-label">错题记录</span>
          </div>
          <div class="exam-card">
            <div class="exam-icon green">
              <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20M15,18V16H6V18H15M18,14V12H6V14H18Z" fill="currentColor" /></svg>
            </div>
            <span class="exam-label">模考记录</span>
          </div>
          <div class="exam-card">
            <div class="exam-icon purple">
              <svg viewBox="0 0 24 24"><path d="M12,3L1,9L12,15L21,10.09V17H23V9M5,13.18V17.18L12,21L19,17.18V13.18L12,17L5,13.18Z" fill="currentColor" /></svg>
            </div>
            <span class="exam-label">模拟考试</span>
          </div>
        </div>
      </div>

    <!-- 考题训练弹窗 -->
    <div v-if="showExamModal" class="modal-overlay" @click="showExamModal = false">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <span class="modal-title">港机设备-起重机械</span>
        </div>
        <div class="question-types">
          <div class="type-item" @click="startPractice('single')">
            <span>单选题</span>
            <span class="count">(20)</span>
          </div>
          <div class="type-item" @click="startPractice('multiple')">
            <span>多选题</span>
            <span class="count">(20)</span>
          </div>
          <div class="type-item" @click="startPractice('essay')">
            <span>论述题</span>
            <span class="count">(XX)</span>
          </div>
          <div class="type-item" @click="startPractice('other')">
            <span>XXX</span>
            <span class="count">(2)</span>
          </div>
        </div>
        <button class="cancel-btn" @click="showExamModal = false">取消</button>
      </div>
    </div>

      <!-- 笔记 -->
      <div v-if="activeTab === 'notes'" class="notes-content">
        <div class="chapter-header">
          <span class="chapter-title">港机设备-起重机械</span>
          <button class="switch-course-btn">
            切换课程
            <svg viewBox="0 0 24 24" class="dropdown-icon"><path d="M7,10L12,15L17,10H7Z" /></svg>
          </button>
        </div>
        <div class="notes-text">
          <p>在港口作业场景中，装载机(常被俗称"装载车")属于场内专用工程机械车辆,而非道路货运车辆。主要用于港口散货、件杂货的铲装、转运、堆垛作业。</p>
          <p>其车型定位和组分可以从以下维度明确:</p>
          <p>1.技功能属性属于土方工程机械范畴，核心特征是配备前端铲斗工作装置，区别于货的货物运输功能。装载机的核心作用是货物装卸与场内短驳，不承担长运输任务。</p>
          <p>2.按作业场景属性港口使用的装载机多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为多为</p>
        </div>
      </div>

      <!-- 附件 -->
      <div v-if="activeTab === 'attachments'" class="attachments-content">
        <div class="chapter-header">
          <span class="chapter-title">港机设备-起重机械</span>
        </div>
        <div
          v-for="file in attachments"
          :key="file.id"
          class="attachment-item"
        >
          <div class="file-icon pdf">
            <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" fill="white" /></svg>
          </div>
          <div class="file-info">
            <div class="file-name">{{ file.name }}</div>
            <div class="file-size">{{ file.size }}</div>
          </div>
          <button class="download-file-btn">下载</button>
        </div>
        <div class="file-type-icons">
          <div class="type-icon ppt">
            <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2Z" fill="white" /></svg>
          </div>
          <div class="type-icon word">
            <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2Z" fill="white" /></svg>
          </div>
          <div class="type-icon excel">
            <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2Z" fill="white" /></svg>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeTab = ref('catalog')
const showExamModal = ref(false)

const tabs = [
  { label: '目录', value: 'catalog' },
  { label: '考题', value: 'exam' },
  { label: '笔记', value: 'notes' },
  { label: '附件', value: 'attachments' }
]

const startPractice = (type: string) => {
  showExamModal.value = false
  router.push('/exam_practice/1')
}

const lessons = ref([
  {
    id: 1,
    title: '港机设备-起重机械',
    type: '正在播放',
    duration: '09:10',
    isPlaying: true
  },
  {
    id: 2,
    title: '港机设备-起重机械',
    type: '视频',
    duration: '34:20',
    isPlaying: false
  },
  {
    id: 3,
    title: '港机设备-起重机械',
    type: '视频',
    duration: '34:20',
    isPlaying: false
  },
  {
    id: 4,
    title: '港机设备-起重机械',
    type: '视频',
    duration: '34:20',
    isPlaying: false
  }
])

const attachments = ref([
  {
    id: 1,
    name: 'XXXXX资料.pdf',
    size: '219.8KB'
  },
  {
    id: 2,
    name: 'XXXXX资料.pdf',
    size: '219.8KB'
  }
])
</script>

<style scoped>
.course-content-page {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
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
  width: 22px;
  height: 22px;
  fill: #666;
}

.course-banner {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
}

.banner-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.play-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0, 0, 0, 0.3);
}

.play-button {
  width: 60px;
  height: 60px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.play-button svg {
  width: 30px;
  height: 30px;
  margin-left: 4px;
}

.course-info-section {
  background-color: #fff;
  padding: 16px;
}

.course-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0 0 8px 0;
}

.course-tag {
  display: inline-block;
  background-color: #e3f2fd;
  color: #2196f3;
  padding: 4px 12px;
  border-radius: 4px;
  font-size: 13px;
}

.tabs {
  display: flex;
  background-color: #fff;
  border-bottom: 1px solid #f0f0f0;
}

.tab-item {
  flex: 1;
  text-align: center;
  padding: 14px 0;
  font-size: 15px;
  color: #666;
  cursor: pointer;
  position: relative;
}

.tab-item.active {
  color: #3b82f6;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 30px;
  height: 3px;
  background-color: #3b82f6;
  border-radius: 2px;
}

.content-area {
  background-color: #fff;
  min-height: 400px;
}

.chapter-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
}

.chapter-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.switch-course-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 14px;
  cursor: pointer;
}

.dropdown-icon {
  width: 16px;
  height: 16px;
  fill: currentColor;
}

.catalog-content {
  padding-bottom: 16px;
}

.lesson-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  border-bottom: 1px solid #f5f5f5;
}

.lesson-number {
  width: 40px;
  height: 40px;
  background-color: #e3f2fd;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: 600;
  color: #2196f3;
  flex-shrink: 0;
}

.lesson-info {
  flex: 1;
}

.lesson-title {
  font-size: 15px;
  color: #333;
  margin-bottom: 4px;
}

.lesson-meta {
  display: flex;
  gap: 8px;
  font-size: 13px;
  color: #999;
}

.play-btn {
  width: 36px;
  height: 36px;
  background: none;
  border: none;
  cursor: pointer;
}

.play-btn svg {
  width: 100%;
  height: 100%;
  fill: #999;
}

.play-btn.playing svg {
  fill: #3b82f6;
}

.exam-content {
  padding: 16px;
}

.exam-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 12px;
  padding: 16px 0;
}

.exam-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  padding: 24px;
  background-color: #f8f9fa;
  border-radius: 12px;
  cursor: pointer;
}

.exam-card:nth-child(4) {
  grid-column: 1 / -1;
}

.exam-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.exam-icon.blue {
  background-color: #e3f2fd;
  color: #2196f3;
}

.exam-icon.red {
  background-color: #ffebee;
  color: #f44336;
}

.exam-icon.green {
  background-color: #e8f5e9;
  color: #4caf50;
}

.exam-icon.purple {
  background-color: #f3e5f5;
  color: #9c27b0;
}

.exam-icon svg {
  width: 24px;
  height: 24px;
}

.exam-label {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.notes-content {
  padding: 16px;
}

.notes-text {
  padding: 16px 0;
  line-height: 1.8;
}

.notes-text p {
  font-size: 14px;
  color: #333;
  margin: 0 0 12px 0;
}

.attachments-content {
  padding: 16px;
}

.attachment-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid #f5f5f5;
}

.file-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.file-icon.pdf {
  background-color: #f44336;
}

.file-icon svg {
  width: 24px;
  height: 24px;
}

.file-info {
  flex: 1;
}

.file-name {
  font-size: 14px;
  color: #333;
  margin-bottom: 4px;
}

.file-size {
  font-size: 12px;
  color: #999;
}

.download-file-btn {
  padding: 6px 16px;
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 14px;
  cursor: pointer;
}

.file-type-icons {
  display: flex;
  gap: 16px;
  padding: 24px 0;
  justify-content: center;
}

.type-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.type-icon.ppt {
  background-color: #ff6b00;
}

.type-icon.word {
  background-color: #2196f3;
}

.type-icon.excel {
  background-color: #4caf50;
}

.type-icon svg {
  width: 28px;
  height: 28px;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: flex-end;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: #fff;
  border-radius: 16px 16px 0 0;
  width: 100%;
  max-width: 100%;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f0f0;
  margin-bottom: 0;
}

.modal-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.switch-modal-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  color: #3b82f6;
  font-size: 14px;
  cursor: pointer;
}

.question-types {
  display: flex;
  flex-direction: column;
}

.type-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;
  font-size: 15px;
  color: #333;
}

.type-item:active {
  background: #f5f5f5;
}

.count {
  color: #999;
}

.cancel-btn {
  width: 100%;
  padding: 14px;
  background: none;
  border: none;
  border-top: 1px solid #f0f0f0;
  font-size: 15px;
  color: #333;
  cursor: pointer;
  margin-top: 8px;
}
</style>
