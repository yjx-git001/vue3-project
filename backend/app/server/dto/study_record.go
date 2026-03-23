package dto

type AddStudyRecordReq struct {
	CourseEk int64 `json:"courseEk"`
	Duration int   `json:"duration" binding:"required,min=1"`
}

type StudyStatsResp struct {
	TodayDuration int `json:"todayDuration"`
	TotalDuration int `json:"totalDuration"`
}

type CourseDurationReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

type CourseDurationResp struct {
	CourseEk      int64 `json:"courseEk"`
	TotalDuration int   `json:"totalDuration"` // 秒
}
