<template>
  <div class="wrong-question-detail-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="goBack">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">错题记录</span>
      <span class="nav-placeholder"></span>
    </header>

    <!-- 加载中 -->
    <div v-if="loading" class="loading-wrap"><p>加载中...</p></div>

    <!-- 空状态 -->
    <div v-else-if="questions.length === 0" class="loading-wrap"><p>暂无错题</p></div>

    <template v-else>
      <main class="detail-content">
        <div class="header-section">
        <!-- 课程名称 -->
        <div class="course-name">{{ courseName }}</div>

        <!-- 题目类型标签 -->
        <div class="question-tabs">
          <div
            v-for="tab in tabs"
            :key="tab.type"
            :class="['tab-item', { active: activeType === tab.type }]"
            @click="switchType(tab.type)"
          >
            {{ tab.label }}({{ tab.count }})
          </div>
        </div>

        <!-- 题目内容 -->
        </div>
        <div class="question-card">
          <div class="question-header">
            <span class="question-type">{{ typeLabel }}</span>
            <span class="question-status error">错误</span>
          </div>
          <div class="question-text">
            {{ currentIndex + 1 }}.{{ currentQuestion.content }}
          </div>
        </div>

        <!-- 判断题选项 -->
        <div v-if="currentQuestion.questionType === 3" class="options-list">
          <div v-for="opt in judgeOptions" :key="opt.value" class="option-item">
            <div class="option-radio"></div>
            <span class="option-text">{{ opt.label }}</span>
          </div>
        </div>

        <!-- 单选/多选选项 -->
        <div v-else class="options-list">
          <div v-for="opt in currentOptions" :key="opt.key" class="option-item">
            <div class="option-radio"></div>
            <span class="option-label">{{ opt.key }}</span>
            <span class="option-text">{{ opt.text }}</span>
          </div>
        </div>

        <!-- 正确答案 + 解析 -->
        <div class="answer-section">
          <div class="correct-answer">
            <span class="label">正确答案：</span>
            <span class="value">{{ formatAnswerDisplay(currentQuestion.answer) }}</span>
          </div>
          <div v-if="currentQuestion.analysis" class="answer-analysis">
            <span class="label">错题解析：</span>
            <p class="analysis-text">{{ currentQuestion.analysis }}</p>
          </div>
        </div>
      </main>

      <!-- 底部导航 -->
      <footer class="bottom-nav">
        <button class="nav-btn prev" :disabled="currentIndex === 0" @click="prev">
          <svg viewBox="0 0 24 24"><path d="M15.41,16.58L10.83,12L15.41,7.41L14,6L8,12L14,18L15.41,16.58Z" /></svg>
        </button>
        <button class="nav-btn answer-card" @click="showAnswerCard = true">
          <svg viewBox="0 0 24 24"><path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19M17,17H7V15H17V17M17,13H7V11H17V13M17,9H7V7H17V9Z" /></svg>
          <span>答题卡</span>
        </button>
        <button class="nav-btn next" :disabled="currentIndex === currentTypeQuestions.length - 1" @click="next">
          <svg viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" /></svg>
        </button>
      </footer>

      <!-- 答题卡弹窗 -->
      <div v-if="showAnswerCard" class="answer-card-modal" @click="showAnswerCard = false">
        <div class="modal-content" @click.stop>
          <h3 class="modal-title">答题卡</h3>
          <div class="question-grid">
            <button
              v-for="(_, idx) in currentTypeQuestions"
              :key="idx"
              :class="['question-number', 'answered', { current: currentIndex === idx }]"
              @click="goTo(idx)"
            >
              {{ idx + 1 }}
            </button>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { mockExamApi, courseApi } from '@/api'

const route = useRoute()
const router = useRouter()

const goBack = () => {
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push(`/course_content/${route.params.id}`)
  }
}
const loading = ref(true)
const questions = ref<any[]>([])
const courseName = ref('')
const activeType = ref(1)
const currentIndex = ref(0)
const showAnswerCard = ref(false)

const judgeOptions = [
  { value: 'T', label: '正确（T）' },
  { value: 'F', label: '错误（F）' }
]

const formatAnswerDisplay = (answer: string) => {
  const normalized = String(answer || '').toUpperCase()
  if (normalized === 'T') return '正确'
  if (normalized === 'F') return '错误'
  return answer || ''
}

