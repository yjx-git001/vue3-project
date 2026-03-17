import request from './request'

// 用户相关接口
export const userApi = {
  register(username: string, phone: string, password: string) {
    return request.post('/user/register', { username, phone, password })
  },
  login(phone: string, password: string) {
    return request.post('/user/login', { phone, password })
  },
  getUserInfo() {
    return request.get('/user/info')
  }
}

// 课程相关接口
export const courseApi = {
  getCourseList(params?: { tab?: string; filter?: string }) {
    return request.get('/courses', { params })
  },
  getCourseDetail(id: number) {
    return request.get(`/courses/${id}`)
  },
  getUserCourses() {
    return request.get('/user/courses')
  }
}

// 订单相关接口
export const orderApi = {
  createOrder(courseId: number, couponId?: number) {
    return request.post('/orders', { courseId, couponId })
  },
  getOrders() {
    return request.get('/orders')
  }
}

// 考试相关接口
export const examApi = {
  getExamList() {
    return request.get('/exams')
  },
  getExamDetail(id: number) {
    return request.get(`/exams/${id}`)
  },
  submitExam(id: number, answers: Record<number, string>) {
    return request.post(`/exams/${id}/submit`, { answers })
  }
}

// 错题相关接口
export const wrongQuestionApi = {
  getWrongQuestions() {
    return request.get('/wrong-questions')
  },
  getWrongQuestionDetail(id: number) {
    return request.get(`/wrong-questions/${id}`)
  }
}

// 证书相关接口
export const certificateApi = {
  getCertificates() {
    return request.get('/certificates')
  },
  getCertificateDetail(id: number) {
    return request.get(`/certificates/${id}`)
  }
}

// 优惠券相关接口
export const couponApi = {
  getCoupons() {
    return request.get('/coupons')
  }
}
