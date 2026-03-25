<template>
  <div class="course-orders-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程订单</span>
      <span class="nav-placeholder"></span>
    </header>

    <div class="order-tabs">
      <button
        v-for="tab in tabs"
        :key="tab.value"
        :class="['order-tab', { active: activeTab === tab.value }]"
        @click="activeTab = tab.value"
      >
        {{ tab.label }}
      </button>
    </div>

    <main class="orders-content">
      <div v-if="!orders.length" class="empty">暂无订单</div>
      <div v-for="order in orders" :key="order.id" class="order-card">
        <div class="card-header">
          <img
            :src="order.courseImage ? imageBaseUrl + order.courseImage : '/images/course-detail-banner.png'"
            alt="课程图片"
            class="course-image"
          />
          <div class="course-info">
            <h3 class="course-name">{{ order.courseName }}</h3>
            <div class="course-category-tag">{{ order.courseCategoryStr || (order.courseCategory === 2 ? '体系课程' : '单门课程') }}</div>
            <div class="course-price">¥{{ (order.price / 100).toFixed(2) }}</div>
          </div>
          <div class="order-status-badge" :class="statusClass(order.status)">{{ statusText(order) }}</div>
        </div>
        <div class="order-details">
          <div class="detail-row" v-if="order.payTypeStr">
            <span class="label">支付方式：</span>
            <span class="value">{{ order.payTypeStr }}</span>
          </div>
          <div class="detail-row">
            <span class="label">订单编号：</span>
            <span class="value">{{ order.orderNo }}</span>
          </div>
          <div class="detail-row">
            <span class="label">购买时间：</span>
            <span class="value">{{ order.createdAt }}</span>
          </div>
        </div>
        <div class="card-footer">
          <template v-if="order.status === 1">
            <button class="view-course-btn" @click="$router.push('/course/' + order.courseEk)">查看详情</button>
            <button class="pay-btn" @click="$router.push('/purchase?ek=' + order.courseEk)">立即支付</button>
          </template>
          <template v-else>
            <button class="view-course-btn" @click="$router.push('/course/' + order.courseEk)">查看课程</button>
          </template>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { orderApi } from '@/api'

const orders = ref<any[]>([])
const activeTab = ref<'all' | 'pending' | 'completed' | 'cancelled'>('all')
const imageBaseUrl = 'http://localhost:8080'
const tabs = [
  { label: '全部', value: 'all' },
  { label: '待支付', value: 'pending' },
  { label: '已完成', value: 'completed' },
  { label: '已取消', value: 'cancelled' }
] as const

const fetchOrders = async () => {
  const status =
    activeTab.value === 'pending' ? 1 :
    activeTab.value === 'completed' ? 2 :
    activeTab.value === 'cancelled' ? 3 : undefined
  try {
    const res: any = await orderApi.getMyOrders(status)
    orders.value = res.data || []
  } catch (e) {}
}

onMounted(fetchOrders)
watch(activeTab, fetchOrders)

function statusClass(status: number) {
  if (status === 1) return 'pending'
  if (status === 2) return 'completed'
  return 'cancelled'
}

function statusText(order: any) {
  if (order.status === 1) return '待支付'
  if (order.status === 2) return '已完成'
  return '已取消'
}
</script>

<style scoped>
.course-orders-page {
  background-color: #f5f5f5;
  position: fixed;
  inset: 0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
}

.top-nav {
  display: grid;
  grid-template-columns: 24px 1fr 24px;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  flex-shrink: 0;
}

.order-tabs {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  align-items: center;
  background: #fff;
  border-bottom: 1px solid #edf0f6;
  flex-shrink: 0;
}

.order-tab {
  background: none;
  border: none;
  padding: 12px 0;
  font-size: 15px;
  color: #666;
  cursor: pointer;
  position: relative;
}

.order-tab.active {
  color: #3b82f6;
  font-weight: 600;
}

.order-tab.active::after {
  content: '';
  position: absolute;
  left: 50%;
  bottom: 0;
  width: 24px;
  height: 3px;
  border-radius: 99px;
  transform: translateX(-50%);
  background: #3b82f6;
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
  text-align: center;
}

.nav-placeholder {
  width: 24px;
  height: 24px;
  flex-shrink: 0;
}

.orders-content {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
  overscroll-behavior-y: contain;
  padding: 12px 16px;
}

.empty {
  text-align: center;
  color: #999;
  padding: 40px 0;
}

.order-card {
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 12px;
}

.card-header {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  position: relative;
  align-items: flex-start;
}

.course-image {
  width: 80px;
  height: 60px;
  border-radius: 8px;
  object-fit: cover;
  flex-shrink: 0;
}

.course-info {
  flex: 1;
}

.course-name {
  font-size: 15px;
  font-weight: 600;
  color: #333;
  margin: 0 0 6px 0;
  line-height: 1.4;
}

.course-category-tag {
  display: inline-block;
  font-size: 12px;
  color: #3b82f6;
  background: #eaf2ff;
  border-radius: 6px;
  padding: 3px 8px;
  margin-bottom: 6px;
}

.course-price {
  font-size: 16px;
  color: #333;
  font-weight: 600;
}

.order-status-badge {
  position: absolute;
  top: 0;
  right: 0;
  padding: 4px 12px;
  border-radius: 4px;
  font-size: 13px;
  font-weight: 500;
}

.order-status-badge.completed {
  background-color: #e8f5e9;
  color: #4caf50;
  border: 1px solid #4caf50;
}

.order-status-badge.pending {
  background-color: #fff3e0;
  color: #ff9800;
  border: 1px solid #ff9800;
}

.order-status-badge.cancelled {
  background-color: #f5f5f5;
  color: #999;
  border: 1px solid #ccc;
}

.order-details {
  border-top: 1px solid #f0f0f0;
  padding-top: 12px;
  margin-bottom: 12px;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 13px;
}

.detail-row .label { color: #999; }
.detail-row .value { color: #333; }

.card-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.view-course-btn {
  padding: 8px 32px;
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 10px;
  font-size: 14px;
  cursor: pointer;
}

.pay-btn {
  padding: 8px 32px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  cursor: pointer;
}
</style>
