package dto

// 用户筛选请求
type SystemUserFilterRequest struct {
	Username          *string `form:"username" json:"username"`
	Name              *string `form:"name" json:"name"`
	Email             *string `form:"email" json:"email"`
	Phone             *string `form:"phone" json:"phone"`
	Status            *uint   `form:"status" json:"status"`
	Gender            *uint   `form:"gender" json:"gender"`
	SystemDepartment  *uint   `form:"systemDepartment" json:"systemDepartment"`
	SystemJobPosition *uint   `form:"systemJobPosition" json:"systemJobPosition"`
	SystemRole        *uint   `form:"systemRole" json:"systemRole"`
	PageNumber        *uint   `form:"pageNumber" json:"pageNumber"`
	PageSize          *uint   `form:"pageSize" json:"pageSize"`
	IsPagination      *bool   `form:"isPagination" json:"isPagination"`
}

// 添加用户请求
type SystemUserAddRequest struct {
	Username           *string `form:"username" json:"username"`
	Password           *string `form:"password" json:"password"`
	CNName             *string `form:"cnName" json:"cnName"`
	ENName             *string `form:"enName" json:"enName"`
	Email              *string `form:"email" json:"email"`
	Phone              *string `form:"phone" json:"phone"`
	HidePhone          *uint   `form:"hidePhone" json:"hidePhone"`
	Gender             *uint   `form:"gender" json:"gender"`
	SystemDepartments  []uint  `form:"systemDepartments" json:"systemDepartments"`
	SystemJobPositions []uint  `form:"systemJobPositions" json:"systemJobPositions"`
	SystemRole         *uint   `form:"systemRole" json:"systemRole"`
	Description        *string `form:"description" json:"description"`
}

// 批量添加用户请求
type SystemUserMutiAddRequest struct {
	Username           *string `form:"username" json:"username"`
	Password           *string `form:"password" json:"password"`
	CNName             *string `form:"cnName" json:"cnName"`
	ENName             *string `form:"enName" json:"enName"`
	Email              *string `form:"email" json:"email"`
	Phone              *string `form:"phone" json:"phone"`
	HidePhone          *string `form:"hidePhone" json:"hidePhone"`
	Gender             *string `form:"gender" json:"gender"`
	SystemDepartments  *string `form:"systemDepartments" json:"systemDepartments"`
	SystemJobPositions *string `form:"systemJobPositions" json:"systemJobPositions"`
	SystemRole         *string `form:"systemRole" json:"systemRole"`
	Description        *string `form:"description" json:"description"`
}

// 修改用户状态请求
type SystemUserStatusModifyRequest struct {
	Id      uint   `form:"id" json:"id"`
	Operate string `form:"operate" json:"operate"`
}

// 批量修改用户状态请求
type SystemUserStatusMutiModifyRequest struct {
	Ids     []uint `form:"ids" json:"ids"`
	Operate string `form:"operate" json:"operate"`
}
