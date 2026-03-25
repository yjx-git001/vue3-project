<template>
  <div class="coupon-list-page">
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24" aria-hidden="true">
          <path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" />
        </svg>
      </button>
      <span class="nav-title">优惠卡券</span>
      <span class="nav-placeholder"></span>
    </header>

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
      <div v-if="activeTab === 'unused'" class="coupon-list">
        <template v-for="coupon in unusedCoupons" :key="coupon.id">
          <article class="coupon-item">
            <section class="coupon-card">
              <section class="coupon-left">
                <div class="coupon-amount">
                  <span class="currency">¥</span>
                  <span class="amount-value">{{ coupon.amount }}</span>
                </div>
                <div class="coupon-type">{{ coupon.type }}</div>
              </section>

              <section class="coupon-right">
                <h3 class="coupon-title">{{ coupon.title }}</h3>
                <p class="coupon-validity">有效期：{{ coupon.validUntil }}</p>
                <button
                  type="button"
                  class="coupon-rules"
                  @click="toggleRules(coupon.id)"
                >
                  查看详细规则 {{ expandedRuleCouponId === coupon.id ? '^' : '>' }}
                </button>
              </section>

              <section class="coupon-actions">
                <button class="use-btn">去使用</button>
              </section>

              <span class="corner-tag">
                <em>未使用</em>
              </span>
            </section>

            <section v-if="expandedRuleCouponId === coupon.id" class="coupon-notice">
              <p>1.学习卡可兑换任意一门单门课程。</p>
              <p>2.退换课程需在有效期内完成使用。</p>
              <p>3.结算环节选择学习卡兑换，勾选可用的学习卡，即可0元兑换对应课程。</p>
            </section>
          </article>
        </template>
      </div>

      <div v-if="activeTab === 'used'" class="coupon-list">
        <template v-if="usedCoupons.length">
          <article
            v-for="coupon in usedCoupons"
            :key="coupon.id"
            class="coupon-card expired"
          >
            <section class="coupon-left">
              <div class="coupon-amount">
                <span class="currency">¥</span>
                <span class="amount-value">{{ coupon.amount }}</span>
              </div>
              <div class="coupon-type">{{ coupon.type }}</div>
            </section>

            <section class="coupon-right">
              <h3 class="coupon-title">{{ coupon.title }}</h3>
              <p class="coupon-validity">有效期：{{ coupon.validUntil }}</p>
              <a class="coupon-rules">查看详细规则 ></a>
            </section>

            <section class="coupon-actions">
              <button class="used-btn">已使用</button>
            </section>

            <span class="corner-tag expired">
              <em>已使用</em>
            </span>
          </article>
        </template>
        <section v-else class="empty-state">
          <svg viewBox="0 0 24 24" class="empty-icon" aria-hidden="true">
            <path
              d="M21.47,4.35L20.13,3.79V12.82L22.56,6.96C22.97,5.94 22.5,4.77 21.47,4.35M1.97,8.05L6.93,20C7.24,20.77 7.97,21.24 8.74,21.26C9,21.26 9.27,21.21 9.53,21.1L16.9,18.05C17.65,17.74 18.11,17 18.13,16.26C18.14,16 18.09,15.71 18,15.45L13,3.5C12.71,2.73 11.97,2.26 11.19,2.25C10.93,2.25 10.67,2.31 10.42,2.4L3.06,5.45C2.04,5.87 1.55,7.04 1.97,8.05M18.12,4.25A2,2 0 0,0 16.12,2.25H14.67L18.12,10.59"
            />
          </svg>
          <p class="empty-text">暂无已使用的卡券</p>
        </section>
      </div>

      <div v-if="activeTab === 'expired'" class="coupon-list">
        <template v-if="expiredCoupons.length">
          <article
            v-for="coupon in expiredCoupons"
            :key="coupon.id"
            class="coupon-card expired"
          >
            <section class="coupon-left">
              <div class="coupon-amount">
                <span class="currency">¥</span>
                <span class="amount-value">{{ coupon.amount }}</span>
              </div>
              <div class="coupon-type">{{ coupon.type }}</div>
            </section>

            <section class="coupon-right">
              <h3 class="coupon-title">{{ coupon.title }}</h3>
              <p class="coupon-validity">有效期：{{ coupon.validUntil }}</p>
              <a class="coupon-rules">查看详细规则 ></a>
            </section>

            <section class="coupon-actions">
              <button class="expired-btn">已失效</button>
            </section>

            <span class="corner-tag expired">
              <em>已失效</em>
            </span>
          </article>
        </template>
        <section v-else class="empty-state">
          <p class="empty-text">暂无已过期的卡券</p>
        </section>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { couponApi } from '@/api'

