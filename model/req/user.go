package req

// SignUp 注册
type SignUp struct {
	Account  string `json:"account" binding:"required"`  // 账号
	Password string `json:"password" binding:"required"` // 密码
	Name     string `json:"name"`                        // 名称
}

// Login 登录
type Login struct {
	Account  string `json:"account" binding:"required"`  // 账号
	Password string `json:"password" binding:"required"` // 密码
}

// DeleteUser 删除用户
type DeleteUser struct {
	Account string `uri:"account" binding:"required"`
}

type EditUserInfo struct {
	Name string `json:"name"`
}
