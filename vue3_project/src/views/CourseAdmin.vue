<template>
  <div class="course-admin">
    <h1>课程管理</h1>

    <!-- 页面主体：左侧竖排导航 + 右侧内容 -->
    <div class="admin-layout">
      <!-- 左侧竖排 Tab 导航 -->
      <nav class="admin-sidenav">
        <div
          v-for="item in adminTabs"
          :key="item.key"
          :class="['sidenav-item', { active: activeTab === item.key }]"
          @click="activeTab = item.key"
        >
          <span class="sidenav-icon">{{ item.icon }}</span>
          <span class="sidenav-label">{{ item.label }}</span>
        </div>
      </nav>

      <!-- 右侧内容区 -->
      <div class="admin-content">

    <!-- Banner 管理 -->
    <section v-show="activeTab === 'banner'" class="form-section">
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
    <section v-show="activeTab === 'flashsale'" class="form-section">
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
    <section v-show="activeTab === 'system'" class="form-section">
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
    <section v-show="activeTab === 'single'" class="form-section">
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

    <!-- 题目配置 -->
    <section v-show="activeTab === 'question'" class="form-section">
      <h2>题目配置</h2>
      <p class="section-tip">💡 题目绑定到「单门课程」，题目数量由创建课程时配置的各题型数量决定。</p>

      <!-- 选择单门课程 -->
      <div class="form-group">
        <label>选择单门课程：</label>
        <select v-model.number="qCourseEk" @change="fetchQuestionList(true)">
          <option :value="0" disabled>请选择单门课程</option>
          <option v-for="opt in singleCourseOptions" :key="opt.ek" :value="opt.ek">
            {{ opt.courseName }}
          </option>
        </select>
      </div>

      <!-- 加载中 -->
      <div v-if="qCourseDetailLoading" class="q-loading">加载课程配置中...</div>

      <!-- 已选课程但无配额 -->
      <div v-else-if="qCourseEk > 0 && qTypeTabs.length === 0" class="q-empty">
        该课程未配置题目数量，请先在「创建单门课程」中设置各题型数量
      </div>

      <template v-else-if="qCourseEk > 0 && qTypeTabs.length > 0">

        <!-- ===== 视频标题配置 ===== -->
        <div class="content-config-block">
          <div class="content-config-header">
            <span class="content-config-title">📹 视频目录配置</span>
            <span class="content-config-tip">共 {{ videoTitles.length }} 个视频</span>
          </div>
          <div
            v-for="(v, idx) in videoTitles"
            :key="idx"
            class="content-config-row"
          >
            <span class="content-config-num">{{ idx + 1 }}</span>
            <input
              v-model="videoTitles[idx]"
              type="text"
              :placeholder="`第 ${idx + 1} 个视频标题`"
              class="content-config-input"
            />
            <div class="duration-inputs">
              <input
                v-model="videoMinutes[idx]"
                type="number"
                min="0"
                placeholder="分"
                class="duration-min"
              />
              <span class="duration-sep">:</span>
              <input
                v-model="videoSeconds[idx]"
                type="number"
                min="0"
                max="59"
                placeholder="秒"
                class="duration-sec"
              />
            </div>
          </div>
          <button
            class="content-save-btn"
            :disabled="videoSaveLoading"
            @click="handleSaveVideos"
          >
            {{ videoSaveLoading ? '保存中...' : '保存视频标题' }}
          </button>
          <div v-if="videoSaveMsg" class="message" :class="videoSaveMsg.type">{{ videoSaveMsg.text }}</div>
        </div>

        <!-- ===== 附件名称配置 ===== -->
        <div class="content-config-block">
          <div class="content-config-header">
            <span class="content-config-title">📎 附件配置</span>
            <span class="content-config-tip">共 {{ attachmentNames.length }} 个附件</span>
          </div>
          <div
            v-for="(a, idx) in attachmentNames"
            :key="idx"
            class="content-config-row"
          >
            <span class="content-config-num">{{ idx + 1 }}</span>
            <input
              v-model="attachmentNames[idx]"
              type="text"
              :placeholder="`请输入第 ${idx + 1} 个附件名称`"
              class="content-config-input"
            />
          </div>
          <button
            class="content-save-btn"
            :disabled="attachmentSaveLoading"
            @click="handleSaveAttachments"
          >
            {{ attachmentSaveLoading ? '保存中...' : '保存附件名称' }}
          </button>
          <div v-if="attachmentSaveMsg" class="message" :class="attachmentSaveMsg.type">{{ attachmentSaveMsg.text }}</div>
        </div>

        <!-- ===== 笔记配置 ===== -->
        <div class="content-config-block">
          <div class="content-config-header">
            <span class="content-config-title">📝 课程笔记配置</span>
            <span class="content-config-tip">体系课程笔记会汇总所有单门课程的笔记内容</span>
          </div>
          <textarea
            v-model="notesContent"
            placeholder="请输入该单门课程的笔记内容..."
            class="notes-textarea"
            rows="8"
          ></textarea>
          <button
            class="content-save-btn"
            :disabled="notesSaveLoading"
            @click="handleSaveNotes"
          >
            {{ notesSaveLoading ? '保存中...' : '保存笔记' }}
          </button>
          <div v-if="notesSaveMsg" class="message" :class="notesSaveMsg.type">{{ notesSaveMsg.text }}</div>
        </div>

        <!-- ===== 模拟考试配置 ===== -->
        <div class="content-config-block">
          <div class="content-config-header">
            <span class="content-config-title">📋 模拟考试配置</span>
            <span class="content-config-tip">自定义模拟考试各题型出题数量（从题库随机抽取）</span>
          </div>
          <div class="mock-exam-counts">
            <div class="mock-count-item" v-if="qCourseQuota.singleQuestion > 0">
              <label>单选题数量：</label>
              <input type="number" v-model.number="mockSingleCount" min="0" :max="qCount(1)" />
              <span class="mock-count-tip">
                题库已录入 <b :class="qCount(1) < qCourseQuota.singleQuestion ? 'tip-warn' : 'tip-ok'">{{ qCount(1) }}</b> / {{ qCourseQuota.singleQuestion }} 道
                <span v-if="qCount(1) < qCourseQuota.singleQuestion" class="tip-warn">（题库不足，最多可出 {{ qCount(1) }} 道）</span>
              </span>
            </div>
            <div class="mock-count-item" v-if="qCourseQuota.multipleQuestion > 0">
              <label>多选题数量：</label>
              <input type="number" v-model.number="mockMultipleCount" min="0" :max="qCount(2)" />
              <span class="mock-count-tip">
                题库已录入 <b :class="qCount(2) < qCourseQuota.multipleQuestion ? 'tip-warn' : 'tip-ok'">{{ qCount(2) }}</b> / {{ qCourseQuota.multipleQuestion }} 道
                <span v-if="qCount(2) < qCourseQuota.multipleQuestion" class="tip-warn">（题库不足，最多可出 {{ qCount(2) }} 道）</span>
              </span>
            </div>
            <div class="mock-count-item" v-if="qCourseQuota.judgeQuestion > 0">
              <label>判断题数量：</label>
              <input type="number" v-model.number="mockJudgeCount" min="0" :max="qCount(3)" />
              <span class="mock-count-tip">
                题库已录入 <b :class="qCount(3) < qCourseQuota.judgeQuestion ? 'tip-warn' : 'tip-ok'">{{ qCount(3) }}</b> / {{ qCourseQuota.judgeQuestion }} 道
                <span v-if="qCount(3) < qCourseQuota.judgeQuestion" class="tip-warn">（题库不足，最多可出 {{ qCount(3) }} 道）</span>
              </span>
            </div>
            <div v-if="qCourseQuota.singleQuestion === 0 && qCourseQuota.multipleQuestion === 0 && qCourseQuota.judgeQuestion === 0" class="q-empty-type">
              请先选择单门课程
            </div>
          </div>
          <button class="content-save-btn" :disabled="mockSaveLoading" @click="handleSaveMockConfig">
            {{ mockSaveLoading ? '保存中...' : '保存模拟考试配置' }}
          </button>
          <div v-if="mockSaveMsg" class="message" :class="mockSaveMsg.type">{{ mockSaveMsg.text }}</div>
        </div>

        <!-- ===== 考题配置 ===== -->
        <!-- 考题配置：左侧竖排 Tab + 右侧内容 -->
        <div class="q-layout">
          <!-- 题型 Tab（竖排） -->
          <div class="q-type-tabs">
            <div
              v-for="tab in qTypeTabs"
              :key="tab.type"
              :class="['q-tab', { active: qActiveType === tab.type }]"
              @click="switchQType(tab.type)"
            >
              <span>{{ tab.label }}</span>
              <span class="q-tab-count" :class="{ full: qCount(tab.type) >= tab.quota }">
                {{ qCount(tab.type) }}/{{ tab.quota }}
              </span>
            </div>
          </div>

          <!-- 右侧内容区 -->
          <div class="q-content-area">

            <!-- 当前题型的题目列表 -->
            <div class="question-list-wrap">
          <div v-if="questionList.filter(q => q.questionType === qActiveType).length === 0" class="q-empty-type">
            暂无{{ ['', '单选', '多选', '判断'][qActiveType] }}题，请在下方添加
          </div>
          <div
            v-for="(q, idx) in questionList.filter(q => q.questionType === qActiveType)"
            :key="q.id"
            class="q-item"
          >
            <div class="q-header">
              <span class="q-index">第 {{ idx + 1 }} 题</span>
              <button class="q-del-btn" @click="handleDeleteQuestion(q.id)">删除</button>
            </div>
            <div class="q-content">{{ q.content }}</div>
            <div v-if="qActiveType !== 3" class="q-options">
              <span v-if="q.optionA">A. {{ q.optionA }}</span>
              <span v-if="q.optionB">B. {{ q.optionB }}</span>
              <span v-if="q.optionC">C. {{ q.optionC }}</span>
              <span v-if="q.optionD">D. {{ q.optionD }}</span>
            </div>
            <div class="q-answer">✓ 正确答案：{{ q.answer === 'T' ? '正确' : q.answer === 'F' ? '错误' : q.answer }}</div>
            <div v-if="q.analysis" class="q-analysis">解析：{{ q.analysis }}</div>
          </div>
        </div>

        <!-- 添加题目表单 -->
        <div class="add-question-form">
          <div class="add-q-header">
            <h3>添加{{ ['', '单选', '多选', '判断'][qActiveType] }}题</h3>
            <span v-if="qIsFull" class="q-full-tip">已达上限 {{ qCurrentQuota }} 道，无法继续添加</span>
            <span v-else class="q-remain-tip">还需添加 {{ qCurrentQuota - qCurrentCount }} 道</span>
          </div>

          <template v-if="!qIsFull">
            <div class="form-group">
              <label>题干：</label>
              <textarea v-model="qForm.content" placeholder="请输入题目内容" rows="3"></textarea>
            </div>

            <!-- 单选/多选的选项 -->
            <template v-if="qActiveType !== 3">
              <div class="form-group">
                <label>选项 A：</label>
                <input v-model="qForm.optionA" type="text" placeholder="选项A内容" />
              </div>
              <div class="form-group">
                <label>选项 B：</label>
                <input v-model="qForm.optionB" type="text" placeholder="选项B内容" />
              </div>
              <div class="form-group">
                <label>选项 C：</label>
                <input v-model="qForm.optionC" type="text" placeholder="选项C内容" />
              </div>
              <div class="form-group">
                <label>选项 D：</label>
                <input v-model="qForm.optionD" type="text" placeholder="选项D内容" />
              </div>
              <div class="form-group">
                <label>正确答案：</label>
                <!-- 单选 -->
                <div v-if="qActiveType === 1" class="answer-radio-group">
                  <label v-for="opt in availableOptions" :key="opt" class="checkbox-label">
                    <input type="radio" :value="opt" v-model="qForm.answer" />
                    {{ opt }}
                  </label>
                </div>
                <!-- 多选 -->
                <div v-else class="answer-radio-group">
                  <label v-for="opt in availableOptions" :key="opt" class="checkbox-label">
                    <input type="checkbox" :value="opt" v-model="qFormMultiAnswer" />
                    {{ opt }}
                  </label>
                  <span class="q-answer-preview">已选：{{ qFormMultiAnswer.slice().sort().join('') || '（未选）' }}</span>
                </div>
              </div>
            </template>

            <!-- 判断题 -->
            <template v-else>
              <div class="form-group">
                <label>正确答案：</label>
                <div class="answer-radio-group">
                  <label class="checkbox-label">
                    <input type="radio" value="T" v-model="qForm.answer" /> 正确（T）
                  </label>
                  <label class="checkbox-label">
                    <input type="radio" value="F" v-model="qForm.answer" /> 错误（F）
                  </label>
                </div>
              </div>
            </template>

            <div class="form-group">
              <label>解析（可选）：</label>
              <textarea v-model="qForm.analysis" placeholder="解题解析，可不填" rows="2"></textarea>
            </div>
            <button type="button" :disabled="qLoading" @click="handleAddQuestion">
              {{ qLoading ? '添加中...' : '添加题目' }}
            </button>
            <div v-if="qMessage" class="message" :class="qMessage.type">{{ qMessage.text }}</div>
          </template>
        </div>

          </div><!-- /.q-content-area -->
        </div><!-- /.q-layout -->
      </template>
    </section>

    <!-- 优惠卡券管理 -->
    <section v-show="activeTab === 'coupon'" class="admin-section">
      <h2>优惠卡券管理</h2>
      <div class="form-group">
        <label>发放数量：</label>
        <input v-model.number="couponIssueForm.count" type="number" min="1" max="100" placeholder="1-100" />
      </div>
      <div class="form-group">
        <label>优惠金额（元）：</label>
        <input v-model.number="couponIssueForm.amount" type="number" min="1" placeholder="如 20" />
      </div>
      <div class="form-group">
        <label>卡券标题：</label>
        <input v-model.trim="couponIssueForm.title" type="text" placeholder="默认：学习卡兑换" />
      </div>
      <div class="form-group">
        <label>卡券类型：</label>
        <select v-model="couponIssueForm.couponType">
          <option value="单门课程">单门课程</option>
          <option value="体系课程">体系课程</option>
        </select>
      </div>
      <div class="form-group">
        <label>有效期：</label>
        <input v-model="couponIssueForm.validUntil" type="date" />
      </div>
      <div class="form-group">
        <label>指定用户：</label>
        <div class="user-select-combobox">
          <input
            v-model.trim="couponUserKeyword"
            type="text"
            placeholder="搜索并选择用户（必填）"
            @focus="handleCouponUserFocus"
            @blur="handleCouponUserBlur"
            @input="handleCouponUserInput"
          />
          <div v-if="showCouponUserDropdown" class="user-dropdown">
            <div
              v-for="u in couponUserOptions"
              :key="u.id"
              class="user-option"
              @mousedown.prevent="selectCouponUser(u)"
            >
              {{ formatCouponUserOption(u) }}
            </div>
            <div v-if="couponUserOptions.length === 0" class="user-option user-option-empty">
              未找到匹配用户
            </div>
          </div>
        </div>
        <div class="issue-tip">指定用户为必填，可通过姓名或手机号搜索后选择</div>
      </div>
      <button @click="handleIssueCoupons" :disabled="couponLoading">
        {{ couponLoading ? '发放中...' : '发放优惠卡券' }}
      </button>

      <div v-if="couponMessage" class="message" :class="couponMessage.type">{{ couponMessage.text }}</div>

      <table v-if="couponList.length" class="coupon-table">
        <thead>
          <tr>
            <th>卡券码</th>
            <th>金额(元)</th>
            <th>标题</th>
            <th>类型</th>
            <th>有效期</th>
            <th>用户ID</th>
            <th>用户</th>
            <th>手机号</th>
            <th>状态</th>
            <th>发放时间</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="c in couponList" :key="c.id">
            <td class="code-cell">{{ c.code }}</td>
            <td>{{ c.amount }}</td>
            <td>{{ c.title }}</td>
            <td>{{ c.couponType }}</td>
            <td>{{ c.validUntil }}</td>
            <td>{{ c.userId }}</td>
            <td>{{ c.userName || '-' }}</td>
            <td>{{ c.userPhone || '-' }}</td>
            <td>{{ c.statusStr }}</td>
            <td>{{ c.createdAt }}</td>
          </tr>
        </tbody>
      </table>
    </section>

    <!-- 卡密管理 -->
    <section v-show="activeTab === 'cardkey'" class="admin-section">
      <h2>卡密管理（用于卡密支付）</h2>
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

      </div><!-- /.admin-content -->
    </div><!-- /.admin-layout -->
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { courseAdminApi, uploadApi, bannerAdminApi, flashSaleAdminApi, courseApi, cardKeyApi, couponAdminApi, userApi, questionApi, courseContentApi, mockExamApi } from '@/api'

