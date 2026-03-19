package models

type PayType int8

const (
	PayTypeWechat PayType = iota + 1
	PayTypeCard
)

var PayTypeMap = map[PayType]string{
	PayTypeWechat: "微信支付",
	PayTypeCard:   "卡密支付",
}

type OrderStatus int8

const (
	OrderStatusPending   OrderStatus = iota + 1 // 待支付
	OrderStatusPaid                             // 已支付
	OrderStatusCancelled                        // 已取消
)

var OrderStatusMap = map[OrderStatus]string{
	OrderStatusPending:   "待支付",
	OrderStatusPaid:      "已支付",
	OrderStatusCancelled: "已取消",
}

type Order struct {
	BaseModel
	OrderNo    string      `json:"orderNo" gorm:"type:varchar(32);uniqueIndex;comment:订单编号"`
	UserID     uint        `json:"userId" gorm:"index;comment:用户ID"`
	CourseEk   int64       `json:"courseEk" gorm:"comment:课程业务ID"`
	CourseName string      `json:"courseName" gorm:"type:varchar(200);comment:课程名称"`
	Price      int         `json:"price" gorm:"comment:支付金额(分)"`
	PayType    PayType     `json:"payType" gorm:"comment:支付方式"`
	Status     OrderStatus `json:"status" gorm:"default:2;comment:状态"`
}

func (Order) TableName() string {
	return "orders"
}
