package dto

import (
	"backend/app/models"
	"backend/pkg/db_query"
)

// 课程分页查询请求
type CourseGetPageReq struct {
	db_query.QueryBase
	CourseCategory models.CourseCategory `form:"courseCategory" column:"course_category" filter:"type:exact"` // 课程类型：1-单门课程 2-体系课程，不传则查询全部
}

// 课程列表响应
type CoursePageResp struct {
	ID                uint                  `json:"id"`                // 课程ID
	Ek                int64                 `json:"ek"`                // 业务ID
	Title             string                `json:"title"`             // 课程标题
	CourseTime        int                   `json:"courseTime"`        // 课时
	Price             int64                 `json:"price"`             // 价格(分)
	OriginalPrice     int64                 `json:"originalPrice"`     // 原价(分)
	Image             string                `json:"image"`             // 课程图片
	CourseCategory    models.CourseCategory `json:"courseCategory"`    // 课程类型：1-单门课程 2-体系课程
	CourseCategoryStr string                `json:"courseCategoryStr"` // 课程类型字符串
	CourseTag         string                `json:"courseTag"`         // 课程标签
}

// 课程详情查询请求
type CourseDetailReq struct {
	Ek             int64                 `form:"ek"`             // 业务ID
	CourseCategory models.CourseCategory `form:"courseCategory"` // 课程类型：1-单门课程 2-体系课程
}

// 体系课程详情响应
type CourseSystemDetailResp struct {
	ID                 uint                         `json:"id" gorm:"-"`        // 课程ID
	Ek                 int64                        `json:"ek"`                 // 业务ID
	Code               string                       `json:"code"`               // 编号
	CourseSystemName   string                       `json:"courseSystemName"`   // 课程名称
	CourseTime         int                          `json:"courseTime"`         // 课时
	Price              int64                        `json:"price"`              // 价格(分)
	OriginalPrice      int64                        `json:"originalPrice"`      // 原价(分)
	Image              string                       `json:"image"`              // 课程图片
	CourseCategory     models.CourseCategory        `json:"courseCategory"`     // 课程类型：1-单门课程 2-体系课程
	CourseTag          string                       `json:"courseTag"`          // 课程标签
	CourseDetail       string                       `json:"courseDetail"`       // 课程详情
	CourseContent      []models.SingleCourseTagList `json:"courseContent"`      // 课程内容（包含的单门课程列表）
	CourseIntroduction string                       `json:"courseIntroduction"` // 课程介绍图片
}

// 单门课程详情响应
type CourseSingleDetailResp struct {
	ID                   uint                  `json:"id" gorm:"-"`          // 课程ID
	Ek                   int64                 `json:"ek"`                   // 业务ID
	Code                 string                `json:"code"`                 // 编号
	SingleCourseName     string                `json:"singleCourseName"`     // 单门课程名称
	CourseTime           int                   `json:"courseTime"`           // 课时
	Price                int64                 `json:"price"`                // 价格(分)
	OriginalPrice        int64                 `json:"originalPrice"`        // 原价(分)
	Image                string                `json:"image"`                // 课程图片
	CourseTag            string                `json:"courseTag"`            // 课程标签
	CourseDetail         string                `json:"courseDetail"`         // 课程详情
	CourseCategory       models.CourseCategory `json:"courseCategory"`       // 课程类型：1-单门课程 2-体系课程
	TheoreticalQuestions models.Num            `json:"theoreticalQuestions"` // 理论题数量
	VideoQuestions       int                   `json:"videoQuestions"`       // 视频题数量
	Attachment           int                   `json:"attachment"`           // 附件数量
}

// 创建体系课程请求
type CourseSystemCreateReq struct {
	CourseSystemName   string                       `json:"courseSystemName"`   // 课程名称
	CourseTime         int                          `json:"courseTime"`         // 课时
	Price              int64                        `json:"price"`              // 价格(分)
	OriginalPrice      int64                        `json:"originalPrice"`      // 原价(分)
	Image              string                       `json:"image"`              // 课程图片
	CourseCategory     models.CourseCategory        `json:"courseCategory"`     // 课程类型：1-单门课程 2-体系课程
	CourseTag          string                       `json:"courseTag"`          // 课程标签
	CourseDetail       string                       `json:"courseDetail"`       // 课程详情
	CourseContent      []models.SingleCourseTagList `json:"courseContent"`      // 课程内容（包含的单门课程列表）
	CourseIntroduction string                       `json:"courseIntroduction"` // 课程介绍图片
}

// 创建单门课程请求
type CourseSingleCreateReq struct {
	SingleCourseName     string                `json:"singleCourseName"`     // 单门课程名称
	CourseTime           int                   `json:"courseTime"`           // 课时
	Price                int64                 `json:"price"`                // 价格(分)
	OriginalPrice        int64                 `json:"originalPrice"`        // 原价(分)
	Image                string                `json:"image"`                // 课程图片
	CourseTag            string                `json:"courseTag"`            // 课程标签
	CourseDetail         string                `json:"courseDetail"`         // 课程详情
	CourseCategory       models.CourseCategory `json:"courseCategory"`       // 课程类型：1-单门课程 2-体系课程
	TheoreticalQuestions models.Num            `json:"theoreticalQuestions"` // 理论题数量
	VideoQuestions       int                   `json:"videoQuestions"`       // 视频题数量
	Attachment           int                   `json:"attachment"`           // 附件数量
}