// ===== 页面级竖排 Tab =====
const activeTab = ref('banner')
const adminTabs = [
  { key: 'banner',    icon: '🖼️', label: 'Banner管理' },
  { key: 'flashsale', icon: '⚡', label: '限时秒杀' },
  { key: 'system',    icon: '📚', label: '体系课程' },
  { key: 'single',    icon: '📖', label: '单门课程' },
  { key: 'question',  icon: '📝', label: '题目配置' },
  { key: 'coupon',    icon: '🎟️', label: '优惠卡券' },
  { key: 'cardkey',   icon: '🔑', label: '卡密管理' },
]

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
// ==================== 视频 & 附件配置 ====================
const videoTitles = ref<string[]>([])
const videoMinutes = ref<string[]>([])  // 分钟
const videoSeconds = ref<string[]>([])  // 秒
const attachmentNames = ref<string[]>([])
const videoSaveLoading = ref(false)
const attachmentSaveLoading = ref(false)
const videoSaveMsg = ref<{ type: string; text: string } | null>(null)
const attachmentSaveMsg = ref<{ type: string; text: string } | null>(null)

const handleSaveVideos = async () => {
  videoSaveLoading.value = true
  videoSaveMsg.value = null
  try {
    const videos = videoTitles.value.map((title, i) => {
      const m = String(videoMinutes.value[i] ?? 0).padStart(2, '0')
      const s = String(videoSeconds.value[i] ?? 0).padStart(2, '0')
      return { sort: i + 1, title, duration: `${m}:${s}` }
    })
    await courseContentApi.saveVideos(qCourseEk.value, videos)
    videoSaveMsg.value = { type: 'success', text: '视频标题保存成功！' }
  } catch (e: any) {
    videoSaveMsg.value = { type: 'error', text: e.message || '保存失败' }
  } finally {
    videoSaveLoading.value = false
  }
}

