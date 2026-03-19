package dto

type AddStudyRecordReq struct {
	Duration int `json:"duration" binding:"required,min=1"`
}

type StudyStatsResp struct {
	TodayDuration int `json:"todayDuration"`
	TotalDuration int `json:"totalDuration"`
}
