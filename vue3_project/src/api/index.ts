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
  getCertificateDetail(id: number | string) {
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
  addRecord(duration: number, courseEk?: number) {
    return request.post('/study/record', { duration, courseEk: courseEk || 0 })
  },
  getStats() {
    return request.get('/study/stats')
  },
  getCourseDuration(courseEk: number) {
    return request.get('/study/course_duration', { params: { courseEk } })
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

// 视频与附件配置接口
export const courseContentApi = {
  // 获取课程的视频列表
  getVideos(courseEk: number) {
    return request.get('/course_content/videos', { params: { courseEk } })
  },
  // 保存视频标题列表（全量覆盖）
  saveVideos(courseEk: number, videos: { sort: number; title: string; duration: string }[]) {
    return request.post('/course_content/videos', { courseEk, videos })
  },
  // 获取课程的附件列表
  getAttachments(courseEk: number) {
    return request.get('/course_content/attachments', { params: { courseEk } })
  },
  // 保存附件列表（全量覆盖）
  saveAttachments(courseEk: number, attachments: { sort: number; name: string }[]) {
    return request.post('/course_content/attachments', { courseEk, attachments })
  },
  // 获取课程笔记
  getNotes(courseEk: number) {
    return request.get('/course_content/notes', { params: { courseEk } })
  },
  // 保存课程笔记
  saveNotes(courseEk: number, content: string) {
    return request.post('/course_content/notes', { courseEk, content })
  }
}

// 题目管理接口
export const questionApi = {
  // 获取课程的题目列表
  getList(courseEk: number) {
    return request.get('/questions/list', { params: { courseEk } })
  },
  // 新增题目
  create(data: {
    courseEk: number
    questionType: number   // 1=单选 2=多选 3=判断
    content: string        // 题干
    optionA: string
    optionB: string
    optionC?: string
    optionD?: string
    answer: string         // 单选/判断: 'A'/'B'/'C'/'D'/'T'/'F'; 多选: 'AB' 'ACD' 等
    analysis?: string      // 解析
  }) {
    return request.post('/questions/create', data)
  },
  // 删除题目
  delete(id: number) {
    return request.delete(`/questions/delete?id=${id}`)
  },
  // 按题型获取题目（用于前台练习）
  getByType(courseEk: number, questionType: number) {
    return request.get('/questions/by_type', { params: { courseEk, questionType } })
  }
}

// 模拟考试配置接口
export const mockExamApi = {
  getConfig(courseEk: number) {
    return request.get('/course_content/mock_exam_config', { params: { courseEk } })
  },
  saveConfig(data: { courseEk: number; singleCount: number; multipleCount: number; judgeCount: number }) {
    return request.post('/course_content/mock_exam_config', data)
  },
  saveRecord(data: { courseEk: number; score: number; total: number; correct: number }) {
    return request.post('/course_content/mock_exam_record', data)
  },
  getStats(courseEk: number) {
    return request.get('/course_content/mock_exam_stats', { params: { courseEk } })
  },
  saveWrongQuestions(data: { courseEk: number; questionIds: number[] }) {
    return request.post('/course_content/wrong_questions', data)
  },
  getWrongQuestionCourses() {
    return request.get('/course_content/wrong_questions/courses')
  },
  getWrongQuestionList(courseEk: number) {
    return request.get('/course_content/wrong_questions/list', { params: { courseEk } })
  }
}