const handleSaveAttachments = async () => {
  attachmentSaveLoading.value = true
  attachmentSaveMsg.value = null
  try {
    const attachments = attachmentNames.value.map((name, i) => ({ sort: i + 1, name }))
    await courseContentApi.saveAttachments(qCourseEk.value, attachments)
    attachmentSaveMsg.value = { type: 'success', text: '附件名称保存成功！' }
  } catch (e: any) {
    attachmentSaveMsg.value = { type: 'error', text: e.message || '保存失败' }
  } finally {
    attachmentSaveLoading.value = false
  }
}

const notesContent = ref('')
const notesSaveLoading = ref(false)
const notesSaveMsg = ref<{ type: string; text: string } | null>(null)

const handleSaveNotes = async () => {
  notesSaveLoading.value = true
  notesSaveMsg.value = null
  try {
    await courseContentApi.saveNotes(qCourseEk.value, notesContent.value)
    notesSaveMsg.value = { type: 'success', text: '笔记保存成功！' }
  } catch (e: any) {
    notesSaveMsg.value = { type: 'error', text: e.message || '保存失败' }
  } finally {
    notesSaveLoading.value = false
  }
}

// ==================== 题目配置 ====================
const qCourseEk = ref(0)
const questionList = ref<any[]>([])
const qLoading = ref(false)
const qMessage = ref<{ type: string; text: string } | null>(null)
const qFormMultiAnswer = ref<string[]>([])
// 当前题型 Tab：1=单选 2=多选 3=判断
const qActiveType = ref(1)

