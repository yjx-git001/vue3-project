package dto

import (
	"backend/app/models"
	"backend/app/pkg/db_query"
)

type CourseGetPageReq struct {
	db_query.QueryBase
	Category string `form:"category" table:"courses" column:"category" filter:"type:exact"`
}

type CoursePageResp struct {
	ID            uint                `json:"id"`
	Title         string              `json:"title"`
	Description   string              `json:"description"`
	Price         int64               `json:"price"`
	OriginalPrice int64               `json:"originalPrice"`
	Image         string              `json:"image"`
	Tag           string              `json:"tag"`
	Category      string              `json:"category"`
	Status        models.CourseStatus `json:"status"`
	StatusName    string              `json:"statusName" gorm:"-"`
}

type CourseDetailReq struct {
	ID uint `form:"id" binding:"required"`
}

type CourseDetailResp struct {
	CoursePageResp
}
