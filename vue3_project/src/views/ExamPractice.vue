<template>
  <div class="exam-practice-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">考题训练</span>
      <span class="nav-placeholder"></span>
    </header>

    <!-- 加载中 -->
    <div v-if="loading" class="loading-wrap">
      <p>加载题目中...</p>
    </div>

    <!-- 无题目 -->
    <div v-else-if="questions.length === 0" class="empty-wrap">
      <p>该题型暂无题目</p>
      <button class="back-text-btn" @click="$router.back()">返回</button>
    </div>

    <template v-else>
      <!-- 课程信息区 -->
      <div class="course-info-section">
        <!-- 第一行：课程名 -->
        <h2 class="course-name">{{ courseName || (isMockMode ? '模拟考试' : '考题训练') }}</h2>

        <!-- 第二行：横排题型 tab（左）+ 答题进度（右），左右齐平 -->
        <div class="tabs-progress-row">
          <div class="question-tabs">
            <div
              v-for="tab in questionTypeTabs"
              :key="tab.type"
              :class="['tab-item', { active: activeType === tab.type }]"
              @click="switchType(tab.type)"
            >
              {{ tab.label }}({{ tab.count }})
            </div>
          </div>
          <div class="total-progress-badge">
            <span class="progress-answered">{{ totalAnswered }}</span>
            <span class="progress-total">/{{ allQuestions.length }}</span>
          </div>
        </div>
      </div>

      <main class="exam-content">
        <!-- 白色题目卡片：题型标签 + 答题结果 + 题目文字 -->
        <div class="question-card">
          <div class="question-card-header">
            <span class="question-type-tag">{{ typeLabel }}</span>
            <div v-if="answerResult !== null" class="status-badge" :class="answerResult ? 'correct' : 'error'">
              {{ answerResult ? '正确' : '错误' }}
            </div>
          </div>
          <p class="question-text">
            {{ currentIndex + 1 }}.{{ currentQuestion.content }}
          </p>
        </div>

        <!-- 判断题选项 -->
        <div v-if="activeType === 3" class="options-list">
          <div
            v-for="opt in judgeOptions"
            :key="opt.value"
            :class="['option-item', getOptionClass(opt.value)]"
            @click="selectAnswer(opt.value)"
          >
            <div class="option-content">
              <div class="option-radio" :class="{ selected: isSelected(opt.value), correct: isCorrectOption(opt.value), wrong: isWrongOption(opt.value) }">
                <svg v-if="isSelected(opt.value)" viewBox="0 0 24 24" class="radio-icon">
                  <circle cx="12" cy="12" r="5" />
                </svg>
              </div>
              <span class="option-text">{{ opt.label }}</span>
            </div>
          </div>
        </div>

        <!-- 单选/多选选项 -->
        <div v-else class="options-list">
          <div
            v-for="opt in currentOptions"
            :key="opt.key"
            :class="['option-item', getOptionClass(opt.key)]"
            @click="selectAnswer(opt.key)"
          >
            <div class="option-content">
              <div class="option-radio" :class="{ selected: isSelected(opt.key), correct: isCorrectOption(opt.key), wrong: isWrongOption(opt.key) }">
                <svg v-if="isSelected(opt.key)" viewBox="0 0 24 24" class="radio-icon">
                  <circle cx="12" cy="12" r="5" />
                </svg>
              </div>
              <span class="option-label">{{ opt.key }}</span>
              <span class="option-text">{{ opt.text }}</span>
            </div>
          </div>
        </div>

        <!-- 答案解析区（答题后显示） -->
        <div v-if="answerResult !== null" class="answer-section">
          <div class="correct-answer">
            <span class="label">正确答案：</span>
            <span class="value">{{ currentQuestion.answer }}</span>
          </div>
          <div v-if="currentQuestion.analysis" class="answer-analysis">
            <span class="label">题目解析：</span>
            <p class="analysis-text">{{ currentQuestion.analysis }}</p>
          </div>
        </div>

        <!-- 多选题提交按钮 -->
        <div v-if="activeType === 2 && answerResult === null && multiSelected.length > 0" class="submit-wrap">
          <button class="submit-btn" @click="submitMultiple">提交答案</button>
        </div>
      </main>

      <!-- 答题卡弹窗 -->
      <div v-if="showAnswerCard" class="answer-card-modal" @click="showAnswerCard = false">
        <div class="modal-content" @click.stop>
          <h3 class="modal-title">答题卡</h3>
          <div class="filter-tabs">
            <button class="clear-btn" @click="clearAnswers">
              <svg viewBox="0 0 24 24" class="clear-icon">
                <path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19Z" />
              </svg>
              清空答题卡
            </button>
            <div class="status-indicators">
              <div class="status-item">
                <span class="status-dot answered-dot"></span>
                <span class="status-text">已答</span>
              </div>
              <div class="status-item">
                <span class="status-dot unanswered-dot"></span>
                <span class="status-text">未答</span>
              </div>
            </div>
          </div>
          <div class="question-grid">
            <button
              v-for="(_, idx) in questions"
              :key="idx"
              :class="['question-number', answers[idx] !== undefined ? 'answered' : '']"
              @click="goToQuestion(idx)"
            >
              {{ idx + 1 }}
            </button>
          </div>
        </div>
      </div>

      <!-- 底部导航 -->
      <footer class="bottom-nav">
        <button class="nav-btn prev" @click="previousQuestion" :disabled="currentIndex === 0">
          <svg viewBox="0 0 24 24"><path d="M15.41,16.58L10.83,12L15.41,7.41L14,6L8,12L14,18L15.41,16.58Z" /></svg>
        </button>
        <button class="nav-btn answer-card" @click="showAnswerCard = true">
          <svg viewBox="0 0 24 24"><path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19M17,17H7V15H17V17M17,13H7V11H17V13M17,9H7V7H17V9Z" /></svg>
          <span>答题卡</span>
        </button>
        <button class="nav-btn next" @click="nextQuestion" :disabled="currentIndex === questions.length - 1">
          <svg viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" /></svg>
        </button>
      </footer>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { questionApi, courseApi, mockExamApi, studyApi } from '@/api'

