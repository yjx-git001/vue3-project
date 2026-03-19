<template>
  <div class="course-purchase">
    <header class="purchase-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程购买</span>
      <div class="nav-placeholder"></div>
    </header>

    <main class="purchase-content">
      <!-- 课程信息 -->
      <section class="course-section">
        <img :src="courseDetail.image ? imageBaseUrl + courseDetail.image : '/images/course-detail-banner.png'" alt="Course" class="course-thumb" />
        <div class="course-details">
          <h2 class="course-title">{{ courseDetail.courseName }}</h2>
          <span class="course-tag">{{ courseDetail.courseCategory === 2 ? '体系课程' : '单门课程' }}</span>
        </div>
      </section>

      <!-- 价格信息 -->
      <section class="price-section">
        <div class="price-row">
          <span class="price-label">课程价格</span>
          <span class="price-value">¥{{ (price / 100).toFixed(2) }}</span>
        </div>
        <div class="price-row total">
          <span class="price-value total-price">合计¥{{ (price / 100).toFixed(2) }}</span>
        </div>
      </section>

      <!-- 支付方式 -->
      <section class="payment-section">
        <h3 class="section-title">支付方式</h3>
        <div class="payment-option" :class="{ active: paymentMethod === 'wechat' }" @click="paymentMethod = 'wechat'">
          <div class="payment-row">
            <div class="payment-left">
              <svg viewBox="0 0 24 24" class="payment-icon wechat"><path d="M9.5,4C5.36,4 2,6.69 2,10C2,11.89 3.08,13.56 4.78,14.66L4,17L6.5,15.5C7.39,15.81 8.37,16 9.41,16C9.15,15.37 9,14.7 9,14C9,10.69 12.13,8 16,8C16.19,8 16.38,8 16.56,8.03C15.54,5.69 12.78,4 9.5,4M6.5,6.5C7.33,6.5 8,7.17 8,8C8,8.83 7.33,9.5 6.5,9.5C5.67,9.5 5,8.83 5,8C5,7.17 5.67,6.5 6.5,6.5M11.5,6.5C12.33,6.5 13,7.17 13,8C13,8.83 12.33,9.5 11.5,9.5C10.67,9.5 10,8.83 10,8C10,7.17 10.67,6.5 11.5,6.5M16,9C12.69,9 10,11.24 10,14C10,16.76 12.69,19 16,19C16.67,19 17.31,18.92 17.91,18.75L20,20L19.38,18.13C20.95,17.22 22,15.71 22,14C22,11.24 19.31,9 16,9M14,11.5C14.55,11.5 15,11.95 15,12.5C15,13.05 14.55,13.5 14,13.5C13.45,13.5 13,13.05 13,12.5C13,11.95 13.45,11.5 14,11.5M18,11.5C18.55,11.5 19,11.95 19,12.5C19,13.05 18.55,13.5 18,13.5C17.45,13.5 17,13.05 17,12.5C17,11.95 17.45,11.5 18,11.5Z" /></svg>
              <span class="payment-name">微信支付</span>
            </div>
            <div class="payment-check" v-if="paymentMethod === 'wechat'">
              <svg viewBox="0 0 24 24" class="check-icon"><path d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z" /></svg>
            </div>
            <div class="payment-uncheck" v-else></div>
          </div>
        </div>
        <div class="payment-option" :class="{ active: paymentMethod === 'card' }" @click="paymentMethod = 'card'">
          <div class="payment-row">
            <div class="payment-left">
              <svg viewBox="0 0 24 24" class="payment-icon card"><path d="M20,8H4V6H20M20,18H4V12H20M20,4H4C2.89,4 2,4.89 2,6V18A2,2 0 0,0 4,20H20A2,2 0 0,0 22,18V6C22,4.89 21.1,4 20,4Z" /></svg>
              <span class="payment-name">学习卡</span>
            </div>
            <div class="payment-check" v-if="paymentMethod === 'card'">
              <svg viewBox="0 0 24 24" class="check-icon"><path d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z" /></svg>
            </div>
            <div class="payment-uncheck" v-else></div>
          </div>
          <div v-if="paymentMethod === 'card'" class="card-input-section" @click.stop>
            <input type="text" class="card-input" placeholder="请输入卡密" v-model="cardNumber" />
          </div>
        </div>
      </section>
    </main>

    <footer class="purchase-footer">
      <div class="footer-info">
        <div class="total-amount">¥{{ (price / 100).toFixed(2) }}</div>
      </div>
      <div v-if="payError" style="color:#ef4444;font-size:12px;margin-right:8px">{{ payError }}</div>
      <button class="submit-btn" @click="handleSubmit" :disabled="paying">{{ paying ? '处理中...' : '提交订单' }}</button>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { courseApi, orderApi } from '@/api'

