package models

import (
	"backend/app/pkg/logger"
	"backend/config"
)

func Init() error {
	logger.Sugar.Info("models auto migrate")
	err := config.DB.AutoMigrate(
		&User{},
		&Course{},
	)
	if err != nil {
		logger.Sugar.Errorf("migrate models err: %s", err)
		return err
	}
	return nil
}
