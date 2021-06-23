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
		v1.Use(Cors()) // 跨域
		v1.POST("/signup", controller.SignUp)
		v1.PUT("/login", controller.Login)
	}
}
