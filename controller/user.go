package controller

import (
	"courses-gin/global"
	"courses-gin/model/dao"
	"courses-gin/model/req"
	"courses-gin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

// SignUp 注册
func SignUp(c *gin.Context) {
	// 绑定参数
	signUpReq := req.SignUp{}
	err := c.ShouldBindJSON(&signUpReq)
	if err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}

	// 判断账号是否注册
	num, err := dao.CountUserByAccount(signUpReq.Account)
	if err != nil {
		respError("30000", "注册失败", c)
		return
	}
	if num != 0 {
		respError("30000", "账号已被注册", c)
		return
	}

	// 注册流程
	salt := util.GetRandomSalt()
	password := util.EncryptPassword(signUpReq.Password, salt)
	err = dao.CreateStudent(signUpReq.Account, password, salt)
	if err != nil {
		respError("30000", "注册失败", c)
		return
	}
	respSuccess(nil, c)
}

// Login 登录
func Login(c *gin.Context) {
	loginReq := req.Login{}

	// 绑定参数
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}

	// 查询用户
	user, err := dao.QueryUserByAccount(loginReq.Account)
	if err != nil {
		respError("30000", "登录失败", c)
		return
	}

	// 校验密码是否正确
	encryptPassword := util.EncryptPassword(loginReq.Password, user.Salt)
	if strings.Compare(encryptPassword, user.Password) != 0 {
		respError("20000", "密码错误", c)
	}

	// 生成sessionID
	sessionId := uuid.NewV4().String()
	err = global.RedisSetEX(fmt.Sprintf("session_%s", sessionId), loginReq.Account, time.Hour)
	if err != nil {
		respError("30000", "登录失败", c)
		return
	}
	respSuccess(sessionId, c)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	deleteReq := req.DeleteUser{}
	if err := c.ShouldBindUri(&deleteReq); err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}
	if err := dao.DeleteUserByAccount(deleteReq.Account); err != nil {
		respError("30000", "删除用户失败", c)
		return
	}
	respSuccess(nil, c)
}

// UpdateUserInfo 编辑用户信息
func UpdateUserInfo(c *gin.Context) {
	// 校验用户id
	account := c.Param("account")
	_, err := dao.QueryUserByAccount(account)
	if err != nil {
		respError("30000", "查询用户失败", c)
		return
	}

	// 绑定参数
	userInfo := req.EditUserInfo{}
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		global.Error("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}

	if err := dao.UpdateUserByAccount(dao.User{Account: account, Name: userInfo.Name}); err != nil {
		respError("30000", "更新失败", c)
		return
	}
	respSuccess(nil, c)
}
