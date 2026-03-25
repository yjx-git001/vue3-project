package models

type CouponStatus int8

const (
	CouponStatusUnused  CouponStatus = iota + 1 // 未使用
	CouponStatusUsed                            // 已使用（预留）
	CouponStatusExpired                         // 已过期（按有效期自动计算）
)

var CouponStatusMap = map[CouponStatus]string{
	CouponStatusUnused:  "未使用",
	CouponStatusUsed:    "已使用",
	CouponStatusExpired: "已过期",
}

type Coupon struct {
	BaseModel
	Code       string       `json:"code" gorm:"type:varchar(40);uniqueIndex;comment:卡券码"`
	UserID     uint         `json:"userId" gorm:"index;comment:用户ID"`
	Title      string       `json:"title" gorm:"type:varchar(100);comment:卡券标题"`
	CouponType string       `json:"couponType" gorm:"type:varchar(50);comment:卡券类型"`
	Amount     int          `json:"amount" gorm:"comment:优惠金额(元)"`
	ValidUntil string       `json:"validUntil" gorm:"type:varchar(20);comment:有效期"`
	Status     CouponStatus `json:"status" gorm:"default:1;comment:状态"`
}

func (Coupon) TableName() string {
	return "coupons"
}
