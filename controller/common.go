package controller

import (
	"courses-gin/model/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func respSuccess(data interface{}, c *gin.Context) {
	baseResp := &resp.BaseResponse{
		Status:  "00000",
		Message: "success",
		Data:    data,
	}
	c.Set("baseResp", *baseResp)
	c.Set("status", baseResp.Status)
	c.JSON(http.StatusOK, baseResp)
}
func respError(status, message string, c *gin.Context) {
	baseResp := &resp.BaseResponse{
		Status:  status,
		Message: message,
	}
	c.Set("baseResp", *baseResp)
	c.Set("status", baseResp.Status)
	c.JSON(http.StatusOK, baseResp)
}
