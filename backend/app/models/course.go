package models

import (
	"backend/pkg/utils"
	"fmt"

	"gorm.io/gorm"
)

type CourseCategory int8

const (
	CourseCategorySingle CourseCategory = iota + 1
	CourseCategorySystem
)

var CourseCategoryMap = map[CourseCategory]string{
	CourseCategorySingle: "单门课程",
	CourseCategorySystem: "体系课程",
}

type CourseSystem struct {
	BaseModel
	Ek                 int64                      `json:"ek" gorm:"type:bigint;comment:体系课程ek"`
	Code               string                     `json:"code" gorm:"type:varchar(30);comment:编号"`
	CourseSystemName   string                     `json:"courseSystemName" gorm:"type:varchar(100);comment:课程名称"`
	CourseTime         int                        `json:"courseTime" gorm:"type:int;comment:课时"`
	Price              int64                      `json:"price" gorm:"type:bigint;comment:价格(分)"`
	OriginalPrice      int64                      `json:"originalPrice" gorm:"type:bigint;comment:原价(分)"`
	Image              string                     `json:"image" gorm:"type:varchar(255);comment:课程图片"`
	CourseClass        CourseCategory             `json:"courseCategory" gorm:"type:tinyint;comment:课程类型 1-单门课程 2-体系课程"`
	CourseTag          string                     `json:"courseTag" gorm:"type:varchar(100);comment:课程标签"`
	CourseDetail       string                     `json:"courseDetail" gorm:"type:varchar(255);comment:课程详情"`
	CourseContent      Array[SingleCourseTagList] `json:"courseContent" gorm:"type:text;comment:课程内容"`
	CourseIntroduction string                     `json:"courseIntroduction" gorm:"type:varchar(255);comment:课程介绍图片"`
}
type SingleCourseTagList struct {
	Ek               int64  `json:"ek" gorm:"type:bigint;comment:单门课程ek"`
	SingleCourseName string `json:"singleCourseName" gorm:"type:varchar(100);comment:单门课程名称"`
}

func (CourseSystem) TableName() string {
	return "courses_system"
}

func (m *CourseSystem) AfterCreate(tx *gorm.DB) (err error) {
	if m.Ek == 0 {
		err = tx.Model(m).Updates(map[string]interface{}{
			"ek":   utils.GetForeignKey(m.ID),
			"code": fmt.Sprintf("%s%04d", utils.GetPlatformCode("COURSE-"), m.ID%10000),
		}).Error
	}
	return
}

type CourseSingle struct {
	BaseModel
	Ek                   int64          `json:"ek" gorm:"type:bigint;comment:单门课程ek"`
	Code                 string         `json:"code" gorm:"type:varchar(30);comment:编号"`
	SingleCourseName     string         `json:"singleCourseName" gorm:"type:varchar(100);comment:单门课程名称"`
	CourseTime           int            `json:"courseTime" gorm:"type:int;comment:课时"`
	Price                int64          `json:"price" gorm:"type:bigint;comment:价格(分)"`
	OriginalPrice        int64          `json:"originalPrice" gorm:"type:bigint;comment:原价(分)"`
	Image                string         `json:"image" gorm:"type:varchar(255);comment:课程图片"`
	CourseTag            string         `json:"courseTag" gorm:"type:varchar(100);comment:课程标签"`
	CourseDetail         string         `json:"courseDetail" gorm:"type:varchar(255);comment:课程详情"`
	CourseClass          CourseCategory `json:"courseCategory" gorm:"type:tinyint;comment:课程类型 1-单门课程 2-体系课程"`
	TheoreticalQuestions Array[Num]     `json:"theoreticalQuestions" gorm:"type:json;comment:理论题数量"`
	VideoQuestions       int            `json:"videoQuestions" gorm:"type:int;comment:视频题数量"`
	Attachment           int            `json:"attachment" gorm:"type:int;comment:附件数量"`
}

type Num struct {
	SingleQuestionNum   int `json:"singleQuestion"`   // 题库中单选题数量
	MultipleQuestionNum int `json:"multipleQuestion"` // 题库中多选题数量
	JudgeQuestionNum    int `json:"judgeQuestion"`    // 题库中判断题数量
}

func (CourseSingle) TableName() string {
	return "courses_single"
}

func (m *CourseSingle) AfterCreate(tx *gorm.DB) (err error) {
	if m.Ek == 0 {
		err = tx.Model(m).Updates(map[string]interface{}{
			"ek":   utils.GetForeignKey(m.ID),
			"code": fmt.Sprintf("%s%04d", utils.GetPlatformCode("COURSE-"), m.ID%10000),
		}).Error
	}
	return
}
