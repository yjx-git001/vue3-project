package dto

type IssueCouponReq struct {
	Count      int    `json:"count" binding:"required,min=1,max=100"`
	UserID     uint   `json:"userId"`
	Phone      string `json:"phone"`
	Amount     int    `json:"amount" binding:"required,min=1"`
	Title      string `json:"title"`
	CouponType string `json:"couponType"`
	ValidUntil string `json:"validUntil" binding:"required"`
}

type CouponAdminResp struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	UserID     uint   `json:"userId"`
	UserName   string `json:"userName"`
	UserPhone  string `json:"userPhone"`
	Title      string `json:"title"`
	CouponType string `json:"couponType"`
	Amount     int    `json:"amount"`
	ValidUntil string `json:"validUntil"`
	Status     int8   `json:"status"`
	StatusStr  string `json:"statusStr"`
	CreatedAt  string `json:"createdAt"`
}

type CouponResp struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	Title      string `json:"title"`
	CouponType string `json:"couponType"`
	Amount     int    `json:"amount"`
	ValidUntil string `json:"validUntil"`
	Status     int8   `json:"status"`
	StatusStr  string `json:"statusStr"`
}
