<template>
  <div class="certificate-detail-page">
    <!-- 顶部导航 -->
    <header class="top-nav">
      <button class="back-btn" @click="$router.back()">
        <svg viewBox="0 0 24 24"><path d="M20,11V13H8L13.5,18.5L12.08,19.92L4.16,12L12.08,4.08L13.5,5.5L8,11H20Z" /></svg>
      </button>
      <span class="nav-title">学习证书</span>
    </header>

    <main class="certificate-content">
      <!-- 成功提示 -->
      <div class="success-section">
        <div class="success-icon">
          <svg viewBox="0 0 24 24">
            <path d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z" fill="#4caf50" />
          </svg>
        </div>
        <h2 class="success-title">恭喜您！</h2>
        <p class="success-text">已成功完成本课程的学习</p>
      </div>

      <!-- 证书展示 -->
      <div class="certificate-card">
        <div class="certificate-header">
          <h1 class="certificate-title">学时证书</h1>
          <p class="certificate-subtitle">LEARNING HOURS CERTIFICATE</p>
        </div>

        <div class="certificate-body">
          <div class="info-row">
            <span class="info-label">学员姓名/Student Name</span>
            <span class="info-value">张建国</span>
          </div>

          <div class="info-row">
            <span class="info-label">课程名称/Course Name</span>
            <span class="info-value">自动化集装箱码头操作实务(高级)</span>
          </div>

          <div class="stats-row">
            <div class="stat-item">
              <span class="stat-label">累计学时</span>
              <span class="stat-value">24.5学时</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">最高成绩</span>
              <span class="stat-value">96分</span>
            </div>
          </div>

          <div class="issuer-section">
            <div class="issuer-left">
              <div class="issuer-info">
                <span class="issuer-label">颁发机构</span>
                <div class="issuer-row">
                  <div class="issuer-logo">
                    <svg viewBox="0 0 24 24" class="logo-icon">
                      <path d="M12,3L1,9L12,15L21,10.09V17H23V9M5,13.18V17.18L12,21L19,17.18V13.18L12,17L5,13.18Z" fill="#3b82f6" />
                    </svg>
                  </div>
                  <span class="issuer-name">港口学堂</span>
                </div>
              </div>
              <div class="issue-date">
                <span class="date-label">发证日期：</span>
                <span class="date-value">2024年10月24日</span>
              </div>
            </div>
            <div class="stamp">
              <svg viewBox="0 0 100 100" class="stamp-svg">
                <circle cx="50" cy="50" r="45" fill="none" stroke="#d32f2f" stroke-width="3"/>
                <text x="50" y="35" text-anchor="middle" fill="#d32f2f" font-size="12" font-weight="bold">港口认证专用章</text>
                <text x="50" y="65" text-anchor="middle" fill="#d32f2f" font-size="20" font-weight="bold">★</text>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- 底部操作按钮 -->
    <footer class="action-footer">
      <button class="share-btn">分享</button>
      <button class="download-btn">下载</button>
    </footer>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.certificate-detail-page {
  background: linear-gradient(to bottom, #e3f2fd 0%, #f5f5f5 40%);
  min-height: 100vh;
  padding-bottom: 80px;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: transparent;
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

.certificate-content {
  padding: 20px 16px;
}

.success-section {
  text-align: center;
  margin-bottom: 30px;
}

.success-icon {
  width: 60px;
  height: 60px;
  margin: 0 auto 16px;
  background-color: #e8f5e9;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.success-icon svg {
  width: 36px;
  height: 36px;
}

.success-title {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin: 0 0 8px 0;
}

.success-text {
  font-size: 15px;
  color: #666;
  margin: 0;
}

.certificate-card {
  background-color: #fff;
  border-radius: 26px;
  padding: 30px 24px;
  border: 2px solid #3b82f6;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.certificate-header {
  text-align: center;
  margin-bottom: 30px;
  padding-bottom: 20px;
}

.certificate-title {
  font-size: 28px;
  font-weight: 700;
  color: #3b82f6;
  margin: 0 0 8px 0;
  letter-spacing: 2px;
}

.certificate-subtitle {
  font-size: 14px;
  color: #3b82f6;
  margin: 0;
  letter-spacing: 1px;
}

.certificate-body {
  padding: 20px 0;
}

.info-row {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #f0f0f0;
}

.info-label {
  font-size: 13px;
  color: #999;
  margin-bottom: 6px;
}

.info-value {
  font-size: 16px;
  color: #333;
  font-weight: 500;
}

.stats-row {
  display: flex;
  gap: 40px;
  margin: 30px 0;
  padding-bottom: 30px;
  border-bottom: 1px solid #f0f0f0;
}

.stat-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.stat-label {
  font-size: 13px;
  color: #999;
  margin-bottom: 8px;
}

.stat-value {
  font-size: 20px;
  color: #3b82f6;
  font-weight: 700;
}

.issuer-section {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  margin: 30px 0 0;
}

.issuer-left {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.issuer-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.issuer-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.issuer-logo {
  width: 50px;
  height: 50px;
  background-color: #e3f2fd;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo-icon {
  width: 30px;
  height: 30px;
}

.issuer-text {
  display: flex;
  flex-direction: column;
}

.issuer-label {
  font-size: 12px;
  color: #999;
  margin-bottom: 4px;
}

.issuer-name {
  font-size: 16px;
  color: #333;
  font-weight: 600;
}

.stamp {
  width: 80px;
  height: 80px;
}

.stamp-svg {
  width: 100%;
  height: 100%;
}

.issue-date {
  text-align: left;
  padding-top: 0;
}

.date-label {
  font-size: 13px;
  color: #999;
}

.date-value {
  font-size: 13px;
  color: #666;
}

.action-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 12px 16px;
  background-color: #fff;
  display: flex;
  gap: 12px;
  box-shadow: 0 -2px 8px rgba(0,0,0,0.05);
}

.share-btn {
  flex: 1;
  padding: 14px;
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
}

.download-btn {
  flex: 1;
  padding: 14px;
  background-color: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
}
</style>
