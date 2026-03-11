<template>
  <div class="coupon-list-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">优惠卡券</span>
      <div class="nav-actions">
        <button class="icon-btn">
          <svg viewBox="0 0 24 24"><path d="M12,16A2,2 0 0,1 14,18A2,2 0 0,1 12,20A2,2 0 0,1 10,18A2,2 0 0,1 12,16M12,10A2,2 0 0,1 14,12A2,2 0 0,1 12,14A2,2 0 0,1 10,12A2,2 0 0,1 12,10M12,4A2,2 0 0,1 14,6A2,2 0 0,1 12,8A2,2 0 0,1 10,6A2,2 0 0,1 12,4Z" /></svg>
        </button>
        <button class="icon-btn">
          <svg viewBox="0 0 24 24"><path d="M18,16.08C17.24,16.08 16.56,16.38 16.04,16.85L8.91,12.7C8.96,12.47 9,12.24 9,12C9,11.76 8.96,11.53 8.91,11.3L15.96,7.19C16.5,7.69 17.21,8 18,8A3,3 0 0,0 21,5A3,3 0 0,0 18,2C16.36,2 15,3.36 15,5C15,5.24 15.04,5.47 15.09,5.7L8.04,9.81C7.5,9.31 6.79,9 6,9A3,3 0 0,0 3,12A3,3 0 0,0 6,15C6.79,15 7.5,14.69 8.04,14.19L15.16,18.35C15.11,18.56 15.08,18.78 15.08,19C15.08,20.61 16.39,21.92 18,21.92C19.61,21.92 20.92,20.61 20.92,19A2.92,2.92 0 0,0 18,16.08Z" /></svg>
        </button>
      </div>
    </header>

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

    <main class="coupon-content">
      <!-- 未使用 -->
      <div v-if="activeTab === 'unused'" class="coupon-list">
        <div
          v-for="coupon in unusedCoupons"
          :key="coupon.id"
          class="coupon-card"
        >
          <div class="coupon-left">
            <div class="coupon-amount">¥{{ coupon.amount }}</div>
            <div class="coupon-type">{{ coupon.type }}</div>
          </div>
          <div class="coupon-right">
            <div class="coupon-header">
              <h3 class="coupon-title">{{ coupon.title }}</h3>
            </div>
            <div class="coupon-validity">有效期：{{ coupon.validUntil }}</div>
            <a class="coupon-rules">查看详细规则 ></a>
            <button class="use-btn">去使用</button>
          </div>
          <span class="course-tag">未使用</span>
        </div>

        <div class="coupon-notice">
          <p>1.学习卡可兑换任意一门单门课程。</p>
          <p>2.退换课程需在有效期内完成使用。</p>
          <p>3.结算环节选择学习卡兑换，勾选可用的学习卡，即可兑换对应课程。</p>
        </div>
      </div>

      <!-- 已使用 -->
      <div v-if="activeTab === 'used'" class="coupon-list">
        <div class="empty-state">
          <svg viewBox="0 0 24 24" class="empty-icon">
            <path d="M21.47,4.35L20.13,3.79V12.82L22.56,6.96C22.97,5.94 22.5,4.77 21.47,4.35M1.97,8.05L6.93,20C7.24,20.77 7.97,21.24 8.74,21.26C9,21.26 9.27,21.21 9.53,21.1L16.9,18.05C17.65,17.74 18.11,17 18.13,16.26C18.14,16 18.09,15.71 18,15.45L13,3.5C12.71,2.73 11.97,2.26 11.19,2.25C10.93,2.25 10.67,2.31 10.42,2.4L3.06,5.45C2.04,5.87 1.55,7.04 1.97,8.05M18.12,4.25A2,2 0 0,0 16.12,2.25H14.67L18.12,10.59" />
          </svg>
          <p class="empty-text">暂无已使用的卡券</p>
        </div>
      </div>

      <!-- 已过期 -->
      <div v-if="activeTab === 'expired'" class="coupon-list">
        <div
          v-for="coupon in expiredCoupons"
          :key="coupon.id"
          class="coupon-card expired"
        >
          <div class="coupon-left">
            <div class="coupon-amount">¥{{ coupon.amount }}</div>
            <div class="coupon-type">{{ coupon.type }}</div>
          </div>
          <div class="coupon-right">
            <div class="coupon-header">
              <h3 class="coupon-title">{{ coupon.title }}</h3>
              <span class="expired-tag">已过期</span>
            </div>
            <div class="coupon-validity">有效期：{{ coupon.validUntil }}</div>
            <a class="coupon-rules">查看详细规则 ></a>
            <button class="expired-btn">已失效</button>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const activeTab = ref('unused')

