<template>
  <div class="course-admin">
    <h1>课程管理</h1>

    <!-- Banner 管理 -->
    <section class="form-section">
      <h2>首页 Banner 管理</h2>
      <form @submit.prevent="handleCreateBanner">
        <div class="form-group">
          <label>Banner 图片：</label>
          <input type="file" accept="image/*" @change="handleBannerImageUpload" required />
          <span v-if="bannerForm.image" class="upload-success">✓ 已上传</span>
        </div>
        <div class="form-group">
          <label>排序（数字越小越靠前）：</label>
          <input v-model.number="bannerForm.sort" type="number" />
        </div>
        <button type="submit" :disabled="bannerLoading">
          {{ bannerLoading ? '创建中...' : '添加 Banner' }}
        </button>
      </form>
      <div v-if="bannerMessage" class="message" :class="bannerMessage.type">
        {{ bannerMessage.text }}
      </div>

      <!-- Banner 列表 -->
      <div v-if="bannerList.length > 0" class="banner-list">
        <h3>当前 Banner 列表</h3>
        <div v-for="item in bannerList" :key="item.id" class="banner-item">
          <img :src="'http://localhost:8080' + item.image" class="banner-preview" />
          <div class="banner-info">
            <span>排序：{{ item.sort }}</span>
            <span v-if="item.link">链接：{{ item.link }}</span>
          </div>
          <button class="delete-btn" type="button" @click="handleDeleteBanner(item.id)">删除</button>
        </div>
      </div>
    </section>

    <!-- 限时秒杀管理 -->
    <section class="form-section">
      <h2>限时秒杀管理</h2>
      <form @submit.prevent="handleCreateFlashSale">
        <div class="form-group">
          <label>课程类型：</label>
          <select v-model.number="flashSaleForm.courseCategory" @change="flashSaleForm.courseEk = 0">
            <option :value="1">单门课程</option>
            <option :value="2">体系课程</option>
          </select>
        </div>
        <div class="form-group">
          <label>选择课程：</label>
          <select v-model.number="flashSaleForm.courseEk" required>
            <option :value="0" disabled>请选择课程</option>
            <option v-for="option in filteredCourseOptions" :key="option.ek" :value="option.ek">
              {{ option.courseName }}
            </option>
          </select>
        </div>
        <div class="form-group">
          <label>秒杀价格(元)：</label>
          <input v-model.number="flashSaleForm.priceYuan" type="number" step="0.01" min="0" required />
        </div>
        <div class="form-group">
          <label>开始时间：</label>
          <div class="datetime-row">
            <input v-model="flashSaleForm.startDate" type="date" required class="date-input" />
            <select v-model="flashSaleForm.startHour" class="time-select">
              <option v-for="h in hours" :key="h" :value="h">{{ h }}</option>
            </select>
            <span>:</span>
            <select v-model="flashSaleForm.startMinute" class="time-select">
              <option v-for="m in minutes" :key="m" :value="m">{{ m }}</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label>结束时间：</label>
          <div class="datetime-row">
            <input v-model="flashSaleForm.endDate" type="date" required class="date-input" />
            <select v-model="flashSaleForm.endHour" class="time-select">
              <option v-for="h in hours" :key="h" :value="h">{{ h }}</option>
            </select>
            <span>:</span>
            <select v-model="flashSaleForm.endMinute" class="time-select">
              <option v-for="m in minutes" :key="m" :value="m">{{ m }}</option>
            </select>
          </div>
        </div>
        <button type="submit" :disabled="flashSaleLoading">
          {{ flashSaleLoading ? '创建中...' : '上架秒杀' }}
        </button>
      </form>
      <div v-if="flashSaleMessage" class="message" :class="flashSaleMessage.type">
        {{ flashSaleMessage.text }}
      </div>

      <!-- 秒杀列表 -->
      <div v-if="flashSaleList.length > 0" class="flash-sale-list">
        <h3>秒杀列表</h3>
        <div v-for="item in flashSaleList" :key="item.id" class="flash-sale-item">
          <div class="flash-sale-info">
            <span class="fs-name">{{ item.course?.courseName || '-' }}</span>
            <span>秒杀价：¥{{ (item.price / 100).toFixed(2) }}</span>
            <span>{{ formatTime(item.startTime) }} ~ {{ formatTime(item.endTime) }}</span>
            <span :class="['fs-status', item.enable ? 'on' : 'off']">{{ item.enable ? '已上架' : '已下架' }}</span>
          </div>
          <div class="flash-sale-actions">
            <button type="button" class="enable-btn" @click="handleToggleEnable(item)">
              {{ item.enable ? '下架' : '上架' }}
            </button>
            <button type="button" class="delete-btn" @click="handleDeleteFlashSale(item.id)">删除</button>
          </div>
        </div>
      </div>
    </section>

    <!-- 创建体系课程 -->
    <section class="form-section">
      <h2>创建体系课程</h2>
      <form @submit.prevent="handleCreateSystem">
        <div class="form-group">
          <label>课程名称：</label>
          <input v-model="systemForm.courseName" type="text" required />
        </div>
        <div class="form-group">
          <label>课时：</label>
          <input v-model.number="systemForm.courseTime" type="number" required />
        </div>
        <div class="form-group">
          <label>价格(分)：</label>
          <input v-model.number="systemForm.price" type="number" required />
        </div>
        <div class="form-group">
          <label>原价(分)：</label>
          <input v-model.number="systemForm.originalPrice" type="number" required />
        </div>
        <div class="form-group">
          <label>课程图片：</label>
          <input type="file" accept="image/*" @change="handleSystemImageUpload" required />
          <span v-if="systemForm.image" class="upload-success">✓ 已上传</span>
        </div>
        <div class="form-group">
          <label>课程详情：</label>
          <textarea v-model="systemForm.courseDetail" required></textarea>
        </div>
        <div class="form-group">
          <label>课程介绍图片：</label>
          <input type="file" accept="image/*" @change="handleSystemIntroUpload" required />
          <span v-if="systemForm.courseIntroduction" class="upload-success">✓ 已上传</span>
        </div>
        <div class="form-group">
          <label>课程标签分类：</label>
          <div class="checkbox-group">
            <label v-for="tag in tagOptions" :key="tag.value" class="checkbox-label">
              <input type="checkbox" :value="tag.value" v-model="systemForm.courseTagClass" />
              {{ tag.text }}
            </label>
          </div>
        </div>
        <button type="submit" :disabled="systemLoading">
          {{ systemLoading ? '创建中...' : '创建体系课程' }}
        </button>
      </form>
      <div v-if="systemMessage" class="message" :class="systemMessage.type">
        {{ systemMessage.text }}
      </div>
    </section>

    <!-- 创建单门课程 -->
    <section class="form-section">
      <h2>创建单门课程</h2>
      <form @submit.prevent="handleCreateSingle">
        <div class="form-group">
          <label>课程名称：</label>
          <input v-model="singleForm.courseName" type="text" required />
        </div>
        <div class="form-group">
          <label>所属体系课程：</label>
          <select v-model.number="singleForm.parentEk" required>
            <option value="" disabled>请选择体系课程</option>
            <option v-for="option in systemOptions" :key="option.ek" :value="option.ek">
              {{ option.courseName }}
            </option>
          </select>
        </div>
        <div class="form-group">
          <label>课时：</label>
          <input v-model.number="singleForm.courseTime" type="number" required />
        </div>
        <div class="form-group">
          <label>价格(分)：</label>
          <input v-model.number="singleForm.price" type="number" required />
        </div>
        <div class="form-group">
          <label>原价(分)：</label>
          <input v-model.number="singleForm.originalPrice" type="number" required />
        </div>
        <div class="form-group">
          <label>课程图片：</label>
          <input type="file" accept="image/*" @change="handleSingleImageUpload" required />
          <span v-if="singleForm.image" class="upload-success">✓ 已上传</span>
        </div>
        <div class="form-group">
          <label>课程详情：</label>
          <textarea v-model="singleForm.courseDetail" required></textarea>
        </div>
        <div class="form-group">
          <label>课程标签分类：</label>
          <div class="checkbox-group">
            <label v-for="tag in tagOptions" :key="tag.value" class="checkbox-label">
              <input type="checkbox" :value="tag.value" v-model="singleForm.courseTagClass" />
              {{ tag.text }}
            </label>
          </div>
        </div>
        <div class="form-group">
          <label>单选题数量：</label>
          <input v-model.number="singleForm.theoreticalQuestions.singleQuestion" type="number" required />
        </div>
        <div class="form-group">
          <label>多选题数量：</label>
          <input v-model.number="singleForm.theoreticalQuestions.multipleQuestion" type="number" required />
        </div>
        <div class="form-group">
          <label>判断题数量：</label>
          <input v-model.number="singleForm.theoreticalQuestions.judgeQuestion" type="number" required />
        </div>
        <div class="form-group">
          <label>视频题数量：</label>
          <input v-model.number="singleForm.videoQuestions" type="number" required />
        </div>
        <div class="form-group">
          <label>附件数量：</label>
          <input v-model.number="singleForm.attachment" type="number" required />
        </div>
        <button type="submit" :disabled="singleLoading">
          {{ singleLoading ? '创建中...' : '创建单门课程' }}
        </button>
      </form>
      <div v-if="singleMessage" class="message" :class="singleMessage.type">
        {{ singleMessage.text }}
      </div>
    </section>

    <!-- 卡密管理 -->
    <section class="admin-section">
      <h2>卡密管理</h2>
      <div class="form-group">
        <label>生成数量：</label>
        <input v-model.number="cardKeyCount" type="number" min="1" max="100" placeholder="1-100" />
        <button @click="handleGenerateCardKeys" :disabled="cardKeyLoading">
          {{ cardKeyLoading ? '生成中...' : '生成卡密' }}
        </button>
      </div>
      <div v-if="cardKeyMessage" class="message" :class="cardKeyMessage.type">{{ cardKeyMessage.text }}</div>

      <table v-if="cardKeyList.length" class="card-key-table">
        <thead>
          <tr>
            <th>卡密</th>
            <th>状态</th>
            <th>使用者ID</th>
            <th>使用者姓名</th>
            <th>生成时间</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="k in cardKeyList" :key="k.id">
            <td class="code-cell">
              {{ k.code }}
              <button class="copy-btn" @click="copyCode(k.code)">{{ copiedCode === k.code ? '已复制!' : '复制' }}</button>
            </td>
            <td :class="k.used ? 'used' : 'unused'">{{ k.used ? '已使用' : '未使用' }}</td>
            <td>{{ k.userId || '-' }}</td>
            <td>{{ k.userName || '-' }}</td>
            <td>{{ k.createdAt }}</td>
          </tr>
        </tbody>
      </table>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { courseAdminApi, uploadApi, bannerAdminApi, flashSaleAdminApi, courseApi, cardKeyApi } from '@/api'

