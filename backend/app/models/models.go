package models

import (
	"backend/config"
	"backend/pkg/logger"
)

func Init() error {
	logger.Sugar.Info("models auto migrate")
	err := config.DB.AutoMigrate(
		&User{},             //用户表
		&Course{},           //课程表
		&Banner{},           //轮播图表
		&FlashSale{},        //秒杀活动表
		&StudyRecord{},      //学习记录表
		&Order{},            //订单表
		&CardKey{},          //卡密表
		&Question{},         //题目表
		&CourseVideo{},      //视频目录表
		&CourseAttachment{}, //附件表
		&CourseNotes{},      //课程笔记表
		&MockExamConfig{},   //模拟考试配置表
		&MockExamRecord{},   //模拟考试记录表
		&WrongQuestionRecord{}, //错题记录表
	)
	if err != nil {
		logger.Sugar.Errorf("migrate models err: %s", err)
		return err
	}
	return nil
}
