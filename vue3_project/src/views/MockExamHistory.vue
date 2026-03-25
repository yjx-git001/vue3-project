<template>
  <div class="mock-history-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="goBack">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">模考记录</span>
      <span class="nav-placeholder"></span>
    </header>

    <!-- 课程名称 -->
    <div class="course-name-row">{{ courseName }}</div>

    <!-- 理论 / 实操 tab -->
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

    <div v-if="loading" class="loading-wrap">
      <p>加载中...</p>
    </div>

    <div v-else-if="list.length === 0" class="empty-wrap">
      <p>暂无模考记录</p>
    </div>

    <main v-else class="record-list">
      <div
        v-for="item in list"
        :key="item.id"
        class="record-card"
      >
        <!-- 第一行：题库名称 + 合格/不合格 -->
        <div class="card-row-top">
          <span class="lib-name">{{ courseName }}</span>
          <span class="pass-badge" :class="item.score >= 60 ? 'pass' : 'fail'">
            {{ item.score >= 60 ? '合格' : '不合格' }}
          </span>
        </div>

        <div class="card-info-block">
          <!-- 第二行：用时 得分 总分 -->
          <div class="card-row-mid">
            <span class="meta-item">用时：{{ formatMinutes(item.duration) }}</span>
            <span class="meta-item">得分：{{ item.score }}</span>
            <span class="meta-item">总分：100</span>
          </div>

          <!-- 第三行：时间 -->
          <div class="card-row-time">
            <span>时间：{{ formatDate(item.createdAt) }}</span>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
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

const courseName = ref('')
const activeTab = ref('theory')
const loading = ref(true)

const tabs = [
  { label: '理论', value: 'theory' },
  { label: '实操', value: 'practice' }
]

const list = ref<Array<{
  id: number
  score: number
  total: number
  correct: number
  duration: number
  createdAt: string
}>>([])

// 秒 → "X分" 或 "X分X秒"
const formatMinutes = (seconds: number) => {
  if (!seconds || seconds <= 0) return '0分'
  const m = Math.floor(seconds / 60)
  const s = seconds % 60
  if (m === 0) return `${s}秒`
  if (s === 0) return `${m}分`
  return `${m}分${s}秒`
}

// "2006-01-02 15:04:05" → "2006.01.02 15:04"
const formatDate = (dateStr: string) => {
  if (!dateStr) return ''
  // 替换日期部分的 - 为 .，去掉秒
  return dateStr.replace(/^(\d{4})-(\d{2})-(\d{2}) (\d{2}:\d{2}).*$/, '$1.$2.$3 $4')
}

onMounted(async () => {
  const courseEk = Number(route.params.id)
  try {
    const [nameRes, histRes]: any[] = await Promise.all([
      courseApi.getCourseDetail(courseEk),
      mockExamApi.getHistory(courseEk)
    ])
    courseName.value = nameRes.data?.courseName || '模考记录'
    list.value = histRes.data || []
  } catch {
    list.value = []
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.mock-history-page {
  background-color: #f2f2f7;
  min-height: 100vh;
}

/* 顶部导航 */
.top-nav {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
}

.back-btn {
  background: none;
  border: none;
  padding: 0;
  cursor: pointer;
  flex-shrink: 0;
  display: flex;
  align-items: center;
}

.back-btn svg {
  width: 24px;
  height: 24px;
  fill: #333;
}

.nav-title {
  font-size: 17px;
  font-weight: 600;
  color: #333;
  flex: 1;
  text-align: center;
}

.nav-placeholder {
  width: 24px;
  flex-shrink: 0;
}

/* 课程名称行 */
.course-name-row {
  text-align: center;
  font-size: 15px;
  color: #333;
  padding: 12px 16px 4px;
  background-color: #fff;
}

/* Tab */
.tabs {
  display: flex;
  background-color: #fff;
  border-bottom: 1px solid #e8e8e8;
}

.tab-item {
  flex: 1;
  text-align: center;
  padding: 12px 0;
  font-size: 15px;
  color: #999;
  cursor: pointer;
  position: relative;
}

.tab-item.active {
  color: #1677ff;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 28px;
  height: 2px;
  background-color: #1677ff;
  border-radius: 2px;
}

/* 空态/加载 */
.loading-wrap,
.empty-wrap {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 40vh;
  color: #999;
  font-size: 15px;
}

/* 记录列表 */
.record-list {
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

/* 卡片 */
.record-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
}

/* 第一行：题库名 + 合格状态 */
.card-row-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
}

.lib-name {
  font-size: 15px;
  font-weight: 700;
  color: #1a1a1a;
  flex: 1;
  margin-right: 8px;
}

.pass-badge {
  font-size: 14px;
  font-weight: 600;
  flex-shrink: 0;
}

.pass-badge.pass {
  color: #1677ff;
}

.pass-badge.fail {
  color: #f57c00;
}

.card-info-block {
  background-color: #f5f6fb;
  border-radius: 10px;
  padding: 10px 12px;
}

/* 第二行：用时 得分 总分 */
.card-row-mid {
  display: flex;
  gap: 22px;
  margin-bottom: 8px;
}

.meta-item {
  font-size: 13px;
  color: #555;
}

/* 第三行：时间 */
.card-row-time {
  font-size: 13px;
  color: #555;
}
</style>