// 限时秒杀
const hours = Array.from({ length: 24 }, (_, i) => String(i).padStart(2, '0'))
const minutes = Array.from({ length: 60 }, (_, i) => String(i).padStart(2, '0'))
const allCourseOptions = ref<Array<{ ek: number; courseName: string; courseCategory: number }>>([])
const flashSaleList = ref<any[]>([])
const flashSaleForm = ref({
  courseEk: 0,
  priceYuan: 0,
  startDate: '',
  startHour: '00',
  startMinute: '00',
  endDate: '',
  endHour: '23',
  endMinute: '59',
  courseCategory: 1
})
const flashSaleLoading = ref(false)
const flashSaleMessage = ref<{ type: string; text: string } | null>(null)

const filteredCourseOptions = computed(() =>
  allCourseOptions.value.filter(c => c.courseCategory === flashSaleForm.value.courseCategory)
)

const formatTime = (t: string) => t ? t.replace('T', ' ').slice(0, 16) : '-'

const fetchAllCourseOptions = async () => {
  try {
    const res: any = await courseApi.getCourseList({ pageIndex: 1, pageSize: 100 })
    allCourseOptions.value = (res.data?.list || []).map((c: any) => ({
      ek: c.ek,
      courseName: c.title,
      courseCategory: c.courseCategory
    }))
  } catch {}
}

