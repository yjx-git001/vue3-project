package dto

type FlashSaleCreateReq struct {
	CourseEk  int64  `json:"courseEk" binding:"required"`
	Price     int64  `json:"price" binding:"required"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
}

type FlashSaleSetEnableReq struct {
	Enable bool `json:"enable"`
}
