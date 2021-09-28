package router

import (
	"courses-gin/global"
	"courses-gin/model/resp"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "SessionID"},
		AllowMethods:    []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:          time.Hour,
	})
}

// PermissionMiddleWare 校验session
func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := c.GetHeader("SessionID")
		isExist, account, err := global.RedisGet(fmt.Sprintf("session_%s", sessionId))
		if err != nil || !isExist {
			c.AbortWithStatusJSON(http.StatusForbidden, resp.BaseResponse{Status: "10000", Message: "用户未登录"})
			return
		}
		c.Request.Header.Set("uid", account)
		c.Next()
	}
}
