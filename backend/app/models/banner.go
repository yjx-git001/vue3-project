package models

type Banner struct {
	BaseModel
	Image  string `json:"image" gorm:"type:varchar(255);comment:banner图片"`
	Sort   int    `json:"sort" gorm:"type:int;default:0;comment:排序(越小越靠前)"`
	Enable bool   `json:"enable" gorm:"type:tinyint(1);default:1;comment:是否启用"`
}

func (Banner) TableName() string {
	return "banners"
}