const fetchFlashSales = async () => {
  try {
    const res: any = await flashSaleAdminApi.getList()
    flashSaleList.value = res.data || []
  } catch {}
}

const handleCreateFlashSale = async () => {
  flashSaleLoading.value = true
  flashSaleMessage.value = null
  try {
    // 先下架已上架的
    const active = flashSaleList.value.find(i => i.enable)
    if (active) {
      await flashSaleAdminApi.setEnable(active.id, false)
    }
    const f = flashSaleForm.value
    const data = {
      courseEk: f.courseEk,
      price: Math.round(f.priceYuan * 100),
      startTime: `${f.startDate} ${f.startHour}:${f.startMinute}:00`,
      endTime: `${f.endDate} ${f.endHour}:${f.endMinute}:00`
    }
    await flashSaleAdminApi.create(data)
    // 上架新创建的
    await fetchFlashSales()
    const newest = flashSaleList.value[0]
    if (newest) await flashSaleAdminApi.setEnable(newest.id, true)
    flashSaleMessage.value = { type: 'success', text: '上架成功！' }
    flashSaleForm.value = { courseEk: 0, priceYuan: 0, startDate: '', startHour: '00', startMinute: '00', endDate: '', endHour: '23', endMinute: '59', courseCategory: 1 }
    await fetchFlashSales()
  } catch (error: any) {
    flashSaleMessage.value = { type: 'error', text: error.message || '操作失败' }
  } finally {
    flashSaleLoading.value = false
  }
}

