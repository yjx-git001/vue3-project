<template>
  <div class="course-content-page">
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程详情</span>
      <div class="nav-placeholder"></div>
    </header>

    <div class="course-banner">
      <img :src="courseDetail.image ? imageBaseUrl + courseDetail.image : '/images/course-detail-banner.png'" alt="课程横幅" class="banner-image" />
      <div class="play-overlay">
        <div class="play-button">
          <svg viewBox="0 0 24 24"><path d="M8,5.14V19.14L19,12.14L8,5.14Z" fill="white" /></svg>
        </div>
      </div>
    </div>

    <!-- 课程信息 -->
    <div class="course-info-section">
      <h2 class="course-title">{{ courseDetail.courseName || '课程名称' }}</h2>
      <span class="course-tag">{{ courseDetail.courseCategory === 2 ? '体系课程' : '单门课程' }}</span>
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
        <div v-if="lessons.length === 0" class="empty-tip">暂无视频目录</div>
        <div
          v-for="lesson in lessons"
          :key="lesson.id"
          class="lesson-item"
          :class="{ 'lesson-item--playing': lesson.isPlaying }"
          @click="playLesson(lesson.id)"
        >
          <div class="lesson-number" :class="{ 'lesson-number--playing': lesson.isPlaying }">{{ String(lesson.id).padStart(2, '0') }}</div>
          <div class="lesson-info">
            <div class="lesson-title">{{ lesson.title }}</div>
            <div class="lesson-meta">
              <span class="lesson-type" :class="{ 'lesson-type--playing': lesson.isPlaying }">{{ lesson.type }}</span>
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
          <span class="chapter-title">{{ currentLessonName }}</span>
          <button class="switch-course-btn" @click="showSwitchModal = true">
            切换课程
            <svg viewBox="0 0 24 24" class="dropdown-icon"><path d="M7,10L12,15L17,10H7Z" /></svg>
          </button>
        </div>
        <div class="exam-grid">
          <div class="exam-card" @click="openExamModal">
            <img src="/images/exam-practice-icon.png" class="exam-icon-img" />
            <span class="exam-label">考题训练</span>
          </div>
          <div class="exam-card" @click="goWrongQuestions">
            <img src="/images/error-record-icon.png" class="exam-icon-img" />
            <span class="exam-label">错题记录</span>
          </div>
          <div class="exam-card" @click="goMockHistory">
            <img src="/images/exam-history-icon.png" class="exam-icon-img" />
            <span class="exam-label">模考记录</span>
          </div>
        </div>
        <div class="exam-card-wide" @click="startMockExam">
          <img src="/images/mock-exam-icon.png" class="exam-icon-img" />
          <span class="exam-label">模拟考试</span>
        </div>
      </div>

    <!-- 切换课程弹窗 -->
    <div v-if="showSwitchModal" class="modal-overlay" @click="showSwitchModal = false">
      <div class="modal-content" @click.stop>
        <div class="modal-header-row">
          <span class="modal-title-text">切换课程</span>
          <button class="modal-close-btn" @click="showSwitchModal = false">✕</button>
        </div>
        <div class="question-types">
          <div
            v-for="lesson in lessons"
            :key="lesson.id"
            class="type-item"
            :class="{ 'type-item-active': selectedLessonId === lesson.id }"
            @click="selectLesson(lesson)"
          >
            <span>{{ lesson.title }}</span>
            <span class="count">{{ lesson.duration }}</span>
          </div>
        </div>
        <button class="cancel-btn" @click="showSwitchModal = false">取消</button>
      </div>
    </div>

    <!-- 考题训练弹窗 -->
    <div v-if="showExamModal" class="modal-overlay" @click="showExamModal = false">
      <div class="modal-content" @click.stop>
        <div class="question-types">
          <div
            v-for="tab in examTypeTabs"
            :key="tab.type"
            class="type-item"
            @click="startPractice(tab.type)"
          >
            <span>{{ tab.label }}</span>
            <span class="count">({{ tab.count }})</span>
          </div>
        </div>
        <button class="cancel-btn" @click="showExamModal = false">取消</button>
      </div>
    </div>

      <!-- 笔记 -->
      <div v-if="activeTab === 'notes'" class="notes-content">
        <div v-if="!notesContent" class="empty-tip">暂无笔记内容</div>
        <div v-else class="notes-text" style="white-space: pre-wrap;">{{ notesContent }}</div>
      </div>

      <!-- 附件 -->
      <div v-if="activeTab === 'attachments'">
        <div class="attachments-content">
          <div v-if="attachments.length === 0" class="empty-tip">暂无附件</div>
          <div
            v-for="file in attachments"
            :key="file.id"
            class="attachment-item"
          >
            <div class="file-icon pdf">
              <svg viewBox="0 0 48 52" class="pdf-file-svg" aria-hidden="true">
                <path d="M9 3h24l9 9v31a5 5 0 0 1-5 5H9a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5z" fill="#ef5350" />
                <path d="M33 3v9h9z" fill="#d94744" />
                <path d="M16.8 36.2c3.4-1.1 6.8-3 9.9-6.3c2.7.4 5.1 1.4 7.2 2.6" fill="none" stroke="#fff" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
                <path d="M27.3 20.6c1 3.3 2.6 6.9 4.3 9.9c-2.5-.4-5-.5-7.6-.3c1.3-3 2.4-6.2 3.3-9.6z" fill="none" stroke="#fff" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
                <circle cx="16.8" cy="36.2" r="1.2" fill="#fff" />
                <circle cx="33.9" cy="32.6" r="1.2" fill="#fff" />
              </svg>
            </div>
            <div class="file-info">
              <div class="file-name">{{ file.name }}</div>
              <div class="file-size">{{ file.size || '219.8KB' }}</div>
            </div>
            <button class="download-file-btn">下载</button>
          </div>
        </div>
        <div class="file-type-icons">
          <div class="type-icon">
            <svg viewBox="0 0 40 46" aria-hidden="true">
              <path d="M6 2h20l8 8v30a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4z" fill="#ef765e"/>
              <path d="M26 2v8h8z" fill="#d45e48"/>
              <text x="20" y="25" text-anchor="middle" dominant-baseline="middle" fill="#fff" font-size="16" font-weight="700">p</text>
            </svg>
          </div>
          <div class="type-icon">
            <svg viewBox="0 0 40 46" aria-hidden="true">
              <path d="M6 2h20l8 8v30a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4z" fill="#4f97db"/>
              <path d="M26 2v8h8z" fill="#3f7eb9"/>
              <text x="20" y="25" text-anchor="middle" dominant-baseline="middle" fill="#fff" font-size="16" font-weight="700">w</text>
            </svg>
          </div>
          <div class="type-icon">
            <svg viewBox="0 0 40 46" aria-hidden="true">
              <path d="M6 2h20l8 8v30a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4z" fill="#58c37a"/>
              <path d="M26 2v8h8z" fill="#45a863"/>
              <text x="20" y="25" text-anchor="middle" dominant-baseline="middle" fill="#fff" font-size="16" font-weight="700">x</text>
            </svg>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { courseApi, questionApi, courseContentApi, mockExamApi } from '@/api'

