package services

import (
	"backend/app/server/dao"
	"backend/app/server/dto"
	"backend/pkg/db"
	"fmt"
	"time"
)

type certificateService struct{}

var CertificateService = new(certificateService)

func (s certificateService) GetDetail(userID uint, courseEk int64) (*dto.CertificateDetailResp, error) {
	course, err := dao.CourseDao.GetByEk(db.Db, courseEk)
	if err != nil {
		return nil, err
	}

	user, _ := dao.UserDao.GetByID(db.Db, userID)

	totalDuration, _ := dao.StudyRecordDao.GetCourseDuration(db.Db, userID, courseEk)
	mockCount, highestScore, _ := dao.MockExamRecordDao.GetStats(db.Db, userID, courseEk)

	qualified := totalDuration >= 180 && mockCount > 0 && highestScore >= 60
	issueDate := ""
	if qualified {
		issueDate = s.getIssueDate(userID, courseEk)
	}

	studentName := user.Name
	if studentName == "" {
		studentName = user.Nickname
	}
	if studentName == "" {
		studentName = user.Phone
	}

	issuerName := user.Organization
	if issuerName == "" {
		issuerName = user.Company
	}
	if issuerName == "" {
		issuerName = "港口学堂"
	}

	resp := &dto.CertificateDetailResp{
		CourseEk:     courseEk,
		CourseName:   course.CourseName,
		StudentName:  studentName,
		StudyHours:   formatStudyHours(totalDuration),
		HighestScore: formatScore(highestScore),
		IssuerName:   issuerName,
		IssueDate:    issueDate,
		Qualified:    qualified,
	}
	return resp, nil
}

func (s certificateService) getIssueDate(userID uint, courseEk int64) string {
	studyDate, _ := dao.StudyRecordDao.GetCourseQualifiedDate(db.Db, userID, courseEk, 180)
	passDate, _ := dao.MockExamRecordDao.GetFirstPassDate(db.Db, userID, courseEk, 60)

	if studyDate == nil || passDate == nil {
		return ""
	}

	studyDay := dateOnly(*studyDate)
	passDay := dateOnly(*passDate)
	if passDay.After(studyDay) {
		return formatDateCN(passDay)
	}
	return formatDateCN(studyDay)
}

func formatStudyHours(seconds int) string {
	if seconds < 0 {
		return "—"
	}
	hours := float64(seconds) / 3600
	return fmt.Sprintf("%.1f学时", hours)
}

func formatScore(score int) string {
	if score <= 0 {
		return "—"
	}
	return fmt.Sprintf("%d分", score)
}

func dateOnly(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func formatDateCN(t time.Time) string {
	return t.Format("2006年01月02日")
}
