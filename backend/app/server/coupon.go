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
	"strings"
	"time"
)

type couponService struct{}

var CouponService = new(couponService)

func (s couponService) Issue(req *dto.IssueCouponReq) ([]string, error) {
	if req.UserID == 0 && strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("请填写用户ID或手机号")
	}

	user, err := s.getUser(req.UserID, req.Phone)
	if err != nil {
		return nil, err
	}

	validUntil, err := normalizeCouponDate(req.ValidUntil)
	if err != nil {
		return nil, err
	}

	title := strings.TrimSpace(req.Title)
	if title == "" {
		title = "学习卡兑换"
	}

	couponType := strings.TrimSpace(req.CouponType)
	if couponType == "" {
		couponType = "单门课程"
	}

	coupons := make([]models.Coupon, req.Count)
	codes := make([]string, req.Count)
	for i := 0; i < req.Count; i++ {
		b := make([]byte, 6)
		rand.Read(b)
		code := fmt.Sprintf("CPN%s-%s", time.Now().Format("20060102"), strings.ToUpper(hex.EncodeToString(b)))
		coupons[i] = models.Coupon{
			Code:       code,
			UserID:     user.ID,
			Title:      title,
			CouponType: couponType,
			Amount:     req.Amount,
			ValidUntil: validUntil,
			Status:     models.CouponStatusUnused,
		}
		codes[i] = code
	}

	if err := dao.CouponDao.BatchCreate(db.Db, coupons); err != nil {
		return nil, err
	}
	return codes, nil
}

func (s couponService) GetAdminList() ([]dto.CouponAdminResp, error) {
	list, err := dao.CouponDao.GetList(db.Db)
	if err != nil {
		return nil, err
	}

	userMap := s.getUserMap(list)
	result := make([]dto.CouponAdminResp, 0, len(list))
	for _, c := range list {
		u := userMap[c.UserID]
		status, statusStr := s.getCouponStatus(c)
		createdAt := ""
		if c.CreatedAt != nil {
			createdAt = c.CreatedAt.Format("2006-01-02 15:04:05")
		}
		result = append(result, dto.CouponAdminResp{
			ID:         c.ID,
			Code:       c.Code,
			UserID:     c.UserID,
			UserName:   u.Name,
			UserPhone:  u.Phone,
			Title:      c.Title,
			CouponType: c.CouponType,
			Amount:     c.Amount,
			ValidUntil: c.ValidUntil,
			Status:     int8(status),
			StatusStr:  statusStr,
			CreatedAt:  createdAt,
		})
	}
	return result, nil
}

func (s couponService) GetByUserID(userID uint) ([]dto.CouponResp, error) {
	list, err := dao.CouponDao.GetByUserID(db.Db, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.CouponResp, 0, len(list))
	for _, c := range list {
		status, statusStr := s.getCouponStatus(c)
		result = append(result, dto.CouponResp{
			ID:         c.ID,
			Code:       c.Code,
			Title:      c.Title,
			CouponType: c.CouponType,
			Amount:     c.Amount,
			ValidUntil: c.ValidUntil,
			Status:     int8(status),
			StatusStr:  statusStr,
		})
	}
	return result, nil
}

func (s couponService) getUser(userID uint, phone string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	if userID > 0 {
		user, err = dao.UserDao.GetByID(db.Db, userID)
	} else {
		user, err = dao.UserDao.GetByPhone(db.Db, strings.TrimSpace(phone))
	}
	if err != nil || user.ID == 0 {
		return models.User{}, errors.New("用户不存在")
	}
	return user, nil
}

type couponUserInfo struct {
	Name  string
	Phone string
}

func (s couponService) getUserMap(coupons []models.Coupon) map[uint]couponUserInfo {
	userIDs := make([]uint, 0, len(coupons))
	exists := make(map[uint]bool)
	for _, c := range coupons {
		if c.UserID > 0 && !exists[c.UserID] {
			exists[c.UserID] = true
			userIDs = append(userIDs, c.UserID)
		}
	}

	userMap := make(map[uint]couponUserInfo)
	if len(userIDs) == 0 {
		return userMap
	}

	var users []models.User
	db.Db.Select("id, name, nickname, phone").Where("id IN ?", userIDs).Find(&users)
	for _, u := range users {
		name := u.Name
		if name == "" {
			name = u.Nickname
		}
		if name == "" {
			name = u.Phone
		}
		userMap[u.ID] = couponUserInfo{Name: name, Phone: u.Phone}
	}
	return userMap
}

func (s couponService) getCouponStatus(c models.Coupon) (models.CouponStatus, string) {
	if c.Status == models.CouponStatusUsed {
		return models.CouponStatusUsed, models.CouponStatusMap[models.CouponStatusUsed]
	}
	if isCouponExpired(c.ValidUntil) {
		return models.CouponStatusExpired, models.CouponStatusMap[models.CouponStatusExpired]
	}
	return models.CouponStatusUnused, models.CouponStatusMap[models.CouponStatusUnused]
}

func normalizeCouponDate(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("请输入有效期")
	}
	layouts := []string{
		"2006-01-02",
		"2006.01.02",
		"2006/01/02",
		"2006-01-02 15:04:05",
		"2006.01.02 15:04:05",
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, raw, time.Local); err == nil {
			return t.Format("2006.01.02"), nil
		}
	}
	return "", errors.New("有效期格式错误，请使用 YYYY-MM-DD")
}

func isCouponExpired(validUntil string) bool {
	if validUntil == "" {
		return false
	}
	date, err := normalizeCouponDate(validUntil)
	if err != nil {
		return false
	}
	today := time.Now().Format("2006.01.02")
	return today > date
}
