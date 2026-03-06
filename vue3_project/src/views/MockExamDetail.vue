<template>
  <div class="mock-exam-detail-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">模拟考试</span>
    </header>

    <!-- 课程信息区 -->
    <div class="course-info-section">
      <h2 class="course-name">港口特种设备检修课程</h2>
      <div class="exam-status">
        <div class="timer">
          <svg viewBox="0 0 24 24" class="timer-icon">
            <path d="M12,20A7,7 0 0,1 5,13A7,7 0 0,1 12,6A7,7 0 0,1 19,13A7,7 0 0,1 12,20M19.03,7.39L20.45,5.97C20,5.46 19.55,5 19.04,4.56L17.62,6C16.07,4.74 14.12,4 12,4A9,9 0 0,0 3,13A9,9 0 0,0 12,22C17,22 21,18 21,13C21,10.88 20.26,8.93 19.03,7.39M11,14H13V8H11M15,1H9V3H15V1Z" />
          </svg>
          <span class="time-text">{{ timeRemaining }}</span>
        </div>
        <div class="progress-info">
          <span class="progress-text">{{ currentQuestion }}/{{ totalQuestions }}</span>
        </div>
      </div>
      <div class="progress-bar">
        <div class="progress-fill" :style="{ width: progressPercentage + '%' }"></div>
      </div>
    </div>

    <main class="exam-content">
      <!-- 题目类型 -->
      <div class="question-type-link">
        <a href="#" class="type-link">单选题</a>
      </div>

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
          :class="['option-item', { selected: selectedAnswer === option.id }]"
          @click="selectAnswer(option.id)"
        >
          <div class="option-content">
            <div class="option-radio">
              <svg v-if="selectedAnswer === option.id" viewBox="0 0 24 24" class="check-icon">
                <path d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z" />
              </svg>
            </div>
            <span class="option-label">{{ option.label }}</span>
            <span class="option-text">{{ option.text }}</span>
          </div>
        </div>
      </div>
    </main>

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

    <!-- 答题卡弹窗 -->
    <div v-if="showAnswerCard" class="answer-card-modal" @click="showAnswerCard = false">
      <div class="modal-content" @click.stop>
        <h3 class="modal-title">答题卡</h3>

        <!-- 筛选标签 -->
        <div class="filter-tabs">
          <button :class="['filter-btn', { active: filterType === 'all' }]" @click="filterType = 'all'">
            <svg viewBox="0 0 24 24" class="filter-icon">
              <path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19Z" />
            </svg>
            清空答题卡
          </button>
          <button :class="['filter-btn', { active: filterType === 'answered' }]" @click="filterType = 'answered'">
            已答
          </button>
          <button :class="['filter-btn', { active: filterType === 'unanswered' }]" @click="filterType = 'unanswered'">
            未答
          </button>
        </div>

        <!-- 题号网格 -->
        <div class="question-grid">
          <button
            v-for="num in totalQuestions"
            :key="num"
            :class="['question-number', getQuestionStatus(num)]"
            @click="goToQuestion(num)"
          >
            {{ num }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const timeRemaining = ref('29:45')
const currentQuestion = ref(10)
const totalQuestions = ref(20)
const selectedAnswer = ref('A')
const showAnswerCard = ref(false)
const filterType = ref('all')

const progressPercentage = computed(() => {
  return (currentQuestion.value / totalQuestions.value) * 100
})

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

// 答题记录
const answers = ref<Record<number, string>>({
  1: 'A',
  2: 'B',
  3: 'C'
})

const selectAnswer = (answerId: string) => {
  selectedAnswer.value = answerId
  answers.value[currentQuestion.value] = answerId
}

const previousQuestion = () => {
  if (currentQuestion.value > 1) {
    currentQuestion.value--
  }
}

const nextQuestion = () => {
  if (currentQuestion.value < totalQuestions.value) {
    currentQuestion.value++
  }
}

const goToQuestion = (num: number) => {
  currentQuestion.value = num
  showAnswerCard.value = false
}

const getQuestionStatus = (num: number) => {
  if (answers.value[num]) {
    return 'answered'
  }
  return ''
}
</script>

<style scoped>
.mock-exam-detail-page {
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

.exam-status {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.timer {
  display: flex;
  align-items: center;
  gap: 4px;
}

.timer-icon {
  width: 18px;
  height: 18px;
  fill: #ff4d4f;
}

.time-text {
  font-size: 16px;
  color: #ff4d4f;
  font-weight: 600;
}

.progress-info {
  font-size: 16px;
  color: #3b82f6;
  font-weight: 600;
}

.progress-bar {
  height: 8px;
  background-color: #e0e0e0;
  border-radius: 4px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: #3b82f6;
  transition: width 0.3s;
}

.exam-content {
  padding: 16px;
}

.question-type-link {
  margin-bottom: 16px;
}

.type-link {
  color: #3b82f6;
  text-decoration: none;
  font-size: 15px;
  font-weight: 500;
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
}

.option-item {
  background-color: #fff;
  border-radius: 12px;
  padding: 14px 16px;
  cursor: pointer;
  transition: all 0.2s;
}

.option-item.selected {
  background-color: #e3f2fd;
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
  border: 2px solid #3b82f6;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.check-icon {
  width: 14px;
  height: 14px;
  fill: #3b82f6;
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