const activeTab = ref('unused')

const tabs = [
  { label: '未使用', value: 'unused' },
  { label: '已使用', value: 'used' },
  { label: '已过期', value: 'expired' }
]

const expandedRuleCouponId = ref<number | null>(null)
const coupons = ref<any[]>([])

const toggleRules = (couponId: number) => {
  expandedRuleCouponId.value = expandedRuleCouponId.value === couponId ? null : couponId
}

watch(activeTab, (tab) => {
  if (tab !== 'unused') {
    expandedRuleCouponId.value = null
  }
})

const normalizeCoupon = (item: any) => ({
  id: item.id,
  amount: Number(item.amount || 0),
  type: item.couponType || '单门课程',
  title: item.title || '学习卡兑换',
  validUntil: item.validUntil || '-',
  status: Number(item.status || 0)
})

const unusedCoupons = computed(() =>
  coupons.value.filter(c => Number(c.status) === 1).map(normalizeCoupon)
)

const usedCoupons = computed(() =>
  coupons.value.filter(c => Number(c.status) === 2).map(normalizeCoupon)
)

const expiredCoupons = computed(() =>
  coupons.value.filter(c => Number(c.status) === 3).map(normalizeCoupon)
)

const fetchCoupons = async () => {
  try {
    const res: any = await couponApi.getCoupons()
    coupons.value = res.data || []
  } catch {
    coupons.value = []
  }
}

onMounted(() => {
  fetchCoupons()
})
</script>

<style scoped>
.coupon-list-page {
  position: fixed;
  inset: 0;
  background: #eef1f6;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.top-nav {
  height: 52px;
  padding: 0 16px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  position: relative;
  background: #fff;
}

.back-btn {
  width: 24px;
  height: 24px;
  border: none;
  padding: 0;
  background: transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  flex-shrink: 0;
}

.back-btn svg {
  width: 24px;
  height: 24px;
  fill: #262b31;
}

.nav-title {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  font-size: 16px;
  font-weight: 700;
  color: #20262f;
  letter-spacing: 0.2px;
  text-align: center;
  white-space: nowrap;
}

.nav-placeholder {
  margin-left: auto;
  width: 24px;
  height: 24px;
  flex-shrink: 0;
}

.tabs {
  background: #fff;
  display: flex;
  padding: 0 8px;
}

.tab-item {
  flex: 1;
  text-align: center;
  padding: 14px 0 12px;
  font-size: 15px;
  color: #4d5561;
  position: relative;
  cursor: pointer;
}

.tab-item.active {
  color: #2f7df2;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  left: 50%;
  bottom: 0;
  width: 72px;
  height: 3px;
  transform: translateX(-50%);
  background: #2f7df2;
  border-radius: 2px;
}

.coupon-content {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  padding: 10px 10px 14px;
}

.coupon-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.coupon-item {
  background: #f8f9fb;
  border-radius: 12px;
  box-shadow: 0 1px 5px rgba(31, 40, 51, 0.04);
  overflow: hidden;
}

.coupon-card {
  position: relative;
  background: #f8f9fb;
  border-radius: 12px;
  display: grid;
  grid-template-columns: 108px 1fr auto;
  align-items: center;
  gap: 10px;
  min-height: 94px;
  padding: 0 12px 0 0;
  box-shadow: 0 1px 5px rgba(31, 40, 51, 0.04);
  overflow: hidden;
}

