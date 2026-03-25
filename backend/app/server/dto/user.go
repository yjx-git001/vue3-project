package dto

type RegisterReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Token string       `json:"token"`
	User  UserInfoResp `json:"user"`
}

type UpdateProfileReq struct {
	Nickname     string `json:"nickname"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	City         string `json:"city"`
	Organization string `json:"organization"`
	Company      string `json:"company"`
	Language     string `json:"language"`
	Avatar       string `json:"avatar"`
}

type UserInfoResp struct {
	ID           uint   `json:"id"`
	Nickname     string `json:"nickname"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Avatar       string `json:"avatar"`
	City         string `json:"city"`
	Organization string `json:"organization"`
	Company      string `json:"company"`
	Language     string `json:"language"`
	CreatedAt    string `json:"createdAt"`
}

type AdminUserOptionResp struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
