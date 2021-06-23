package controller

import (
	"courses-gin/global"
	"courses-gin/model/dao"
	"courses-gin/model/req"
	"courses-gin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
	"time"
)

// 注册
func SignUp(c *gin.Context) {
	req := req.SignUp{}

	// 绑定参数
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		log.Println("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}

	// 判断账号是否注册
	num, err := dao.CountUserByAccount(req.Account)
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
	password := util.EncryptPassword(req.Password, salt)
	err = dao.CreateStudent(req.Account, password, salt)
	if err != nil {
		respError("30000", "注册失败", c)
		return
	}
	respSuccess(nil, c)
}

// 登录
func Login(c *gin.Context) {
	req := req.Login{}

	// 绑定参数
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		log.Println("bind param", err)
		respError("20000", "参数校验失败", c)
		return
	}

	// 查询用户
	user, err := dao.QueryUserByAccount(req.Account)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			respError("30000", "账号不存在", c)
			return
		}
		respError("30000", "登录失败", c)
		return
	}

	// 校验密码是否正确
	encryptPassword := util.EncryptPassword(req.Password, user.Salt)
	if strings.Compare(encryptPassword, user.Password) != 0 {
		respError("20000", "密码错误", c)
	}

	// 生成sessionID
	sessionId := uuid.NewV4().String()
	err = global.RedisSetEX(fmt.Sprintf("course_%s", req.Account), sessionId, time.Hour)
	if err != nil {
		respError("30000", "登录失败", c)
		return
	}
	respSuccess(sessionId, c)
}