const route = useRoute()
const router = useRouter()
const paymentMethod = ref('wechat')
const cardNumber = ref('')
const courseDetail = ref<any>({})
const imageBaseUrl = 'http://localhost:8080'
const paying = ref(false)
const payError = ref('')

const price = computed(() => courseDetail.value.price || 0)

onMounted(async () => {
  const ek = Number(route.query.ek)
  if (!ek) return
  try {
    const res: any = await courseApi.getCourseDetail(ek)
    courseDetail.value = res.data
  } catch (e) {}
})

const handleSubmit = async () => {
  paying.value = true
  payError.value = ''
  try {
    const ek = Number(route.query.ek)
    const payType = paymentMethod.value === 'wechat' ? 1 : 2
    await orderApi.create(ek, payType, cardNumber.value || undefined)
    router.push('/course/' + ek)
  } catch (e: any) {
    payError.value = e.response?.data?.msg || '支付失败'
  } finally {
    paying.value = false
  }
}
</script>

<style scoped>
.course-purchase {
  background-color: #f5f7fa;
  min-height: 100vh;
  padding-bottom: 80px;
  overflow-x: hidden;
  max-width: 100vw;
  margin: 0 auto;
}

.purchase-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
}

.nav-placeholder {
  width: 24px;
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
  width: 20px;
  height: 20px;
  fill: #666;
}

.purchase-content {
  padding: 16px;
}

.course-section {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
}

.course-thumb {
  width: 80px;
  height: 60px;
  border-radius: 8px;
  object-fit: cover;
}

.course-details {
  flex: 1;
}

.course-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.course-tag {
  display: inline-block;
  background-color: #e3f2fd;
  color: #1976d2;
  font-size: 12px;
  padding: 4px 12px;
  border-radius: 12px;
}

.price-section {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 16px;
}

.price-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
}

.price-row.total {
  border-top: 1px solid #eee;
  margin-top: 8px;
  padding-top: 16px;
  justify-content: flex-end;
}

.price-label {
  font-size: 15px;
  color: #666;
}

.total-label {
  color: #333;
  font-weight: 600;
}

.price-value {
  font-size: 15px;
  color: #333;
  font-weight: 500;
}

.price-value.discount {
  color: #333;
}

.price-value.total-price {
  font-size: 20px;
  color: #333;
  font-weight: 700;
}

.payment-section {
  border-radius: 12px;
  padding: 16px;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
}

.payment-option {
  display: flex;
  flex-direction: column;
  padding: 16px;
  border: 1px solid #e0e0e0;
  border-radius: 16px;
  margin-bottom: 12px;
  cursor: pointer;
  transition: all 0.2s;
}

.payment-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.payment-option.active {
  border-color: #e0e0e0;
  background-color: #fff;
}

.payment-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.payment-icon {
  width: 24px;
  height: 24px;
}

.payment-icon.wechat {
  fill: #07c160;
}

.payment-icon.card {
  fill: #ff6a00;
}

.payment-name {
  font-size: 15px;
  color: #333;
  font-weight: 500;
}

.payment-check {
  width: 20px;
  height: 20px;
  background-color: #07c160;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.check-icon {
  width: 14px;
  height: 14px;
  fill: #fff;
}

.payment-uncheck {
  width: 20px;
  height: 20px;
  border: 2px solid #ccc;
  border-radius: 50%;
}

.card-input-section {
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #eee;
}

.card-input {
  width: 100%;
  padding: 4px 0;
  border: none;
  outline: none;
  font-size: 15px;
  color: #999;
  background: transparent;
}

.card-input:focus {
  border-color: #ff6a00;
}

.purchase-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 -2px 8px rgba(0,0,0,0.1);
}

.footer-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.total-amount {
  font-size: 24px;
  color: #ff6a00;
  font-weight: 700;
}

.saved-amount {
  font-size: 12px;
  color: #999;
}

.submit-btn {
  padding: 12px 40px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 25px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
}
</style>