const route = useRoute()

const loading = ref(true)
const courseEk = ref(0)
const courseName = ref('')
const isMockMode = ref(false)  // 是否模拟考试模式
const activeType = ref(1)  // 1=单选 2=多选 3=判断

// 全部题目（按题型分组）
const allQuestions = ref<any[]>([])

// 当前题型的题目
const questions = computed(() =>
  allQuestions.value.filter(q => q.questionType === activeType.value)
)

// 题型tab（只显示有题的）
const questionTypeTabs = computed(() => {
  const map = [
    { type: 1, label: '单选题' },
    { type: 2, label: '多选题' },
    { type: 3, label: '判断题' }
  ]
  return map
    .map(t => ({ ...t, count: allQuestions.value.filter(q => q.questionType === t.type).length }))
    .filter(t => t.count > 0)
})

const typeLabel = computed(() => {
  const m: Record<number, string> = { 1: '单选题', 2: '多选题', 3: '判断题' }
  return m[activeType.value] || ''
})

// 当前题目索引
const currentIndex = ref(0)
const currentQuestion = computed(() => (questions.value[currentIndex.value] || {}) as any)

// 当前题目选项（单选/多选）
const currentOptions = computed(() => {
  const q = currentQuestion.value as any
  const opts: { key: string; text: string }[] = []
  if (q.optionA) opts.push({ key: 'A', text: q.optionA })
  if (q.optionB) opts.push({ key: 'B', text: q.optionB })
  if (q.optionC) opts.push({ key: 'C', text: q.optionC })
  if (q.optionD) opts.push({ key: 'D', text: q.optionD })
  return opts
})

// 判断题选项
const judgeOptions = [
  { value: 'T', label: '正确（T）' },
  { value: 'F', label: '错误（F）' }
]

// 答题记录：按题型分组，每种题型内以 index 为 key
// 多选题临时已选项
const multiSelected = ref<string[]>([])
// 当前题目答题结果
const answerResult = ref<boolean | null>(null)

const showAnswerCard = ref(false)

// 切换题型时重置
const switchType = (type: number) => {
  activeType.value = type
  currentIndex.value = 0
  answerResult.value = null
  multiSelected.value = []
}

// 分题型答题记录
const typeAnswers = ref<Record<number, Record<number, string>>>({ 1: {}, 2: {}, 3: {} })

