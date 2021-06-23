package req

// 注册
type SignUp struct {
	// 账号
	Account string `json:"account" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
}

// 登录
type Login struct {
	// 账号
	Account string `json:"account" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
}
