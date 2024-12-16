package dto

// 用户筛选请求
type SystemUserFilterRequest struct {
	Username     *string `form:"username" json:"username"`
	Name         *string `form:"name" json:"name"`
	Email        *string `form:"email" json:"email"`
	Phone        *string `form:"phone" json:"phone"`
	Status       *uint   `form:"status" json:"status"`
	Gender       *uint   `form:"gender" json:"gender"`
	Department   *uint   `form:"department" json:"department"`
	JobPosition  *uint   `form:"jobPosition" json:"jobPosition"`
	Role         *uint   `form:"role" json:"role"`
	PageNumber   *uint   `form:"pageNumber" json:"pageNumber"`
	PageSize     *uint   `form:"pageSize" json:"pageSize"`
	IsPagination *bool   `form:"isPagination" json:"isPagination"`
}

// 添加用户请求
type SystemUserAddRequest struct {
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	CNName       *string `json:"cnName"`
	ENName       *string `json:"enName"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	Gender       *uint   `json:"gender"`
	Departments  []uint  `json:"departments"`
	JobPositions []uint  `json:"jobPositions"`
	Roles        []uint  `json:"roles"`
	Creator      *string `json:"creator"`
}