// 只显示单门课程（category=1）
const singleCourseOptions = computed(() =>
  allCourseOptions.value.filter(c => c.courseCategory === 1)
)

// 当前选中课程的配额（从课程详情获取）
const qCourseQuota = ref({ singleQuestion: 0, multipleQuestion: 0, judgeQuestion: 0 })
const qCourseDetailLoading = ref(false)

// 各题型已配置数量
const qCount = (type: number) => questionList.value.filter(q => q.questionType === type).length

// 当前题型配额
const qCurrentQuota = computed(() => {
  if (qActiveType.value === 1) return qCourseQuota.value.singleQuestion
  if (qActiveType.value === 2) return qCourseQuota.value.multipleQuestion
  return qCourseQuota.value.judgeQuestion
})
// 当前题型已录数量
const qCurrentCount = computed(() => qCount(qActiveType.value))
// 是否已满额
const qIsFull = computed(() => qCurrentCount.value >= qCurrentQuota.value && qCurrentQuota.value > 0)

// 题型 Tab 列表
const qTypeTabs = computed(() => [
  { type: 1, label: '单选题', quota: qCourseQuota.value.singleQuestion },
  { type: 2, label: '多选题', quota: qCourseQuota.value.multipleQuestion },
  { type: 3, label: '判断题', quota: qCourseQuota.value.judgeQuestion },
].filter(t => t.quota > 0))

