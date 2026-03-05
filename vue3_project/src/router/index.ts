import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import CourseDetail from '@/views/CourseDetail.vue'
import SingleCourseDetail from '@/views/SingleCourseDetail.vue'
import CoursePurchase from '@/views/CoursePurchase.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/course/:id',
      name: 'courseDetail',
      component: CourseDetail
    },
    {
      path: '/single-course/:id',
      name: 'singleCourseDetail',
      component: SingleCourseDetail
    },
    {
      path: '/purchase',
      name: 'coursePurchase',
      component: CoursePurchase
    }
  ]
})

export default router
