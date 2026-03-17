package models

import (
	"backend/pkg/logger"
	"backend/config"
)

func Init() error {
	logger.Sugar.Info("models auto migrate")
	err := config.DB.AutoMigrate(
		&User{},
		&CourseSystem{},
		&CourseSingle{},
	)
	if err != nil {
		logger.Sugar.Errorf("migrate models err: %s", err)
		return err
	}
	return nil
}