const router = useRouter()
const route = useRoute()
const activeTab = ref('catalog')
const showExamModal = ref(false)
const showSwitchModal = ref(false)
const courseDetail = ref<any>({})
const imageBaseUrl = 'http://localhost:8080'

// 笔记内容（所有子课程拼接）
const notesContent = ref('')
const selectedLessonId = ref<number | null>(null)
const currentLessonName = computed(() => {
  const l = lessons.value.find(l => l.id === selectedLessonId.value)
  return l ? l.title : (lessons.value[0]?.title || '请选择课程')
})

// 当前课程的题目统计（按题型）
const examQuestions = ref<any[]>([])
const examTypeTabs = computed(() => {
  const single = examQuestions.value.filter(q => q.questionType === 1).length
  const multiple = examQuestions.value.filter(q => q.questionType === 2).length
  const judge = examQuestions.value.filter(q => q.questionType === 3).length
  const tabs = []
  if (single > 0) tabs.push({ label: '单选题', type: 1, count: single })
  if (multiple > 0) tabs.push({ label: '多选题', type: 2, count: multiple })
  if (judge > 0) tabs.push({ label: '判断题', type: 3, count: judge })
  return tabs
})

const fetchExamQuestions = async (courseEk: number) => {
  try {
    const res: any = await questionApi.getList(courseEk)
    examQuestions.value = res.data || []
  } catch {
    examQuestions.value = []
  }
}

