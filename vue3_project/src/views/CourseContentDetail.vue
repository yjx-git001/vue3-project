<template>
  <div class="course-content-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程详情</span>
    </header>

    <!-- 课程横幅 -->
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
          <div class="exam-card">
            <img src="/images/error-record-icon.png" class="exam-icon-img" />
            <span class="exam-label">错题记录</span>
          </div>
          <div class="exam-card">
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
              <svg viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" fill="white" /></svg>
            </div>
            <div class="file-info">
              <div class="file-name">{{ file.name }}</div>
              <div class="file-size">{{ file.size }}</div>
            </div>
            <button class="download-file-btn">下载</button>
          </div>
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
      selectedLessonId.value = lessons.value[0].id
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
  router.push(`/exam_practice/${mockEk}?mode=mock`)
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
  text-align: center;
  padding: 14px 16px;
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
  background-color: #f5f5f5;
  min-height: 400px;
  padding: 12px;
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
  padding-bottom: 16px;
}

.lesson-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background: #fff;
  border-radius: 12px;
  margin-bottom: 10px;
}

.lesson-number {
  width: 40px;
  height: 40px;
  background-color: #e3f2fd;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 600;
  color: #2196f3;
  flex-shrink: 0;
}

.lesson-item:first-child .lesson-number {
  background-color: #3b82f6;
  color: #fff;
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
  width: 32px;
  height: 32px;
  background: #e0e0e0;
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
  width: 18px;
  height: 18px;
  fill: #999;
}

.play-btn.playing {
  background: #3b82f6;
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
  padding: 12px;
}

.attachment-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background-color: #fff;
  border-radius: 12px;
  margin-bottom: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08);
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
  padding: 24px 16px;
  justify-content: flex-start;
  margin-top: 100px;
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