// 当前题型的答题记录（快捷访问）
const answers = computed(() => typeAnswers.value[activeType.value] || {})

// 所有题型合计已答数
const totalAnswered = computed(() =>
  Object.values(typeAnswers.value).reduce((sum, rec) => sum + Object.keys(rec).length, 0)
)

// 选择答案（单选/判断）
const selectAnswer = (key: string) => {
  if (answerResult.value !== null) return  // 已答过，不能再选
  if (activeType.value === 2) {
    // 多选：toggle
    const idx = multiSelected.value.indexOf(key)
    if (idx >= 0) {
      multiSelected.value.splice(idx, 1)
    } else {
      multiSelected.value.push(key)
    }
    return
  }
  // 单选/判断：立即判断
  const correct = (currentQuestion.value as any).answer
  ;(typeAnswers.value[activeType.value] as Record<number, string>)[currentIndex.value] = key
  answerResult.value = key === correct
}

// 多选提交
const submitMultiple = () => {
  const userAnswer = multiSelected.value.sort().join('')
  const correct = (currentQuestion.value as any).answer
  ;(typeAnswers.value[activeType.value] as Record<number, string>)[currentIndex.value] = userAnswer
  answerResult.value = userAnswer === correct
}

// 选项样式
const isSelected = (key: string) => {
  if (activeType.value === 2) return multiSelected.value.includes(key)
  return answers.value[currentIndex.value] === key
}
const isCorrectOption = (key: string) => {
  if (answerResult.value === null) return false
  const correct = currentQuestion.value.answer || ''
  return correct.includes(key)
}
const isWrongOption = (key: string) => {
  if (answerResult.value === null || answerResult.value) return false
  return isSelected(key) && !isCorrectOption(key)
}
const getOptionClass = (key: string) => {
  if (answerResult.value === null) {
    return { selected: isSelected(key) }
  }
  return {
    correct: isCorrectOption(key),
    wrong: isWrongOption(key)
  }
}

const previousQuestion = () => {
  if (currentIndex.value > 0) {
    currentIndex.value--
    const rec = typeAnswers.value[activeType.value]?.[currentIndex.value]
    const correctAns: string = (questions.value[currentIndex.value] as any)?.answer ?? ''
    answerResult.value = rec !== undefined ? rec === correctAns : null
    multiSelected.value = rec ? rec.split('') : []
  }
}

const nextQuestion = () => {
  if (currentIndex.value < questions.value.length - 1) {
    currentIndex.value++
    const rec = typeAnswers.value[activeType.value]?.[currentIndex.value]
    const correctAns: string = (questions.value[currentIndex.value] as any)?.answer ?? ''
    answerResult.value = rec !== undefined ? rec === correctAns : null
    multiSelected.value = rec ? rec.split('') : []
  }
}

const goToQuestion = (idx: number) => {
  currentIndex.value = idx
  const rec = typeAnswers.value[activeType.value]?.[idx]
  const correctAns: string = (questions.value[idx] as any)?.answer ?? ''
  answerResult.value = rec !== undefined ? rec === correctAns : null
  multiSelected.value = rec ? rec.split('') : []
  showAnswerCard.value = false
}

const clearAnswers = () => {
  typeAnswers.value = { 1: {}, 2: {}, 3: {} }
  answerResult.value = null
  multiSelected.value = []
}

// 切换题目时重置答题状态
watch(currentIndex, () => {
  if ((typeAnswers.value[activeType.value]?.[currentIndex.value]) === undefined) {
    answerResult.value = null
    multiSelected.value = []
  }
})

