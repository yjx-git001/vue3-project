<template>
  <div class="single-course-detail">
    <header class="detail-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">课程详情</span>
      <div class="nav-placeholder"></div>
    </header>

    <main class="detail-content">
      <!-- 顶部视频区 -->
      <section class="video-section">
        <img src="/images/single-course.png" alt="Course" class="course-banner" />
        <div class="play-overlay">
          <div class="play-btn"></div>
        </div>
      </section>

      <!-- 课程信息区 -->
      <section class="course-info">
        <h1 class="course-name">【生产装卸】散矿</h1>
        <div class="price-section">
          <span class="price">¥50</span>
        </div>
        <p class="course-desc">课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述课程描述。述课程描述课程描述课程描述课程描述。述课程描述课程描述课程描述课程描述。述课程描述课程描述课程描述课程描述。</p>
        <button class="more-btn">查看更多</button>
      </section>

      <!-- 功能卡片区 -->
      <section class="cards-section">
        <div class="card-row">
          <!-- 理论题库 -->
          <div class="card theory-card">
            <div class="card-corner">
              <span class="card-title">理论题库</span>
            </div>
            <div class="card-icon">
              <svg viewBox="0 0 24 24" class="icon-svg"><path d="M19,20H5V4H7V7H17V4H19M12,2A1,1 0 0,1 13,3A1,1 0 0,1 12,4A1,1 0 0,1 11,3A1,1 0 0,1 12,2M19,11H5V20H19V11Z" /></svg>
            </div>
            <div class="card-content">
              <div class="stat-line">
                <span class="stat-label">单选题</span>
                <span class="stat-value orange">10</span>
                <span class="stat-label">多选题</span>
                <span class="stat-value orange">5</span>
              </div>
              <div class="stat-line">
                <span class="stat-label">判断题</span>
                <span class="stat-value orange">8</span>
              </div>
            </div>
          </div>

          <!-- 视频课程 -->
          <div class="card video-card">
            <div class="card-corner">
              <span class="card-title">视频课程</span>
            </div>
            <div class="card-icon">
              <svg viewBox="0 0 24 24" class="icon-svg"><path d="M10,16.5V7.5L16,12M18.4,12L14,8.4V15.6M6.6,12L11,8.4V15.6H6.6" /></svg>
            </div>
            <div class="card-content">
              <div class="stat-line">
                <span class="stat-label">数量</span>
                <span class="stat-value purple">1</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 附件 -->
        <div class="card attachment-card">
          <div class="card-corner">
            <span class="card-title">附件</span>
          </div>
          <div class="card-icon">
            <svg viewBox="0 0 24 24" class="icon-svg"><path d="M14,3H10V5H6V9H4V15H6V19H10V21H14V19H18V15H20V9H16V5H14V3M8,17H6V15H8V17M8,13H6V11H8V13M8,9H6V7H8V9M16,17H14V15H16V17M16,13H14V11H16V13M16,9H14V7H16V9Z" /></svg>
          </div>
          <div class="card-content">
            <div class="stat-line">
              <span class="stat-label">数量</span>
              <span class="stat-value teal">1</span>
            </div>
          </div>
        </div>
      </section>
    </main>

    <!-- 底部购买栏 -->
    <footer class="buy-footer">
      <button class="buy-btn" @click="$router.push('/purchase')">立即购买</button>
    </footer>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.single-course-detail {
  background-color: #f5f7fa;
  min-height: 100vh;
  padding-bottom: 80px;
  overflow-x: hidden;
  max-width: 100vw;
  margin: 0 auto;
}

.detail-nav {
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

.detail-content {
  padding: 10px 16px;
}

.video-section {
  position: relative;
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 16px;
}

.course-banner {
  width: 100%;
  display: block;
}

.play-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(0,0,0,0.15);
}

.play-btn {
  width: 60px;
  height: 60px;
  background-color: rgba(255,255,255,0.9);
  border-radius: 50%;
  position: relative;
}

.play-btn::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-35%, -50%);
  width: 0;
  height: 0;
  border-top: 12px solid transparent;
  border-bottom: 12px solid transparent;
  border-left: 18px solid #3b82f6;
}

.course-info {
  background-color: #fff;
  padding: 16px;
  border-radius: 12px;
  margin-bottom: 16px;
}

.course-name {
  font-size: 22px;
  font-weight: 600;
  color: #222;
  margin-bottom: 12px;
}

.price-section {
  margin-bottom: 12px;
}

.price {
  font-size: 26px;
  color: #ff6a00;
  font-weight: 700;
}

.course-desc {
  font-size: 15px;
  color: #555;
  line-height: 1.7;
  margin-bottom: 16px;
}

.more-btn {
  background-color: #f5f7fa;
  border: none;
  color: #888;
  font-size: 14px;
  padding: 8px 24px;
  border-radius: 20px;
  width: 100%;
  cursor: pointer;
}

.cards-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.card-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.card {
  background-color: #fff;
  border-radius: 12px;
  padding: 12px;
  position: relative;
  overflow: hidden;
}

.theory-card {
  background: linear-gradient(135deg, #ff9500 0%, #fff3e0 100%);
}

.video-card {
  background: linear-gradient(135deg, #9c27b0 0%, #f3e5f5 100%);
}

.attachment-card {
  background: linear-gradient(135deg, #26a69a 0%, #e0f7fa 100%);
}

.card-corner {
  position: absolute;
  top: 0;
  left: 0;
  background-color: rgba(255,255,255,0.2);
  padding: 6px 16px 6px 12px;
  border-bottom-right-radius: 20px;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: #fff;
}

.card-icon {
  position: absolute;
  top: 8px;
  right: 12px;
  opacity: 0.3;
}

.icon-svg {
  width: 36px;
  height: 36px;
  fill: #fff;
}

.card-content {
  margin-top: 40px;
}

.stat-line {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #666;
}

.stat-value {
  font-size: 20px;
  font-weight: 700;
}

.orange {
  color: #ff6a00;
}

.purple {
  color: #9c27b0;
}

.teal {
  color: #26a69a;
}

.buy-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 12px 16px;
  background-color: #f5f7fa;
}

.buy-btn {
  width: 100%;
  padding: 14px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 20px;
  font-weight: 600;
  cursor: pointer;
}
</style>