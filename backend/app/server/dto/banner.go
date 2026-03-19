package dto

type BannerCreateReq struct {
	Image string `json:"image"` // banner图片路径
	Sort  int    `json:"sort"`  // 排序
}
