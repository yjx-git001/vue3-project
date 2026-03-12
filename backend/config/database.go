package config

import (
	"backend/pkg/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "root:yjx020613@@tcp(127.0.0.1:3306)/port_learning?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		db.Db = DB
	}
	return err
}
