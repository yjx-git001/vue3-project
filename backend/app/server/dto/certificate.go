package dto

// CertificateDetailResp 学时证书详情返回
type CertificateDetailResp struct {
	CourseEk     int64  `json:"courseEk"`
	CourseName   string `json:"courseName"`
	StudentName  string `json:"studentName"`
	StudyHours   string `json:"studyHours"`
	HighestScore string `json:"highestScore"`
	IssuerName   string `json:"issuerName"`
	IssueDate    string `json:"issueDate"`
	Qualified    bool   `json:"qualified"`
}