const selectLesson = (lesson: any) => {
  selectedLessonId.value = lesson.id
  showSwitchModal.value = false
  // 拉取该课程对应的题目
  fetchExamQuestions(lesson.courseEk || Number(route.params.id))
}

const openExamModal = () => {
  if (examTypeTabs.value.length === 0) {
    alert('该课程暂无配置题目')
    return
  }
  showExamModal.value = true
}

const playLesson = (lessonId: number) => {
  lessons.value = lessons.value.map(lesson => ({
    ...lesson,
    isPlaying: lesson.id === lessonId,
    type: lesson.id === lessonId ? '正在播放' : '视频'
  }))
  selectedLessonId.value = lessonId
}

onMounted(async () => {
  const ek = Number(route.params.id)
  if (ek) {
    try {
      const res: any = await courseApi.getCourseDetail(ek)
      courseDetail.value = res.data
    } catch (e) {
      console.error('获取课程详情失败:', e)
    }
    const isSystem = courseDetail.value.courseCategory === 2
    const subCourses: any[] = isSystem ? (courseDetail.value.courseContent || []) : []

    // 加载视频目录
    try {
      if (isSystem) {
        // 体系课程：聚合所有子课程的视频，带上各自的 courseEk
        const allVideos = await Promise.all(
          subCourses.map(async (sub: any) => {
            try {
              const vRes: any = await courseContentApi.getVideos(sub.ek)
              const vList: { sort: number; title: string; duration: string }[] = vRes.data || []
              return vList.map(v => ({
                subEk: sub.ek,
                subName: sub.courseName,
                sort: v.sort,
                title: v.title,
                duration: v.duration || ''
              }))
            } catch {
              return []
            }
          })
        )
        let globalIdx = 0
        lessons.value = allVideos.flat().map((v, i) => ({
          id: ++globalIdx,
          title: v.title,
          type: i === 0 ? '正在播放' : '视频',
          duration: v.duration,
          isPlaying: i === 0,
          courseEk: v.subEk
        }))
      } else {
        // 单门课程
        const vRes: any = await courseContentApi.getVideos(ek)
        const vList: { sort: number; title: string; duration: string }[] = vRes.data || []
        lessons.value = vList.map((v, i) => ({
          id: v.sort,
          title: v.title,
          type: i === 0 ? '正在播放' : '视频',
          duration: v.duration || '',
          isPlaying: i === 0,
          courseEk: ek
        }))
      }
    } catch {
      // 保留默认占位数据
    }

    // 加载附件列表
    try {
      if (isSystem) {
        // 体系课程：聚合所有子课程附件
        const allAttachments = await Promise.all(
          subCourses.map(async (sub: any) => {
            try {
              const aRes: any = await courseContentApi.getAttachments(sub.ek)
              return (aRes.data || []) as { sort: number; name: string }[]
            } catch {
              return []
            }
          })
        )
        let aIdx = 0
        attachments.value = allAttachments.flat().map(a => ({
          id: ++aIdx,
          name: a.name,
          size: ''
        }))
      } else {
        // 单门课程
        const aRes: any = await courseContentApi.getAttachments(ek)
        const aList: { sort: number; name: string }[] = aRes.data || []
        attachments.value = aList.map(a => ({
          id: a.sort,
          name: a.name,
          size: ''
        }))
      }
    } catch {
      // 保留默认占位数据
    }

    // 加载笔记
    try {
      if (isSystem) {
        // 体系课程：聚合所有子课程笔记拼接
        const notesResults = await Promise.all(
          subCourses.map(async (sub: any) => {
            try {
              const nRes: any = await courseContentApi.getNotes(sub.ek)
              return nRes.data?.content || ''
            } catch {
              return ''
            }
          })
        )
        notesContent.value = notesResults.filter(c => c).join('\n\n')
      } else {
        // 单门课程
        const nRes: any = await courseContentApi.getNotes(ek)
        notesContent.value = nRes.data?.content || ''
      }
    } catch {
      notesContent.value = ''
    }

    // 默认加载题目：体系课程加载第一个子课程的题目，单门课程直接加载
    const defaultEk = isSystem && subCourses.length > 0 ? subCourses[0].ek : ek
    await fetchExamQuestions(defaultEk)
    if (lessons.value.length > 0) {
      const firstLesson = lessons.value[0]
      if (firstLesson) selectedLessonId.value = firstLesson.id
    }
  }
})

