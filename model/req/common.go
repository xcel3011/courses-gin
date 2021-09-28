package req

// PageInfo 分页
type PageInfo struct {
	PageIndex int `form:"pageIndex" binding:"required"`
	PageSize  int `form:"pageSize" binding:"required"`
}
