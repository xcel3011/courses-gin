package resp

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CourseInfo struct {
	Id             uint    `json:"id"`
	Name           string  `json:"name"`
	AuthorId       uint    `json:"authorId"`
	CourseDuration int     `json:"courseDuration"`
	LecturesNumber int     `json:"lecturesNumber"`
	Rate           int     `json:"rate"`
	CoursePrice    float64 `json:"coursePrice"`
	MonthlyPrice   float64 `json:"monthlyPrice"`
}
