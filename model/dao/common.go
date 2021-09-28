package dao

type Role int8

const (
	Student Role = 1 // 学生
	Teacher Role = 2 // 教师
	Admin   Role = 3 // 管理员
)