const qForm = ref({
  content: '',
  optionA: '',
  optionB: '',
  optionC: '',
  optionD: '',
  answer: '',
  analysis: ''
})

// 答案选项固定 A B C D
const availableOptions = ['A', 'B', 'C', 'D']

const resetQForm = () => {
  qForm.value = { content: '', optionA: '', optionB: '', optionC: '', optionD: '', answer: '', analysis: '' }
  qFormMultiAnswer.value = []
  qMessage.value = null
}

const switchQType = (type: number) => {
  qActiveType.value = type
  resetQForm()
}

const fetchQuestionList = async (resetTab = false) => {
  if (!qCourseEk.value) return
  qCourseDetailLoading.value = true
  videoSaveMsg.value = null
  attachmentSaveMsg.value = null
  try {
    // 获取课程详情拿配额
    const detailRes: any = await courseApi.getCourseDetail(qCourseEk.value, 1)
    const detail = detailRes.data
    const tq = detail?.theoreticalQuestions
    qCourseQuota.value = {
      singleQuestion: tq?.singleQuestion || 0,
      multipleQuestion: tq?.multipleQuestion || 0,
      judgeQuestion: tq?.judgeQuestion || 0,
    }
    // 初始化视频标题数组（按配额数量）
    const videoCount = detail?.videoQuestions || 0
    const attachmentCount = detail?.attachment || 0

    // 加载已保存的视频标题
    try {
      const vRes: any = await courseContentApi.getVideos(qCourseEk.value)
      const saved: { sort: number; title: string; duration: string }[] = vRes.data || []
      videoTitles.value = Array.from({ length: videoCount }, (_, i) => {
        const found = saved.find(v => v.sort === i + 1)
        return found ? found.title : ''
      })
      videoMinutes.value = Array.from({ length: videoCount }, (_, i) => {
        const found = saved.find(v => v.sort === i + 1)
        if (found?.duration) {
          const [mm = ''] = found.duration.split(':')
          return mm.replace(/^0/, '') || '0'
        }
        return ''
      })
      videoSeconds.value = Array.from({ length: videoCount }, (_, i) => {
        const found = saved.find(v => v.sort === i + 1)
        if (found?.duration) {
          const [, ss = ''] = found.duration.split(':')
          return ss.replace(/^0/, '') || '0'
        }
        return ''
      })
    } catch {
      videoTitles.value = Array(videoCount).fill('')
      videoMinutes.value = Array(videoCount).fill('')
      videoSeconds.value = Array(videoCount).fill('')
    }

    // 加载已保存的附件名称
    try {
      const aRes: any = await courseContentApi.getAttachments(qCourseEk.value)
      const saved: { sort: number; name: string }[] = aRes.data || []
      attachmentNames.value = Array.from({ length: attachmentCount }, (_, i) => {
        const found = saved.find(a => a.sort === i + 1)
        return found ? found.name : ''
      })
    } catch {
      attachmentNames.value = Array(attachmentCount).fill('')
    }

    // 加载已保存的笔记
    try {
      const nRes: any = await courseContentApi.getNotes(qCourseEk.value)
      notesContent.value = nRes.data?.content || ''
    } catch {
      notesContent.value = ''
    }
    notesSaveMsg.value = null

    // 加载模拟考试配置
    try {
      const mRes: any = await mockExamApi.getConfig(qCourseEk.value)
      mockSingleCount.value = mRes.data?.singleCount || 0
      mockMultipleCount.value = mRes.data?.multipleCount || 0
      mockJudgeCount.value = mRes.data?.judgeCount || 0
    } catch {
      mockSingleCount.value = 0
      mockMultipleCount.value = 0
      mockJudgeCount.value = 0
    }
    mockSaveMsg.value = null

    // 默认激活第一个有配额的题型（仅切换课程时重置）
    if (resetTab) {
      const firstTab = [
        { type: 1, quota: qCourseQuota.value.singleQuestion },
        { type: 2, quota: qCourseQuota.value.multipleQuestion },
        { type: 3, quota: qCourseQuota.value.judgeQuestion },
      ].find(t => t.quota > 0)
      if (firstTab) qActiveType.value = firstTab.type
    }
  } catch {
    qCourseQuota.value = { singleQuestion: 0, multipleQuestion: 0, judgeQuestion: 0 }
    videoTitles.value = []
    videoMinutes.value = []
    videoSeconds.value = []
    attachmentNames.value = []
    notesContent.value = ''
    mockSingleCount.value = 0
    mockMultipleCount.value = 0
    mockJudgeCount.value = 0
  } finally {
    qCourseDetailLoading.value = false
  }
  try {
    const res: any = await questionApi.getList(qCourseEk.value)
    questionList.value = res.data || []
  } catch {
    questionList.value = []
  }
  resetQForm()
}

