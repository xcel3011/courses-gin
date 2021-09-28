package resp

import uuid "github.com/satori/go.uuid"

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CourseInfo struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	AuthorId       uuid.UUID `json:"authorId"`
	CourseDuration int       `json:"courseDuration"`
	LecturesNumber int       `json:"lecturesNumber"`
	Rate           int       `json:"rate"`
	CoursePrice    float64   `json:"coursePrice"`
	MonthlyPrice   float64   `json:"monthlyPrice"`
}
