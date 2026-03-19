package dto

import "backend/app/models"

type CreateOrderReq struct {
	CourseEk int64          `json:"courseEk" binding:"required"`
	PayType  models.PayType `json:"payType" binding:"required"`
	CardCode string         `json:"cardCode"`
}

type CreatePendingReq struct {
	CourseEk int64 `json:"courseEk" binding:"required"`
}

type PayOrderReq struct {
	OrderNo  string         `json:"orderNo" binding:"required"`
	PayType  models.PayType `json:"payType" binding:"required"`
	CardCode string         `json:"cardCode"`
}

type OrderResp struct {
	ID          uint               `json:"id"`
	OrderNo     string             `json:"orderNo"`
	CourseEk    int64              `json:"courseEk"`
	CourseName  string             `json:"courseName"`
	CourseImage string             `json:"courseImage"`
	Price       int                `json:"price"`
	PayType     models.PayType     `json:"payType"`
	PayTypeStr  string             `json:"payTypeStr"`
	Status      models.OrderStatus `json:"status"`
	StatusStr   string             `json:"statusStr"`
	CreatedAt   string             `json:"createdAt"`
}

type GenerateCardKeyReq struct {
	Count int `json:"count" binding:"required,min=1,max=100"`
}

type CardKeyResp struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	Used      bool   `json:"used"`
	UserID    uint   `json:"userId"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
}
