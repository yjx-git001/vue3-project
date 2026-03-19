package models

import (
	"backend/config"
	"backend/pkg/logger"
)

func Init() error {
	logger.Sugar.Info("models auto migrate")
	err := config.DB.AutoMigrate(
		&User{},        //用户表
		&Course{},      //课程表
		&Banner{},      //轮播图表
		&FlashSale{},   //秒杀活动表
		&StudyRecord{}, //学习记录表
		&Order{},       //订单表
		&CardKey{},     //卡密表
	)
	if err != nil {
		logger.Sugar.Errorf("migrate models err: %s", err)
		return err
	}
	return nil
}
