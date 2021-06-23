package dao

import "time"

type BaseModel struct {
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time `sql:"index"` // 删除时间
}

// 角色
type Role int8

const (
	Student Role = 1 // 学生
	Teacher Role = 2 // 教师
	Admin   Role = 3 // 管理员
)
