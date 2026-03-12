package models

import (
	"backend/pkg/utils"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Ek           int64  `json:"ek" gorm:"type:bigint;comment:业务ID"`
	Code         string `json:"code" gorm:"type:varchar(30);comment:编号"`
	Phone        string `json:"phone" gorm:"type:varchar(20);uniqueIndex"` //手机号码
	Password     string `json:"-" gorm:"type:varchar(100)"`                //密码
	Avatar       string `json:"avatar" gorm:"type:varchar(255)"`           //头像
	Nickname     string `json:"nickname" gorm:"type:varchar(50)"`          //昵称`
	Name         string `json:"name" gorm:"type:varchar(50)"`              //姓名
	City         string `json:"city" gorm:"type:varchar(50)"`              //城市
	Organization string `json:"organization" gorm:"type:varchar(100)"`     //机构
	Company      string `json:"company" gorm:"type:varchar(100)"`          //公司
	Language     string `json:"language" gorm:"type:varchar(20)"`          //语言
}

func (User) TableName() string {
	return "users"
}

func (m User) AfterCreate(tx *gorm.DB) (err error) {
	if m.Ek == 0 {
		err = tx.Model(m).Updates(map[string]interface{}{
			"ek":   utils.GetForeignKey(m.ID),
			"code": fmt.Sprintf("%s%04d", utils.GetPlatformCode("USER-"), m.ID%10000),
		}).Error
	}
	return
}
