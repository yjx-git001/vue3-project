<template>
  <div class="mock-exam-detail-page">
    <!-- 加载中 -->
    <div v-if="loading" class="loading-wrap">
      <p>加载题目中...</p>
    </div>

    <!-- 无题目 -->
    <div v-else-if="questions.length === 0" class="empty-wrap">
      <p>该课程暂未配置模拟考试题目</p>
      <button class="back-text-btn" @click="$router.back()">返回</button>
    </div>

    <!-- 交卷结果页 -->
    <div v-else-if="submitted" class="result-page">

      <!-- 顶部导航 -->
      <header class="top-nav">
        <button class="back-btn" @click="$router.back()">
          <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
        </button>
        <span class="nav-title">考试结果</span>
        <span class="nav-placeholder"></span>
      </header>

      <!-- 得分卡片 -->
      <div class="result-score-card">
        <div class="score-big">{{ score }}</div>
        <div class="score-desc">最终得分</div>
        <div class="result-stats">
          <div class="stat-col">
            <span class="stat-num">{{ questions.length }}</span>
            <span class="stat-label">总题数</span>
          </div>
          <div class="stat-col">
            <span class="stat-num">{{ correctCount }}</span>
            <span class="stat-label">正确数</span>
          </div>
          <div class="stat-col">
            <span class="stat-num">{{ questions.length - correctCount }}</span>
            <span class="stat-label">错误数</span>
          </div>
        </div>
      </div>

      <!-- 错题回顾 -->
      <div class="wrong-section">
        <div v-if="wrongQuestions.length > 0">
          <!-- 「错题回顾」标题在所有卡片外面上方 -->
          <div class="wrong-section-title">错题回顾</div>

          <div
            v-for="(item, idx) in wrongQuestions"
            :key="item.q.id || idx"
            class="wrong-item-wrap"
          >
            <!-- 一个白色块：题目 + 选项 + 正确答案 + 解析 -->
            <div class="wrong-item">
              <!-- 单选题（蓝）  错误（红）两端对齐 -->
              <div class="wrong-item-top">
                <span class="wrong-type-tag">{{ typeLabel2(item.q.questionType) }}</span>
                <span class="wrong-badge">错误</span>
              </div>

              <!-- 题目文字 -->
              <p class="wrong-q-text">{{ idx + 1 }}.{{ item.q.content }}</p>

              <!-- 判断题选项 -->
              <div v-if="item.q.questionType === 3" class="wrong-options">
                <div v-for="opt in judgeOptions" :key="opt.value" class="wrong-opt">
                  <span class="wrong-opt-dot"></span>
                  <span>{{ opt.label }}</span>
                </div>
              </div>

              <!-- 单选/多选选项 -->
              <div v-else class="wrong-options">
                <div v-for="opt in getOptions(item.q)" :key="opt.key" class="wrong-opt">
                  <span class="wrong-opt-dot"></span>
                  <span class="wrong-opt-key">{{ opt.key }}</span>
                  <span>{{ opt.text }}</span>
                </div>
              </div>

              <!-- 正确答案 + 解析 -->
              <div class="wrong-answer-card">
                <div class="wrong-correct-ans">正确答案：{{ item.q.answer }}</div>
                <div v-if="item.q.analysis" class="wrong-analysis">错题解析：{{ item.q.analysis }}</div>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="all-correct-tip">🎉 全部答对，太棒了！</div>
      </div>

      <!-- 底部继续考试按钮 -->
      <div class="result-footer">
        <button class="retry-btn" @click="retryExam">继续考试</button>
      </div>
    </div>

    <!-- 答题页 -->
    <template v-else>
      <!-- 顶部导航 -->
      <header class="top-nav">
        <button class="back-btn" @click="confirmBack">
          <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
        </button>
        <span class="nav-title">模拟考试</span>
        <span class="nav-placeholder"></span>
      </header>

      <!-- 课程信息区 -->
      <div class="course-info-section">
        <h2 class="course-name">{{ courseName }}</h2>
        <div class="exam-meta-row">
          <div class="timer" :class="{ urgent: timeLeft <= 300 }">
            <svg viewBox="0 0 24 24" class="timer-icon">
              <path d="M12,20A7,7 0 0,1 5,13A7,7 0 0,1 12,6A7,7 0 0,1 19,13A7,7 0 0,1 12,20M19.03,7.39L20.45,5.97C20,5.46 19.55,5 19.04,4.56L17.62,6C16.07,4.74 14.12,4 12,4A9,9 0 0,0 3,13A9,9 0 0,0 12,22C17,22 21,18 21,13C21,10.88 20.26,8.93 19.03,7.39M11,14H13V8H11M15,1H9V3H15V1Z" />
            </svg>
            <span class="time-text">{{ timeDisplay }}</span>
          </div>
          <div class="progress-info">
            <span class="answered-count">{{ answeredCount }}</span>
            <span class="total-count">/{{ questions.length }}</span>
          </div>
        </div>
        <div class="progress-bar">
          <div class="progress-fill" :style="{ width: (answeredCount / questions.length * 100) + '%' }"></div>
        </div>
      </div>

      <main class="exam-content">
        <!-- 题目卡片 -->
        <div class="question-card">
          <div class="question-card-header">
            <span class="question-type-label">{{ typeLabel }}</span>
            <div v-if="answerResult !== null" class="status-badge" :class="answerResult ? 'correct' : 'error'">
              {{ answerResult ? '正确' : '错误' }}
            </div>
          </div>
          <p class="question-text">{{ currentIndex + 1 }}. {{ currentQuestion.content }}</p>
        </div>

        <!-- 判断题选项 -->
        <div v-if="currentQuestion.questionType === 3" class="options-list">
          <div
            v-for="opt in judgeOptions"
            :key="opt.value"
            :class="['option-item', { selected: answers[currentIndex] === opt.value }]"
            @click="selectAnswer(opt.value)"
          >
            <div class="option-content">
              <div class="option-radio" :class="{ selected: answers[currentIndex] === opt.value }">
                <svg v-if="answers[currentIndex] === opt.value" viewBox="0 0 24 24" class="radio-icon">
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
            :class="['option-item', { selected: isSelected(opt.key) }]"
            @click="selectAnswer(opt.key)"
          >
            <div class="option-content">
              <div class="option-radio" :class="{ selected: isSelected(opt.key) }">
                <svg v-if="isSelected(opt.key)" viewBox="0 0 24 24" class="radio-icon">
                  <circle cx="12" cy="12" r="5" />
                </svg>
              </div>
              <span class="option-label">{{ opt.key }}</span>
              <span class="option-text">{{ opt.text }}</span>
            </div>
          </div>
        </div>

        <!-- 多选题提交按钮 -->
        <div v-if="currentQuestion.questionType === 2 && !answers[currentIndex] && multiSelected.length > 0" class="submit-wrap">
          <button class="submit-btn" @click="submitMultiple">确认选择</button>
        </div>
      </main>

      <!-- 答题卡弹窗 -->
      <div v-if="showAnswerCard" class="answer-card-modal" @click="showAnswerCard = false">
        <div class="modal-content" @click.stop>
          <h3 class="modal-title">答题卡</h3>
          <div class="filter-tabs">
            <button class="clear-btn" @click="clearAnswers">
              <svg viewBox="0 0 24 24" class="clear-icon"><path d="M19,3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V5H19V19Z" /></svg>
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
              :class="['question-number', answers[idx] ? 'answered' : '', currentIndex === idx ? 'current' : '']"
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
        <button v-if="currentIndex < questions.length - 1" class="nav-btn next" @click="nextQuestion">
          <svg viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" /></svg>
        </button>
        <button v-else class="nav-btn submit-final" @click="handleSubmit">交卷</button>
      </footer>

      <!-- 自定义确认弹窗 -->
      <div v-if="dialog.show" class="dialog-mask">
        <div class="dialog-box">
          <p class="dialog-msg">{{ dialog.msg }}</p>
          <div class="dialog-btns">
            <button class="dialog-cancel" @click="dialog.show = false">取消</button>
            <button class="dialog-confirm" @click="dialog.onConfirm">确定</button>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { questionApi, courseApi, mockExamApi, studyApi } from '@/api'

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const submitted = ref(false)