const handleToggleEnable = async (item: any) => {
  try {
    // 上架时先下架其他
    if (!item.enable) {
      const active = flashSaleList.value.find(i => i.enable)
      if (active) await flashSaleAdminApi.setEnable(active.id, false)
    }
    await flashSaleAdminApi.setEnable(item.id, !item.enable)
    await fetchFlashSales()
  } catch {}
}

const handleDeleteFlashSale = async (id: number) => {
  try {
    await flashSaleAdminApi.delete(id)
    await fetchFlashSales()
  } catch {}
}

// 体系课程表单
const systemForm = ref({
  courseName: '',
  courseTime: 0,
  price: 0,
  originalPrice: 0,
  image: '',
  courseDetail: '',
  courseIntroduction: '',
  courseTagClass: [] as number[]
})

// 单门课程表单
const singleForm = ref({
  courseName: '',
  parentEk: 0,
  courseTime: 0,
  price: 0,
  originalPrice: 0,
  image: '',
  courseDetail: '',
  courseTagClass: [] as number[],
  theoreticalQuestions: {
    singleQuestion: 0,
    multipleQuestion: 0,
    judgeQuestion: 0
  },
  videoQuestions: 0,
  attachment: 0
})

// 体系课程选项
const systemOptions = ref<Array<{ ek: number; courseName: string }>>([])

// 标签分类选项
const tagOptions = ref<Array<{ text: string; value: number }>>([])

// 加载状态
const systemLoading = ref(false)
const singleLoading = ref(false)

// 消息提示
const systemMessage = ref<{ type: string; text: string } | null>(null)
const singleMessage = ref<{ type: string; text: string } | null>(null)

// 图片上传处理
const handleSystemImageUpload = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  try {
    const res: any = await uploadApi.uploadFile(file)
    systemForm.value.image = res.data.url
  } catch (error) {
    console.error('上传失败:', error)
  }
}

const handleSystemIntroUpload = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  try {
    const res: any = await uploadApi.uploadFile(file)
    systemForm.value.courseIntroduction = res.data.url
  } catch (error) {
    console.error('上传失败:', error)
  }
}

const handleSingleImageUpload = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  try {
    const res: any = await uploadApi.uploadFile(file)
    singleForm.value.image = res.data.url
  } catch (error) {
    console.error('上传失败:', error)
  }
}

// 获取体系课程选项
const fetchSystemOptions = async () => {
  try {
    const res: any = await courseAdminApi.getSystemOptions()
    systemOptions.value = res.data || []
  } catch (error) {
    console.error('获取体系课程选项失败:', error)
  }
}

