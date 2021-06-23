package global

import (
	"courses-gin/model/req"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	args := Config.DBConn.User + ":" + Config.DBConn.Password + "@tcp(" + Config.DBConn.Address + ")/" + Config.DBConn.DBName + "?charset=utf8&loc=Asia%2FShanghai&&parseTime=true"
	Db, err = gorm.Open("mysql", args)
	if err != nil {
		log.Fatal("connect mysql", err)
	}

	// 数据库设置
	Db.SingularTable(true)
	if Config.DBConn.MaxOpenConns != 0 {
		Db.DB().SetMaxOpenConns(Config.DBConn.MaxOpenConns)

	}
	if Config.DBConn.MaxIdleConns != 0 {
		Db.DB().SetMaxIdleConns(Config.DBConn.MaxIdleConns)
	}
}

func Paging(db *gorm.DB, params req.PageInfo, out interface{}) error {
	return db.Limit(params.PageSize).Offset(params.PageSize * (params.PageIndex - 1)).Find(out).Error
}
