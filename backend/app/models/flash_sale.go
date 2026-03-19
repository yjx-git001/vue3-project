package models

import (
	"time"
)

type FlashSale struct {
	BaseModel
	CourseEk  int64      `json:"courseEk" gorm:"type:bigint;comment:课程ek"`
	Price     int64      `json:"price" gorm:"type:bigint;comment:秒杀价格(分)"`
	StartTime *time.Time `json:"startTime" gorm:"comment:开始时间"`
	EndTime   *time.Time `json:"endTime" gorm:"comment:结束时间"`
	Enable    bool       `json:"enable" gorm:"type:tinyint(1);default:0;comment:是否上架"`
	Course    *Course    `json:"course,omitempty" gorm:"foreignKey:CourseEk;references:Ek;-:migration"`
}

func (FlashSale) TableName() string {
	return "flash_sales"
}
