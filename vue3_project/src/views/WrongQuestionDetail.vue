<template>
  <div class="wrong-question-detail-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">错题记录</span>
    </header>

    <main class="detail-content">
      <!-- 课程名称 -->
      <div class="course-name">港口特种设备检修课程</div>

      <!-- 题目类型标签 -->
      <div class="question-tabs">
        <div
          v-for="tab in questionTabs"
          :key="tab.id"
          :class="['tab-item', { active: activeTab === tab.id }]"
          @click="activeTab = tab.id"
        >
          {{ tab.name }}({{ tab.count }})
        </div>
      </div>

      <!-- 题目内容 -->
      <div class="question-content">
        <div class="question-header">
          <span class="question-type">单选题</span>
          <span class="question-status error">错误</span>
        </div>

        <div class="question-text">
          1.共识定额去度假啊客户收问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题（）。
        </div>

        <div class="options-list">
          <div
            v-for="option in options"
            :key="option.id"
            class="option-item"
          >
            <div class="option-radio"></div>
            <span class="option-label">{{ option.label }}</span>
            <span class="option-text">{{ option.text }}</span>
          </div>
        </div>

        <div class="answer-section">
          <div class="correct-answer">
            <span class="label">正确答案：</span>
            <span class="value">A、B、C、D</span>
          </div>
          <div class="answer-analysis">
            <span class="label">错题解析：</span>
            <p class="analysis-text">
              XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
            </p>
          </div>
        </div>
      </div>
    </main>

    <!-- 底部导航 -->
    <footer class="bottom-nav">
      <button class="nav-btn prev">
        <svg viewBox="0 0 24 24"><path d="M15.41,16.58L10.83,12L15.41,7.41L14,6L8,12L14,18L15.41,16.58Z" /></svg>
      </button>
      <button class="nav-btn answer-card">
        <svg viewBox="0 0 24 24"><path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19M17,17H7V15H17V17M17,13H7V11H17V13M17,9H7V7H17V9Z" /></svg>
        <span>答题卡</span>
      </button>
      <button class="nav-btn next">
        <svg viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" /></svg>
      </button>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const activeTab = ref('single')

const questionTabs = ref([
  { id: 'single', name: '单选题', count: 20 },
  { id: 'multiple', name: '多选题', count: 10 },
  { id: 'judge', name: '判断题', count: 30 }
])

const options = ref([
  { id: 'A', label: 'A', text: '30.5吨' },
  { id: 'B', label: 'B', text: '30.5吨' },
  { id: 'C', label: 'C', text: '30.5吨' },
  { id: 'D', label: 'D', text: '30.5吨' }
])
</script>

<style scoped>
.wrong-question-detail-page {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 80px;
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

.detail-content {
  padding: 15px;
}

.course-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 15px;
}

.question-tabs {
  display: flex;
  gap: 24px;
  background-color: #fff;
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 15px;
}

.tab-item {
  font-size: 14px;
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

.question-content {
  background-color: #fff;
  border-radius: 8px;
  padding: 16px;
}

.question-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.question-type {
  font-size: 15px;
  color: #3b82f6;
  font-weight: 600;
}

.question-status {
  font-size: 14px;
  font-weight: 500;
}

.question-status.error {
  color: #ff4d4f;
}

.question-text {
  font-size: 15px;
  color: #333;
  line-height: 1.6;
  margin-bottom: 20px;
}

.options-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.option-item {
  background-color: #ffe4e6;
  border-radius: 8px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.option-radio {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  border: 2px solid #ff4d4f;
  flex-shrink: 0;
}

.option-label {
  font-size: 15px;
  font-weight: 600;
  color: #333;
}

.option-text {
  font-size: 15px;
  color: #333;
}

.answer-section {
  border-top: 1px solid #eee;
  padding-top: 15px;
}

.correct-answer {
  margin-bottom: 15px;
}

.correct-answer .label {
  font-size: 14px;
  color: #666;
}

.correct-answer .value {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.answer-analysis .label {
  font-size: 14px;
  color: #666;
  display: block;
  margin-bottom: 8px;
}

.analysis-text {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0;
  word-break: break-all;
}

.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  background-color: #fff;
  border-top: 1px solid #eee;
}

.nav-btn {
  background: none;
  border: none;
  padding: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
  color: #666;
}

.nav-btn svg {
  width: 24px;
  height: 24px;
  fill: currentColor;
}

.nav-btn.answer-card {
  flex-direction: column;
  gap: 2px;
}

.nav-btn.answer-card span {
  font-size: 12px;
}

.nav-btn.prev,
.nav-btn.next {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 1px solid #ddd;
  justify-content: center;
}
</style>
