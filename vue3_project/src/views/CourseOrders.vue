<template>
  <div class="course-orders-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程订单</span>
    </header>

    <!-- 订单状态标签 -->
    <div class="order-tabs">
      <div
        v-for="tab in orderTabs"
        :key="tab.id"
        :class="['tab-item', { active: activeTab === tab.id }]"
        @click="activeTab = tab.id"
      >
        {{ tab.name }}
      </div>
    </div>

    <main class="orders-content">
      <div
        v-for="order in ordersList"
        :key="order.id"
        class="order-card"
      >
        <div class="card-header">
          <img :src="order.image" :alt="order.courseName" class="course-image" />
          <div class="course-info">
            <h3 class="course-name">{{ order.courseName }}</h3>
            <div class="course-tag">{{ order.courseType }}</div>
            <div class="course-price">¥{{ order.price }}</div>
          </div>
          <div class="order-status-badge" :class="order.status">
            {{ getStatusText(order.status) }}
          </div>
        </div>

        <div class="order-details">
          <div class="detail-row">
            <span class="label">实际支付：</span>
            <span class="value price">¥{{ order.actualPayment }}</span>
          </div>
          <div class="detail-row">
            <span class="label">支付方式：</span>
            <span class="value">{{ order.paymentMethod }}</span>
          </div>
          <div class="detail-row">
            <span class="label">订单编号：</span>
            <span class="value">{{ order.orderNumber }}</span>
          </div>
          <div class="detail-row">
            <span class="label">购买时间：</span>
            <span class="value">{{ order.purchaseTime }}</span>
          </div>
        </div>

        <div class="card-footer">
          <button
            v-if="order.status === 'completed'"
            class="view-course-btn"
            @click="$router.push('/course_content/' + order.courseId)"
          >
            查看课程
          </button>
          <template v-else-if="order.status === 'pending'">
            <button class="detail-btn" @click="$router.push('/course/' + order.courseId)">查看详情</button>
            <button class="pay-btn" @click="$router.push('/purchase')">立即支付</button>
          </template>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const activeTab = ref('all')

const orderTabs = ref([
  { id: 'all', name: '全部' },
  { id: 'pending', name: '待支付' },
  { id: 'completed', name: '已完成' },
  { id: 'cancelled', name: '已取消' }
])

const ordersList = ref([
  {
    id: 1,
    courseId: 1,
    courseName: '港口特种设备检修课程',
    courseType: '体系课程',
    price: '350.00',
    actualPayment: '129.00',
    paymentMethod: '微信支付',
    orderNumber: '39429309234',
    purchaseTime: '2025.12.25  10:00',
    status: 'completed',
    image: '/images/course-detail-banner.png'
  },
  {
    id: 2,
    courseId: 2,
    courseName: '港口特种设备检修课程',
    courseType: '体系课程',
    price: '350.00',
    actualPayment: '350.00',
    paymentMethod: '',
    orderNumber: '39429309235',
    purchaseTime: '2025.12.26  14:30',
    status: 'pending',
    image: '/images/course-detail-banner.png'
  }
])

const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    completed: '已完成',
    pending: '待支付',
    cancelled: '已取消'
  }
  return statusMap[status] || ''
}
</script>

<style scoped>
.course-orders-page {
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

.order-tabs {
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

.orders-content {
  padding: 12px 16px;
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

.course-tag {
  display: inline-block;
  background-color: #e3f2fd;
  color: #2196f3;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
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
  border-radius: 12px;
  font-size: 13px;
  font-weight: 500;
}

.order-status-badge.completed {
  background-color: #e8f5e9;
  color: #4caf50;
}

.order-status-badge.pending {
  background-color: #fff3e0;
  color: #ff9800;
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

.detail-row .label {
  color: #999;
}

.detail-row .value {
  color: #333;
}

.detail-row .value.price {
  color: #3b82f6;
  font-weight: 600;
  font-size: 14px;
}

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
  border-radius: 20px;
  font-size: 14px;
  cursor: pointer;
}

.detail-btn {
  padding: 8px 24px;
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 20px;
  font-size: 14px;
  cursor: pointer;
}

.pay-btn {
  padding: 8px 32px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 20px;
  font-size: 14px;
  cursor: pointer;
}
</style>