onMounted(async () => {
  courseEk.value = Number(route.params.id)
  const modeParam = route.query.mode as string
  const typeParam = Number(route.query.type)

  isMockMode.value = modeParam === 'mock'

  if (!isMockMode.value && typeParam && [1, 2, 3].includes(typeParam)) {
    activeType.value = typeParam
  }

  try {
    const [questionsRes, detailRes]: any[] = await Promise.all([
      questionApi.getList(courseEk.value),
      courseApi.getCourseDetail(courseEk.value)
    ])
    const allPool: any[] = questionsRes.data || []
    courseName.value = detailRes.data?.courseName || ''

    if (isMockMode.value) {
      // 模拟考试模式：按配置随机抽题
      const cfgRes: any = await mockExamApi.getConfig(courseEk.value)
      const cfg = cfgRes.data || { singleCount: 0, multipleCount: 0, judgeCount: 0 }

      const shuffle = (arr: any[]) => arr.slice().sort(() => Math.random() - 0.5)
      const singles = shuffle(allPool.filter(q => q.questionType === 1)).slice(0, cfg.singleCount)
      const multiples = shuffle(allPool.filter(q => q.questionType === 2)).slice(0, cfg.multipleCount)
      const judges = shuffle(allPool.filter(q => q.questionType === 3)).slice(0, cfg.judgeCount)
      allQuestions.value = [...singles, ...multiples, ...judges]

      // 默认激活第一个有题的题型
      if (singles.length > 0) activeType.value = 1
      else if (multiples.length > 0) activeType.value = 2
      else if (judges.length > 0) activeType.value = 3
    } else {
      allQuestions.value = allPool
    }
  } catch {
    allQuestions.value = []
  } finally {
    loading.value = false
  }
})

let startTime = 0
onMounted(() => { startTime = Date.now() })

onUnmounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) return
  const duration = Math.floor((Date.now() - startTime) / 1000)
  if (duration < 1) return
  try {
    await studyApi.addRecord(duration, courseEk.value)
  } catch {}
})
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

.nav-progress {
  font-size: 14px;
  color: #999;
  min-width: 48px;
  text-align: right;
}

.loading-wrap, .empty-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
  color: #999;
  font-size: 15px;
  gap: 16px;
}

.back-text-btn {
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 8px;
  padding: 10px 24px;
  font-size: 15px;
  cursor: pointer;
}

.course-info-section {
  background-color: #fff;
  padding: 16px;
  margin-bottom: 12px;
}

/* 课程名 + 右侧进度数字 */
.course-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 12px 0;
}

/* tab 与进度数字同行，左右撑满 */
.tabs-progress-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* 右侧进度：已答/总数 */
.total-progress-badge {
  display: flex;
  align-items: baseline;
  gap: 1px;
  flex-shrink: 0;
}
.progress-answered {
  font-size: 18px;
  font-weight: 700;
  color: #3b82f6;
  line-height: 1;
}
.progress-total {
  font-size: 13px;
  color: #999;
}

.question-tabs {
  display: flex;
  gap: 24px;
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

.status-badge.correct {
  color: #52c41a;
  background-color: #f6ffed;
}

.exam-content {
  padding: 16px;
}

/* 白色题目卡片 */
.question-card {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 12px;
}

.question-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.question-type-tag {
  font-size: 14px;
  font-weight: 600;
  color: #3b82f6;
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
  border: 2px solid transparent;
}

.option-item.correct {
  background-color: #f6ffed;
  border-color: #52c41a;
}

.option-item.wrong {
  background-color: #fff1f0;
  border-color: #ff4d4f;
}

.option-item.selected {
  border-color: #3b82f6;
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
  border: 2px solid #ddd;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: border-color 0.2s;
}

.option-radio.selected {
  border-color: #3b82f6;
}

.option-radio.correct {
  border-color: #52c41a;
  background: #52c41a;
}

.option-radio.wrong {
  border-color: #ff4d4f;
  background: #ff4d4f;
}

.radio-icon {
  width: 10px;
  height: 10px;
  fill: #fff;
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
  margin-bottom: 12px;
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

.submit-wrap {
  margin-top: 16px;
  text-align: center;
}

.submit-btn {
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 24px;
  padding: 12px 40px;
  font-size: 15px;
  cursor: pointer;
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

.nav-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
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

.filter-tabs {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.clear-btn {
  padding: 8px 16px;
  background-color: #fff;
  border: 1px solid #ff9800;
  border-radius: 20px;
  font-size: 14px;
  color: #ff9800;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
}

.clear-icon {
  width: 16px;
  height: 16px;
  fill: #ff9800;
  border-radius: 10px;
}

.status-indicators {
  display: flex;
  gap: 16px;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.status-dot {
  width: 12px;
  height: 12px;
  border-radius: 2px;
}

.answered-dot {
  background-color: #3b82f6;
}

.unanswered-dot {
  background-color: #f5f5f5;
  border: 1px solid #ddd;
}

.status-text {
  font-size: 14px;
  color: #666;
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
  font-weight: 600;
}
</style>