// 获取标签分类选项
const fetchTagOptions = async () => {
  try {
    const res: any = await courseAdminApi.getQueryRule()
    tagOptions.value = res.data?.edit?.courseTagClass?.values || []
  } catch (error) {
    console.error('获取标签选项失败:', error)
  }
}

// Banner 管理
const bannerList = ref<any[]>([])
const bannerForm = ref({ image: '', sort: 0 })
const bannerLoading = ref(false)
const bannerMessage = ref<{ type: string; text: string } | null>(null)

const fetchBanners = async () => {
  try {
    const res: any = await bannerAdminApi.getList()
    bannerList.value = res.data || []
  } catch (error) {
    console.error('获取banner失败:', error)
  }
}

const handleBannerImageUpload = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  try {
    const res: any = await uploadApi.uploadFile(file)
    bannerForm.value.image = res.data.url
  } catch (error) {
    console.error('上传失败:', error)
  }
}

const handleCreateBanner = async () => {
  bannerLoading.value = true
  bannerMessage.value = null
  try {
    await bannerAdminApi.create(bannerForm.value)
    bannerMessage.value = { type: 'success', text: 'Banner 创建成功！' }
    bannerForm.value = { image: '', sort: 0 }
    await fetchBanners()
  } catch (error: any) {
    bannerMessage.value = { type: 'error', text: error.message || '创建失败' }
  } finally {
    bannerLoading.value = false
  }
}

const handleDeleteBanner = async (id: number) => {
  try {
    await bannerAdminApi.delete(id)
    await fetchBanners()
  } catch (error) {
    console.error('删除失败:', error)
  }
}

// 创建体系课程
const handleCreateSystem = async () => {
  systemLoading.value = true
  systemMessage.value = null
  try {
    await courseAdminApi.createSystemCourse(systemForm.value)
    systemMessage.value = { type: 'success', text: '体系课程创建成功！' }
    // 重置表单
    systemForm.value = {
      courseName: '',
      courseTime: 0,
      price: 0,
      originalPrice: 0,
      image: '',
      courseDetail: '',
      courseIntroduction: '',
      courseTagClass: []
    }
    // 刷新体系课程选项
    await fetchSystemOptions()
  } catch (error: any) {
    systemMessage.value = { type: 'error', text: error.message || '创建失败' }
  } finally {
    systemLoading.value = false
  }
}

// 创建单门课程
const handleCreateSingle = async () => {
  singleLoading.value = true
  singleMessage.value = null
  try {
    await courseAdminApi.createSingleCourse(singleForm.value)
    singleMessage.value = { type: 'success', text: '单门课程创建成功！' }
    // 重置表单
    singleForm.value = {
      courseName: '',
      parentEk: 0,
      courseTime: 0,
      price: 0,
      originalPrice: 0,
      image: '',
      courseDetail: '',
      courseTagClass: [],
      theoreticalQuestions: {
        singleQuestion: 0,
        multipleQuestion: 0,
        judgeQuestion: 0
      },
      videoQuestions: 0,
      attachment: 0
    }
  } catch (error: any) {
    singleMessage.value = { type: 'error', text: error.message || '创建失败' }
  } finally {
    singleLoading.value = false
  }
}

// 页面加载时获取体系课程选项和标签选项
// 卡密管理
const cardKeyCount = ref<number | null>(null)
const cardKeyLoading = ref(false)
const cardKeyMessage = ref<{ type: string; text: string } | null>(null)
const cardKeyList = ref<any[]>([])

const copyCode = async (code: string) => {
  await navigator.clipboard.writeText(code)
  copiedCode.value = code
  setTimeout(() => copiedCode.value = '', 2000)
}

const copiedCode = ref('')

const fetchCardKeyList = async () => {
  const res: any = await cardKeyApi.getList()
  cardKeyList.value = res.data || []
}

const handleGenerateCardKeys = async () => {
  if (!cardKeyCount.value) {
    cardKeyMessage.value = { type: 'error', text: '请输入生成数量' }
    return
  }
  cardKeyLoading.value = true
  cardKeyMessage.value = null
  try {
    await cardKeyApi.generate(cardKeyCount.value)
    cardKeyMessage.value = { type: 'success', text: `成功生成 ${cardKeyCount.value} 个卡密` }
    fetchCardKeyList()
  } catch (e: any) {
    cardKeyMessage.value = { type: 'error', text: e.response?.data?.msg || '生成失败' }
  } finally {
    cardKeyLoading.value = false
  }
}

