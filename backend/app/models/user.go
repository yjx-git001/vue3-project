package models

type User struct {
	BaseModel
	Phone    string `json:"phone" gorm:"type:varchar(20);uniqueIndex"`
	Password string `json:"-" gorm:"type:varchar(100)"`
	Name     string `json:"name" gorm:"type:varchar(50)"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
