package dao

import (
	"courses-gin/global"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	Name     string // 用户名
	Account  string // 电话
	Password string // 密码
	Salt     string // 盐
	Role     Role
	gorm.Model
}

var _db = global.Db

// CreateStudent 创建学生用户
func CreateStudent(account, password, salt string) error {
	user := User{
		Account:  account,
		Password: password,
		Salt:     salt,
		Role:     Student,
	}
	err := _db.Create(&user).Error
	if err != nil {
		global.Error("create student", err)
		return err
	}
	return nil
}

// CountUserByAccount 根据用户查询数量
func CountUserByAccount(account string) (num int64, err error) {
	err = _db.Table("user").Where("account = ? AND deleted_at IS null", account).Count(&num).Error
	if err != nil {
		global.Error("count user by account", err)
	}
	return
}

// QueryUserByAccount 查询用户
func QueryUserByAccount(account string) (user User, err error) {
	err = _db.Where("account = ?", account).First(&user).Error
	if err != nil {
		global.Error("query user by account", err)
	}
	return
}

// QueryUserById 根据id查询用户
func QueryUserById(id uint) (user User, err error) {
	err = _db.Where("id = ?", id).First(&user).Error
	if err != nil {
		global.Error("query user by id", err)
	}
	return
}

// DeleteUserByAccount 删除用户
func DeleteUserByAccount(account string) error {
	if err := _db.Where("account = ?", account).Delete(&User{}).Error; err != nil {
		global.Error("delete user", err)
		return err
	}
	return nil
}

// UpdateUserByAccount 更新用户信息
func UpdateUserByAccount(user User) error {
	if err := _db.Model(&User{}).
		Where("account = ?", user.Account).
		Updates(map[string]interface{}{"name": user.Name}).Error; err != nil {
		global.Error("update user,", err)
		return err
	}
	return nil
}