const tabs = [
  { label: '目录', value: 'catalog' },
  { label: '考题', value: 'exam' },
  { label: '笔记', value: 'notes' },
  { label: '附件', value: 'attachments' }
]

const startPractice = (questionType: number) => {
  showExamModal.value = false
  // 传当前选中的子课程 ek（如果有），否则用路由 ek（单门课程直接访问时）
  const selectedLesson = lessons.value.find(l => l.id === selectedLessonId.value)
  const practiceEk = selectedLesson?.courseEk || Number(route.params.id)
  router.push(`/exam_practice/${practiceEk}?type=${questionType}`)
}

const goWrongQuestions = () => {
  const selectedLesson = lessons.value.find(l => l.id === selectedLessonId.value)
  const ek = selectedLesson?.courseEk || Number(route.params.id)
  router.push(`/wrong_question_detail/${ek}`)
}

const goMockHistory = () => {
  const selectedLesson = lessons.value.find(l => l.id === selectedLessonId.value)
  const ek = selectedLesson?.courseEk || Number(route.params.id)
  router.push(`/mock_exam_history/${ek}`)
}

const startMockExam = async () => {
  // 获取当前选中的子课程 ek
  const selectedLesson = lessons.value.find(l => l.id === selectedLessonId.value)
  const mockEk = selectedLesson?.courseEk || Number(route.params.id)
  // 检查是否有模拟考试配置
  try {
    const res: any = await mockExamApi.getConfig(mockEk)
    const cfg = res.data
    if (!cfg || (cfg.singleCount === 0 && cfg.multipleCount === 0 && cfg.judgeCount === 0)) {
      alert('该课程暂未配置模拟考试题目数量，请联系管理员')
      return
    }
  } catch {
    alert('获取模拟考试配置失败，请重试')
    return
  }
  router.push(`/mock_exam_detail/${mockEk}`)
}

const lessons = ref<{
  id: number
  title: string
  type: string
  duration: string
  isPlaying: boolean
  courseEk: number | null
}[]>([])

const attachments = ref<{
  id: number
  name: string
  size: string
}[]>([])
</script>

<style scoped>
.course-content-page {
  background-color: #f2f4f8;
  min-height: 100vh;
  max-width: 560px;
  margin: 0 auto;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 16px 12px;
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

.nav-placeholder {
  width: 24px;
  height: 24px;
}

.course-banner {
  position: relative;
  margin: 0 16px;
  border-radius: 16px;
  aspect-ratio: 16 / 9;
  overflow: hidden;
  background: #d9e9ff;
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
  background: rgba(0, 0, 0, 0.14);
}

.play-button {
  width: 76px;
  height: 76px;
  background: rgba(255, 255, 255, 0.72);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.play-button svg {
  width: 34px;
  height: 34px;
  margin-left: 4px;
}

.course-info-section {
  background-color: #fff;
  padding: 14px 16px 10px;
}

.course-title {
  font-size: 18px;
  font-weight: 600;
  color: #343b45;
  margin: 0 0 10px 0;
}

.course-tag {
  display: inline-block;
  background-color: #e7f0ff;
  color: #4d84e9;
  padding: 6px 12px;
  border-radius: 8px;
  font-size: 13px;
  line-height: 1;
}

.tabs {
  display: flex;
  background-color: #fff;
  padding: 0 18px;
  gap: 30px;
  border-bottom: 1px solid #eceff5;
}

.tab-item {
  padding: 14px 0;
  font-size: 16px;
  color: #6f7681;
  font-weight: 500;
  cursor: pointer;
  position: relative;
}

.tab-item.active {
  color: #333;
  font-weight: 700;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 22px;
  height: 4px;
  background-color: #4668ff;
  border-radius: 99px;
}

.content-area {
  background-color: #f2f4f8;
  min-height: 400px;
  padding: 12px 12px 20px;
}

.chapter-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #fff;
  border-radius: 12px;
  margin-bottom: 12px;
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
  padding-bottom: 6px;
}

.lesson-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 14px;
  background: #fff;
  border-radius: 14px;
  margin-bottom: 10px;
  border: 1px solid transparent;
}