const handleAddQuestion = async () => {
  if (qIsFull.value) {
    qMessage.value = { type: 'error', text: `该题型已达配额上限（${qCurrentQuota.value} 道）` }
    return
  }
  if (!qForm.value.content.trim()) {
    qMessage.value = { type: 'error', text: '请填写题干' }
    return
  }
  if (qActiveType.value !== 3) {
    if (!qForm.value.optionA.trim() || !qForm.value.optionB.trim()) {
      qMessage.value = { type: 'error', text: '请至少填写选项A和选项B' }
      return
    }
  }
  // 多选答案组合
  let answer = qForm.value.answer
  if (qActiveType.value === 2) {
    if (qFormMultiAnswer.value.length < 2) {
      qMessage.value = { type: 'error', text: '多选题至少选择2个正确答案' }
      return
    }
    answer = qFormMultiAnswer.value.sort().join('')
  }
  if (!answer) {
    qMessage.value = { type: 'error', text: '请选择正确答案' }
    return
  }
  qLoading.value = true
  qMessage.value = null
  try {
    await questionApi.create({
      courseEk: qCourseEk.value,
      questionType: qActiveType.value,
      content: qForm.value.content,
      optionA: qForm.value.optionA,
      optionB: qForm.value.optionB,
      optionC: qForm.value.optionC || undefined,
      optionD: qForm.value.optionD || undefined,
      answer,
      analysis: qForm.value.analysis || undefined
    })
    qMessage.value = { type: 'success', text: '题目添加成功！' }
    resetQForm()
    await fetchQuestionList()
  } catch (e: any) {
    qMessage.value = { type: 'error', text: e.message || '添加失败' }
  } finally {
    qLoading.value = false
  }
}

const handleDeleteQuestion = async (id: number) => {
  if (!confirm('确认删除这道题目？')) return
  try {
    await questionApi.delete(id)
    await fetchQuestionList()
  } catch {}
}

// ==================== 模拟考试配置 ====================
const mockSingleCount = ref(0)
const mockMultipleCount = ref(0)
const mockJudgeCount = ref(0)
const mockSaveLoading = ref(false)
const mockSaveMsg = ref<{ type: string; text: string } | null>(null)

const handleSaveMockConfig = async () => {
  // 前端校验：出题数不能超过已录入题目数
  if (mockSingleCount.value > qCount(1)) {
    mockSaveMsg.value = { type: 'error', text: `单选题出题数（${mockSingleCount.value}）超过题库已录入数量（${qCount(1)}）` }
    return
  }
  if (mockMultipleCount.value > qCount(2)) {
    mockSaveMsg.value = { type: 'error', text: `多选题出题数（${mockMultipleCount.value}）超过题库已录入数量（${qCount(2)}）` }
    return
  }
  if (mockJudgeCount.value > qCount(3)) {
    mockSaveMsg.value = { type: 'error', text: `判断题出题数（${mockJudgeCount.value}）超过题库已录入数量（${qCount(3)}）` }
    return
  }
  mockSaveLoading.value = true
  mockSaveMsg.value = null
  try {
    await mockExamApi.saveConfig({
      courseEk: qCourseEk.value,
      singleCount: mockSingleCount.value,
      multipleCount: mockMultipleCount.value,
      judgeCount: mockJudgeCount.value
    })
    mockSaveMsg.value = { type: 'success', text: '模拟考试配置保存成功！' }
  } catch (e: any) {
    mockSaveMsg.value = { type: 'error', text: e.message || '保存失败' }
  } finally {
    mockSaveLoading.value = false
  }
}

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

// 优惠卡券管理
const couponLoading = ref(false)
const couponMessage = ref<{ type: string; text: string } | null>(null)
const couponList = ref<any[]>([])
const couponUserKeyword = ref('')
const couponUserOptions = ref<Array<{ id: number; name: string; phone: string }>>([])
const showCouponUserDropdown = ref(false)
const couponIssueForm = ref({
  count: 1,
  amount: 20,
  title: '学习卡兑换',
  couponType: '单门课程',
  validUntil: '',
  userId: 0
})

