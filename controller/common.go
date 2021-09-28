package controller

import (
	"courses-gin/model/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func respSuccess(data interface{}, c *gin.Context) {
	baseResp := &resp.BaseResponse{
		Status:  "00000",
		Message: "success",
		Data:    data,
	}

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

func getUserId(c *gin.Context) uint {
	parseUint, _ := strconv.ParseUint(c.GetHeader("uid"), 10, 64)
	return uint(parseUint)
}
