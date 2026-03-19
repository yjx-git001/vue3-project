package models

type CardKey struct {
	BaseModel
	Code   string `json:"code" gorm:"type:varchar(32);uniqueIndex;comment:卡密"`
	Used   bool   `json:"used" gorm:"default:false;comment:是否已使用"`
	UserID uint   `json:"userId" gorm:"comment:使用者ID"`
}

func (CardKey) TableName() string {
	return "card_keys"
}
