package dto

// 分页信息
type Pagination struct {
	PageNumber   uint  `json:"pageNumber"`   // 页码
	PageSize     uint  `json:"pageSize"`     // 每页数据量
	Total        int64 `json:"total"`        // 数据量
	IsPagination bool  `json:"isPagination"` // 是否需要分页，默认 false，则不需要分页
}

// 数据分页返回格式
type PaginationResponse struct {
	Pagination Pagination  `json:"pagination"`
	List       interface{} `json:"list"`
}

// 分页数据设置
const (
	MaxPageSize     uint = 100 // 每次请求最大的数据量，为了数据安全性
	DefaultPageSize uint = 1   // 默认每页数据量
)

// 分页查询
func (p *Pagination) GetPaginationLimitAndOffset() (limit int, offset int) {
	if !p.IsPagination {
		return 0, 0
	}

	pageSize := p.PageSize
	if pageSize < 1 || pageSize > MaxPageSize {
		pageSize = DefaultPageSize
	}

	pageNumber := p.PageNumber
	if pageNumber < 1 {
		pageNumber = 1
	}

	p.PageSize = pageSize
	p.PageNumber = pageNumber

	limit = int(pageSize)
	offset = int(pageSize * (pageNumber - 1))
	return limit, offset
}
