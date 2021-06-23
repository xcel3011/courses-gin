package dao

import (
	"courses-gin/global"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

// 用户表
type User struct {
	Id       uuid.UUID // 用户id
	Name     string    // 用户名
	Account  string    // 电话
	Password string    // 密码
	Salt     string    // 盐
	Role     Role
	BaseModel
}

// 创建学生用户
func CreateStudent(account, password, salt string) error {
	user := User{
		Id:       uuid.NewV4(),
		Account:  account,
		Password: password,
		Salt:     salt,
		Role:     Student,
		BaseModel: BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err := global.Db.Create(user).Error
	if err != nil {
		log.Println("create student", err)
		return err
	}
	return nil
}

func CountUserByAccount(account string) (num int, err error) {
	err = global.Db.Table("user").Where("account = ? AND deleted_at != null", account).Count(&num).Error
	if err != nil {
		log.Println("count user by account", err)
	}
	return
}

func QueryUserByAccount(account string) (user User, err error) {
	err = global.Db.Where("account = ?", account).Find(&user).Error
	if err != nil {
		log.Println("query user by account", err)
	}
	return

}
