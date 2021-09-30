package controller

import (
	"courses-gin/global"
	"courses-gin/model/dao"
	"courses-gin/model/req"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllCoursesList 获取课程列表
func GetAllCoursesList(c *gin.Context) {
	coursesReq := req.PageInfo{PageIndex: 1, PageSize: 9}
	if err := c.ShouldBindQuery(&coursesReq); err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}
	count, courses, err := dao.QueryAllCourses(coursesReq)
	if err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}
	c.HTML(http.StatusOK, "courses.html", gin.H{"count": count, "courses": courses})
}

// CreateCourse 创建课程
func CreateCourse(c *gin.Context) {
	courseReq := req.CreateCourse{}
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}

	// 判断用户角色
	userId := getUserId(c)
	user, err := dao.QueryUserById(userId)
	if err != nil {
		respError("30000", "查询用户角色失败", c)
		return
	}
	if user.Role == dao.Teacher {
		respError("20000", "学生用户不能创建课程", c)
		return
	}

	// 创建课程
	course := dao.Course{
		Name:           courseReq.Name,
		AuthorId:       userId,
		CourseDuration: courseReq.CourseDuration,
		LecturesNumber: courseReq.LecturesNumber,
		Rate:           5,
		CoursePrice:    courseReq.CoursePrice,
		MonthlyPrice:   courseReq.MonthlyPrice,
	}
	if err := dao.CreateCourse(course); err != nil {
		respError("30000", "创建课程失败", c)
		return
	}
}