const tabs = computed(() => {
  const types = [
    { type: 1, label: '单选题' },
    { type: 2, label: '多选题' },
    { type: 3, label: '判断题' }
  ]
  return types
    .map(t => ({ ...t, count: questions.value.filter(q => q.questionType === t.type).length }))
    .filter(t => t.count > 0)
})

const currentTypeQuestions = computed(() =>
  questions.value.filter(q => q.questionType === activeType.value)
)

const currentQuestion = computed(() => currentTypeQuestions.value[currentIndex.value] || {})

const typeLabel = computed(() => {
  const m: Record<number, string> = { 1: '单选题', 2: '多选题', 3: '判断题' }
  return m[activeType.value] || ''
})

const currentOptions = computed(() => {
  const q = currentQuestion.value
  const opts: { key: string; text: string }[] = []
  if (q.optionA) opts.push({ key: 'A', text: q.optionA })
  if (q.optionB) opts.push({ key: 'B', text: q.optionB })
  if (q.optionC) opts.push({ key: 'C', text: q.optionC })
  if (q.optionD) opts.push({ key: 'D', text: q.optionD })
  return opts
})

const switchType = (type: number) => {
  activeType.value = type
  currentIndex.value = 0
  showAnswerCard.value = false
}

const prev = () => { if (currentIndex.value > 0) currentIndex.value-- }
const next = () => { if (currentIndex.value < currentTypeQuestions.value.length - 1) currentIndex.value++ }
const goTo = (idx: number) => { currentIndex.value = idx; showAnswerCard.value = false }

onMounted(async () => {
  const courseEk = Number(route.params.id)
  try {
    const [wrongRes, detailRes]: any[] = await Promise.all([
      mockExamApi.getWrongQuestionList(courseEk),
      courseApi.getCourseDetail(courseEk)
    ])
    questions.value = wrongRes.data || []
    courseName.value = detailRes.data?.courseName || ''
    // 默认选第一个有题的类型
    const firstTab = [1, 2, 3].find(t => questions.value.some(q => q.questionType === t))
    if (firstTab) activeType.value = firstTab
  } catch {
    questions.value = []
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.wrong-question-detail-page {
  background-color: #f5f5f5;
  position: fixed;
  inset: 0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  flex-shrink: 0;
}

.back-btn {
  background: none;
  border: none;
  padding: 0;
  width: 24px;
  height: 24px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
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

.nav-placeholder {
  width: 24px;
  height: 24px;
  flex-shrink: 0;
}

.loading-wrap {
  flex: 1;
  min-height: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
  font-size: 15px;
}

.detail-content {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  overscroll-behavior-y: contain;
  -webkit-overflow-scrolling: touch;
  padding: 0 15px 15px;
}

.header-section {
  background-color: #fff;
  margin: 0 -15px 12px;
  padding: 8px 15px 0;
}

.course-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 12px;
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

.question-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 12px;
}

.question-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.question-type {
  font-size: 14px;
  font-weight: 600;
  color: #3b82f6;
}

.question-status.error {
  font-size: 13px;
  color: #ff4d4f;
  font-weight: 600;
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
  gap: 10px;
  margin-bottom: 12px;
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
  background-color: #fff;
  border-radius: 12px;
  padding: 14px 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.correct-answer {
}

.correct-answer .label,
.answer-analysis .label {
  font-size: 14px;
  color: #666;
}

.correct-answer .value {
  font-size: 14px;
  color: #3b82f6;
  font-weight: 500;
}

.answer-analysis .label {
  display: block;
  margin-bottom: 4px;
}

.analysis-text {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0;
}

.bottom-nav {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  background-color: #fff;
  border-top: 1px solid #eee;
  flex-shrink: 0;
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

.nav-btn:disabled {
  opacity: 0.3;
}

.answer-card-modal {
  position: fixed;
  inset: 0;
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
  height: 85vh;
  max-height: 85vh;
  overflow-y: auto;
}

.modal-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0 0 16px 0;
  text-align: center;
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
  border-radius: 10px;
  background-color: #f5f5f5;
  color: #333;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.question-number.answered {
  background-color: #e3f2fd;
  color: #3b82f6;
  font-weight: 600;
}

.question-number.current {
  outline: 2px solid #3b82f6;
}
</style>