const fetchCouponList = async () => {
  const res: any = await couponAdminApi.getList()
  couponList.value = res.data || []
}

const formatCouponUserOption = (u: { id: number; name: string; phone: string }) => {
  return `${u.name}（ID:${u.id} / ${u.phone}）`
}

const fetchUserOptions = async () => {
  const res: any = await userApi.getAdminUserOptions(couponUserKeyword.value || undefined)
  couponUserOptions.value = res.data || []
}

const handleCouponUserInput = async () => {
  couponIssueForm.value.userId = 0
  showCouponUserDropdown.value = true
  await fetchUserOptions()
}

const handleCouponUserFocus = async () => {
  showCouponUserDropdown.value = true
  await fetchUserOptions()
}

const handleCouponUserBlur = () => {
  setTimeout(() => {
    showCouponUserDropdown.value = false
  }, 120)
}

const selectCouponUser = (u: { id: number; name: string; phone: string }) => {
  couponIssueForm.value.userId = u.id
  couponUserKeyword.value = formatCouponUserOption(u)
  showCouponUserDropdown.value = false
}

const handleIssueCoupons = async () => {
  const f = couponIssueForm.value
  if (!f.count || f.count < 1) {
    couponMessage.value = { type: 'error', text: '请输入发放数量' }
    return
  }
  if (!f.amount || f.amount < 1) {
    couponMessage.value = { type: 'error', text: '请输入优惠金额' }
    return
  }
  if (!f.validUntil) {
    couponMessage.value = { type: 'error', text: '请选择有效期' }
    return
  }
  if (!f.userId) {
    couponMessage.value = { type: 'error', text: '请选择指定用户' }
    return
  }

  couponLoading.value = true
  couponMessage.value = null
  try {
    await couponAdminApi.issue({
      count: f.count,
      amount: f.amount,
      title: f.title || undefined,
      couponType: f.couponType || undefined,
      validUntil: f.validUntil,
      userId: f.userId
    })
    couponMessage.value = { type: 'success', text: `成功发放 ${f.count} 张优惠卡券` }
    couponIssueForm.value = {
      ...couponIssueForm.value,
      count: 1,
      userId: 0
    }
    couponUserKeyword.value = ''
    fetchCouponList()
    fetchUserOptions()
  } catch (e: any) {
    couponMessage.value = { type: 'error', text: e.response?.data?.msg || '发放失败' }
  } finally {
    couponLoading.value = false
  }
}

onMounted(() => {
  fetchSystemOptions()
  fetchTagOptions()
  fetchBanners()
  fetchAllCourseOptions()
  fetchFlashSales()
  fetchCardKeyList()
  fetchCouponList()
  fetchUserOptions()
})
</script>

<style scoped>
.course-admin {
  max-width: 1100px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
}

/* 页面主体：左侧导航 + 右侧内容 */
.admin-layout {
  display: flex;
  gap: 0;
  align-items: flex-start;
  min-height: 600px;
}

/* 左侧竖排导航栏 */
.admin-sidenav {
  width: 110px;
  flex-shrink: 0;
  background: #f0f2f5;
  border-radius: 10px 0 0 10px;
  border: 1px solid #e0e0e0;
  border-right: none;
  overflow: hidden;
  position: sticky;
  top: 20px;
}

.sidenav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  padding: 14px 8px;
  cursor: pointer;
  font-size: 12px;
  color: #555;
  border-bottom: 1px solid #e0e0e0;
  transition: background 0.2s, color 0.2s;
  user-select: none;
}
.sidenav-item:last-child {
  border-bottom: none;
}
.sidenav-item:hover {
  background: #e6eaf0;
}
.sidenav-item.active {
  background: #3b82f6;
  color: #fff;
}
.sidenav-icon {
  font-size: 18px;
  line-height: 1;
}
.sidenav-label {
  text-align: center;
  line-height: 1.3;
}

/* 右侧内容区 */
.admin-content {
  flex: 1;
  min-width: 0;
  border: 1px solid #e0e0e0;
  border-radius: 0 10px 10px 0;
  background: #fff;
  padding: 24px;
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

.coupon-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 12px;
  font-size: 13px;
}

.issue-tip {
  margin-top: 6px;
  font-size: 12px;
  color: #909399;
}

.user-select-combobox {
  position: relative;
}

.user-dropdown {
  position: absolute;
  left: 0;
  right: 0;
  top: calc(100% + 4px);
  max-height: 220px;
  overflow-y: auto;
  border: 1px solid #dcdfe6;
  border-radius: 6px;
  background: #fff;
  z-index: 20;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.08);
}

.user-option {
  padding: 8px 10px;
  font-size: 13px;
  color: #333;
  cursor: pointer;
}

.user-option:hover {
  background: #f5f7fa;
}

.user-option-empty {
  color: #999;
  cursor: default;
}

.card-key-table th,
.card-key-table td {
  border: 1px solid #e0e0e0;
  padding: 8px 12px;
  text-align: left;
}

.coupon-table th,
.coupon-table td {
  border: 1px solid #e0e0e0;
  padding: 8px 10px;
  text-align: left;
}

