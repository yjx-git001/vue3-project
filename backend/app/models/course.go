package models

import (
	"backend/pkg/utils"
	"fmt"

	"gorm.io/gorm"
)

type CourseStatus int8

const (
	CourseStatusNormal  CourseStatus = iota + 1
	CourseStatusOffline
)

var CourseStatusMap = map[CourseStatus]string{
	CourseStatusNormal:  "正常",
	CourseStatusOffline: "下架",
}

type Course struct {
	BaseModel
	Ek            int64        `json:"ek" gorm:"type:bigint;comment:业务ID"`
	Code          string       `json:"code" gorm:"type:varchar(30);comment:编号"`
	Title         string       `json:"title" gorm:"type:varchar(100);comment:课程名称"`
	Description   string       `json:"description" gorm:"type:varchar(500)"`
	Price         int64        `json:"price" gorm:"type:bigint;comment:价格(分)"`
	OriginalPrice int64        `json:"originalPrice" gorm:"type:bigint;comment:原价(分)"`
	Image         string       `json:"image" gorm:"type:varchar(255)"`
	Tag           string       `json:"tag" gorm:"type:varchar(50)"`
	Category      string       `json:"category" gorm:"type:varchar(50)"`
	Status        CourseStatus `json:"status" gorm:"default:1;type:tinyint"`
}

func (Course) TableName() string {
	return "courses"
}

func (m Course) AfterCreate(tx *gorm.DB) (err error) {
	if m.Ek == 0 {
		err = tx.Model(m).Updates(map[string]interface{}{
			"ek":   utils.GetForeignKey(m.ID),
			"code": fmt.Sprintf("%s%04d", utils.GetPlatformCode("COURSE-"), m.ID%10000),
		}).Error
	}
	return
}