// 自定义确认弹窗
const dialog = ref({ show: false, msg: '', onConfirm: () => {} })
const showDialog = (msg: string, onConfirm: () => void) => {
  dialog.value = { show: true, msg, onConfirm: () => { dialog.value.show = false; onConfirm() } }
}
const courseEk = ref(0)
const courseName = ref('')
const questions = ref<any[]>([])
const currentIndex = ref(0)
const showAnswerCard = ref(false)

// 答题记录：index -> 答案字符串
const answers = ref<Record<number, string>>({})
// 多选临时选项
const multiSelected = ref<string[]>([])
// 当前题目答题结果
const answerResult = ref<boolean | null>(null)
let autoNextTimer: ReturnType<typeof setTimeout> | null = null

// 倒计时：30分钟 = 1800秒
const timeLeft = ref(1800)
let timer: ReturnType<typeof setInterval> | null = null

const timeDisplay = computed(() => {
  const m = Math.floor(timeLeft.value / 60).toString().padStart(2, '0')
  const s = (timeLeft.value % 60).toString().padStart(2, '0')
  return `${m}:${s}`
})

const currentQuestion = computed(() => questions.value[currentIndex.value] || {})

const typeLabel = computed(() => {
  const m: Record<number, string> = { 1: '单选题', 2: '多选题', 3: '判断题' }
  return m[currentQuestion.value.questionType] || ''
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

const judgeOptions = [
  { value: 'T', label: '正确（T）' },
  { value: 'F', label: '错误（F）' }
]

const answeredCount = computed(() => Object.keys(answers.value).length)

const isSelected = (key: string) => {
  if (currentQuestion.value.questionType === 2) return multiSelected.value.includes(key)
  return answers.value[currentIndex.value] === key
}

const clearAutoNextTimer = () => {
  if (autoNextTimer !== null) {
    clearTimeout(autoNextTimer)
    autoNextTimer = null
  }
}

const scheduleAutoNext = () => {
  clearAutoNextTimer()
  autoNextTimer = setTimeout(() => {
    if (currentIndex.value < questions.value.length - 1) {
      nextQuestion()
    }
  }, 1000)
}

const selectAnswer = (key: string) => {
  if (answerResult.value !== null) return
  if (currentQuestion.value.questionType === 2) {
    // 多选：toggle
    const idx = multiSelected.value.indexOf(key)
    if (idx >= 0) multiSelected.value.splice(idx, 1)
    else multiSelected.value.push(key)
    return
  }
  // 单选/判断：显示解析后自动下一题
  const correct = String(currentQuestion.value.answer || '')
  answers.value[currentIndex.value] = key
  answerResult.value = key === correct
  scheduleAutoNext()
}

const submitMultiple = () => {
  if (answerResult.value !== null || multiSelected.value.length === 0) return
  const userAnswer = multiSelected.value.slice().sort().join('')
  const correct = String(currentQuestion.value.answer || '')
  answers.value[currentIndex.value] = userAnswer
  multiSelected.value = userAnswer.split('')
  answerResult.value = userAnswer === correct
  scheduleAutoNext()
}

const previousQuestion = () => {
  if (currentIndex.value > 0) {
    clearAutoNextTimer()
    currentIndex.value--
    syncCurrentQuestionState()
  }
}

const nextQuestion = () => {
  if (currentIndex.value < questions.value.length - 1) {
    clearAutoNextTimer()
    currentIndex.value++
    syncCurrentQuestionState()
  }
}

const goToQuestion = (idx: number) => {
  clearAutoNextTimer()
  currentIndex.value = idx
  syncCurrentQuestionState()
  showAnswerCard.value = false
}

const syncCurrentQuestionState = () => {
  const rec = answers.value[currentIndex.value]
  const correct = String(currentQuestion.value.answer || '')
  answerResult.value = rec !== undefined ? rec === correct : null
  if (currentQuestion.value.questionType === 2 && rec) {
    multiSelected.value = rec.split('')
  } else {
    multiSelected.value = []
  }
}

// 计算得分
const correctCount = computed(() => {
  return questions.value.filter((q, idx) => {
    const userAns = answers.value[idx] || ''
    return userAns === q.answer
  }).length
})

const score = computed(() => {
  if (questions.value.length === 0) return 0
  return Math.round((correctCount.value / questions.value.length) * 100)
})

// 错题列表
const wrongQuestions = computed(() =>
  questions.value
    .map((q, idx) => ({ q, userAns: answers.value[idx] || '' }))
    .filter(item => item.userAns !== item.q.answer)
)

const typeLabel2 = (type: number) => {
  const m: Record<number, string> = { 1: '单选题', 2: '多选题', 3: '判断题' }
  return m[type] || ''
}

const getOptions = (q: any) => {
  const opts: { key: string; text: string }[] = []
  if (q.optionA) opts.push({ key: 'A', text: q.optionA })
  if (q.optionB) opts.push({ key: 'B', text: q.optionB })
  if (q.optionC) opts.push({ key: 'C', text: q.optionC })
  if (q.optionD) opts.push({ key: 'D', text: q.optionD })
  return opts
}

// 继续考试：重置所有状态重新抽题
const retryExam = async () => {
  submitted.value = false
  clearAutoNextTimer()
  answers.value = {}
  multiSelected.value = []
  answerResult.value = null
  currentIndex.value = 0
  timeLeft.value = 1800
  try {
    const [questionsRes, cfgRes]: any[] = await Promise.all([
      (await import('@/api')).questionApi.getList(courseEk.value),
      mockExamApi.getConfig(courseEk.value)
    ])
    const allPool: any[] = questionsRes.data || []
    const cfg = cfgRes.data || { singleCount: 0, multipleCount: 0, judgeCount: 0 }
    const shuffle = (arr: any[]) => arr.slice().sort(() => Math.random() - 0.5)
    const singles = shuffle(allPool.filter(q => q.questionType === 1)).slice(0, cfg.singleCount)
    const multiples = shuffle(allPool.filter(q => q.questionType === 2)).slice(0, cfg.multipleCount)
    const judges = shuffle(allPool.filter(q => q.questionType === 3)).slice(0, cfg.judgeCount)
    questions.value = [...singles, ...multiples, ...judges]
    if (questions.value.length > 0) {
      syncCurrentQuestionState()
      startTimer()
    }
  } catch {
    questions.value = []
  }
}

const clearAnswers = () => {
  clearAutoNextTimer()
  answers.value = {}
  multiSelected.value = []
  answerResult.value = null
}

const doSubmit = async () => {
  stopTimer()
  showAnswerCard.value = false
  // 保存考试记录
  try {
    await mockExamApi.saveRecord({
      courseEk: courseEk.value,
      score: score.value,
      total: questions.value.length,
      correct: correctCount.value
    })
  } catch { /* 静默失败 */ }
  // 保存错题记录
  try {
    const wrongIds = questions.value
      .filter((q, idx) => (answers.value[idx] || '') !== q.answer)
      .map((q: any) => q.id)
      .filter((id: any) => !!id)
    if (wrongIds.length > 0) {
      await mockExamApi.saveWrongQuestions({ courseEk: courseEk.value, questionIds: wrongIds })
    }
  } catch { /* 静默失败 */ }
  submitted.value = true
}

const handleSubmit = () => {
  if (answeredCount.value < questions.value.length) {
    const unanswered = questions.value.length - answeredCount.value
    showDialog(`还有 ${unanswered} 题未作答，确定交卷吗？`, () => {
      doSubmit()
    })
    return
  }
  doSubmit()
}

const confirmBack = () => {
  showDialog('确定退出考试吗？当前答题进度将丢失。', () => {
    stopTimer()
    router.back()
  })
}

const startTimer = () => {
  timer = setInterval(() => {
    timeLeft.value--
    if (timeLeft.value <= 0) {
      doSubmit()
    }
  }, 1000)
}

const stopTimer = () => {
  if (timer) {
    clearInterval(timer)
    timer = null
  }
}

let startTime = 0

onMounted(async () => {
  startTime = Date.now()
  courseEk.value = Number(route.params.id)
  try {
    const [questionsRes, detailRes, cfgRes]: any[] = await Promise.all([
      questionApi.getList(courseEk.value),
      courseApi.getCourseDetail(courseEk.value),
      mockExamApi.getConfig(courseEk.value)
    ])
    const allPool: any[] = questionsRes.data || []
    courseName.value = detailRes.data?.courseName || '模拟考试'
    const cfg = cfgRes.data || { singleCount: 0, multipleCount: 0, judgeCount: 0 }

    const shuffle = (arr: any[]) => arr.slice().sort(() => Math.random() - 0.5)
    const singles = shuffle(allPool.filter(q => q.questionType === 1)).slice(0, cfg.singleCount)
    const multiples = shuffle(allPool.filter(q => q.questionType === 2)).slice(0, cfg.multipleCount)
    const judges = shuffle(allPool.filter(q => q.questionType === 3)).slice(0, cfg.judgeCount)
    questions.value = [...singles, ...multiples, ...judges]

    if (questions.value.length > 0) {
      syncCurrentQuestionState()
      startTimer()
    }
  } catch {
    questions.value = []
  } finally {
    loading.value = false
  }
})

onUnmounted(async () => {
  stopTimer()
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
.mock-exam-detail-page {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 80px;
}

/* 顶部导航 */
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
.nav-placeholder {
  width: 24px;
}

.exam-meta-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
/* 倒计时 */
.timer {
  display: flex;
  align-items: center;
  gap: 4px;
}
.timer-icon {
  width: 16px;
  height: 16px;
  fill: #3b82f6;
}
.time-text {
  font-size: 15px;
  color: #3b82f6;
  font-weight: 700;
}
.timer.urgent .timer-icon { fill: #ff4d4f; }
.timer.urgent .time-text { color: #ff4d4f; }

/* 课程信息区 */
.course-info-section {
  background-color: #fff;
  padding: 16px;
  margin-bottom: 12px;
}
.course-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 10px 0;
}
.progress-info {
  font-size: 14px;
  white-space: nowrap;
}
.answered-count {
  color: #3b82f6;
  font-weight: 700;
}
.total-count {
  color: #999;
}
.progress-bar {
  width: 100%;
  height: 6px;
  background-color: #e0e0e0;
  border-radius: 3px;
  overflow: hidden;
}
.progress-fill {
  height: 100%;
  background-color: #3b82f6;
  transition: width 0.3s;
}

/* 题目内容 */
.exam-content {
  padding: 12px 16px;
}
.question-card {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 12px;
  box-shadow: 0 1px 4px rgba(0,0,0,0.06);
}
.question-type-label {
  font-size: 14px;
  font-weight: 600;
  color: #3b82f6;
  display: block;
  margin-bottom: 12px;
}
.question-text {
  font-size: 15px;
  color: #333;
  line-height: 1.6;
  margin: 0;
}

/* 选项 */
.options-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.option-item {
  background-color: #fff;
  border-radius: 12px;
  padding: 14px 16px;
  cursor: pointer;
  transition: all 0.2s;
  border: 2px solid transparent;
}
.option-item.selected {
  border-color: #3b82f6;
  background-color: #eff6ff;
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
.radio-icon {
  width: 10px;
  height: 10px;
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

/* 多选提交 */
.submit-wrap {
  margin-top: 16px;
  display: flex;
  justify-content: center;
}
.submit-btn {
  padding: 10px 40px;
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  cursor: pointer;
}

/* 底部导航 */
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
.nav-btn.prev:disabled,
.nav-btn.next:disabled {
  opacity: 0.3;
}
.nav-btn.submit-final {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 2px solid #3b82f6;
  color: #3b82f6;
  font-size: 12px;
  font-weight: 600;
  justify-content: center;
}

/* 答题卡弹窗 */
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
  border-radius: 10px;
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
  width: 14px;
  height: 14px;
  border-radius: 50%;
}
.answered-dot { background-color: #3b82f6; }
.unanswered-dot { background-color: #f5f5f5; border: 1px solid #ccc; }
.status-text { font-size: 14px; color: #666; }
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

/* 加载 / 空状态 */
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
  padding: 10px 32px;
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  cursor: pointer;
}

/* 结果页 */
.result-page {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 80px;
}

/* 得分卡片 */
.result-score-card {
  background: #fff;
  border-radius: 12px;
  margin: 12px 16px 10px;
  padding: 32px 24px 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.score-big {
  font-size: 80px;
  font-weight: 700;
  color: #22c55e;
  line-height: 1;
}
.score-desc {
  font-size: 14px;
  color: #999;
  margin: 6px 0 24px;
}
.result-stats {
  display: flex;
  width: 100%;
  justify-content: space-around;
  padding-top: 20px;
}
.stat-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}
.stat-num {
  font-size: 22px;
  font-weight: 600;
  color: #333;
}
.stat-label {
  font-size: 13px;
  color: #999;
}

/* 错题回顾 */
.wrong-section {
  padding: 0 16px;
}
.wrong-section-title {
  font-size: 15px;
  font-weight: 600;
  color: #333;
  margin: 12px 0 10px;
}
.wrong-item-wrap {
  margin-bottom: 12px;
}
.wrong-item {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 8px;
}
.wrong-answer-card {
  border-top: 1px solid #f0f0f0;
  margin-top: 14px;
  padding-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.wrong-item-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
.wrong-type-tag {
  font-size: 13px;
  color: #3b82f6;
  font-weight: 600;
}
.wrong-badge {
  font-size: 13px;
  color: #ff4d4f;
  font-weight: 600;
}
.wrong-q-text {
  font-size: 15px;
  color: #333;
  line-height: 1.6;
  margin: 0 0 14px;
}
.wrong-options {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 14px;
}
.wrong-opt {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 14px;
  border-radius: 10px;
  background: #fff1f0;
  color: #ff4d4f;
  font-size: 14px;
}
.wrong-opt-dot {
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: #ff4d4f;
  flex-shrink: 0;
}
.wrong-opt-key {
  font-weight: 600;
}
.wrong-correct-ans {
  font-size: 14px;
  color: #3b82f6;
  font-weight: 500;
}
.wrong-analysis {
  font-size: 13px;
  color: #333;
  line-height: 1.6;
}
.all-correct-tip {
  text-align: center;
  padding: 40px 0;
  font-size: 16px;
  color: #22c55e;
  font-weight: 600;
}

/* 底部按钮 */
.result-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 12px 16px;
  background: #fff;
  border-top: 1px solid #eee;
}
.retry-btn {
  width: 100%;
  padding: 14px;
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
}

/* 自定义确认弹窗 */
.dialog-mask {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.45);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
}
.dialog-box {
  background: #fff;
  border-radius: 14px;
  padding: 24px 20px 16px;
  width: 280px;
  text-align: center;
}
.dialog-msg {
  font-size: 15px;
  color: #333;
  line-height: 1.6;
  margin: 0 0 20px;
}
.dialog-btns {
  display: flex;
  gap: 12px;
}
.dialog-cancel {
  flex: 1;
  padding: 10px;
  background: #f5f5f5;
  color: #666;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  cursor: pointer;
}
.dialog-confirm {
  flex: 1;
  padding: 10px;
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  cursor: pointer;
}
</style>
