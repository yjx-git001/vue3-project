import { createRouter, createWebHashHistory } from 'vue-router' // 关键：替换为 createWebHashHistory
import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import CourseList from '@/views/CourseList.vue'
import CourseDetail from '@/views/CourseDetail.vue'
import SingleCourseDetail from '@/views/SingleCourseDetail.vue'
import CoursePurchase from '@/views/CoursePurchase.vue'
import Profile from '@/views/Profile.vue'
import CourseStudy from '@/views/CourseStudy.vue'
import MockExam from '@/views/MockExam.vue'
import MockExamDetail from '@/views/MockExamDetail.vue'
import ExamPractice from '@/views/ExamPractice.vue'
import WrongQuestions from '@/views/WrongQuestions.vue'
import WrongQuestionDetail from '@/views/WrongQuestionDetail.vue'
import CourseOrders from '@/views/CourseOrders.vue'
import LearningCertificate from '@/views/LearningCertificate.vue'
import CertificateDetail from '@/views/CertificateDetail.vue'
import CouponList from '@/views/CouponList.vue'
import AccountSettings from '@/views/AccountSettings.vue'
import ChatSupport from '@/views/ChatSupport.vue'
import CourseContentDetail from '@/views/CourseContentDetail.vue'

const router = createRouter({
  // 关键修改：从 createWebHistory 改为 createWebHashHistory
  history: createWebHashHistory(import.meta.env.BASE_URL), 
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
      // 首页 - 显示课程体系、限时秒杀、热门课程等
    },
    {
      path: '/login',
      name: 'login',
      component: Login
      // 登录页面
    },
    {
      path: '/courses',
      name: 'courseList',
      component: CourseList
      // 课程列表页 - 显示所有课程，支持分类和筛选
    },
    {
      path: '/course/:id',
      name: 'courseDetail',
      component: CourseDetail
      // 课程详情页 - 显示单个课程的详细信息
    },
    {
      path: '/course_content/:id',
      name: 'courseContentDetail',
      component: CourseContentDetail
      // 课程内容详情页 - 已购买课程的学习页面
    },
    {
      path: '/single_course/:id',
      name: 'singleCourseDetail',
      component: SingleCourseDetail
      // 单门课程详情页
    },
    {
      path: '/purchase',
      name: 'coursePurchase',
      component: CoursePurchase
      // 课程购买页面
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile
      // 个人中心页 - 显示用户信息、我的课程、功能菜单等
    },
    {
      path: '/course_study',
      name: 'courseStudy',
      component: CourseStudy
      // 课程学习页 - 显示用户正在学习的课程列表和学习进度
    },
    {
      path: '/mock_exam',
      name: 'mockExam',
      component: MockExam
      // 模拟考试页 - 显示可参加的模拟考试列表
    },
    {
      path: '/mock_exam_detail/:id',
      name: 'mockExamDetail',
      component: MockExamDetail
      // 模拟考试答题页 - 进行模拟考试答题
    },
    {
      path: '/exam_practice/:id',
      name: 'examPractice',
      component: ExamPractice
      // 考题训练页 - 进行考题训练，显示答案和解析
    },
    {
      path: '/wrong_questions',
      name: 'wrongQuestions',
      component: WrongQuestions
      // 错题记录页 - 显示用户的错题记录
    },
    {
      path: '/wrong_question_detail/:id',
      name: 'wrongQuestionDetail',
      component: WrongQuestionDetail
      // 错题详情页 - 显示错题的详细内容和解析
    },
    {
      path: '/course_orders',
      name: 'courseOrders',
      component: CourseOrders
      // 课程订单页 - 显示用户的课程订单列表
    },
    {
      path: '/learning_certificate',
      name: 'learningCertificate',
      component: LearningCertificate
      // 学习证书页 - 显示可下载学时证明的课程列表
    },
    {
      path: '/certificate_detail/:id',
      name: 'certificateDetail',
      component: CertificateDetail
      // 学时证书详情页 - 显示学时证书详细信息
    },
    {
      path: '/coupon_list',
      name: 'couponList',
      component: CouponList
      // 优惠卡券页 - 显示用户的优惠卡券列表
    },
    {
      path: '/account_settings',
      name: 'accountSettings',
      component: AccountSettings
      // 账号设置页 - 显示用户账号设置选项
    },
    {
      path: '/chat_support',
      name: 'chatSupport',
      component: ChatSupport
      // 在线客服页 - DeepSeek AI 对话
    }
  ]
})

export default router