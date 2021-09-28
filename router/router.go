package router

import (
	"courses-gin/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(router *gin.Engine) {

	// 静态页面
	static := router.Group("/")
	{
		static.GET("/index.html", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
		static.GET("/login.html", func(c *gin.Context) { c.HTML(http.StatusOK, "login.html", nil) })
		static.GET("/courses.html", func(c *gin.Context) { c.HTML(http.StatusOK, "courses.html", nil) })
	}

	// 接口
	v1 := router.Group("/v1")
	{
		v1.Use(Cors())                        // 跨域
		v1.POST("/signup", controller.SignUp) // 注册
		v1.PUT("/login", controller.Login)    // 登录

		// 用户管理
		user := v1.Group("/user")
		user.Use(PermissionMiddleWare())
		{
			user.DELETE("/:account", controller.DeleteUser)  // 删除用户
			user.PUT("/:account", controller.UpdateUserInfo) // 更新用户信息
		}

		// 课程
		course := v1.Group("/course")
		course.Use(PermissionMiddleWare())
		{
			course.POST("/", controller.CreateCourse)            // 创建课程
			course.GET("/courses", controller.GetAllCoursesList) // 获取所有课程
		}
	}
}
