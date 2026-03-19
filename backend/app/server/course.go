package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"backend/pkg/db"
	"backend/pkg/db_query"
	"errors"
)

type courseService struct{}

var CourseService = new(courseService)

func (s courseService) GetPage(req *dto.CourseGetPageReq, userID uint) (*db_query.PaginationQ, error) {
	var result *db_query.PaginationQ
	var err error

	if req.CourseCategory == models.CourseCategorySingle {
		result, err = dao.CourseDao.GetSinglePage(db.Db, req)
	} else if req.CourseCategory == models.CourseCategorySystem {
		result, err = dao.CourseDao.GetSystemPage(db.Db, req)
	} else {
		result, err = dao.CourseDao.GetAllPage(db.Db, req)
	}

	if err != nil {
		return nil, err
	}
	result = fillCategoryStr(result)

	// 填充购买状态
	if userID > 0 {
		list := result.Data.([]dto.CoursePageResp)
		eks := make([]int64, 0, len(list))
		for _, c := range list {
			eks = append(eks, c.Ek)
		}
		purchasedMap := getPurchasedMap(userID, eks)
		for i := range list {
			list[i].Purchased = purchasedMap[list[i].Ek]
		}
		result.Data = list
	}

	return result, nil
}

func getPurchasedMap(userID uint, eks []int64) map[int64]bool {
	var orders []models.Order
	db.Db.Where("user_id = ? AND course_ek IN ? AND status = ?", userID, eks, models.OrderStatusPaid).Find(&orders)
	result := make(map[int64]bool, len(orders))
	for _, o := range orders {
		result[o.CourseEk] = true
	}
	return result
}

func fillCategoryStr(result *db_query.PaginationQ) *db_query.PaginationQ {
	list := result.Data.([]dto.CoursePageResp)
	for i := range list {
		list[i].CourseCategoryStr = models.CourseCategoryMap[list[i].CourseCategory]
		list[i].CourseTagClassStr = convertTagClassToStr(list[i].CourseTagClass)
	}
	result.Data = list
	return result
}

func convertTagClassToStr(tags models.Array[models.CourseTagClass]) []string {
	result := make([]string, 0, len(tags))
	for _, tag := range tags {
		if str, ok := models.CourseTagClassMap[tag]; ok {
			result = append(result, str)
		}
	}
	return result
}

func (s courseService) GetDetail(req *dto.CourseDetailReq, userID uint) (interface{}, error) {
	if req.CourseCategory == models.CourseCategorySystem {
		return s.getSystemDetail(req.Ek, userID)
	} else if req.CourseCategory == models.CourseCategorySingle {
		return s.getSingleDetail(req.Ek, userID)
	}
	if result, err := s.getSingleDetail(req.Ek, userID); err == nil {
		return result, nil
	}
	return s.getSystemDetail(req.Ek, userID)
}

func (s courseService) getSystemDetail(ek int64, userID uint) (interface{}, error) {
	course, err := dao.CourseDao.GetSystemByEk(db.Db, ek)
	if err != nil {
		return nil, errors.New("课程不存在")
	}
	var children []dto.CourseChildResp
	for _, c := range course.Children {
		children = append(children, dto.CourseChildResp{
			ID:         c.ID,
			Ek:         c.Ek,
			CourseName: c.CourseName,
			Price:      c.Price,
		})
	}
	purchased := false
	if userID > 0 {
		purchased = getPurchasedMap(userID, []int64{ek})[ek]
	}
	return dto.CourseSystemDetailResp{
		ID:                 course.ID,
		Ek:                 course.Ek,
		Code:               course.Code,
		CourseName:         course.CourseName,
		CourseTime:         course.CourseTime,
		Price:              course.Price,
		OriginalPrice:      course.OriginalPrice,
		Image:              course.Image,
		CourseCategory:     course.CourseCategory,
		CourseTagClass:     course.CourseTagClass,
		CourseTagClassStr:  convertTagClassToStr(course.CourseTagClass),
		CourseDetail:       course.CourseDetail,
		CourseContent:      children,
		CourseIntroduction: course.CourseIntroduction,
		Purchased:          purchased,
	}, nil
}