onMounted(() => {
  fetchSystemOptions()
  fetchTagOptions()
  fetchBanners()
  fetchAllCourseOptions()
  fetchFlashSales()
  fetchCardKeyList()
})
</script>

<style scoped>
.course-admin {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
}

.form-section {
  background: #f5f5f5;
  padding: 20px;
  margin-bottom: 30px;
  border-radius: 8px;
}

h2 {
  margin-bottom: 20px;
  color: #333;
}

.form-group {
  margin-bottom: 15px;
}

.datetime-row {
  display: flex;
  align-items: center;
  gap: 6px;
}

.date-input {
  flex: 1;
  padding: 6px 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 14px;
}

.time-select {
  width: 64px;
  padding: 6px 4px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 14px;
  text-align: center;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input, select, textarea {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
}

textarea {
  min-height: 80px;
  resize: vertical;
}

button {
  width: 100%;
  padding: 10px;
  background: #409eff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

button:hover {
  background: #66b1ff;
}

button:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.message {
  margin-top: 15px;
  padding: 10px;
  border-radius: 4px;
}

.message.success {
  background: #f0f9ff;
  color: #67c23a;
  border: 1px solid #67c23a;
}

.message.error {
  background: #fef0f0;
  color: #f56c6c;
  border: 1px solid #f56c6c;
}

.upload-success {
  margin-left: 10px;
  color: #67c23a;
  font-weight: bold;
}

.checkbox-group {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-weight: normal;
}

.checkbox-label input[type="checkbox"] {
  width: auto;
  margin-right: 5px;
}
.flash-sale-list {
  margin-top: 20px;
}

.flash-sale-list h3 {
  margin-bottom: 12px;
  color: #333;
}

.flash-sale-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px;
  background: #fff;
  border-radius: 6px;
  margin-bottom: 10px;
  border: 1px solid #eee;
  gap: 12px;
}

.flash-sale-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 13px;
  color: #666;
  flex: 1;
}

.fs-name {
  font-weight: bold;
  color: #333;
  font-size: 14px;
}

.fs-status.on {
  color: #67c23a;
}

.fs-status.off {
  color: #999;
}

.flash-sale-actions {
  display: flex;
  gap: 8px;
  flex-shrink: 0;
}

.enable-btn {
  width: auto;
  padding: 6px 14px;
  background: #409eff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
}

.enable-btn:hover {
  background: #66b1ff;
}

.banner-list {
  margin-top: 20px;
}

.banner-list h3 {
  margin-bottom: 12px;
  color: #333;
}

.banner-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px;
  background: #fff;
  border-radius: 6px;
  margin-bottom: 10px;
  border: 1px solid #eee;
}

.banner-preview {
  width: 120px;
  height: 60px;
  object-fit: cover;
  border-radius: 4px;
  flex-shrink: 0;
}

.banner-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 13px;
  color: #666;
}

.delete-btn {
  width: auto;
  padding: 6px 14px;
  background: #f56c6c;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
  flex-shrink: 0;
}

.delete-btn:hover {
  background: #e64242;
}

.card-key-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 12px;
  font-size: 13px;
}

.card-key-table th,
.card-key-table td {
  border: 1px solid #e0e0e0;
  padding: 8px 12px;
  text-align: left;
}

.card-key-table th {
  background: #f5f5f5;
  font-weight: 600;
}

.code-cell {
  font-family: monospace;
  letter-spacing: 0.5px;
}

.copy-btn {
  margin-left: 8px;
  padding: 2px 8px;
  font-size: 12px;
  border: 1px solid #3b82f6;
  border-radius: 4px;
  background: #fff;
  color: #3b82f6;
  cursor: pointer;
}

.copy-btn:hover {
  background: #3b82f6;
  color: #fff;
}

.used { color: #999; }
.unused { color: #3b82f6; font-weight: 500; }
</style>
