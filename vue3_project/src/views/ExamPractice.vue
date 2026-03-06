<template>
  <div class="exam-practice-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">考题训练</span>
    </header>

    <!-- 课程信息区 -->
    <div class="course-info-section">
      <h2 class="course-name">港口特种设备检修课程</h2>

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

      <div class="exam-status">
        <div class="question-type-link">
          <a href="#" class="type-link">单选题</a>
        </div>
        <div class="status-badge error">错误</div>
      </div>
    </div>

    <main class="exam-content">
      <!-- 题目内容 -->
      <div class="question-content">
        <p class="question-text">
          {{ currentQuestionData.number }}.{{ currentQuestionData.text }}
        </p>
      </div>

      <!-- 选项列表 -->
      <div class="options-list">
        <div
          v-for="option in currentQuestionData.options"
          :key="option.id"
          :class="['option-item', {
            selected: selectedAnswer === option.id,
            wrong: selectedAnswer === option.id && option.id !== correctAnswer
          }]"
          @click="selectAnswer(option.id)"
        >
          <div class="option-content">
            <div class="option-radio">
              <svg v-if="selectedAnswer === option.id" viewBox="0 0 24 24" class="radio-icon">
                <circle cx="12" cy="12" r="5" />
              </svg>
            </div>
            <span class="option-label">{{ option.label }}</span>
            <span class="option-text">{{ option.text }}</span>
          </div>
        </div>
      </div>

      <!-- 答案解析区 -->
      <div class="answer-section">
        <div class="correct-answer">
          <span class="label">正确答案：</span>
          <span class="value">{{ correctAnswer }}</span>
        </div>
        <div class="answer-analysis">
          <span class="label">错题解析：</span>
          <p class="analysis-text">{{ analysisText }}</p>
        </div>
      </div>
    </main>

    <!-- 答题卡弹窗 -->
    <div v-if="showAnswerCard" class="answer-card-modal" @click="showAnswerCard = false">
      <div class="modal-content" @click.stop>
        <h3 class="modal-title">答题卡</h3>
        <div class="filter-tabs">
          <button :class="['filter-btn', { active: filterType === 'all' }]" @click="filterType = 'all'">
            <svg viewBox="0 0 24 24" class="filter-icon">
              <path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19Z" />
            </svg>
            清空答题卡
          </button>
          <button :class="['filter-btn', { active: filterType === 'answered' }]" @click="filterType = 'answered'">已答</button>
          <button :class="['filter-btn', { active: filterType === 'unanswered' }]" @click="filterType = 'unanswered'">未答</button>
        </div>
        <div class="question-grid">
          <button
            v-for="num in totalQuestions"
            :key="num"
            :class="['question-number', answers[num] ? 'answered' : '']"
            @click="goToQuestion(num)"
          >
            {{ num }}
          </button>
        </div>
      </div>
    </div>

    <!-- 底部导航 -->
    <footer class="bottom-nav">
      <button class="nav-btn prev" @click="previousQuestion">
        <svg viewBox="0 0 24 24"><path d="M15.41,16.58L10.83,12L15.41,7.41L14,6L8,12L14,18L15.41,16.58Z" /></svg>
      </button>
      <button class="nav-btn answer-card" @click="showAnswerCard = true">
        <svg viewBox="0 0 24 24"><path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19M17,17H7V15H17V17M17,13H7V11H17V13M17,9H7V7H17V9Z" /></svg>
        <span>答题卡</span>
      </button>
      <button class="nav-btn next" @click="nextQuestion">
        <svg viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" /></svg>
      </button>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const activeTab = ref('single')
const selectedAnswer = ref('A')
const correctAnswer = ref('B')
const analysisText = ref('XXXXXXXXXXX')
const showAnswerCard = ref(false)
const filterType = ref('all')
const totalQuestions = ref(20)
const answers = ref<Record<number, string>>({ 1: 'A' })

const questionTabs = ref([
  { id: 'single', name: '单选题', count: 20 },
  { id: 'multiple', name: '多选题', count: 10 },
  { id: 'judge', name: '判断题', count: 30 }
])

const currentQuestionData = ref({
  number: 1,
  text: '共识定额去度假啊客户收问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题问题（）。',
  options: [
    { id: 'A', label: 'A', text: '30.5吨' },
    { id: 'B', label: 'B', text: '30.5吨' },
    { id: 'C', label: 'C', text: '30.5吨' },
    { id: 'D', label: 'D', text: '30.5吨' }
  ]
})

const selectAnswer = (answerId: string) => {
  selectedAnswer.value = answerId
}

const previousQuestion = () => {
  // 上一题逻辑
}

const nextQuestion = () => {
  // 下一题逻辑
}

const goToQuestion = (_num: number) => {
  showAnswerCard.value = false
}
</script>

<style scoped>
.exam-practice-page {
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

.course-info-section {
  background-color: #fff;
  padding: 16px;
  margin-bottom: 12px;
}

.course-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 12px 0;
}

.question-tabs {
  display: flex;
  gap: 24px;
  margin-bottom: 12px;
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

.exam-status {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.question-type-link {
  flex: 1;
}

.type-link {
  color: #3b82f6;
  text-decoration: none;
  font-size: 15px;
  font-weight: 500;
}

.status-badge {
  padding: 4px 12px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: 500;
}

.status-badge.error {
  color: #ff4d4f;
  background-color: #fff1f0;
}

.exam-content {
  padding: 16px;
}

.question-content {
  margin-bottom: 20px;
}

.question-text {
  font-size: 15px;
  color: #333;
  line-height: 1.6;
  margin: 0;
}

.options-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.option-item {
  background-color: #fff;
  border-radius: 12px;
  padding: 14px 16px;
  cursor: pointer;
  transition: all 0.2s;
}

.option-item.selected.wrong {
  background-color: #ffe4e6;
}

.option-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.option-radio {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 2px solid #ff4d4f;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.radio-icon {
  width: 10px;
  height: 10px;
  fill: #ff4d4f;
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
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
}

.correct-answer {
  margin-bottom: 12px;
}

.correct-answer .label {
  font-size: 14px;
  color: #3b82f6;
  font-weight: 500;
}

.correct-answer .value {
  font-size: 14px;
  color: #3b82f6;
  font-weight: 600;
}

.answer-analysis .label {
  font-size: 14px;
  color: #666;
  font-weight: 500;
  display: block;
  margin-bottom: 8px;
}

.analysis-text {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0;
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

.answer-card-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0,0,0,0.5);
  display: flex;
  align-items: flex-end;
  z-index: 1000;
}

.modal-content {
  background-color: #fff;
  border-radius: 20px 20px 0 0;
  padding: 20px;
  width: 100%;
  max-height: 80vh;
  overflow-y: auto;
}

.modal-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0 0 16px 0;
  text-align: center;
}

.filter-tabs {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
}

.filter-btn {
  padding: 6px 16px;
  background-color: #f5f5f5;
  border: 1px solid #e0e0e0;
  border-radius: 20px;
  font-size: 14px;
  color: #666;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
}

.filter-btn.active {
  background-color: #fff3e0;
  border-color: #ff9800;
  color: #ff9800;
}

.filter-icon {
  width: 16px;
  height: 16px;
  fill: currentColor;
}

.question-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  gap: 12px;
}

.question-number {
  width: 100%;
  aspect-ratio: 1;
  border: none;
  border-radius: 12px;
  background-color: #f5f5f5;
  color: #333;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.question-number.answered {
  background-color: #e3f2fd;
  color: #3b82f6;
}
</style>

