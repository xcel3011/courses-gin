package req

// CreateCourse 创建课程
type CreateCourse struct {
	Name           string  `json:"name"`
	CourseDuration int     `json:"courseDuration"`
	LecturesNumber int     `json:"lecturesNumber"`
	CoursePrice    float64 `json:"coursePrice"`
	MonthlyPrice   float64 `json:"monthlyPrice"`
}
