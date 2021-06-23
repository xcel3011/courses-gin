package req

// 分页
type PageInfo struct {
	PageIndex int64 `form:"pageIndex" binding:"required"`
	PageSize  int64 `form:"pageSize" binding:"required"`
}
