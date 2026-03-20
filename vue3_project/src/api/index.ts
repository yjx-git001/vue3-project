import request from './request'

// 用户相关接口
export const userApi = {
  register(nickname: string, phone: string, password: string) {
    return request.post('/user/register', { nickname, phone, password })
  },
  login(phone: string, password: string) {
    return request.post('/user/login', { phone, password })
  },
  getUserInfo() {
    return request.get('/user/info')
  },
  updateProfile(data: { nickname?: string; name?: string; city?: string; organization?: string; company?: string; language?: string; avatar?: string }) {
    return request.put('/user/profile', data)
  }
}

// 课程相关接口
export const courseApi = {
  getCourseList(params?: {
    pageIndex?: number
    pageSize?: number
    courseCategory?: number
    courseTagClass?: number
  }) {
    return request.get('/courses/list', { params })
  },
  getHotCourses() {
    return request.get('/courses/hot')
  },
  getCourseDetail(ek: number, courseCategory?: number) {
    return request.get('/courses/detail', { params: { ek, courseCategory } })
  },
  getUserCourses() {
    return request.get('/user/courses')
  }
}

// 课程管理接口（后台）
export const courseAdminApi = {
  // 获取体系课程选项列表
  getSystemOptions() {
    return request.get('/courses/system_options')
  },
  // 获取课程规则选项（标签分类等）
  getQueryRule() {
    return request.get('/courses/query_rule')
  },
  // 创建体系课程
  createSystemCourse(data: {
    courseName: string
    courseTime: number
    price: number
    originalPrice: number
    image: string
    courseDetail: string
    courseIntroduction: string
  }) {
    return request.post('/courses/system', data)
  },
  // 创建单门课程
  createSingleCourse(data: {
    courseName: string
    parentEk: number
    courseTime: number
    price: number
    originalPrice: number
    image: string
    courseDetail: string
    courseTagClass: number[]
    theoreticalQuestions: {
      singleQuestion: number
      multipleQuestion: number
      judgeQuestion: number
    }
    videoQuestions: number
    attachment: number
  }) {
    return request.post('/courses/single', data)
  }
}

// 订单相关接口
export const orderApi = {
  create(courseEk: number, payType: number, cardCode?: string) {
    return request.post('/order/create', { courseEk, payType, cardCode })
  },
  createPending(courseEk: number) {
    return request.post('/order/pending', { courseEk })
  },
  pay(orderNo: string, payType: number, cardCode?: string) {
    return request.post('/order/pay', { orderNo, payType, cardCode })
  },
  getMyOrders() {
    return request.get('/order/my')
  },
  hasPurchased(courseEk: number) {
    return request.get('/order/purchased', { params: { courseEk } })
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

// 上传接口
export const uploadApi = {
  uploadFile(file: File) {
    const formData = new FormData()
    formData.append('file', file)
    return request.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }
}

// 限时秒杀管理接口
export const flashSaleAdminApi = {
  getList() {
    return request.get('/flash_sales/list')
  },
  getActive() {
    return request.get('/flash_sales/active')
  },
  create(data: { courseEk: number; price: number; startTime: string; endTime: string }) {
    return request.post('/flash_sales/create', data)
  },
  setEnable(id: number, enable: boolean) {
    return request.put(`/flash_sales/enable?id=${id}`, { enable })
  },
  delete(id: number) {
    return request.delete(`/flash_sales/delete?id=${id}`)
  }
}
export const bannerAdminApi = {
  getList() {
    return request.get('/banners/list')
  },
  create(data: { image: string; sort: number }) {
    return request.post('/banners/create', data)
  },
  delete(id: number) {
    return request.delete(`/banners/delete?id=${id}`)
  }
}

// 学时接口
export const studyApi = {
  addRecord(duration: number) {
    return request.post('/study/record', { duration })
  },
  getStats() {
    return request.get('/study/stats')
  }
}

// 卡密接口
export const cardKeyApi = {
  generate(count: number) {
    return request.post('/admin/cardkey/generate', { count })
  },
  getList() {
    return request.get('/admin/cardkey/list')
  }
}
