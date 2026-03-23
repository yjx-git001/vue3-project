package dto

// ===== 题目 =====

type QuestionCreateReq struct {
	CourseEk     int64  `json:"courseEk" binding:"required"`
	QuestionType int8   `json:"questionType" binding:"required,oneof=1 2 3"`
	Content      string `json:"content" binding:"required"`
	OptionA      string `json:"optionA"`
	OptionB      string `json:"optionB"`
	OptionC      string `json:"optionC"`
	OptionD      string `json:"optionD"`
	Answer       string `json:"answer" binding:"required"`
	Analysis     string `json:"analysis"`
}

type QuestionListReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

type QuestionDeleteReq struct {
	ID uint `form:"id" binding:"required"`
}

// ===== 视频 =====

type VideoItem struct {
	Sort     int    `json:"sort"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type VideoSaveReq struct {
	CourseEk int64       `json:"courseEk" binding:"required"`
	Videos   []VideoItem `json:"videos"`
}

type VideoListReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

// ===== 附件 =====

type AttachmentItem struct {
	Sort int    `json:"sort"`
	Name string `json:"name"`
}

type AttachmentSaveReq struct {
	CourseEk    int64            `json:"courseEk" binding:"required"`
	Attachments []AttachmentItem `json:"attachments"`
}

type AttachmentListReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

// ===== 笔记 =====

type NotesSaveReq struct {
	CourseEk int64  `json:"courseEk" binding:"required"`
	Content  string `json:"content"`
}

type NotesGetReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

type NotesResp struct {
	CourseEk int64  `json:"courseEk"`
	Content  string `json:"content"`
}

// ===== 模拟考试配置 =====

type MockExamConfigSaveReq struct {
	CourseEk      int64 `json:"courseEk" binding:"required"`
	SingleCount   int   `json:"singleCount" binding:"min=0"`
	MultipleCount int   `json:"multipleCount" binding:"min=0"`
	JudgeCount    int   `json:"judgeCount" binding:"min=0"`
}

type MockExamConfigGetReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

type MockExamConfigResp struct {
	CourseEk      int64 `json:"courseEk"`
	SingleCount   int   `json:"singleCount"`
	MultipleCount int   `json:"multipleCount"`
	JudgeCount    int   `json:"judgeCount"`
}

// ===== 模拟考试记录 =====

type MockExamRecordSaveReq struct {
	CourseEk int64 `json:"courseEk" binding:"required"`
	Score    int   `json:"score"`
	Total    int   `json:"total"`
	Correct  int   `json:"correct"`
}

type MockExamRecordGetReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

type MockExamStatsResp struct {
	MockCount    int `json:"mockCount"`
	HighestScore int `json:"highestScore"`
}

// ===== 错题记录 =====

type WrongQuestionSaveReq struct {
	CourseEk    int64  `json:"courseEk" binding:"required"`
	QuestionIDs []uint `json:"questionIds"`
}

type WrongQuestionListReq struct {
	CourseEk int64 `form:"courseEk" binding:"required"`
}

type WrongQuestionCourseResp struct {
	CourseEk            int64  `json:"courseEk"`
	CourseName          string `json:"courseName"`
	CourseImage         string `json:"courseImage"`
	SingleCount         int    `json:"singleCount"`
	MultipleCount       int    `json:"multipleCount"`
	JudgeCount          int    `json:"judgeCount"`
}

type WrongQuestionItem struct {
	ID           uint   `json:"id"`
	QuestionType int8   `json:"questionType"`
	Content      string `json:"content"`
	OptionA      string `json:"optionA"`
	OptionB      string `json:"optionB"`
	OptionC      string `json:"optionC"`
	OptionD      string `json:"optionD"`
	Answer       string `json:"answer"`
	Analysis     string `json:"analysis"`
}