func (s courseService) getSingleDetail(ek int64, userID uint) (interface{}, error) {
	course, err := dao.CourseDao.GetSingleByEk(db.Db, ek)
	if err != nil {
		return nil, errors.New("课程不存在")
	}
	var theoreticalQuestions models.Num
	if len(course.TheoreticalQuestions) > 0 {
		theoreticalQuestions = course.TheoreticalQuestions[0]
	}
	purchased := false
	if userID > 0 {
		purchased = getPurchasedMap(userID, []int64{ek})[ek]
	}
	return dto.CourseSingleDetailResp{
		ID:                   course.ID,
		Ek:                   course.Ek,
		Code:                 course.Code,
		CourseName:           course.CourseName,
		CourseTime:           course.CourseTime,
		Price:                course.Price,
		OriginalPrice:        course.OriginalPrice,
		Image:                course.Image,
		CourseDetail:         course.CourseDetail,
		CourseCategory:       course.CourseCategory,
		CourseTagClass:       course.CourseTagClass,
		CourseTagClassStr:    convertTagClassToStr(course.CourseTagClass),
		TheoreticalQuestions: theoreticalQuestions,
		VideoQuestions:       course.VideoQuestions,
		Attachment:           course.Attachment,
		Purchased:            purchased,
	}, nil
}

func (s courseService) CreateSystem(req *dto.CourseSystemCreateReq) error {
	course := &models.Course{
		CourseName:         req.CourseName,
		CourseCategory:     models.CourseCategorySystem,
		CourseTime:         req.CourseTime,
		Price:              req.Price,
		OriginalPrice:      req.OriginalPrice,
		Image:              req.Image,
		CourseDetail:       req.CourseDetail,
		CourseIntroduction: req.CourseIntroduction,
		CourseTagClass:     req.CourseTagClass,
	}
	return dao.CourseDao.CreateSystem(db.Db, course)
}

func (s courseService) CreateSingle(req *dto.CourseSingleCreateReq) error {
	course := &models.Course{
		CourseName:           req.CourseName,
		CourseCategory:       models.CourseCategorySingle,
		ParentEk:             req.ParentEk,
		CourseTime:           req.CourseTime,
		Price:                req.Price,
		OriginalPrice:        req.OriginalPrice,
		Image:                req.Image,
		CourseDetail:         req.CourseDetail,
		CourseIntroduction:   req.CourseIntroduction,
		CourseTagClass:       req.CourseTagClass,
		TheoreticalQuestions: models.Array[models.Num]{req.TheoreticalQuestions},
		VideoQuestions:       req.VideoQuestions,
		Attachment:           req.Attachment,
	}
	return dao.CourseDao.CreateSingle(db.Db, course)
}

func (s courseService) GetSystemOptions() ([]dto.CourseSystemOption, error) {
	return dao.CourseDao.GetSystemOptions(db.Db)
}

type FilterData struct {
	Name   string     `json:"name"`
	Label  string     `json:"label"`
	Values []DictData `json:"values"`
}

type DictData struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}

func (s courseService) QueryRule() (map[string]interface{}, error) {
	var ruleMap = make(map[string]interface{})
	filterList := make([]interface{}, 0)
	editList := make(map[string]*FilterData)

	// 课程标签分类筛选
	editList["courseTagClass"] = &FilterData{
		Name:  "courseTagClass",
		Label: "课程标签分类",
		Values: []DictData{
			{Text: "港机设备", Value: int64(models.CourseTagClassGangji)},
			{Text: "安全风险", Value: int64(models.CourseTagClassSafety)},
			{Text: "生产装卸", Value: int64(models.CourseTagClassProduction)},
			{Text: "其他", Value: int64(models.CourseOtherTagClass)},
		},
	}
	filterList = append(filterList, editList["courseTagClass"])

	ruleMap["filter"] = filterList
	ruleMap["edit"] = editList
	return ruleMap, nil
}
