package dao

import (
	"courses-gin/global"
	"courses-gin/model/req"
	"courses-gin/model/resp"
	"gorm.io/gorm"
)

// Course 课程表结构
type Course struct {
	Name           string
	AuthorId       uint
	CourseDuration int
	LecturesNumber int
	Rate           float64
	CoursePrice    float64
	MonthlyPrice   float64
	gorm.Model
}

// QueryAllCourses 查询所有的课程
func QueryAllCourses(params req.PageInfo) (count int64, courses []resp.CourseInfo, err error) {
	q := _db.Model(&Course{}).Count(&count)
	if err = global.Paging(q, params, &courses); err != nil {
		global.Error("query all courses", err)
		return 0, nil, err
	}
	return
}

// CreateCourse 创建课程
func CreateCourse(course Course) error {
	if err := _db.Create(&course).Error; err != nil {
		global.Error("create course", err)
		return err
	}
	return nil
}