.coupon-item .coupon-card {
  border-radius: 0;
  box-shadow: none;
}

.coupon-card.expired {
  background: #f4f4f5;
}

.coupon-left {
  position: relative;
  height: 94px;
  border-radius: 12px 0 0 12px;
  background: linear-gradient(180deg, #fff8ea 0%, #ffeecb 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-right: 1px dashed #f0dcae;
}

.coupon-left::before,
.coupon-left::after {
  content: '';
  position: absolute;
  right: -6px;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #eef1f6;
}

.coupon-left::before {
  top: -6px;
}

.coupon-left::after {
  bottom: -6px;
}

.coupon-card.expired .coupon-left {
  background: linear-gradient(180deg, #f1f1f1 0%, #e7e7e7 100%);
  border-right-color: #dddddd;
}

.coupon-amount {
  display: inline-flex;
  align-items: baseline;
  gap: 1px;
  line-height: 1;
  font-size: 40px;
  font-weight: 700;
  color: #ef5f17;
  letter-spacing: -0.6px;
}

.currency {
  font-size: 0.56em;
  transform: translateY(-1px);
}

.amount-value {
  font-size: 1em;
}

.coupon-card.expired .coupon-amount {
  color: #8f8f8f;
}

.coupon-type {
  margin-top: 8px;
  font-size: 12px;
  color: #ef5f17;
}

.coupon-card.expired .coupon-type {
  color: #8f8f8f;
}

.coupon-right {
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.coupon-title {
  margin: 0;
  font-size: 16px;
  line-height: 1.15;
  color: #353c45;
  font-weight: 700;
}

.coupon-validity {
  margin: 0;
  font-size: 13px;
  line-height: 1.2;
  color: #59606b;
}

.coupon-rules {
  font-size: 12px;
  line-height: 1.2;
  color: #8f96a1;
  text-decoration: none;
  border: none;
  background: transparent;
  padding: 0;
  text-align: left;
  width: fit-content;
  cursor: pointer;
}

.coupon-actions {
  width: 104px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.use-btn,
.expired-btn {
  width: 90px;
  height: 38px;
  border: none;
  border-radius: 999px;
  font-size: 14px;
  font-weight: 600;
}

.use-btn {
  background: #f06d31;
  color: #fff;
  cursor: pointer;
}

.expired-btn {
  background: #d3d3d5;
  color: #f8f8f8;
  cursor: not-allowed;
}

.corner-tag {
  position: absolute;
  top: 8px;
  right: -20px;
  width: 72px;
  height: 16px;
  background: #ff842f;
  clip-path: polygon(8% 0, 92% 0, 100% 100%, 0 100%);
  transform: rotate(45deg);
  transform-origin: center;
  display: flex;
  align-items: center;
  justify-content: center;
  padding-left: 0;
  box-sizing: border-box;
  pointer-events: none;
}

.corner-tag em {
  transform: none;
  color: #fff;
  font-style: normal;
  font-size: 8px;
  font-weight: 600;
  letter-spacing: 0;
  line-height: 1;
  white-space: nowrap;
}

.corner-tag.expired {
  background: #c9c9c9;
}

.coupon-notice {
  background: #f8f9fb;
  border-radius: 0;
  padding: 16px;
  margin-top: 0;
  box-shadow: none;
}

.coupon-notice p {
  margin: 0 0 8px;
  font-size: 13px;
  line-height: 1.45;
  color: #7b838f;
}

.coupon-notice p:last-child {
  margin-bottom: 0;
}

.empty-state {
  background: #f8f9fb;
  border-radius: 12px;
  padding: 56px 18px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.empty-icon {
  width: 72px;
  height: 72px;
  fill: #d5d9de;
  margin-bottom: 12px;
}

.empty-text {
  margin: 0;
  font-size: 14px;
  color: #9098a3;
}
</style>
