package dto

import (
	"backend/app/models"
	"backend/pkg/db_query"
)

// 课程分页查询请求
type CourseGetPageReq struct {
	db_query.QueryBase
	CourseCategory models.CourseCategory `form:"courseCategory" column:"course_category" filter:"type:exact"` // 课程类型筛选
	CourseTagClass models.CourseTagClass `form:"courseTagClass"`                                               // 标签分类筛选
}

// 课程列表响应
type CoursePageResp struct {
	ID                   uint                                `json:"id"`                // 课程ID
	Ek                   int64                               `json:"ek"`                // 课程ek
	Title                string                              `json:"title"`             // 课程名称
	CourseTime           int                                 `json:"courseTime"`        // 课时
	Price                int64                               `json:"price"`             // 价格(分)
	OriginalPrice        int64                               `json:"originalPrice"`     // 原价(分)
	Image                string                              `json:"image"`             // 课程图片
	CourseCategory       models.CourseCategory               `json:"courseCategory"`    // 课程类型
	CourseCategoryStr    string                              `json:"courseCategoryStr"` // 课程类型中文
	ParentEk             int64                               `json:"parentEk"`          // 所属体系课程ek
	CourseTagClass       models.Array[models.CourseTagClass] `json:"courseTagClass"`    // 标签分类
	CourseTagClassStr    []string                            `json:"courseTagClassStr"` // 标签分类中文
	Purchased            bool                                `json:"purchased"`         // 是否已购买
}

// 课程详情查询请求
type CourseDetailReq struct {
	Ek             int64                 `form:"ek"`             // 课程ek
	CourseCategory models.CourseCategory `form:"courseCategory"` // 课程类型
}

// 体系课程下的子课程简要信息
type CourseChildResp struct {
	ID         uint   `json:"id"`         // 课程ID
	Ek         int64  `json:"ek"`         // 课程ek
	CourseName string `json:"courseName"` // 课程名称
	Price      int64  `json:"price"`      // 价格(分)
}

// 体系课程选项（用于创建单门课程时选择所属体系）
type CourseSystemOption struct {
	Ek         int64  `json:"ek"`         // 体系课程ek
	CourseName string `json:"courseName"` // 体系课程名称
}

// 体系课程详情响应
type CourseSystemDetailResp struct {
	ID                 uint                                `json:"id"`
	Ek                 int64                               `json:"ek"`
	Code               string                              `json:"code"`
	CourseName         string                              `json:"courseName"`
	CourseTime         int                                 `json:"courseTime"`
	Price              int64                               `json:"price"`
	OriginalPrice      int64                               `json:"originalPrice"`
	Image              string                              `json:"image"`
	CourseCategory     models.CourseCategory               `json:"courseCategory"`
	CourseTagClass     models.Array[models.CourseTagClass] `json:"courseTagClass"`
	CourseTagClassStr  []string                            `json:"courseTagClassStr"`
	CourseDetail       string                              `json:"courseDetail"`
	CourseContent      []CourseChildResp                   `json:"courseContent"`
	CourseIntroduction string                              `json:"courseIntroduction"`
	Purchased          bool                                `json:"purchased"`
}

// 单门课程详情响应
type CourseSingleDetailResp struct {
	ID                   uint                                `json:"id"`
	Ek                   int64                               `json:"ek"`
	Code                 string                              `json:"code"`
	CourseName           string                              `json:"courseName"`
	CourseTime           int                                 `json:"courseTime"`
	Price                int64                               `json:"price"`
	OriginalPrice        int64                               `json:"originalPrice"`
	Image                string                              `json:"image"`
	CourseDetail         string                              `json:"courseDetail"`
	CourseCategory       models.CourseCategory               `json:"courseCategory"`
	CourseTagClass       models.Array[models.CourseTagClass] `json:"courseTagClass"`
	CourseTagClassStr    []string                            `json:"courseTagClassStr"`
	TheoreticalQuestions models.Num                          `json:"theoreticalQuestions"`
	VideoQuestions       int                                 `json:"videoQuestions"`
	Attachment           int                                 `json:"attachment"`
	Purchased            bool                                `json:"purchased"`
}

// 创建体系课程请求
type CourseSystemCreateReq struct {
	CourseName         string                              `json:"courseName"`         // 课程名称
	CourseTime         int                                 `json:"courseTime"`         // 课时
	Price              int64                               `json:"price"`              // 价格(分)
	OriginalPrice      int64                               `json:"originalPrice"`      // 原价(分)
	Image              string                              `json:"image"`              // 课程图片
	CourseDetail       string                              `json:"courseDetail"`       // 课程详情
	CourseIntroduction string                              `json:"courseIntroduction"` // 课程介绍图片
	CourseTagClass     models.Array[models.CourseTagClass] `json:"courseTagClass"`     // 课程标签分类
}

// 创建单门课程请求
type CourseSingleCreateReq struct {
	CourseName           string                              `json:"courseName"`           // 课程名称
	ParentEk             int64                               `json:"parentEk"`             // 所属体系课程ek(0表示独立单门课程)
	CourseTime           int                                 `json:"courseTime"`           // 课时
	Price                int64                               `json:"price"`                // 价格(分)
	OriginalPrice        int64                               `json:"originalPrice"`        // 原价(分)
	Image                string                              `json:"image"`                // 课程图片
	CourseDetail         string                              `json:"courseDetail"`         // 课程详情
	CourseIntroduction   string                              `json:"courseIntroduction"`   // 课程介绍图片
	CourseTagClass       models.Array[models.CourseTagClass] `json:"courseTagClass"`       // 课程标签分类
	TheoreticalQuestions models.Num                          `json:"theoreticalQuestions"` // 理论题数量
	VideoQuestions       int                                 `json:"videoQuestions"`       // 视频题数量
	Attachment           int                                 `json:"attachment"`           // 附件数量
}
