package models

// Question 题目表
type Question struct {
	BaseModel
	CourseEk     int64  `json:"courseEk" gorm:"type:bigint;index;comment:所属单门课程ek"`
	QuestionType int8   `json:"questionType" gorm:"type:tinyint;comment:题目类型 1=单选 2=多选 3=判断"`
	Content      string `json:"content" gorm:"type:text;comment:题干"`
	OptionA      string `json:"optionA" gorm:"type:varchar(500);comment:选项A"`
	OptionB      string `json:"optionB" gorm:"type:varchar(500);comment:选项B"`
	OptionC      string `json:"optionC" gorm:"type:varchar(500);comment:选项C"`
	OptionD      string `json:"optionD" gorm:"type:varchar(500);comment:选项D"`
	Answer       string `json:"answer" gorm:"type:varchar(10);comment:正确答案"`
	Analysis     string `json:"analysis" gorm:"type:text;comment:解析"`
}

func (Question) TableName() string {
	return "questions"
}

// CourseVideo 视频目录表
type CourseVideo struct {
	BaseModel
	CourseEk int64  `json:"courseEk" gorm:"type:bigint;index;comment:所属单门课程ek"`
	Sort     int    `json:"sort" gorm:"type:int;comment:排序"`
	Title    string `json:"title" gorm:"type:varchar(200);comment:视频标题"`
	Duration string `json:"duration" gorm:"type:varchar(20);comment:时长 MM:SS"`
}

func (CourseVideo) TableName() string {
	return "course_videos"
}

// CourseAttachment 附件表
type CourseAttachment struct {
	BaseModel
	CourseEk int64  `json:"courseEk" gorm:"type:bigint;index;comment:所属单门课程ek"`
	Sort     int    `json:"sort" gorm:"type:int;comment:排序"`
	Name     string `json:"name" gorm:"type:varchar(200);comment:附件名称"`
}

func (CourseAttachment) TableName() string {
	return "course_attachments"
}

// CourseNotes 课程笔记表
type CourseNotes struct {
	BaseModel
	CourseEk int64  `json:"courseEk" gorm:"type:bigint;uniqueIndex;comment:所属单门课程ek"`
	Content  string `json:"content" gorm:"type:longtext;comment:笔记内容"`
}

func (CourseNotes) TableName() string {
	return "course_notes"
}

// MockExamConfig 模拟考试配置表（每门单门课程一条）
type MockExamConfig struct {
	BaseModel
	CourseEk       int64 `json:"courseEk" gorm:"type:bigint;uniqueIndex;comment:所属单门课程ek"`
	SingleCount    int   `json:"singleCount" gorm:"type:int;default:0;comment:单选题出题数量"`
	MultipleCount  int   `json:"multipleCount" gorm:"type:int;default:0;comment:多选题出题数量"`
	JudgeCount     int   `json:"judgeCount" gorm:"type:int;default:0;comment:判断题出题数量"`
}

func (MockExamConfig) TableName() string {
	return "mock_exam_configs"
}

// WrongQuestionRecord 错题记录表（每道答错题一条）
type WrongQuestionRecord struct {
	BaseModel
	UserID     uint  `json:"userId" gorm:"type:bigint;index;comment:用户ID"`
	CourseEk   int64 `json:"courseEk" gorm:"type:bigint;index;comment:所属单门课程ek"`
	QuestionID uint  `json:"questionId" gorm:"type:bigint;index;comment:题目ID"`
}

func (WrongQuestionRecord) TableName() string {
	return "wrong_question_records"
}

// MockExamRecord 模拟考试记录表（每次交卷一条）
type MockExamRecord struct {
	BaseModel
	UserID    uint  `json:"userId" gorm:"type:bigint;index;comment:用户ID"`
	CourseEk  int64 `json:"courseEk" gorm:"type:bigint;index;comment:所属单门课程ek"`
	Score     int   `json:"score" gorm:"type:int;default:0;comment:得分(满分100)"`
	Total     int   `json:"total" gorm:"type:int;default:0;comment:总题数"`
	Correct   int   `json:"correct" gorm:"type:int;default:0;comment:答对题数"`
}

func (MockExamRecord) TableName() string {
	return "mock_exam_records"
}
