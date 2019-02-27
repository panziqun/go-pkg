package app

type Page struct {
	PageIndex int `json:"pageIndex"` // 分页索引
	PageSize  int `json:"pageSize"`  // 分页大小
	PageCount int `json:"pageCount"` // 分页数量
	TotalRows int `json:"totalRows"` // 总行数
}