const tabs = [
  { label: '未使用', value: 'unused' },
  { label: '已使用', value: 'used' },
  { label: '已过期', value: 'expired' }
]

const unusedCoupons = ref([
  {
    id: 1,
    amount: 20,
    type: '单门课程',
    title: '学习卡兑换',
    tag: '单门课程',
    validUntil: '2025.12.30'
  },
  {
    id: 2,
    amount: 20,
    type: '单门课程',
    title: '学习卡兑换',
    tag: '单门课程',
    validUntil: '2025.12.30'
  }
])

const expiredCoupons = ref([
  {
    id: 3,
    amount: 20,
    type: '单门课程',
    title: '学习卡兑换',
    validUntil: '2025.12.30'
  }
])
</script>

<style scoped>
.coupon-list-page {
  background-color: #f5f5f5;
  min-height: 100vh;
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

.tabs {
  display: flex;
  background-color: #fff;
  padding: 0 16px;
}

.tab-item {
  flex: 1;
  text-align: center;
  padding: 14px 0;
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

.coupon-content {
  padding: 16px;
}

.coupon-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.coupon-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  display: flex;
  gap: 16px;
  position: relative;
}

.coupon-card.expired {
  opacity: 0.6;
}

.coupon-left {
  background: linear-gradient(135deg, #fff4e6 0%, #ffe7ba 100%);
  border-radius: 8px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-width: 90px;
}

.coupon-card.expired .coupon-left {
  background: linear-gradient(135deg, #f5f5f5 0%, #e0e0e0 100%);
}

.coupon-amount {
  font-size: 28px;
  font-weight: 700;
  color: #ff6b00;
  margin-bottom: 4px;
}

.coupon-card.expired .coupon-amount {
  color: #999;
}

.coupon-type {
  font-size: 12px;
  color: #ff6b00;
}

.coupon-card.expired .coupon-type {
  color: #999;
}

.coupon-right {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.coupon-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.coupon-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.course-tag {
  position: absolute;
  top: 0;
  right: 0;
  background-color: #ff6b00;
  color: #fff;
  padding: 4px 12px;
  border-radius: 0 12px 0 12px;
  font-size: 12px;
  transform: rotate(45deg);
  transform-origin: top right;
  width: 80px;
  text-align: center;
  z-index: 1;
}

.expired-tag {
  position: absolute;
  top: 0;
  right: 0;
  background-color: #999;
  color: #fff;
  padding: 4px 12px;
  border-radius: 0 12px 0 12px;
  font-size: 12px;
  transform: rotate(45deg);
  transform-origin: top right;
  width: 80px;
  text-align: center;
  z-index: 1;
}

.coupon-validity {
  font-size: 13px;
  color: #666;
}

.coupon-rules {
  font-size: 12px;
  color: #999;
  text-decoration: none;
  cursor: pointer;
}

.use-btn {
  align-self: flex-end;
  padding: 6px 20px;
  background-color: #ff6b00;
  color: #fff;
  border: none;
  border-radius: 20px;
  font-size: 13px;
  cursor: pointer;
  font-weight: 500;
}

.expired-btn {
  align-self: flex-end;
  padding: 6px 20px;
  background-color: #e0e0e0;
  color: #999;
  border: none;
  border-radius: 20px;
  font-size: 13px;
  cursor: not-allowed;
  font-weight: 500;
}

.coupon-notice {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-top: 12px;
}

.coupon-notice p {
  font-size: 13px;
  color: #666;
  line-height: 1.6;
  margin: 8px 0;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
}

.empty-icon {
  width: 80px;
  height: 80px;
  fill: #ccc;
  margin-bottom: 16px;
}

.empty-text {
  font-size: 14px;
  color: #999;
}
</style>
