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

type CourseTagClass int8

const (
	CourseTagClassGangji CourseTagClass = iota + 1
	CourseTagClassSafety
	CourseTagClassProduction
	CourseOtherTagClass
)

var CourseTagClassMap = map[CourseTagClass]string{
	CourseTagClassGangji:     "港机设备",
	CourseTagClassSafety:     "安全风险",
	CourseTagClassProduction: "生产装卸",
	CourseOtherTagClass:      "其他",
}

type Course struct {
	BaseModel
	Ek                 int64                 `json:"ek" gorm:"type:bigint;uniqueIndex;comment:课程ek"`
	Code               string                `json:"code" gorm:"type:varchar(30);comment:编号"`
	CourseName         string                `json:"courseName" gorm:"type:varchar(100);comment:课程名称"`
	CourseCategory     CourseCategory        `json:"courseCategory" gorm:"type:tinyint;comment:课程类型 1-单门课程 2-体系课程"`
	CourseTagClass     Array[CourseTagClass] `json:"courseTagClass" gorm:"type:json;comment:课程标签分类"`
	ParentEk           int64                 `json:"parentEk" gorm:"type:bigint;default:0;comment:所属体系课程ek(0=独立课程,非0=所属体系课程ek)"`
	CourseTime         int                   `json:"courseTime" gorm:"type:int;comment:课时"`
	Price              int64                 `json:"price" gorm:"type:bigint;comment:价格(分)"`
	OriginalPrice      int64                 `json:"originalPrice" gorm:"type:bigint;comment:原价(分)"`
	Image              string                `json:"image" gorm:"type:varchar(255);comment:课程图片"`
	CourseDetail       string                `json:"courseDetail" gorm:"type:text;comment:课程详情"`
	CourseIntroduction string                `json:"courseIntroduction" gorm:"type:varchar(255);comment:课程介绍图片"`
	// 仅单门课程使用
	TheoreticalQuestions Array[Num] `json:"theoreticalQuestions" gorm:"type:json;comment:理论题数量（单门课程）"`
	VideoQuestions       int        `json:"videoQuestions" gorm:"type:int;comment:视频题数量（单门课程）"`
	Attachment           int        `json:"attachment" gorm:"type:int;comment:附件数量（单门课程）"`
	// 关联：体系课程下的单门课程列表（不创建数据库外键约束）
	Children []*Course `json:"children,omitempty" gorm:"foreignKey:ParentEk;references:Ek;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;-:migration"`
}

type Num struct {
	SingleQuestionNum   int `json:"singleQuestion"`   // 题库中单选题数量
	MultipleQuestionNum int `json:"multipleQuestion"` // 题库中多选题数量
	JudgeQuestionNum    int `json:"judgeQuestion"`    // 题库中判断题数量
}

func (Course) TableName() string {
	return "courses"
}

func (m *Course) BeforeSave(tx *gorm.DB) error {
	// 体系课程：parent_ek 必须为 0
	if m.CourseCategory == CourseCategorySystem && m.ParentEk != 0 {
		return fmt.Errorf("体系课程的 parent_ek 必须为 0")
	}
	return nil
}

func (m *Course) AfterCreate(tx *gorm.DB) (err error) {
	if m.Ek == 0 {
		err = tx.Model(m).Updates(map[string]interface{}{
			"ek":   utils.GetForeignKey(m.ID),
			"code": fmt.Sprintf("%s%04d", utils.GetPlatformCode("COURSE-"), m.ID%10000),
		}).Error
	}
	return
}
