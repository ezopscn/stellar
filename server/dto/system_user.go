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
	Username     *string `form:"username" json:"username"`
	Password     *string `form:"password" json:"password"`
	CNName       *string `form:"cnName" json:"cnName"`
	ENName       *string `form:"enName" json:"enName"`
	Email        *string `form:"email" json:"email"`
	Phone        *string `form:"phone" json:"phone"`
	HidePhone    *uint   `form:"hidePhone" json:"hidePhone"`
	Gender       *uint   `form:"gender" json:"gender"`
	Departments  []uint  `form:"departments" json:"departments"`
	JobPositions []uint  `form:"jobPositions" json:"jobPositions"`
	Role         *uint   `form:"role" json:"role"`
	Description  *string `form:"description" json:"description"`
}
