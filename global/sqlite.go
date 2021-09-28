package global

import (
	"courses-gin/model/req"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func init() {
	dbLog, _ := os.Create("db.log")
	db, err := gorm.Open(sqlite.Open("course.db"), &gorm.Config{
		Logger: logger.New(
			log.New(dbLog, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Error}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true},
	})
	if err != nil {
		log.Fatal("connect mysql", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(_config.DBConn.MaxIdleConns)
	sqlDB.SetMaxOpenConns(_config.DBConn.MaxOpenConns)
	Db = db
}

func Paging(db *gorm.DB, params req.PageInfo, out interface{}) error {
	return db.Limit(params.PageSize).Offset(params.PageSize * (params.PageIndex - 1)).Find(out).Error
}
