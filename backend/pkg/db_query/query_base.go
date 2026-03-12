package db_query

import "gorm.io/gorm"

type QueryBase struct {
	Page     int    `form:"pageIndex" binding:"gte=1"`
	PageSize int    `form:"pageSize" binding:"gte=1"`
	Order    string `form:"order"`
	Search   string `form:"search"`
}

type PaginationQ struct {
	//Ok    string      `json:""`
	Size  int         `json:"size"`
	Page  int         `json:"page"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func GetPage[T any](req interface{}, db *gorm.DB, model T, isScan bool) (*PaginationQ, error) {
	var dataList []T
	q := NewDataQuery(req)
	db, count, err := q.GenerateQueryDB(db, true)
	if err != nil {
		return nil, err
	}
	if isScan {
		err = db.Scan(&dataList).Error
	} else {
		err = db.Find(&dataList).Error
	}
	if err != nil {
		return nil, err
	}

	page, pageSize := q.Pagination()
	return &PaginationQ{
		//Ok:    "ok",
		Size:  pageSize,
		Page:  page,
		Data:  dataList,
		Total: count,
	}, nil
}

func GetPageNotCount[T any](req interface{}, db *gorm.DB, model T, isScan bool) (*PaginationQ, error) {
	var dataList []T
	q := NewDataQuery(req)
	db, count, err := q.GenerateQueryDB(db, false)
	if err != nil {
		return nil, err
	}
	if isScan {
		err = db.Scan(&dataList).Error
	} else {
		err = db.Find(&dataList).Error
	}
	if err != nil {
		return nil, err
	}

	page, pageSize := q.Pagination()
	return &PaginationQ{
		//Ok:    "ok",
		Size:  pageSize,
		Page:  page,
		Data:  dataList,
		Total: count,
	}, nil
}