.lesson-item--playing {
  background: #dce6fb;
  border-color: #b8c9f3;
}

.lesson-number {
  width: 50px;
  height: 50px;
  background-color: #bfd0f8;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 700;
  color: #2f76e5;
  flex-shrink: 0;
}

.lesson-number--playing {
  background-color: #b2c7f7;
  color: #2f76e5;
}

.lesson-info {
  flex: 1;
}

.lesson-title {
  font-size: 15px;
  color: #40454f;
  margin-bottom: 4px;
  font-weight: 600;
}

.lesson-meta {
  display: flex;
  gap: 8px;
  color: #5d6470;
  align-items: center;
}

.lesson-type {
  font-size: 13px;
  color: #5d6470;
}

.lesson-type--playing {
  color: #4f5f7c;
  font-weight: 600;
}

.lesson-duration {
  font-size: 13px;
  color: #5d6470;
}

.play-btn {
  width: 30px;
  height: 30px;
  background: #b8c1cf;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  padding: 0;
}

.play-btn svg {
  width: 16px;
  height: 16px;
  fill: #fff;
}

.play-btn.playing {
  background: #3d82ed;
}

.play-btn.playing svg {
  fill: #fff;
}

.exam-content {
  padding-bottom: 16px;
}

.exam-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 12px;
  margin-bottom: 12px;
}

.exam-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  padding: 24px 16px;
  background-color: #fff;
  border-radius: 16px;
  cursor: pointer;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08);
}

.exam-card-wide {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 12px;
  padding: 20px 24px;
  background-color: #fff;
  border-radius: 16px;
  cursor: pointer;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08);
  margin: 0 auto;
  width: 90%;
}

.exam-icon-img {
  width: 48px;
  height: 48px;
}

.exam-label {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.notes-content {
  background-color: #fff;
  padding: 16px;
}
.notes-block {
  margin-bottom: 20px;
}
.notes-course-title {
  font-size: 14px;
  font-weight: 600;
  color: #3b82f6;
  margin-bottom: 8px;
  padding-bottom: 6px;
  border-bottom: 1px solid #e8f0fe;
}
.notes-text {
  font-size: 14px;
  color: #333;
  line-height: 1.8;
}

.attachments-content {
  padding: 10px 12px;
}

.attachment-item {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 16px;
  background-color: #fff;
  border-radius: 14px;
  margin-bottom: 10px;
  border: 1px solid #edf0f6;
}

.file-icon {
  width: 40px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.file-icon.pdf {
  background: none;
}

.pdf-file-svg {
  width: 100%;
  height: 100%;
  display: block;
}

.file-info {
  flex: 1;
}

.file-name {
  font-size: 17px;
  color: #40454f;
  margin-bottom: 2px;
  font-weight: 500;
}

.file-size {
  font-size: 14px;
  color: #9aa0ad;
}

.download-file-btn {
  padding: 0;
  background: none;
  border: none;
  color: #3f87ff;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
}

.file-type-icons {
  display: flex;
  gap: 34px;
  padding: 24px 0 0 44px;
  justify-content: flex-start;
  margin-top: 132px;
}

.type-icon {
  width: 40px;
  height: 46px;
}

.type-icon svg {
  width: 100%;
  height: 100%;
  display: block;
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
  justify-content: center;
  align-items: center;
  gap: 4px;
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

.empty-tip {
  text-align: center;
  color: #bbb;
  font-size: 14px;
  padding: 32px 0;
}

.modal-header-row {  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 16px 10px;
  border-bottom: 1px solid #f0f0f0;
}
.modal-title-text {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}
.modal-close-btn {
  background: none;
  border: none;
  font-size: 16px;
  color: #999;
  cursor: pointer;
  padding: 0 4px;
  width: auto;
}
.type-item-active {
  background: #e3f2fd;
  color: #1976d2;
  font-weight: 600;
}
</style>
