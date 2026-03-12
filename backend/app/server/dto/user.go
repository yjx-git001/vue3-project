package dto

type RegisterReq struct {
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

type UserInfoResp struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Avatar string `json:"avatar"`
}
