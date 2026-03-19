package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"backend/pkg/db"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type orderService struct{}

var OrderService = new(orderService)

func (s orderService) CreatePending(userID uint, courseEk int64) (*models.Order, error) {
	course, err := dao.CourseDao.GetByEk(db.Db, courseEk)
	if err != nil {
		return nil, errors.New("课程不存在")
	}

	purchased, err := dao.OrderDao.ExistsByCourseAndUser(db.Db, userID, courseEk)
	if err != nil {
		return nil, err
	}
	if purchased {
		return nil, errors.New("您已购买过该课程")
	}

	// 已有待支付订单则直接返回，不重复创建
	if existing, err := dao.OrderDao.GetPendingByCourseAndUser(db.Db, userID, courseEk); err == nil {
		return existing, nil
	}

	order := &models.Order{
		OrderNo:    fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), time.Now().UnixNano()%100000),
		UserID:     userID,
		CourseEk:   courseEk,
		CourseName: course.CourseName,
		Price:      int(course.Price),
		Status:     models.OrderStatusPending,
	}
	if err := dao.OrderDao.Create(db.Db, order); err != nil {
		return nil, err
	}
	// 3分钟后自动取消
	go func(orderID uint) {
		time.Sleep(3 * time.Minute)
		dao.OrderDao.CancelIfPending(db.Db, orderID)
	}(order.ID)
	return order, nil
}

func (s orderService) Pay(userID uint, orderNo string, req *dto.CreateOrderReq) error {
	order, err := dao.OrderDao.GetByOrderNo(db.Db, orderNo)
	if err != nil {
		return errors.New("订单不存在")
	}
	if order.UserID != userID {
		return errors.New("无权操作")
	}
	if order.Status != models.OrderStatusPending {
		return errors.New("订单状态异常")
	}

	if req.PayType == models.PayTypeCard {
		if req.CardCode == "" {
			return errors.New("请输入卡密")
		}
		key, err := dao.CardKeyDao.GetByCode(db.Db, req.CardCode)
		if err != nil {
			return errors.New("卡密不存在")
		}
		if key.Used {
			return errors.New("卡密已被使用")
		}
		if err := dao.CardKeyDao.MarkUsed(db.Db, key.ID, userID); err != nil {
			return err
		}
	}

	return dao.OrderDao.UpdateStatus(db.Db, order.ID, req.PayType, models.OrderStatusPaid)
}

func (s orderService) Create(userID uint, req *dto.CreateOrderReq) error {
	course, err := dao.CourseDao.GetByEk(db.Db, req.CourseEk)
	if err != nil {
		return errors.New("课程不存在")
	}

	purchased, err := dao.OrderDao.ExistsByCourseAndUser(db.Db, userID, req.CourseEk)
	if err != nil {
		return err
	}
	if purchased {
		return errors.New("您已购买过该课程")
	}

	if req.PayType == models.PayTypeCard {
		if req.CardCode == "" {
			return errors.New("请输入卡密")
		}
		key, err := dao.CardKeyDao.GetByCode(db.Db, req.CardCode)
		if err != nil {
			return errors.New("卡密不存在")
		}
		if key.Used {
			return errors.New("卡密已被使用")
		}
		if err := dao.CardKeyDao.MarkUsed(db.Db, key.ID, userID); err != nil {
			return err
		}
	}

	order := &models.Order{
		OrderNo:    fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), time.Now().UnixNano()%100000),
		UserID:     userID,
		CourseEk:   req.CourseEk,
		CourseName: course.CourseName,
		Price:      int(course.Price),
		PayType:    req.PayType,
		Status:     models.OrderStatusPaid,
	}
	return dao.OrderDao.Create(db.Db, order)
}

func (s orderService) GetByUserID(userID uint) ([]dto.OrderResp, error) {
	list, err := dao.OrderDao.GetByUserID(db.Db, userID)
	if err != nil {
		return nil, err
	}

	// 批量查课程图片
	ekSet := make([]int64, 0, len(list))
	for _, o := range list {
		ekSet = append(ekSet, o.CourseEk)
	}
	imageMap := make(map[int64]string)
	if len(ekSet) > 0 {
		var courses []models.Course
		db.Db.Select("ek, image").Where("ek IN ?", ekSet).Find(&courses)
		for _, c := range courses {
			imageMap[c.Ek] = c.Image
		}
	}

	result := make([]dto.OrderResp, 0, len(list))
	for _, o := range list {
		result = append(result, dto.OrderResp{
			ID:          o.ID,
			OrderNo:     o.OrderNo,
			CourseEk:    o.CourseEk,
			CourseName:  o.CourseName,
			CourseImage: imageMap[o.CourseEk],
			Price:       o.Price,
			PayType:     o.PayType,
			PayTypeStr:  models.PayTypeMap[o.PayType],
			Status:      o.Status,
			StatusStr:   models.OrderStatusMap[o.Status],
			CreatedAt:   o.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}

func (s orderService) HasPurchased(userID uint, courseEk int64) (bool, error) {
	return dao.OrderDao.ExistsByCourseAndUser(db.Db, userID, courseEk)
}

func (s orderService) GenerateCardKeys(req *dto.GenerateCardKeyReq) ([]string, error) {
	keys := make([]models.CardKey, req.Count)
	codes := make([]string, req.Count)
	for i := 0; i < req.Count; i++ {
		b := make([]byte, 8)
		rand.Read(b)
		code := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), hex.EncodeToString(b))
		keys[i] = models.CardKey{Code: code}
		codes[i] = code
	}
	if err := dao.CardKeyDao.BatchCreate(db.Db, keys); err != nil {
		return nil, err
	}
	return codes, nil
}

func (s orderService) GetCardKeyList() ([]dto.CardKeyResp, error) {
	list, err := dao.CardKeyDao.GetList(db.Db)
	if err != nil {
		return nil, err
	}

	// 收集所有使用者ID
	userIDs := make([]uint, 0)
	for _, k := range list {
		if k.Used && k.UserID > 0 {
			userIDs = append(userIDs, k.UserID)
		}
	}

	// 批量查用户真实姓名
	nameMap := make(map[uint]string)
	if len(userIDs) > 0 {
		var users []models.User
		db.Db.Select("id, name").Where("id IN ?", userIDs).Find(&users)
		for _, u := range users {
			nameMap[u.ID] = u.Name
		}
	}

	result := make([]dto.CardKeyResp, 0, len(list))
	for _, k := range list {
		result = append(result, dto.CardKeyResp{
			ID:        k.ID,
			Code:      k.Code,
			Used:      k.Used,
			UserID:    k.UserID,
			UserName:  nameMap[k.UserID],
			CreatedAt: k.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}
