package req

// PageInfo 分页
type PageInfo struct {
	PageIndex int `json:"pageIndex" form:"pageIndex" binding:"required"`
	PageSize  int `json:"pageSize" form:"pageSize" binding:"required"`
}