.card-key-table th {
  background: #f5f5f5;
  font-weight: 600;
}

.coupon-table th {
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

/* 视频 & 附件配置 */
.content-config-block {
  background: #fff;
  border: 1px solid #e8e8e8;
  border-radius: 8px;
  padding: 14px 16px;
  margin-bottom: 16px;
}
.content-config-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 12px;
}
.content-config-title {
  font-size: 14px;
  font-weight: 600;
  color: #333;
}
.content-config-tip {
  font-size: 12px;
  color: #999;
  background: #f5f5f5;
  padding: 2px 8px;
  border-radius: 10px;
}
.content-config-row {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 8px;
}
.content-config-num {
  width: 24px;
  height: 24px;
  background: #3b82f6;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
  flex-shrink: 0;
}
.content-config-input {
  flex: 1;
  padding: 7px 10px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  box-sizing: border-box;
}
.duration-inputs {
  display: flex;
  align-items: center;
  gap: 4px;
  flex-shrink: 0;
}
.duration-min {
  width: 52px;
  padding: 7px 6px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  text-align: center;
  box-sizing: border-box;
}
.duration-sep {
  font-size: 16px;
  font-weight: 600;
  color: #666;
}
.duration-sec {
  width: 52px;
  padding: 7px 6px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  text-align: center;
  box-sizing: border-box;
}
.content-save-btn {
  width: auto;
  margin-top: 8px;
  padding: 8px 20px;
  background: #3b82f6;
  color: #fff;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
}
.content-save-btn:hover { background: #2563eb; }
.content-save-btn:disabled { background: #ccc; cursor: not-allowed; }
.mock-count-tip {
  font-size: 12px;
  color: #888;
  margin-left: 8px;
}
.tip-ok {
  color: #22c55e;
  font-weight: 700;
}
.tip-warn {
  color: #ef4444;
  font-weight: 700;
}
.notes-textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  line-height: 1.7;
  resize: vertical;
  box-sizing: border-box;
  margin-bottom: 10px;
  color: #333;
}

/* 题目配置样式 */
.section-tip {
  font-size: 13px;
  color: #888;
  background: #fffbe6;
  border: 1px solid #ffe58f;
  border-radius: 6px;
  padding: 8px 12px;
  margin-bottom: 16px;
}
.q-loading {
  color: #999;
  font-size: 14px;
  padding: 12px 0;
}
.q-layout {
  display: flex;
  gap: 12px;
  align-items: flex-start;
}
.q-content-area {
  flex: 1;
  min-width: 0;
}
.q-type-tabs {
  display: flex;
  flex-direction: column;
  gap: 0;
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  flex-shrink: 0;
  width: 90px;
}
.q-tab {
  padding: 12px 8px;
  text-align: center;
  cursor: pointer;
  background: #f9f9f9;
  font-size: 14px;
  color: #555;
  border-bottom: 1px solid #ddd;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  transition: background 0.2s;
}
.q-tab:last-child { border-bottom: none; }
.q-tab.active {
  background: #3b82f6;
  color: #fff;
}
.q-tab-count {
  font-size: 12px;
  font-weight: 600;
  background: rgba(0,0,0,0.08);
  border-radius: 10px;
  padding: 1px 8px;
}
.q-tab.active .q-tab-count { background: rgba(255,255,255,0.25); }
.q-tab-count.full { color: #ff4d4f; }
.q-tab.active .q-tab-count.full { color: #ffd6d6; }

.question-list-wrap {
  margin-bottom: 4px;
}
.q-empty-type {
  color: #bbb;
  font-size: 13px;
  text-align: center;
  padding: 12px 0;
}
.q-item {
  background: #fff;
  border-radius: 8px;
  padding: 12px 14px;
  margin-bottom: 10px;
  border: 1px solid #e8e8e8;
}
.q-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}
.q-index {
  font-size: 13px;
  color: #999;
  flex: 1;
}
.q-del-btn {
  width: auto;
  padding: 3px 10px;
  font-size: 12px;
  background: #f56c6c;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.q-content {
  font-size: 14px;
  color: #333;
  line-height: 1.6;
  margin-bottom: 8px;
}
.q-options {
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 13px;
  color: #555;
  margin-bottom: 6px;
}
.q-answer {
  font-size: 13px;
  color: #3b82f6;
  font-weight: 500;
}
.q-analysis {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}
.q-empty {
  color: #999;
  font-size: 14px;
  text-align: center;
  padding: 20px 0;
}
.add-question-form {
  border-top: 1px solid #e0e0e0;
  padding-top: 16px;
  margin-top: 8px;
}
.add-q-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}
.add-q-header h3 {
  font-size: 15px;
  color: #333;
  margin: 0;
}
.q-full-tip {
  font-size: 13px;
  color: #f56c6c;
  font-weight: 500;
}
.q-remain-tip {
  font-size: 13px;
  color: #52c41a;
  font-weight: 500;
}
.answer-radio-group {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  align-items: center;
}
.q-answer-preview {
  font-size: 13px;
  color: #3b82f6;
  font-weight: 500;
}
</style>
